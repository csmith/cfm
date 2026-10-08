package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	stdhtml "html"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	gmhtml "github.com/yuin/goldmark/v2/renderer/html"
	"github.com/yuin/goldmark/v2/text"
	"github.com/yuin/goldmark/v2/util"
)

const (
	specURL    = "https://spec.commonmark.org/0.31.2/spec.json"
	specSHA256 = "d431b29d97b6f73e69d547109cf5081578fac931e72afe95639ebe766c1b2a20"
	specVer    = "0.31.2"
)

var (
	atxLineRe              = regexp.MustCompile(`^ {0,3}#{1,6}(?:[ \t]|$)`)
	atxClosingRe           = regexp.MustCompile(`^ {0,3}#{1,6}[ \t].*[ \t]#+[ \t]*$`)
	atxEmptyRe             = regexp.MustCompile(`^ {0,3}#{1,6}[ \t]*$`)
	atxClosingOnlyRe       = regexp.MustCompile(`^ {0,3}#{1,6}[ \t]+#+[ \t]*$`)
	thematicBreakRe        = regexp.MustCompile(`^ {0,3}(?:-[ \t]*){3,}$|^ {0,3}(?:\*[ \t]*){3,}$|^ {0,3}(?:_[ \t]*){3,}$`)
	thematicBreakHyphensRe = regexp.MustCompile(`^ {0,3}-{3,}[ \t]*$`)
	refDefRe               = regexp.MustCompile(`^\[[^ \]][^\]]*\]:`)
	listMarkerRe           = regexp.MustCompile(`^[-+*][ \t]+|^[0-9]{1,9}[.)][ \t]+`)
	blankRe                = regexp.MustCompile(`^[ \t]*$`)
)

type specExample struct {
	Markdown  string `json:"markdown"`
	HTML      string `json:"html"`
	Example   int    `json:"example"`
	Section   string `json:"section"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
}

type manifestEntry struct {
	Example int `json:"example"`

	Section  string `json:"section"`
	Included bool   `json:"included"`

	Rationale  string   `json:"rationale"`
	Reasons    []string `json:"excluded_reason,omitempty"`
	Constructs []string `json:"constructs,omitempty"`

	MarkdownFile string `json:"markdown_file,omitempty"`
	HTMLFile     string `json:"html_file,omitempty"`
	Adapted      bool   `json:"adapted_html,omitempty"`
	AdaptedFile  string `json:"adapted_html_file,omitempty"`

	// ExpectedSource names where the test expectation comes from:
	// "spec", "site goldmark", or "hand-authored".
	ExpectedSource string `json:"expected_source,omitempty"`
}

type manifest struct {
	SourceURL    string           `json:"source_url"`
	SourceSHA256 string           `json:"source_sha256"`
	SpecVersion  string           `json:"spec_version"`
	License      string           `json:"license"`
	LicenseURL   string           `json:"license_url"`
	Variant      string           `json:"site_variant"`
	Generator    string           `json:"generator"`
	Total        int              `json:"total_examples"`
	Included     int              `json:"included_examples"`
	Adapted      int              `json:"adapted_examples"`
	Examples     []*manifestEntry `json:"examples"`
}

type flags struct {
	indentedCode   bool
	setextHeading  bool
	emptyATX       bool
	closingHash    bool
	refDefinition  bool
	image          bool
	breakNonHyphen bool
}

func (f flags) reasons() []string {
	var rs []string
	add := func(b bool, s string) {
		if b {
			rs = append(rs, s)
		}
	}
	add(f.indentedCode, "indented code block")
	add(f.setextHeading, "setext heading")
	add(f.emptyATX, "empty ATX heading")
	add(f.closingHash, "ATX closing sequence")
	add(f.refDefinition, "link reference definition")
	add(f.image, "image")
	add(f.breakNonHyphen, "thematic break other than contiguous hyphens")
	return rs
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "generate:", err)
		os.Exit(1)
	}
}

func run() error {
	dir, err := filepath.Abs("..")
	if err != nil {
		return err
	}
	specBytes, err := os.ReadFile(filepath.Join(dir, "spec.json"))
	if err != nil {
		return err
	}
	if err := verifySpec(specBytes); err != nil {
		return err
	}
	var specs []specExample
	if err := json.Unmarshal(specBytes, &specs); err != nil {
		return fmt.Errorf("parse spec.json: %w", err)
	}
	if len(specs) != 652 {
		return fmt.Errorf("spec.json has %d examples, want 652", len(specs))
	}
	m, audit, err := generate(dir, specs)
	if err != nil {
		return err
	}
	printSummary(m, audit)
	return nil
}

func verifySpec(specBytes []byte) error {
	sum := sha256.Sum256(specBytes)
	if got := hex.EncodeToString(sum[:]); got != specSHA256 {
		return fmt.Errorf("spec.json sha256 %s, want %s", got, specSHA256)
	}
	return nil
}

// generate writes examples/NNN.{md,html}, any NNN.adapted.html, and
// manifest.json into dir, and returns the manifest plus audit lines for
// manual review.
func generate(dir string, specs []specExample) (*manifest, []string, error) {
	overrides, err := loadOverrides(dir)
	if err != nil {
		return nil, nil, err
	}

	examplesDir := filepath.Join(dir, "examples")
	if err := os.RemoveAll(examplesDir); err != nil {
		return nil, nil, err
	}
	if err := os.MkdirAll(examplesDir, 0o755); err != nil {
		return nil, nil, err
	}

	site := siteGoldmark()
	m := &manifest{
		SourceURL:    specURL,
		SourceSHA256: specSHA256,
		SpecVersion:  specVer,
		License:      "CC-BY-SA-4.0",
		LicenseURL:   "https://creativecommons.org/licenses/by-sa/4.0/",
		Variant: "goldmark v2.0.2 with typographer, tables, strikethrough, " +
			"linkify, footnotes, auto heading IDs and unsafe HTML; the " +
			"indented code block parser is disabled",
		Generator: "testdata/commonmark/generator (go run .)",
		Examples:  make([]*manifestEntry, 0, len(specs)),
	}

	var audit []string
	for i := range specs {
		entry, line, err := buildEntry(dir, site, overrides, &specs[i])
		if err != nil {
			return nil, nil, err
		}
		if line != "" {
			audit = append(audit, line)
		}
		m.Examples = append(m.Examples, entry)
	}
	if err := rejectUnusedOverrides(overrides, m.Examples); err != nil {
		return nil, nil, err
	}
	if err := checkManifest(dir, m, specs); err != nil {
		return nil, nil, err
	}

	m.Total = len(m.Examples)
	for _, e := range m.Examples {
		if e.Included {
			m.Included++
		}
		if e.Adapted {
			m.Adapted++
		}
	}
	if err := writeManifest(dir, m); err != nil {
		return nil, nil, err
	}
	return m, audit, nil
}

func buildEntry(
	dir string, site markdown, overrides map[int]string, sp *specExample,
) (*manifestEntry, string, error) {
	fl, constructs := classify(sp.Markdown)
	rs := fl.reasons()
	entry := &manifestEntry{
		Example:  sp.Example,
		Section:  sp.Section,
		Included: len(rs) == 0,
	}
	if len(rs) == 0 {
		audit, err := writeIncluded(dir, site, overrides, sp, constructs, entry)
		return entry, audit, err
	}
	entry.Reasons = rs
	entry.Rationale = "Excluded: example involves " + strings.Join(rs, ", ") +
		", which is outside the corpus scope."
	if !strings.Contains(sp.Markdown, "]:") {
		return entry, "", nil
	}
	return entry, fmt.Sprintf(
		"review excluded example %d (%s): markdown contains ]:",
		sp.Example, sp.Section), nil
}

func writeIncluded(
	dir string, site markdown, overrides map[int]string,
	sp *specExample, constructs []string, entry *manifestEntry,
) (string, error) {
	entry.MarkdownFile = fmt.Sprintf("examples/%03d.md", sp.Example)
	entry.HTMLFile = fmt.Sprintf("examples/%03d.html", sp.Example)
	if err := os.WriteFile(filepath.Join(dir, entry.MarkdownFile), []byte(sp.Markdown), 0o644); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(dir, entry.HTMLFile), []byte(sp.HTML), 0o644); err != nil {
		return "", err
	}
	entry.ExpectedSource = "spec"
	entry.Constructs = constructs

	siteHTML, err := render(site, sp.Markdown)
	if err != nil {
		return "", fmt.Errorf("example %d: %w", sp.Example, err)
	}
	override, hasOverride := overrides[sp.Example]
	switch {
	case hasOverride:
		entry.ExpectedSource = "hand-authored"
		if override != sp.HTML {
			entry.Adapted = true
			entry.AdaptedFile = fmt.Sprintf("examples/%03d.adapted.html", sp.Example)
			if err := os.WriteFile(filepath.Join(dir, entry.AdaptedFile), []byte(override), 0o644); err != nil {
				return "", err
			}
			entry.Rationale = "Included; the expected output is hand-authored " +
				"because chameth.com treats indentation explicitly: " +
				"indentation-hidden headings and thematic breaks render as " +
				"such, and indented content is never dropped."
		} else {
			entry.Rationale = "Included; the expected output is hand-authored " +
				"and matches the spec HTML."
		}
	case siteHTML != sp.HTML:
		entry.Adapted = true
		entry.AdaptedFile = fmt.Sprintf("examples/%03d.adapted.html", sp.Example)
		if err := os.WriteFile(filepath.Join(dir, entry.AdaptedFile), []byte(siteHTML), 0o644); err != nil {
			return "", err
		}
		entry.ExpectedSource = "site goldmark"
		entry.Rationale = "Included; the expected output is the chameth.com " +
			"goldmark rendering, which differs from the spec HTML."
	default:
		entry.Rationale = "Included; the chameth.com goldmark rendering " +
			"matches the spec HTML."
	}

	problem := dropCheck(sp.Markdown, siteHTML)
	if problem == "" {
		return "", nil
	}
	if hasOverride {
		return fmt.Sprintf("DATA LOSS example %d (%s): %s (accepted: hand-authored override)",
			sp.Example, sp.Section, problem), nil
	}
	return "", fmt.Errorf("example %d (%s): %s; "+
		"add a hand-authored override at overrides/%03d.html if the loss is intentional",
		sp.Example, sp.Section, problem, sp.Example)
}

func checkManifest(dir string, m *manifest, specs []specExample) error {
	seen := map[int]bool{}
	for _, e := range m.Examples {
		if e.Example < 1 || e.Example > len(specs) || seen[e.Example] {
			return fmt.Errorf("spec example %d accounted more than once", e.Example)
		}
		seen[e.Example] = true
		if !e.Included {
			continue
		}
		for _, f := range []string{e.MarkdownFile, e.HTMLFile} {
			if f == "" {
				return fmt.Errorf("spec example %d: missing file record", e.Example)
			}
			if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
				return fmt.Errorf("spec example %d: %w", e.Example, err)
			}
		}
	}
	for i := range specs {
		if !seen[i+1] {
			return fmt.Errorf("spec example %d missing from manifest", i+1)
		}
	}
	return nil
}

func rejectUnusedOverrides(overrides map[int]string, entries []*manifestEntry) error {
	used := map[int]bool{}
	for _, e := range entries {
		if e.ExpectedSource == "hand-authored" {
			used[e.Example] = true
		}
	}
	var unused []string
	for n := range overrides {
		if !used[n] {
			unused = append(unused, fmt.Sprintf("%03d.html", n))
		}
	}
	if len(unused) == 0 {
		return nil
	}
	sort.Strings(unused)
	return fmt.Errorf("unused overrides: %s", strings.Join(unused, ", "))
}

func writeManifest(dir string, m *manifest) error {
	manifestBytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	manifestBytes = append(manifestBytes, '\n')
	return os.WriteFile(filepath.Join(dir, "manifest.json"), manifestBytes, 0o644)
}

func printSummary(m *manifest, audit []string) {
	fmt.Printf("total %d, included %d, adapted %d, excluded %d\n",
		m.Total, m.Included, m.Adapted, m.Total-m.Included)
	counts := map[string]int{}
	for _, e := range m.Examples {
		for _, r := range e.Reasons {
			counts[r]++
		}
	}
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("  excluded for %-56s %d\n", k+":", counts[k])
	}
	for _, a := range audit {
		fmt.Printf("AUDIT %s\n", a)
	}
}

var overrideNameRe = regexp.MustCompile(`^([0-9]{3})\.html$`)

// loadOverrides reads hand-authored expectations from the overrides
// directory of the corpus dir. File names must be exactly NNN.html, keyed
// by spec example number; a missing directory means there are no overrides.
func loadOverrides(corpusDir string) (map[int]string, error) {
	overridesDir := filepath.Join(corpusDir, "overrides")
	entries, err := os.ReadDir(overridesDir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	overrides := make(map[int]string, len(entries))
	for _, e := range entries {
		match := overrideNameRe.FindStringSubmatch(e.Name())
		if match == nil {
			return nil, fmt.Errorf("bad override file name %q: want NNN.html", e.Name())
		}
		n, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("bad override file name %q: %w", e.Name(), err)
		}
		b, err := os.ReadFile(filepath.Join(overridesDir, e.Name()))
		if err != nil {
			return nil, err
		}
		overrides[n] = string(b)
	}
	return overrides, nil
}

// markdown pairs a goldmark parser with the HTML renderer for its output.
type markdown struct {
	parser   parser.Parser
	renderer gmhtml.Renderer
}

func defaultGoldmark() markdown {
	return markdown{
		parser:   parser.New(),
		renderer: gmhtml.New(),
	}
}

func render(md markdown, src string) (string, error) {
	var buf bytes.Buffer
	doc := md.parser.ParseStringSource(src)
	if err := md.renderer.RenderStringSource(&buf, src, doc); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// siteGoldmark mirrors the chameth.com rendering: default parsers minus the
// indented code block parser, with the site extensions, auto heading IDs
// and unsafe HTML.
func siteGoldmark() markdown {
	return markdown{
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
		renderer: gmhtml.New(
			gmhtml.WithUnsafe(),
			gmhtml.WithExtensions(
				extension.TableHTMLRenderer,
				extension.StrikethroughHTMLRenderer,
				extension.FootnoteHTMLRenderer,
			),
		),
	}
}

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

func classify(src string) (flags, []string) {
	var f flags
	var constructs []string

	srcBytes := []byte(src)
	ctx := parser.NewContext()
	doc := parser.New().Parse(srcBytes, parser.WithContext(ctx))
	// Definitions consumed during parsing are registered in the parse
	// context even where their AST nodes are not retained; count them
	// directly so no definition can escape classification.
	if len(ctx.LinkDefinitions()) > 0 {
		f.refDefinition = true
	}

	starts := lineStarts(srcBytes)
	lineOf := func(off int) int {
		return sort.Search(len(starts), func(i int) bool { return starts[i] > off }) - 1
	}
	covered, hasBreak := walkNodes(doc, srcBytes, starts, lineOf, &f, &constructs)
	scanUncovered(splitLines(srcBytes), covered, hasBreak, &f)

	if constructs == nil {
		constructs = []string{"paragraph"}
	}
	return f, constructs
}

func walkNodes(
	doc ast.Node, srcBytes []byte, starts []int,
	lineOf func(int) int, f *flags, constructs *[]string,
) (map[int]bool, bool) {
	// Lines belong to rendered blocks (paragraphs, code, HTML, consumed
	// link reference definitions); source-level scans only consider
	// uncovered lines, so literal text that merely resembles excluded
	// syntax is not misclassified.
	covered := map[int]bool{}
	hasBreak := false
	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch n.Kind() {
		case ast.KindCodeBlock:
			if n.(*ast.CodeBlock).CodeBlockKind == ast.CodeBlockKindIndented {
				f.indentedCode = true
			} else {
				*constructs = append(*constructs, "fenced code block")
			}
		case ast.KindImage:
			f.image = true
			*constructs = append(*constructs, "image")
		case ast.KindLink:
			*constructs = append(*constructs, "link")
		case ast.KindAutoLink:
			*constructs = append(*constructs, "autolink")
		case ast.KindLinkReferenceDefinition:
			f.refDefinition = true
		case ast.KindHeading:
			if segs := blockSegments(n); len(segs) > 0 {
				line := decontainer(string(lineAt(srcBytes, starts, lineOf(segs[0].Start))))
				classifyHeading(line, f, constructs)
			}
		case ast.KindThematicBreak:
			hasBreak = true
			*constructs = append(*constructs, "thematic break")
		case ast.KindCodeSpan:
			*constructs = append(*constructs, "code span")
		case ast.KindEmphasis, ast.KindStrong:
			*constructs = append(*constructs, "emphasis")
		case ast.KindBlockquote:
			*constructs = append(*constructs, "block quote")
		case ast.KindList:
			*constructs = append(*constructs, "list")
		case ast.KindHTMLBlock:
			*constructs = append(*constructs, "HTML block")
		case ast.KindRawHTML:
			*constructs = append(*constructs, "inline HTML")
		}
		for _, seg := range blockSegments(n) {
			covered[lineOf(seg.Start)] = true
		}
		return ast.WalkContinue, nil
	})
	return covered, hasBreak
}

// blockSegments returns the source lines a block node holds: the raw
// content of code and HTML blocks, or the inline source of other blocks.
// Inline nodes hold none.
func blockSegments(n ast.Node) []text.Segment {
	switch b := n.(type) {
	case *ast.CodeBlock:
		return b.Value.Segments()
	case *ast.HTMLBlock:
		return b.Value.Segments()
	case ast.BlockNode:
		return b.Source()
	}
	return nil
}

func classifyHeading(line string, f *flags, constructs *[]string) {
	if !atxLineRe.MatchString(line) {
		f.setextHeading = true
		return
	}
	if atxClosingRe.MatchString(line) {
		f.closingHash = true
	}
	*constructs = append(*constructs, "ATX heading")
}

func scanUncovered(lines []string, covered map[int]bool, hasBreak bool, f *flags) {
	for i, line := range lines {
		if blankRe.MatchString(line) || covered[i] {
			continue
		}
		trim := strings.TrimLeft(line, " ")
		decon := decontainer(trim)
		switch {
		case atxEmptyRe.MatchString(trim) || atxEmptyRe.MatchString(decon) ||
			atxClosingOnlyRe.MatchString(trim) || atxClosingOnlyRe.MatchString(decon):
			f.emptyATX = true
		case hasBreak && breakLineMatches(trim, decon):
			if !thematicBreakHyphensRe.MatchString(trim) &&
				!thematicBreakHyphensRe.MatchString(decon) {
				f.breakNonHyphen = true
			}
		case refDefRe.MatchString(decon):
			f.refDefinition = true
		}
	}
}

func lineStarts(src []byte) []int {
	starts := []int{0}
	for i, c := range src {
		if c == '\n' {
			starts = append(starts, i+1)
		}
	}
	return starts
}

func lineAt(src []byte, starts []int, ln int) []byte {
	end := len(src)
	if ln+1 < len(starts) {
		end = starts[ln+1]
	}
	return bytes.TrimRight(src[starts[ln]:end], "\r\n")
}

func splitLines(src []byte) []string {
	lines := make([]string, 0, bytes.Count(src, []byte("\n"))+1)
	for s := range strings.SplitSeq(string(src), "\n") {
		lines = append(lines, strings.TrimSuffix(s, "\r"))
	}
	return lines
}

// breakLineMatches reports whether a line renders the AST's thematic
// break. It is only consulted when the AST contains a break, so plain text
// is never misclassified; the container-stripped form handles breaks
// nested in block quotes and list items.
func breakLineMatches(trim, decon string) bool {
	if thematicBreakRe.MatchString(trim) {
		return true
	}
	if strings.HasPrefix(trim, ">") && thematicBreakRe.MatchString(decon) {
		return true
	}
	if listMarkerRe.MatchString(trim) {
		rest := listMarkerRe.ReplaceAllString(trim, "")
		return thematicBreakRe.MatchString(rest)
	}
	return false
}

func decontainer(line string) string {
	s := strings.TrimLeft(line, " ")
	for {
		switch {
		case strings.HasPrefix(s, ">"):
			s = strings.TrimLeft(s[1:], " \t")
		case listMarkerRe.MatchString(s):
			s = listMarkerRe.ReplaceAllString(s, "")
		default:
			return s
		}
	}
}

// dropCheck detects content the site variant loses relative to default
// goldmark (which never drops content), comparing normalised text. A
// non-empty result names the content missing from the site rendering.
func dropCheck(src, siteHTML string) string {
	defHTML, err := render(defaultGoldmark(), src)
	if err != nil {
		return fmt.Sprintf("default render failed: %v", err)
	}
	want := normaliseText(defHTML)
	got := normaliseText(siteHTML)
	if want == got || strings.Contains(got, want) {
		return ""
	}
	var missing strings.Builder
	for w := range strings.FieldsSeq(want) {
		if !strings.Contains(got, w) {
			missing.WriteString(w)
			missing.WriteByte(' ')
		}
	}
	if missing.Len() == 0 {
		return ""
	}
	return fmt.Sprintf("site output is missing content: %q", strings.TrimSpace(missing.String()))
}

// normaliseText canonicalises the visible text of rendered HTML so renders
// differing only in markup or typographic substitutions compare equal.
func normaliseText(rendered string) string {
	s := regexp.MustCompile(`<[^>]*>`).ReplaceAllString(rendered, " ")
	s = stdhtml.UnescapeString(s)
	repl := strings.NewReplacer(
		"‘", "'", "’", "'", "“", "\"", "”", "\"",
		"–", "--", "—", "---", "…", "...",
		"«", "<<", "»", ">>",
	)
	s = repl.Replace(s)
	return strings.Join(strings.Fields(s), " ")
}
