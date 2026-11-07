package cfm

import (
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func renderBlocks(source string, inline func(string, bool) string) string {
	dc := newDocContext(false)
	p := newBlockParser(dc)
	p.parse(source)
	prepassInline(p.root, dc, inline)
	return p.render(Renderer{})
}

func renderBlockFixture(t *testing.T, source string) string {
	t.Helper()
	return renderBlocks(
		source,
		func(s string, _ bool) string {
			return escapeHTML(s)
		},
	)

}

// assertGolden compares got with the test's golden file under
// testdata/golden. Run the tests with -update to rewrite the golden files.
func assertGolden(t *testing.T, got string) {
	t.Helper()
	goldie.New(t, goldie.WithFixtureDir("testdata/golden"), goldie.WithNameSuffix(".html")).
		Assert(t, t.Name(), []byte(got))
}

func TestBlockParagraphs(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"empty", "", ""},
		{"blank lines only", " \n\t\n\n  ", ""},
		{"single paragraph", "hello", "<p>hello</p>\n"},
		{"soft break", "hello\nworld", "<p>hello\nworld</p>\n"},
		{"excess blanks collapse", "one\n\n\n\n\ntwo", "<p>one</p>\n<p>two</p>\n"},
		{"blank with whitespace", "one\n   \ntwo", "<p>one</p>\n<p>two</p>\n"},
		{"leading indentation stripped", "  two spaces\n      six spaces\n\ttab", "<p>two spaces\nsix spaces\ntab</p>\n"},
		{"unicode paragraph", "Grüße aus München", "<p>Grüße aus München</p>\n"},
		{"trailing whitespace only line splits", "one\n  \ntwo", "<p>one</p>\n<p>two</p>\n"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockParagraphBoundaryPreservesTrailingSpaces(t *testing.T) {
	var got []string
	out := renderBlocks(
		"one  \ntwo  \n",
		func(s string, _ bool) string {
			got = append(got, s)
			return escapeHTML(s)
		},
	)

	assert.Equal(t, "<p>one  \ntwo</p>\n", out)
	require.Len(t, got, 1)
	// Interior trailing spaces survive for hard-break handling; only the
	// whitespace at the very end of the block is dropped.
	assert.Equal(t, "one  \ntwo", got[0])
}

func TestBlockHeadings(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{"levels", "# Level one\n## Level two\n### Level three\n#### Level four\n##### Level five\n###### Level six\n"},
		{"collisions", "# Duplicate\n\n## Duplicate\n\n### duplicate\n\n# Hello World\n\n# hello world\n"},
		{"non-ascii dropped", "# Café Üben Tokyo\n\n# 東京の見出し\n\n# naïve Übung\n"},
		{"raw markup in id", "# foo *bar* \\*baz\\*"},
		{"no space is paragraph", "#NoSpaceHere\n\n###Also bad"},
		{"trailing spaces", "# Heading with trailing spaces   \n\n##  Two spaces after marker\n"},
		{"tab after marker", "#\tFoo"},
		{"inner padding", "#                  foo                     "},
		{"closing hashes kept", "## Enclosed Heading ##\n\n### Deeper Enclosure ###\n\n# Just One #\n"},
		{"mid-line hashes kept", "### foo ### b"},
		{"hash without space kept", "# foo#"},
		{"empty markers are paragraphs", "#\n\n##\n\n######\n\nNot markers: # and ## inside text."},
		{"marker with only spaces is paragraph", "# \n##  "},
		{"seven hashes", "####### foo"},
		{"interrupts paragraph", "Foo bar\n# baz\nBar foo"},
		{"indented heading recognized", "    # Buried Heading"},
		{"indented heading interrupts paragraph", "foo\n    # bar"},
		{"up to three spaces", " ### foo\n  ## foo\n   # foo"},
		{"escaped hash is paragraph", "\\## foo"},
		{"escaped closing hashes keep id hyphen", "### foo \\###"},
		{"ids registered from quotes", "> # Solo"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assertGolden(t, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockThematicBreaks(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"hyphens", "before\n\n---\n\nafter\n\n--------------\n\ntail",
			"<p>before</p>\n<hr>\n<p>after</p>\n<hr>\n<p>tail</p>\n"},
		{"two hyphens are paragraph", "Foo\n--", "<p>Foo\n--</p>\n"},
		{"asterisks literal", "***\n\n___", "<p>***</p>\n<p>___</p>\n"},
		{"plus literal", "+++", "<p>+++</p>\n"},
		{"equals literal", "===", "<p>===</p>\n"},
		{"spaced hyphens not a break", "- - -",
			"<ul>\n<li>\n<ul>\n<li>\n<ul>\n<li></li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n"},
		{"interrupts paragraph", "Title Line\n----------", "<p>Title Line</p>\n<hr>\n"},
		{"indented interrupts paragraph", "Foo\n    ---", "<p>Foo</p>\n<hr>\n"},
		{"indented asterisks do not", "Foo\n    ***", "<p>Foo\n***</p>\n"},
		{"trailing spaces", "---   ", "<hr>\n"},
		{"inside list item", "- foo\n  ---", "<ul>\n<li>foo\n<hr>\n</li>\n</ul>\n"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockFences(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{"language", "```go\nif a < b && c != d {\n\treturn \"ok\"\n}\n```"},
		{"plain", "```\nplain fenced text <tag> & entity\n```"},
		{"tilde", "~~~html\n<p>tilde fenced</p>\n~~~"},
		{"unclosed", "```go\nnever closed before EOF"},
		{"longer closing", "````\ninner ``` fence stays content\n````"},
		{"info extra words", "```go title=main.go\npackage main\n```"},
		{"empty", "```\n```"},
		{"blank lines kept", "```\na\n\n\nb\n```"},
		{"backtick in info is paragraph", "```a```\ntext"},
		{"tilde info may hold backticks", "~~~ ```\nx\n~~~"},
		{"escape in info", "```foo\\+bar\nx\n```"},
		{"entity in info", "```f&ouml;\nx\n```"},
		{"unknown entity in info", "```&notit;"},
		{"numeric entity in info", "```C&#43;+\nx\n```"},
		{"numeric entity needs hash", "```&00;"},
		{"invalid numeric prefix", "```&z123;"},
		{"numeric zero becomes replacement", "```&#0;"},
		{"hex entity digit limit", "```&#x0000000;"},
		{"escape before entity stays literal", "```\\&amp;\nx\n```"},
		{"interrupts paragraph", "text\n```js\nvar x;\n```"},
		{"content dedented to opener", "  ```go\n  x()\n  ```"},
		{"content beyond opener indent kept", "```\n    four spaces inside\n```"},
		{"indented opener and closer", "    ```go\n    fmt.Println(\"hi\")\n    ```"},
		{"indented would-be closer stays content", "```\naaa\n    ```"},
		{"closing fence may be three deeper", "```\naaa\n   ```"},
		{"shorter closing fence stays content", "````\na\n```\nb\n````"},
		{"closing fence needs same char", "```\na\n~~~\n```"},
		{"inside quote", "> ```\n> fenced in quote\n> ```"},
		{"inside list item", "- item with code:\n\n  ```go\n  fmt.Println(\"hi\")\n  ```\n\n- next item"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assertGolden(t, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockHTML(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"type six block", "<div class=\"note\">\n  <p>Raw HTML is passed through.</p>\n</div>\n\nafter",
			"<div class=\"note\">\n  <p>Raw HTML is passed through.</p>\n</div>\n<p>after</p>\n"},
		{"interior indentation preserved", "<div class=\"tools\">\n  <ul>\n    <li>one</li>\n  </ul>\n</div>",
			"<div class=\"tools\">\n  <ul>\n    <li>one</li>\n  </ul>\n</div>\n"},
		{"script", "<script>\nif (a < b && c > d) { alert(\"raw\"); }\n</script>",
			"<script>\nif (a < b && c > d) { alert(\"raw\"); }\n</script>\n"},
		{"script ends mid-line", "<script>x</script>\npara", "<script>x</script>\n<p>para</p>\n"},
		{"comment", "<!-- a comment ignores **markup** -->", "<!-- a comment ignores **markup** -->\n"},
		{"processing instruction", "<?php\nx;\n?>\nafter", "<?php\nx;\n?>\n<p>after</p>\n"},
		{"declaration", "<!DOCTYPE html>\nafter", "<!DOCTYPE html>\n<p>after</p>\n"},
		{"cdata", "<![CDATA[x\ny]]>\nafter", "<![CDATA[x\ny]]>\n<p>after</p>\n"},
		{"type seven alone", "<a href=\"x\">\ntext", "<a href=\"x\">\ntext\n"},
		{"type seven cannot interrupt", "para\n<span>x</span>", "<p>para\n&lt;span&gt;x&lt;/span&gt;</p>\n"},
		{"type six interrupts", "para\n<div>\nx\n</div>", "<p>para</p>\n<div>\nx\n</div>\n"},
		{"inline angle brackets stay paragraph", "a < not a tag > b and </ dangling.",
			"<p>a &lt; not a tag &gt; b and &lt;/ dangling.</p>\n"},
		{"uri scheme not a block", "<http://foo.bar.baz>", "<p>&lt;http://foo.bar.baz&gt;</p>\n"},
		{"colon tag not a block", "<m:abc>", "<p>&lt;m:abc&gt;</p>\n"},
		{"dot tag not a block", "<foo.bar.baz>", "<p>&lt;foo.bar.baz&gt;</p>\n"},
		{"adjacent attributes not a block", "<a href='bar'title=title>", "<p>&lt;a href='bar'title=title&gt;</p>\n"},
		{"adjacent tags not a block", "<a><bar>", "<p>&lt;a&gt;&lt;bar&gt;</p>\n"},
		{"value forms make a block", "<a href=foo title='t' data-x=\"y\"/>\ntext",
			"<a href=foo title='t' data-x=\"y\"/>\ntext\n"},
		{"type six ends at blank", "<div>\nline\n\npara", "<div>\nline\n<p>para</p>\n"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockLists(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"tight", "- one\n- two\n- three", "<ul>\n<li>one</li>\n<li>two</li>\n<li>three</li>\n</ul>\n"},
		{"loose", "- one\n\n- two\n\n- three",
			"<ul>\n<li>\n<p>one</p>\n</li>\n<li>\n<p>two</p>\n</li>\n<li>\n<p>three</p>\n</li>\n</ul>\n"},
		{"ordered start", "3. three\n4. four\n5. five",
			"<ol start=\"3\">\n<li>three</li>\n<li>four</li>\n<li>five</li>\n</ol>\n"},
		{"ordered start one omitted", "1. a\n2. b", "<ol>\n<li>a</li>\n<li>b</li>\n</ol>\n"},
		{"ordered zero", "0. ok", "<ol start=\"0\">\n<li>ok</li>\n</ol>\n"},
		{"ordered padded", "003. ok", "<ol start=\"3\">\n<li>ok</li>\n</ol>\n"},
		{"ordered nine digits", "123456789. ok", "<ol start=\"123456789\">\n<li>ok</li>\n</ol>\n"},
		{"ordered ten digits literal", "1234567890. not ok", "<p>1234567890. not ok</p>\n"},
		{"negative literal", "-1. not ok", "<p>-1. not ok</p>\n"},
		{"tight nested", "- foo\n  - bar\n    - baz\n      - boo",
			"<ul>\n<li>foo\n<ul>\n<li>bar\n<ul>\n<li>baz\n<ul>\n<li>boo</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n" +
				"</li>\n</ul>\n"},
		{"sibling indents", "- foo\n - bar\n  - baz\n   - boo",
			"<ul>\n<li>foo</li>\n<li>bar</li>\n<li>baz</li>\n<li>boo</li>\n</ul>\n"},
		{"deep sibling indents", "- a\n - b\n  - c\n   - d\n  - e\n - f\n- g",
			"<ul>\n<li>a</li>\n<li>b</li>\n<li>c</li>\n<li>d</li>\n<li>e</li>\n<li>f</li>\n<li>g</li>\n</ul>\n"},
		{"bullet change splits", "- foo\n- bar\n+ baz",
			"<ul>\n<li>foo</li>\n<li>bar</li>\n</ul>\n<ul>\n<li>baz</li>\n</ul>\n"},
		{"delimiter change splits", "1. foo\n2. bar\n3) baz",
			"<ol>\n<li>foo</li>\n<li>bar</li>\n</ol>\n<ol start=\"3\">\n<li>baz</li>\n</ol>\n"},
		{"ordered does not interrupt unless one", "The number of windows in my house is\n14.  The number of doors is 6.",
			"<p>The number of windows in my house is\n14.  The number of doors is 6.</p>\n"},
		{"ordered one interrupts", "The number of windows in my house is\n1.  The number of doors is 6.",
			"<p>The number of windows in my house is</p>\n<ol>\n<li>The number of doors is 6.</li>\n</ol>\n"},
		{"bullet interrupts", "Foo\n- bar\n- baz", "<p>Foo</p>\n<ul>\n<li>bar</li>\n<li>baz</li>\n</ul>\n"},
		{"continuation line", "- item with\n  continuation line\n- next",
			"<ul>\n<li>item with\ncontinuation line</li>\n<li>next</li>\n</ul>\n"},
		{"multiple indented digits still siblings", "1. a\n\n  2. b\n\n   3. c",
			"<ol>\n<li>\n<p>a</p>\n</li>\n<li>\n<p>b</p>\n</li>\n<li>\n<p>c</p>\n</li>\n</ol>\n"},
		{"loose multi blank", "- foo\n\n- bar\n\n\n- baz",
			"<ul>\n<li>\n<p>foo</p>\n</li>\n<li>\n<p>bar</p>\n</li>\n<li>\n<p>baz</p>\n</li>\n</ul>\n"},
		{"loose by inner blank", "- a\n- b\n\n  c\n- d",
			"<ul>\n<li>\n<p>a</p>\n</li>\n<li>\n<p>b</p>\n<p>c</p>\n</li>\n<li>\n<p>d</p>\n</li>\n</ul>\n"},
		{"lazy item continuation", "- foo\nbar", "<ul>\n<li>foo\nbar</li>\n</ul>\n"},
		{"one space below content column closes", "- one\n\n two", "<ul>\n<li>one</li>\n</ul>\n<p>two</p>\n"},
		{"two spaces continues", "- one\n\n  two", "<ul>\n<li>\n<p>one</p>\n<p>two</p>\n</li>\n</ul>\n"},
		{"wide marker", " -    one\n\n      two", "<ul>\n<li>\n<p>one</p>\n<p>two</p>\n</li>\n</ul>\n"},
		{"five spaces after marker", "-     foo\n\n  bar", "<ul>\n<li>foo</li>\n</ul>\n<p>bar</p>\n"},
		{"marker only spaces then content", "-   \n  foo", "<ul>\n<li>foo</li>\n</ul>\n"},
		{"empty item closed by blank then text", "-\n\n  foo", "<ul>\n<li></li>\n</ul>\n<p>foo</p>\n"},
		{"empty items tight", "- foo\n-\n- bar", "<ul>\n<li>foo</li>\n<li></li>\n<li>bar</li>\n</ul>\n"},
		{"empty item with spaces", "- foo\n-   \n- bar", "<ul>\n<li>foo</li>\n<li></li>\n<li>bar</li>\n</ul>\n"},
		{"lone marker", "*", "<ul>\n<li></li>\n</ul>\n"},
		{"markers cannot interrupt when empty", "foo\n*\n\nfoo\n1.", "<p>foo\n*</p>\n<p>foo\n1.</p>\n"},
		{"no space after marker", "-one\n\n2.two", "<p>-one</p>\n<p>2.two</p>\n"},
		{"nested via wide indent", "- a\n      - b", "<ul>\n<li>a\n- b</li>\n</ul>\n"},
		{"nested at three deeper", "- a\n     - b", "<ul>\n<li>a\n<ul>\n<li>b</li>\n</ul>\n</li>\n</ul>\n"},
		{"spec d-e chain", "- a\n - b\n  - c\n   - d\n    - e",
			"<ul>\n<li>a</li>\n<li>b</li>\n<li>c</li>\n<li>d\n- e</li>\n</ul>\n"},
		{"sibling after nested", "- a\n   - d\n    - e", "<ul>\n<li>a\n<ul>\n<li>d</li>\n<li>e</li>\n</ul>\n</li>\n</ul>\n"},
		{"nested chain wide", "- a\n  - b\n      - c",
			"<ul>\n<li>a\n<ul>\n<li>b\n<ul>\n<li>c</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n"},
		{"ordered nested bullet", "10) foo\n    - bar",
			"<ol start=\"10\">\n<li>foo\n<ul>\n<li>bar</li>\n</ul>\n</li>\n</ol>\n"},
		{"ordered then bullet sibling", "10) foo\n   - bar",
			"<ol start=\"10\">\n<li>foo</li>\n</ol>\n<ul>\n<li>bar</li>\n</ul>\n"},
		{"nested bullet ordered", "1. - 2. foo",
			"<ol>\n<li>\n<ul>\n<li>\n<ol start=\"2\">\n<li>foo</li>\n</ol>\n</li>\n</ul>\n</li>\n</ol>\n"},
		{"marker line item", "- - foo", "<ul>\n<li>\n<ul>\n<li>foo</li>\n</ul>\n</li>\n</ul>\n"},
		{"loose nested inner only", "- a\n  - b\n\n    c\n- d",
			"<ul>\n<li>a\n<ul>\n<li>\n<p>b</p>\n<p>c</p>\n</li>\n</ul>\n</li>\n<li>d</li>\n</ul>\n"},
		{"tight with quote", "* a\n  > b\n  >\n* c",
			"<ul>\n<li>a\n<blockquote>\n<p>b</p>\n</blockquote>\n</li>\n<li>c</li>\n</ul>\n"},
		{"tight with quote and fence", "- a\n  > b\n  ```\n  c\n  ```\n- d",
			"<ul>\n<li>a\n<blockquote>\n<p>b</p>\n</blockquote>\n<pre><code>c\n</code></pre>\n</li>\n<li>d</li>\n</ul>\n"},
		{"fence in tight item", "- a\n- ```\n  b\n\n\n  ```\n- c",
			"<ul>\n<li>a</li>\n<li>\n<pre><code>b\n\n\n</code></pre>\n</li>\n<li>c</li>\n</ul>\n"},
		{"fence then para item", "1. ```\n   foo\n   ```\n\n   bar",
			"<ol>\n<li>\n<pre><code>foo\n</code></pre>\n<p>bar</p>\n</li>\n</ol>\n"},
		{"loose with nested", "* foo\n  * bar\n\n  baz",
			"<ul>\n<li>\n<p>foo</p>\n<ul>\n<li>bar</li>\n</ul>\n<p>baz</p>\n</li>\n</ul>\n"},
		{"two loose sublists", "- a\n  - b\n  - c\n\n- d\n  - e\n  - f",
			"<ul>\n<li>\n<p>a</p>\n<ul>\n<li>b</li>\n<li>c</li>\n</ul>\n</li>\n<li>\n<p>d</p>\n<ul>\n<li>e</li>\n" +
				"<li>f</li>\n</ul>\n</li>\n</ul>\n"},
		{"html between lists", "- foo\n- bar\n\n<!-- -->\n\n- baz\n- bim",
			"<ul>\n<li>foo</li>\n<li>bar</li>\n</ul>\n<!-- -->\n<ul>\n<li>baz</li>\n<li>bim</li>\n</ul>\n"},
		{"checkbox literal", "- [ ] unchecked task\n- [x] checked task",
			"<ul>\n<li>[ ] unchecked task</li>\n<li>[x] checked task</li>\n</ul>\n"},
		{"indented list recognized", "    - buried item", "<ul>\n<li>buried item</li>\n</ul>\n"},
		{"tab content column", "  - foo\n\n\tbar", "<ul>\n<li>\n<p>foo</p>\n<p>bar</p>\n</li>\n</ul>\n"},
		{"tab nested", " - foo\n   - bar\n\t - baz",
			"<ul>\n<li>foo\n<ul>\n<li>bar\n<ul>\n<li>baz</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n"},
		{"deep nested blank", "- foo\n  - bar\n    - baz\n\n\n      bim",
			"<ul>\n<li>foo\n<ul>\n<li>bar\n<ul>\n<li>\n<p>baz</p>\n<p>bim</p>\n</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n"},
		{"ordered nested paragraph", "  1.  A paragraph\n    with two lines.",
			"<ol>\n<li>A paragraph\nwith two lines.</li>\n</ol>\n"},
		{"after paragraph blank low indent", "   foo\n\nbar", "<p>foo</p>\n<p>bar</p>\n"},
		{"quote interrupted by heading", "- foo\n# heading", "<ul>\n<li>foo</li>\n</ul>\n<h1 id=\"heading\">heading</h1>\n"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockQuotes(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{"lazy continuation", "> First quoted line\nlazy continuation of the same paragraph."},
		{"nested", "> outer\n>\n> > inner\n>\n> back to outer"},
		{"rich", "> # Heading in quote\n>\n> - a list item\n>\n> ```\n> fenced in quote\n> ```"},
		{"lazy through list", "> - item\nlazy"},
		{"heading closes quote", "> para\n# heading"},
		{"quote after item", "- foo\n> quote"},
		{"lazy in nested quotes and list", "> 1. > Blockquote\ncontinued here."},
		{"prefixed lazy in nested quotes and list", "> 1. > Blockquote\n> continued here."},
		{"double prefix then spaced lazy", ">>- one\n>>\n  >  > two"},
		{"deep quoted ordered item", "   > > 1.  one\n>>\n>>     two"},
		{"indented deeper quote opens fresh", "> para\n      > deeper"},
		{"blank closes quote", "> a\n\nb"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assertGolden(t, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockIndentedRecognition(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"indented prose", "    indented four spaces\n\n    second block\n    continuation line",
			"<p>indented four spaces</p>\n<p>second block\ncontinuation line</p>\n"},
		{"indented heading then list", "    # Buried Heading\n\n    - buried item",
			"<h1 id=\"buried-heading\">Buried Heading</h1>\n<ul>\n<li>buried item</li>\n</ul>\n"},
		{"indented fence", "    ```go\n    fmt.Println(\"hi\")\n    ```",
			"<pre><code class=\"language-go\">fmt.Println(&quot;hi&quot;)\n</code></pre>\n"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockInlineBoundary(t *testing.T) {
	calls := 0
	out := renderBlocks(
		"para one\n\n# Head\n",
		func(s string, _ bool) string {
			calls++
			return "<" + s + ">"
		},
	)

	assert.Equal(t, "<p><para one></p>\n<h1 id=\"head\"><Head></h1>\n", out)
	assert.Equal(t, 2, calls)

	// Fenced code and raw HTML never reach the inline callback.
	calls = 0
	out = renderBlocks(
		"```go\nx()\n```\n\n<div>\ny\n</div>\n",
		func(s string, _ bool) string {
			calls++
			return escapeHTML(s)
		},
	)

	assert.Equal(t, "<pre><code class=\"language-go\">x()\n</code></pre>\n<div>\ny\n</div>\n", out)
	assert.Zero(t, calls)
}

func TestGenerateID(t *testing.T) {
	ids := &headingIDs{}
	assert.Equal(t, "hello-world", ids.generate("Hello World"))
	assert.Equal(t, "a--b", ids.generate("a_ b"))
	assert.Equal(t, "heading", ids.generate(" 東京 "))
	assert.Equal(t, "---", ids.generate("---"))
	assert.Equal(t, "x1-y2", ids.generate("X1! Y2?"))
	assert.Equal(t, "caf", ids.generate("Café"))
}

func TestGenerateIDCollisions(t *testing.T) {
	ids := &headingIDs{}
	assert.Equal(t, "a", ids.generate("a"))
	assert.Equal(t, "a-2", ids.generate("a 2"))
	assert.Equal(t, "a-1", ids.generate("a"))
	assert.Equal(t, "a-3", ids.generate("a"))
	assert.Equal(t, "a-4", ids.generate("a"))
	assert.Equal(t, "a-5", ids.generate("a 5"))
	assert.Equal(t, "a-6", ids.generate("a"))
	assert.Equal(t, "a-1-1", ids.generate("a 1"))
}

func TestBlockTables(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{"basic", "| One | Two |\n|-----|-----|\n| 1 | 2 |"},
		{"after paragraph prefix", "intro\n| a | b |\n| --- | --- |\n| 1 | 2 |"},
		{"stops at interruption", "| a | b |\n| --- | --- |\n| 1 | 2 |\n# Head"},
		{"count mismatch stays paragraph", "no | table | here\n:--- | :---"},
		{"paragraph before blank pair stays text", "just | some | text\n\nno | table | here\n:--- | :---"},
		{"rows fit the delimiter columns", "| A | B | C |\n|---|---|---|\n| short |\n| x | y | z | extra |"},
		{"multiple prefixed tables", "intro\n| a |\n|-|\n\noutro\n| b |\n|-|"},
		{"no table without delimiter", "| just | one | row |"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assertGolden(t, renderBlockFixture(t, tt.source))
			},
		)
	}
}

func TestBlockTableInlineContext(t *testing.T) {
	var plain, cells []string
	out := renderBlocks(
		"| h |\n|---|\n| c |\n\npara",
		func(s string, table bool) string {
			if table {
				cells = append(cells, s)
			} else {
				plain = append(plain, s)
			}
			return escapeHTML(s)
		},
	)

	assert.Equal(
		t,
		"<table>\n<thead>\n<tr>\n<th>h</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>c</td>\n</tr>\n</tbody>\n"+
			"</table>\n<p>para</p>\n",
		out,
	)
	assert.Equal(t, []string{"h", "c"}, cells)
	assert.Equal(t, []string{"para"}, plain)
}

func TestBlockTableCellCodeSpan(t *testing.T) {
	assert.Equal(
		t,
		"<table>\n<thead>\n<tr>\n<th>a</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td><code>x | y</code></td>\n"+
			"</tr>\n</tbody>\n</table>\n",
		Markdown("| a |\n|---|\n| `x \\| y` |"),
	)
}
