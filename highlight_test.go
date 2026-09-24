package cfm_test

import (
	"strings"
	"testing"

	"chameth.com/cfm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type highlightCall struct {
	code     string
	language string
}

func renderHighlight(t *testing.T, source string, hl cfm.Highlighter) (string, []highlightCall) {
	t.Helper()
	var calls []highlightCall
	r := cfm.Renderer{Highlight: func(code, language string) (string, error) {
		calls = append(calls, highlightCall{code, language})
		return hl(code, language)
	}}
	return r.Markdown(source), calls
}

func TestHighlightSuccess(t *testing.T) {
	out, calls := renderHighlight(t, "```go\nfmt.Println(\"hi\")\n```\n", func(code, language string) (string, error) {
		return "<pre class=\"chroma hl-" + language + "\">" + strings.ToUpper(code) + "</pre>\n", nil
	})
	assert.Equal(t, "<pre class=\"chroma hl-go\">FMT.PRINTLN(\"HI\")\n</pre>\n", out)
	require.Len(t, calls, 1)
	assert.Equal(t, "fmt.Println(\"hi\")\n", calls[0].code)
	assert.Equal(t, "go", calls[0].language)
}

func TestHighlightCallbackOwnsNewline(t *testing.T) {
	out, _ := renderHighlight(t, "```\nx\n```\n\npara", func(code, language string) (string, error) {
		return "<code data-lang=\"" + language + "\">" + code + "</code>", nil
	})
	assert.Equal(t, "<code data-lang=\"\">x\n</code>\n<p>para</p>\n", out)
}

func TestHighlightErrorFallsBackToEscapedBlock(t *testing.T) {
	out, _ := renderHighlight(t, "```go\na < b && c > \"d\"\n```\n", func(code, language string) (string, error) {
		return "", assertError("no highlighter for " + language)
	})
	assert.Equal(t, "<pre><code class=\"language-go\">a &lt; b &amp;&amp; c &gt; &quot;d&quot;\n</code></pre>\n", out)
}

func TestHighlightUnknownLanguage(t *testing.T) {
	out, calls := renderHighlight(t, "```weirdlang\nx\n```\n", func(code, language string) (string, error) {
		if language != "go" {
			return "", assertError("unsupported language")
		}
		return "<hl>" + code + "</hl>\n", nil
	})
	assert.Equal(t, "<pre><code class=\"language-weirdlang\">x\n</code></pre>\n", out)
	require.Len(t, calls, 1)
	assert.Equal(t, "weirdlang", calls[0].language)
}

func TestHighlightNoLanguage(t *testing.T) {
	out, calls := renderHighlight(t, "~~~\nplain\n~~~\n", func(code, language string) (string, error) {
		return "<code>" + language + "|" + code + "</code>\n", nil
	})
	assert.Equal(t, "<code>|plain\n</code>\n", out)
	require.Len(t, calls, 1)
	assert.Empty(t, calls[0].language)
}

func TestHighlightInfoWordDecoded(t *testing.T) {
	out, calls := renderHighlight(t, "```C++&#32;x\ny\n```\n", func(code, language string) (string, error) {
		return "<pre>" + language + "</pre>\n", nil
	})
	assert.Equal(t, "<pre>C++ x</pre>\n", out)
	require.Len(t, calls, 1)
	assert.Equal(t, "C++ x", calls[0].language)
}

func TestHighlightNilFallsBack(t *testing.T) {
	out := (cfm.Renderer{}).Markdown("```go\nx < y\n```\n")
	assert.Equal(t, "<pre><code class=\"language-go\">x &lt; y\n</code></pre>\n", out)
}

func TestHighlightNesting(t *testing.T) {
	source := "- item:\n\n  ```go\n  in list\n  ```\n\n> ```\n> in quote\n> ```\n"
	out, calls := renderHighlight(t, source, func(code, language string) (string, error) {
		return "<hl>" + language + ":" + code + "</hl>", nil
	})
	assert.Equal(t,
		"<ul>\n<li>\n<p>item:</p>\n<hl>go:in list\n</hl></li>\n</ul>\n<blockquote>\n<hl>:in quote\n</hl></blockquote>\n",
		out)
	require.Len(t, calls, 2)
	assert.Equal(t, "in list\n", calls[0].code)
	assert.Equal(t, "go", calls[0].language)
	assert.Equal(t, "in quote\n", calls[1].code)
	assert.Empty(t, calls[1].language)
}

type assertError string

func (e assertError) Error() string { return string(e) }
