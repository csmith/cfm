package cfm_test

import (
	"testing"

	"chameth.com/cfm"
	"github.com/stretchr/testify/assert"
)

func TestUntrusted(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"inline tags escaped", "a <b>x</b>", "<p>a &lt;b&gt;x&lt;/b&gt;</p>\n"},
		{"processing instruction escaped", "<?php x ?>", "<p>&lt;?php x ?&gt;</p>\n"},
		{"HTML block is Markdown", "<div>\n*hi*\n</div>", "<p>&lt;div&gt;\n<em>hi</em>\n&lt;/div&gt;</p>\n"},
		{"script block escaped", "<script>alert(1)</script>", "<p>&lt;script&gt;alert(1)&lt;/script&gt;</p>\n"},
		{"HTML does not interrupt a paragraph", "a\n<div>b</div>", "<p>a\n&lt;div&gt;b&lt;/div&gt;</p>\n"},
		{"HTML in a list item", "- <div>\n  x", "<ul>\n<li>&lt;div&gt;\nx</li>\n</ul>\n"},
		{"HTML in a quote", "> <p>q</p>", "<blockquote>\n<p>&lt;p&gt;q&lt;/p&gt;</p>\n</blockquote>\n"},
		{"URL autolink kept", "<https://x.com>", `<p><a href="https://x.com">https://x.com</a></p>` + "\n"},
		{"email autolink kept", "<me@x.com>", `<p><a href="mailto:me@x.com">me@x.com</a></p>` + "\n"},
		{"unsafe autolink", "<javascript:alert(1)>", "<p><a>javascript:alert(1)</a></p>\n"},
		{"javascript link", "[x](javascript:alert(1))", "<p><a>x</a></p>\n"},
		{"mixed case scheme", "[x](JaVaScRiPt:1)", "<p><a>x</a></p>\n"},
		{"numeric reference in scheme", "[x](java&#x73;cript:1)", "<p><a>x</a></p>\n"},
		{"named reference for colon", "[x](javascript&colon;1)", "<p><a>x</a></p>\n"},
		{"backslash escape in scheme", `[x](javascript\:1)`, "<p><a>x</a></p>\n"},
		{"data link keeps title", `[x](data:text/html,y "t")`, `<p><a title="t">x</a></p>` + "\n"},
		{"unknown scheme", "[x](foo:bar)", "<p><a>x</a></p>\n"},
		{"http link", "[x](http://x)", `<p><a href="http://x">x</a></p>` + "\n"},
		{"upper case https link", "[x](HTTPS://x)", `<p><a href="HTTPS://x">x</a></p>` + "\n"},
		{"mailto link", "[x](mailto:a@b)", `<p><a href="mailto:a@b">x</a></p>` + "\n"},
		{"relative path", "[x](/p)", `<p><a href="/p">x</a></p>` + "\n"},
		{"colon after path", "[x](./a:b)", `<p><a href="./a:b">x</a></p>` + "\n"},
		{"colon in query", "[x](?q=a:b)", `<p><a href="?q=a:b">x</a></p>` + "\n"},
		{"fragment", "[x](#f)", `<p><a href="#f">x</a></p>` + "\n"},
		{"linkify kept", "www.x.com", `<p><a href="http://www.x.com">www.x.com</a></p>` + "\n"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, cfm.Renderer{Untrusted: true}.Markdown(tt.source), "Markdown(%q)", tt.source)
			},
		)
	}
}

func TestUntrustedNestedContent(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"table cell", "| a |\n|---|\n| <b>x</b> [l](javascript:1) |", "<td>&lt;b&gt;x&lt;/b&gt; <a>l</a></td>"},
		{"footnote body", "x[^1]\n\n[^1]: <b>n</b> [l](javascript:1)", "<p>&lt;b&gt;n&lt;/b&gt; <a>l</a>&#160;"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Contains(t, cfm.Renderer{Untrusted: true}.Markdown(tt.source), tt.want)
			},
		)
	}
}

func TestUntrustedNeverEmitsUnsafeSchemes(t *testing.T) {
	// Control characters and entities a browser would decode or skip in the
	// scheme must not hide it.
	sources := []string{
		"[x](java&#9;script:1)",
		"[x](java&#10;script:1)",
		"[x](&#32;javascript:1)",
		"[x](<java script:1>)",
		"[x](<\tjavascript:1>)",
		"[x](&#x6A;avascript&#x3A;1)",
	}
	for _, source := range sources {
		t.Run(
			source,
			func(t *testing.T) {
				assert.NotRegexp(t, `href="[\s\x00-\x1f]*javascript:`, cfm.Renderer{Untrusted: true}.Markdown(source))
			},
		)
	}
}

func TestTrustedKeepsHTMLAndSchemes(t *testing.T) {
	assert.Equal(
		t,
		`<p><b>x</b> <a href="javascript:1">y</a></p>`+"\n",
		cfm.Markdown("<b>x</b> [y](javascript:1)"),
	)
}
