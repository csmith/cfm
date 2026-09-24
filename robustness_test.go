package cfm_test

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"chameth.com/cfm"
	"github.com/stretchr/testify/assert"
)

func FuzzMarkdown(f *testing.F) {
	for _, seed := range []string{
		"",
		"plain paragraph text",
		"a\r\nb\rc\r\nd",
		"**bold *nested strong* tail** and _snake __deep snake___ end",
		"[a [nested (parens)](inner)](outer) [titled](/url \"title\") [empty]()",
		"> quoted\n> > double\n> > > triple\n> back to one",
		"- outer\n  - middle\n    - inner\n      1. deep ordered",
		"<div class=\"raw\">\n<p>inline <b>bold</b> text</p>\n<!-- comment -->\n</div>",
		"text <script>alert(\"x\")</script> more <br/> <img src=x onerror=alert(1)>",
		"a\x00b",
		"\x00",
		"line\x00\x00\x00end",
		"invalid \xff\xfe utf8 \x80 mid",
		"truncated \xe2\x82",
		"lone \xc3 continuation",
		"ref[^1] again[^1] named[^long-label]\n\n[^1]: note one\n\n[^long-label]: note two",
		"orphan definition\n\n[^unused]: nobody cites me",
		"```go\nfunc main() { println(\"hi\") }\n```\nafter fence",
		"~~~\nfenced tilde ~~~ inside\n~~~",
		"````md\nnested ``` fences inside\n````\n",
		"unclosed fence\n```\nnever closed",
		"| left | right |\n| --- | ---: |\n| a | b |",
		"|broken|table|\n|-|\n| row |",
		"# Head\n## Head\n### Head\ntext [link](x) **em**\n\n[^n]: n\n\nref[^n]",
		"***\n___\n* * *\n- - -",
		"<a href=\"never closed", "[unclosed](href", "*unclosed ** deeper",
		"    indented code\n    more\n\ntext",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, source string) {
		want := cfm.Markdown(source)
		assert.Equal(t, want, cfm.Markdown(source))
		assert.Equal(t, want, (cfm.Renderer{}).Markdown(source))
		assert.NotContains(t, want, "\x00")
	})
}

const concurrentDoc = `# Overview

Body one.[^note]

## Overview

Body two.[^note][^note]

~~~go
sum := a + b
~~~

[^note]: Repeated note.
`

const concurrentWant = `<h1 id="overview">Overview</h1>
<p>Body one.<sup id="fnref:1"><a href="#fn:1" class="footnote-ref" role="doc-noteref">1</a></sup></p>
<h2 id="overview-1">Overview</h2>
<p>Body two.<sup id="fnref1:1"><a href="#fn:1" class="footnote-ref" role="doc-noteref">1</a></sup>` +
	`<sup id="fnref2:1"><a href="#fn:1" class="footnote-ref" role="doc-noteref">1</a></sup></p>
<pre><code class="language-go highlighted">sum := a + b
</code></pre>
<div class="footnotes" role="doc-endnotes">
<hr>
<ol>
<li id="fn:1">
<p>Repeated note.&#160;<a href="#fnref:1" class="footnote-backref" role="doc-backlink">&#x21a9;&#xfe0e;</a>&#160;` +
	`<a href="#fnref1:1" class="footnote-backref" role="doc-backlink">&#x21a9;&#xfe0e;</a>&#160;` +
	`<a href="#fnref2:1" class="footnote-backref" role="doc-backlink">&#x21a9;&#xfe0e;</a></p>
</li>
</ol>
</div>
`

func TestConcurrentRendererIndependentCalls(t *testing.T) {
	renderer := cfm.Renderer{Highlight: func(code, language string) (string, error) {
		return `<pre><code class="language-` + language + ` highlighted">` + code + "</code></pre>\n", nil
	}}
	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 25 {
				assert.Equal(t, concurrentWant, renderer.Markdown(concurrentDoc))
			}
		})
	}
	wg.Wait()
}

func TestHighlightCallbackPerCallState(t *testing.T) {
	var calls []string
	shared := cfm.Renderer{Highlight: func(code, language string) (string, error) {
		calls = append(calls, language+"\x1f"+code)
		return `<pre class="hl" data-lang="` + language + `">` + code + `</pre>`, nil
	}}

	first := shared.Markdown("```go\na := 1\n```\n")
	assert.Equal(t, []string{"go\x1fa := 1\n"}, calls)
	assert.Contains(t, first, "<pre class=\"hl\" data-lang=\"go\">a := 1\n</pre>")

	second := shared.Markdown("```rust\nlet x = b;\n```\n")
	assert.Equal(t, []string{"go\x1fa := 1\n", "rust\x1flet x = b;\n"}, calls)
	assert.Contains(t, second, "<pre class=\"hl\" data-lang=\"rust\">let x = b;\n</pre>")
	assert.NotContains(t, second, `data-lang="go"`)

	other := cfm.Renderer{Highlight: func(code, language string) (string, error) {
		return `<code class="other ` + language + `">` + code + `</code>`, nil
	}}
	otherOut := other.Markdown("```go\na := 1\n```\n")
	assert.Contains(t, otherOut, "<code class=\"other go\">a := 1\n</code>")
	assert.NotContains(t, otherOut, `class="hl"`)

	fallback := cfm.Renderer{Highlight: func(code, language string) (string, error) {
		return "", errors.New("highlighter unavailable")
	}}
	fb := fallback.Markdown("```html\n<script>a && b</script>\n```\n")
	assert.Contains(t, fb, "&lt;script&gt;a &amp;&amp; b&lt;/script&gt;")
	assert.NotContains(t, fb, "<script>")
}

func TestAdversarialBoundedInputs(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{"delimiter run of 2000", strings.Repeat("*", 2000)},
		{"alternating delimiters 2000 tokens", strings.Repeat("*a*", 1000)},
		{"nested link brackets 2000 tokens", strings.Repeat("[", 1000) + "x" + strings.Repeat("](y)", 1000)},
		{"blockquote depth 100", strings.Repeat("> ", 100) + "deep *text*\n" +
			strings.Repeat("> ", 100) + "> - list item"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := cfm.Markdown(tt.source)
			assert.Equal(t, want, cfm.Markdown(tt.source))
			assert.Equal(t, want, (cfm.Renderer{}).Markdown(tt.source))
			assert.NotContains(t, want, "\x00")
		})
	}
}
