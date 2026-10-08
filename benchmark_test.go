package cfm_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"chameth.com/cfm"
	"github.com/yuin/goldmark/v2"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
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

type disableCodeBlocks struct{}

func (disableCodeBlocks) SetParserOption(config *parser.Config) {
	config.BlockParsers = config.BlockParsers.Remove(parser.NewCodeBlockParser())
}

func newGoldmark() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			extension.Typographer,
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.Footnote,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			disableCodeBlocks{},
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
}

func goldmarkConvert(md goldmark.Markdown, input string) (string, error) {
	var out strings.Builder
	if err := md.Convert([]byte(input), &out); err != nil {
		return "", err
	}
	return out.String(), nil
}

func workloadDiagnostic(md goldmark.Markdown, name, input string) string {
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
