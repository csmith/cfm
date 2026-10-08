package cfm_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"chameth.com/cfm"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
	"github.com/yuin/goldmark/v2/util"
)

var benchFixtures = []string{
	"small-prose.md",
	"typical-post.md",
	"html-heavy.md",
	"tables-footnotes.md",
	"adversarial-delimiters.md",
	"large-repeated.md",
}

var benchSink string

func BenchmarkMarkdown(b *testing.B) {
	reference := newGoldmark()
	inputs := make(map[string]string, len(benchFixtures))
	var unfair []string
	for _, name := range benchFixtures {
		input := loadFixture(b, name)
		inputs[name] = input
		if diag := workloadDiagnostic(reference, name, input); diag != "" {
			unfair = append(unfair, diag)
		}
	}
	if unfair != nil {
		b.Fatalf("rendered outputs must match before benchmarking:\n\n%s",
			strings.Join(unfair, "\n\n"))
	}
	for _, name := range benchFixtures {
		input := inputs[name]
		b.Run(name, func(b *testing.B) {
			b.Run("cfm", func(b *testing.B) {
				b.SetBytes(int64(len(input)))
				b.ReportAllocs()
				for b.Loop() {
					benchSink = cfm.Markdown(input)
				}
			})
			b.Run("goldmark", func(b *testing.B) {
				b.SetBytes(int64(len(input)))
				b.ReportAllocs()
				var err error
				for b.Loop() {
					var out string
					out, err = goldmarkConvert(reference, input)
					benchSink = out
				}
				if err != nil {
					b.Fatal(err)
				}
			})
		})
	}
}

type goldmarkReference struct {
	parser   parser.Parser
	renderer html.Renderer
}

// commonMarkWithoutCodeBlocks mirrors parser.CommonMark but leaves out the
// indented code block parser, which cfm does not support.
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

func newGoldmark() goldmarkReference {
	return goldmarkReference{
		parser: parser.New(
			parser.WithDefaultParsers(false),
			parser.WithExtensions(
				commonMarkWithoutCodeBlocks{},
				extension.TypographerParser,
				extension.TableParser,
				extension.StrikethroughParser,
				extension.LinkifyParser,
				extension.FootnoteParser,
			),
		),
		renderer: html.New(
			html.WithUnsafe(),
			html.WithExtensions(
				extension.TableHTMLRenderer,
				extension.StrikethroughHTMLRenderer,
				extension.FootnoteHTMLRenderer,
			),
		),
	}
}

func goldmarkConvert(md goldmarkReference, input string) (string, error) {
	var out strings.Builder
	doc := md.parser.ParseStringSource(input)
	if err := md.renderer.RenderStringSource(&out, input, doc); err != nil {
		return "", err
	}
	return out.String(), nil
}

func workloadDiagnostic(md goldmarkReference, name, input string) string {
	want, err := goldmarkConvert(md, input)
	if err != nil {
		return fmt.Sprintf("testdata/bench/%s: goldmark failed: %v", name, err)
	}
	got := cfm.Markdown(input)
	if got == want {
		return ""
	}
	at := 0
	for at < len(got) && at < len(want) && got[at] == want[at] {
		at++
	}
	return fmt.Sprintf(
		"testdata/bench/%s: first output difference at byte %d\n  cfm:      %q\n  goldmark: %q",
		name, at, clip(got, at), clip(want, at))
}

func clip(s string, at int) string {
	lo := max(0, at-100)
	hi := min(len(s), at+120)
	for lo > 0 && s[lo]&0xC0 == 0x80 {
		lo--
	}
	for hi < len(s) && s[hi]&0xC0 == 0x80 {
		hi++
	}
	return s[lo:hi]
}

func loadFixture(b *testing.B, name string) string {
	b.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "bench", name))
	if err != nil {
		b.Fatal(err)
	}
	return string(data)
}
