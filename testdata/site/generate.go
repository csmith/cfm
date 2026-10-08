//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
	"github.com/yuin/goldmark/v2/util"
)

// commonMarkWithoutCodeBlocks mirrors parser.CommonMark but leaves out the
// indented code block parser.
type commonMarkWithoutCodeBlocks struct{}

func (commonMarkWithoutCodeBlocks) ParserOptions(*parser.Config) []parser.Option {
	heading := parser.WithAutoHeadingID()
	return []parser.Option{
		parser.WithBlockParsers(
			util.Prioritized(parser.NewSetextHeadingParser(heading), 100),
			util.Prioritized(parser.NewThematicBreakParser(), 200),
			util.Prioritized(parser.NewListParser(), 300),
			util.Prioritized(parser.NewListItemParser(), 400),
			util.Prioritized(parser.NewATXHeadingParser(heading), 600),
			util.Prioritized(parser.NewFencedCodeBlockParser(), 700),
			util.Prioritized(parser.NewBlockquoteParser(), 800),
			util.Prioritized(parser.NewHTMLBlockParser(), 900),
			util.Prioritized(parser.NewParagraphParser(), 1000),
		),
		parser.WithInlineParsers(
			util.Prioritized(parser.NewCodeSpanParser(), 100),
			util.Prioritized(parser.NewLinkParser(), 200),
			util.Prioritized(parser.NewAutoLinkParser(), 300),
			util.Prioritized(parser.NewRawHTMLParser(), 400),
			util.Prioritized(parser.NewEmphasisParser(), 500),
		),
		parser.WithParagraphTransformers(
			util.Prioritized(parser.LinkReferenceParagraphTransformer, 100),
		),
	}
}

var mdParser = parser.New(
	parser.WithDefaultParsers(false),
	parser.WithExtensions(
		commonMarkWithoutCodeBlocks{},
		extension.TypographerParser,
		extension.TableParser,
		extension.StrikethroughParser,
		extension.LinkifyParser,
		extension.FootnoteParser,
	),
)

var mdRenderer = html.New(
	html.WithUnsafe(),
	html.WithExtensions(
		extension.TableHTMLRenderer,
		extension.StrikethroughHTMLRenderer,
		extension.FootnoteHTMLRenderer,
	),
)

type fixture struct {
	name string
	md   string
}

var fixtures = []fixture{
	// Headings and automatic IDs.
	{"001-heading-levels", "# Level one\n## Level two\n### Level three\n#### Level four\n##### Level five\n###### Level six\n"},
	{"002-heading-id-collisions", "# Duplicate\n\n## Duplicate\n\n### duplicate\n\n# Hello World\n\n# hello world\n"},
	{"003-heading-nonascii", "# Café Üben Tokyo\n\n# 東京の見出し\n\n# naïve Übung\n"},
	{"004-heading-inline-markup", "# *Emphasised* heading\n\n## `Code` and [link](/where)\n"},
	{"005-heading-no-space", "#NoSpaceHere\n\n###Also bad\n"},
	{"006-heading-trailing-spaces", "# Heading with trailing spaces   \n\n##  Two spaces after marker\n"},

	// Typographer, and its exceptions inside code.
	{"010-typographer-quotes", "\"Double quoted\" and 'single quoted' and \"nested 'inner' outer\".\n"},
	{"011-typographer-dashes-ellipsis", "An en -- dash, an em --- dash and an elipsis... done.\n"},
	{"012-typographer-code-span", "Code keeps raw text: `\"quoted\" -- 'single' ...`\n"},
	{"013-typographer-code-block", "```\n\"quoted\" and 'single' -- --- ...\n```\n"},
	{"014-typographer-apostrophes", "Don't can't won't it's a party.\n"},
	{"015-typographer-around-link", "\"[linked](/page)\" and 'plain'.\n"},

	// Tables.
	{"020-table-basic-alignment", "| Left | Center | Right |\n|:-----|:------:|------:|\n| a | b | c |\n| wider cell | m | z |\n"},
	{"021-table-no-alignment", "| One | Two |\n|-----|-----|\n| 1 | 2 |\n"},
	{"022-table-inline-content", "| Format | Example |\n|--------|---------|\n| bold | **bold** |\n| italic | *italic* |\n| code | `code()` |\n| struck | ~~gone~~ |\n| link | [site](/home) |\n"},
	{"023-table-escaped-pipe", "| Expression | Meaning |\n|------------|---------|\n| a \\| b | literal pipe |\n| `x \\| y` | pipe in code span |\n"},
	{"024-table-irregular-rows", "| A | B | C |\n|---|---|---|\n| short row |\n| x | y | z | extra cell |\n"},
	{"025-table-empty-cells", "| | | |\n|---|---|---|\n| | filled | |\n"},
	{"026-table-not-table", "just | some | text\n\nno | table | here\n:--- | :---\n"},

	// Strikethrough.
	{"030-strikethrough-single-and-double", "~single~ and ~~double~~ both strike.\n"},
	{"031-strikethrough-triple-literal", "Text with ~~~triple tildes~~~ inside.\n"},
	{"032-strikethrough-rich", "~~struck **bold** and *emph*~~\n"},
	{"033-strikethrough-unclosed", "~~never closed here\n"},

	// Linkify.
	{"040-linkify-url", "See https://example.com/page for details.\n"},
	{"041-linkify-punctuation", "Ends with dot https://example.com. Wrapped (https://example.com/paren) too.\n"},
	{"042-linkify-query-escape", "Query https://example.com/search?q=x&y=1 works.\n"},
	{"043-linkify-email", "Mail fred.bloggs+tag@example.co.uk now.\n"},
	{"044-linkify-no-scheme", "Bare example.com is not linkified.\n"},

	// Footnotes.
	{"050-footnote-basic", "A statement[^1].\n\n[^1]: The clarifying note.\n"},
	{"051-footnote-repeated-reference", "First use[^1], second use[^1].\n\n[^1]: Referenced twice.\n"},
	{"052-footnote-multiblock", "Complicated[^2].\n\n[^2]: Opening paragraph.\n\n    Second paragraph.\n\n    Third and final paragraph.\n"},
	{"053-footnote-many", "Order three[^3], one[^1], two[^2].\n\n[^1]: First.\n\n[^2]: Second.\n\n[^3]: Third.\n"},
	{"054-footnote-inline-literal", "An inline note^[with *emphasis* inside] here.\n"},
	{"055-footnote-undefined", "Unknown note[^99] stays literal.\n"},

	// Raw HTML (renderer runs with unsafe enabled).
	{"060-html-block", "<div class=\"note\">\n  <p>Raw HTML is passed through.</p>\n</div>\n"},
	{"061-html-indented", "<div class=\"tools\">\n  <ul>\n    <li>one</li>\n  </ul>\n</div>\n"},
	{"062-html-inline", "Inline <b>bold</b> and a <br>break.\n"},
	{"063-html-script", "<script>\nif (a < b && c > d) { alert(\"raw\"); }\n</script>\n"},
	{"064-html-comment", "<!-- a comment ignores **markup** -->\n"},
	{"065-html-escaping-around", "Before <span title=\"a&amp;b\">tag</span> after & marker.\n"},

	// Lists.
	{"070-list-tight", "- one\n- two\n- three\n"},
	{"071-list-loose", "- one\n\n- two\n\n- three\n"},
	{"072-list-ordered-start", "3. three\n4. four\n5. five\n"},
	{"073-list-nested-mixed", "- alpha\n  - alpha one\n  - alpha two\n- beta\n  1. beta one\n  2. beta two\n"},
	{"074-list-checkbox-plain", "- [ ] unchecked task\n- [x] checked task\n"},
	{"075-list-fenced-code", "- item with code:\n\n  ```go\n  fmt.Println(\"hi\")\n  ```\n\n- next item\n"},
	{"076-list-blockquote", "- item\n\n  > quoted inside item\n\n- next\n"},
	{"077-list-continuation", "- item with\n  continuation line\n- next\n"},

	// Block quotes.
	{"080-blockquote-lazy", "> First quoted line\nlazy continuation of the same paragraph.\n"},
	{"081-blockquote-nested", "> outer\n>\n> > inner\n>\n> back to outer\n"},
	{"082-blockquote-rich", "> # Heading in quote\n>\n> - a list item\n>\n> ```\n> fenced in quote\n> ```\n"},

	// Fenced code blocks (no Chroma: plain pre/code output).
	{"090-fence-language", "```go\nif a < b && c != d {\n\treturn \"ok\"\n}\n```\n"},
	{"091-fence-plain", "```\nplain fenced text <tag> & entity\n```\n"},
	{"092-fence-tilde", "~~~html\n<p>tilde fenced</p>\n~~~\n"},
	{"093-fence-unclosed", "```go\nnever closed before EOF\n"},
	{"094-fence-nested-backticks", "````\ninner ``` fence stays content\n````\n"},
	{"095-fence-info-extra", "```go title=main.go\npackage main\n```\n"},

	// Inline markup.
	{"100-emphasis-intraword", "in*tra*word works, in__ter__ underscore does not.\n"},
	{"101-emphasis-nested", "***all three*** and **bold *nested***.\n"},
	{"102-backslash-escapes", "\\*not emphasised\\* and \\_under\\_ and \\[bracket\\]\n"},
	{"103-hard-breaks", "trailing two spaces  \nnext line.\n\nbackslash break\\\nsecond line.\n"},
	{"104-entities", "Known &copy; &amp; &lt;tag&gt; &#8212; unknown &fakeentity; raw & alone.\n"},
	{"105-autolink-angle", "<https://example.com/auto> and <fred@example.com>\n"},
	{"106-link-inline", "[inline link](/dest \"the title\") and [**styled** text](/other).\n"},

	// Document-level behaviour.
	{"120-thematic-break-hyphens", "before\n\n---\n\nafter\n\n--------------\n\ntail\n"},
	{"121-crlf-line-endings", "first paragraph\r\n\r\nsecond paragraph\r\nsecond line\r\n"},
	{"122-unicode-paragraph", "Grüße aus München, 東京 and emoji 🎉 — ça va très bien.\n"},
	{"123-empty-input", ""},
	{"124-whitespace-only", "\n\n   \n\t\n"},
	{"125-paragraph-soft-break", "line one\nline two\n"},
	{"126-trailing-space-paragraph", "paragraph ends with trailing spaces   \n"},

	// Malformed input.
	{"130-malformed-unclosed-emphasis", "**bold never closed and *neither is this\n"},
	{"131-malformed-link-no-paren", "[text](missing-paren\n"},
	{"132-malformed-stray-brackets", "]stray[ and a[rbitrary] text ]here[\n"},
	{"133-malformed-table-solo", "| just | one | row |\n"},
	{"134-malformed-html-tag", "a < not a tag > b and </ dangling.\n"},
}

func main() {
	for _, f := range fixtures {
		var buf bytes.Buffer
		doc := mdParser.ParseStringSource(f.md)
		if err := mdRenderer.RenderStringSource(&buf, f.md, doc); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", f.name, err)
			os.Exit(1)
		}
		if err := os.WriteFile(filepath.Join("testdata", "site", f.name+".md"), []byte(f.md), 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := os.WriteFile(filepath.Join("testdata", "site", f.name+".html"), buf.Bytes(), 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Printf("wrote %d generated fixtures to testdata/site\n", len(fixtures))
}
