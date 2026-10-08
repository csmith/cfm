package cfm

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func footnoteSupHTML(id string, n int) string {
	num := strconv.Itoa(n)
	return `<sup id="` + id + `"><a href="#fn:` + num +
		`" class="footnote-ref" role="doc-noteref">` + num + `</a></sup>`
}

func footnoteBackrefHTML(id string) string {
	return `&#160;<a href="#` + id + `" class="footnote-backref" role="doc-backlink">&#x21a9;&#xfe0e;</a>`
}

func footnoteItemHTML(n int, inner string) string {
	return "<li id=\"fn:" + strconv.Itoa(n) + "\">\n" + inner + "</li>\n"
}

func footnoteWrapperHTML(items ...string) string {
	var b strings.Builder
	b.WriteString("<div class=\"footnotes\" role=\"doc-endnotes\">\n<hr>\n<ol>\n")
	for _, item := range items {
		b.WriteString(item)
	}
	b.WriteString("</ol>\n</div>\n")
	return b.String()
}

// footnoteFixedHTML is a render callback that ignores the label.
func footnoteFixedHTML(html string) func(string) string {
	return func(string) string { return html }
}

func footnoteLazyAll(string) bool { return true }

func TestFootnoteParse(t *testing.T) {
	tests := []struct {
		name         string
		lines        []string
		lazy         func(string) bool
		wantLabel    string
		wantBody     string
		wantConsumed int
	}{
		{
			name:  "no lines",
			lines: nil,
		},
		{
			name:  "plain text",
			lines: []string{"just text"},
		},
		{
			name:  "missing caret",
			lines: []string{"[1]: x"},
		},
		{
			name:  "missing colon",
			lines: []string{"[^1] x"},
		},
		{
			name:  "space before colon",
			lines: []string{"[^1] :x"},
		},
		{
			name:  "empty label",
			lines: []string{"[^]: x"},
		},
		{
			name:  "blank label",
			lines: []string{"[^ ]: x"},
		},
		{
			name:  "inner bracket rejects closure",
			lines: []string{"[^a[b]: x"},
		},
		{
			name:  "unclosed label",
			lines: []string{"[^1: x"},
		},
		{
			name:  "four space indent is code",
			lines: []string{"    [^1]: x"},
		},
		{
			name:  "tab indent is code",
			lines: []string{"\t[^1]: x"},
		},
		{
			name:         "basic",
			lines:        []string{"[^1]: The clarifying note."},
			wantLabel:    "1",
			wantBody:     "The clarifying note.",
			wantConsumed: 1,
		},
		{
			name:         "up to three leading spaces allowed",
			lines:        []string{"   [^1]: x"},
			wantLabel:    "1",
			wantBody:     "x",
			wantConsumed: 1,
		},
		{
			name:         "no space after colon",
			lines:        []string{"[^1]:x"},
			wantLabel:    "1",
			wantBody:     "x",
			wantConsumed: 1,
		},
		{
			name:         "body indentation after colon preserved",
			lines:        []string{"[^1]:    code"},
			wantLabel:    "1",
			wantBody:     "    code",
			wantConsumed: 1,
		},
		{
			name:         "empty body",
			lines:        []string{"[^1]:"},
			wantLabel:    "1",
			wantBody:     "",
			wantConsumed: 1,
		},
		{
			name:         "blank remainder behaves as empty first line",
			lines:        []string{"[^1]:   ", "    b"},
			wantLabel:    "1",
			wantBody:     "b",
			wantConsumed: 2,
		},
		{
			name:         "continuation lines",
			lines:        []string{"[^1]: one", "    two"},
			wantLabel:    "1",
			wantBody:     "one\ntwo",
			wantConsumed: 2,
		},
		{
			name:         "continuation dedents four columns",
			lines:        []string{"[^1]: a", "     b"},
			wantLabel:    "1",
			wantBody:     "a\n b",
			wantConsumed: 2,
		},
		{
			name:         "significant code indentation kept",
			lines:        []string{"[^1]: a", "", "        code"},
			wantLabel:    "1",
			wantBody:     "a\n\n    code",
			wantConsumed: 3,
		},
		{
			name:         "tab counts as four columns",
			lines:        []string{"[^1]: a", "\tb"},
			wantLabel:    "1",
			wantBody:     "a\nb",
			wantConsumed: 2,
		},
		{
			name:         "short indent stops continuation",
			lines:        []string{"[^1]: a", "   b"},
			wantLabel:    "1",
			wantBody:     "a",
			wantConsumed: 1,
		},
		{
			name:         "lazy continuation joins the paragraph",
			lines:        []string{"[^1]: a", "b", "c"},
			lazy:         footnoteLazyAll,
			wantLabel:    "1",
			wantBody:     "a\nb\nc",
			wantConsumed: 3,
		},
		{
			name:         "lazy callback can reject a line",
			lines:        []string{"[^1]: a", "# b"},
			lazy:         func(string) bool { return false },
			wantLabel:    "1",
			wantBody:     "a",
			wantConsumed: 1,
		},
		{
			name:         "lazy continuation skips leading indent",
			lines:        []string{"[^1]: a", "  b"},
			lazy:         footnoteLazyAll,
			wantLabel:    "1",
			wantBody:     "a\nb",
			wantConsumed: 2,
		},
		{
			name:         "lazy continuation keeps trailing spaces",
			lines:        []string{"[^1]: a  ", "b  "},
			lazy:         footnoteLazyAll,
			wantLabel:    "1",
			wantBody:     "a  \nb  ",
			wantConsumed: 2,
		},
		{
			name:         "lazy does not cross blank lines",
			lines:        []string{"[^1]: a", "", "b"},
			lazy:         footnoteLazyAll,
			wantLabel:    "1",
			wantBody:     "a",
			wantConsumed: 2,
		},
		{
			name:         "lazy needs an open paragraph",
			lines:        []string{"[^1]:", "b"},
			lazy:         footnoteLazyAll,
			wantLabel:    "1",
			wantBody:     "",
			wantConsumed: 1,
		},
		{
			name: "blank separated paragraphs",
			lines: []string{
				"[^2]: Opening paragraph.", "",
				"    Second paragraph.", "",
				"    Third and final paragraph.",
			},
			wantLabel:    "2",
			wantBody:     "Opening paragraph.\n\nSecond paragraph.\n\nThird and final paragraph.",
			wantConsumed: 5,
		},
		{
			name: "stops before unindented next content",
			lines: []string{
				"[^2]: Opening paragraph.",
				"",
				"    Second paragraph.",
				"",
				"    Third and final paragraph.",
				"Next",
			},
			wantLabel:    "2",
			wantBody:     "Opening paragraph.\n\nSecond paragraph.\n\nThird and final paragraph.",
			wantConsumed: 5,
		},
		{
			name:         "blank before unindented content is consumed",
			lines:        []string{"[^1]: a", "", "next"},
			wantLabel:    "1",
			wantBody:     "a",
			wantConsumed: 2,
		},
		{
			name:         "trailing blanks at end of input not consumed",
			lines:        []string{"[^1]: a", "", ""},
			wantLabel:    "1",
			wantBody:     "a",
			wantConsumed: 1,
		},
		{
			name:         "label with spaces",
			lines:        []string{"[^my note]: x"},
			wantLabel:    "my note",
			wantBody:     "x",
			wantConsumed: 1,
		},
		{
			name:         "escaped bracket stays in label",
			lines:        []string{"[^a\\]b]: x"},
			wantLabel:    "a\\]b",
			wantBody:     "x",
			wantConsumed: 1,
		},
		{
			name:         "label characters are raw",
			lines:        []string{"[^A-b_2]: x"},
			wantLabel:    "A-b_2",
			wantBody:     "x",
			wantConsumed: 1,
		},
		{
			name:         "trailing body whitespace kept for hard breaks",
			lines:        []string{"[^1]: a  ", "    b"},
			wantLabel:    "1",
			wantBody:     "a  \nb",
			wantConsumed: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				label, body, consumed := parseFootnoteDefinition(tt.lines, tt.lazy)
				assert.Equal(t, tt.wantLabel, label)
				assert.Equal(t, tt.wantBody, body)
				assert.Equal(t, tt.wantConsumed, consumed)
			},
		)
	}
}

func TestFootnoteReference(t *testing.T) {
	t.Run(
		"sup markup",
		func(t *testing.T) {
			var f footnotes
			f.define("1", "note")
			got, ok := f.reference("1")
			require.True(t, ok)
			assert.Equal(t, footnoteSupHTML("fnref:1", 1), got)
		},
	)

	t.Run(
		"undefined is literal",
		func(t *testing.T) {
			var f footnotes
			got, ok := f.reference("99")
			assert.False(t, ok)
			assert.Empty(t, got)
		},
	)

	t.Run(
		"numbering follows encounter order",
		func(t *testing.T) {
			var f footnotes
			f.define("3", "Third.")
			f.define("1", "First.")
			f.define("2", "Second.")
			got3, ok := f.reference("3")
			require.True(t, ok)
			got1, ok := f.reference("1")
			require.True(t, ok)
			got2, ok := f.reference("2")
			require.True(t, ok)
			assert.Equal(t, footnoteSupHTML("fnref:1", 1), got3)
			assert.Equal(t, footnoteSupHTML("fnref:2", 2), got1)
			assert.Equal(t, footnoteSupHTML("fnref:3", 3), got2)
		},
	)

	t.Run(
		"repeated references get unique ref ids",
		func(t *testing.T) {
			var f footnotes
			f.define("1", "Referenced twice.")
			got1, ok := f.reference("1")
			require.True(t, ok)
			got2, ok := f.reference("1")
			require.True(t, ok)
			got3, ok := f.reference("1")
			require.True(t, ok)
			assert.Equal(t, footnoteSupHTML("fnref:1", 1), got1)
			assert.Equal(t, footnoteSupHTML("fnref1:1", 1), got2)
			assert.Equal(t, footnoteSupHTML("fnref2:1", 1), got3)
		},
	)

	t.Run(
		"undefined labels do not consume numbers",
		func(t *testing.T) {
			var f footnotes
			f.define("1", "note")
			_, ok := f.reference("99")
			assert.False(t, ok)
			got, ok := f.reference("1")
			require.True(t, ok)
			assert.Equal(t, footnoteSupHTML("fnref:1", 1), got)
		},
	)

	t.Run(
		"labels are case sensitive",
		func(t *testing.T) {
			var f footnotes
			f.define("a", "lower")
			f.define("A", "upper")
			gotA, ok := f.reference("A")
			require.True(t, ok)
			gota, ok := f.reference("a")
			require.True(t, ok)
			assert.Equal(t, footnoteSupHTML("fnref:1", 1), gotA)
			assert.Equal(t, footnoteSupHTML("fnref:2", 2), gota)
		},
	)

	t.Run(
		"first definition wins",
		func(t *testing.T) {
			var f footnotes
			assert.True(t, f.define("x", "first"))
			assert.False(t, f.define("x", "second"))
			_, ok := f.reference("x")
			require.True(t, ok)
			var got string
			f.render(
				func(label string) string {
					got = f.defs[label]
					return ""
				},
			)
			assert.Equal(t, "first", got)
		},
	)
}

func TestFootnoteRender(t *testing.T) {
	// The render callback receives the note label and returns the complete
	// list item content, backlinks included: production resolves bodies in a
	// source-order pass before rendering.
	tests := []struct {
		name   string
		define [][2]string
		refs   []string
		inner  func(label string) string
		want   string
	}{
		{
			name:   "basic",
			define: [][2]string{{"1", "The clarifying note."}},
			refs:   []string{"1"},
			inner: footnoteFixedHTML(
				"<p>The clarifying note." +
					footnoteBackrefHTML("fnref:1") + "</p>\n",
			),
			want: footnoteWrapperHTML(
				footnoteItemHTML(
					1,
					"<p>The clarifying note."+footnoteBackrefHTML("fnref:1")+"</p>\n",
				),
			),
		},
		{
			name:   "repeated reference gets two backlinks",
			define: [][2]string{{"1", "Referenced twice."}},
			refs:   []string{"1", "1"},
			inner: footnoteFixedHTML(
				"<p>Referenced twice." +
					footnoteBackrefHTML("fnref:1") + footnoteBackrefHTML("fnref1:1") + "</p>\n",
			),
			want: footnoteWrapperHTML(
				footnoteItemHTML(
					1,
					"<p>Referenced twice."+footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n",
				),
			),
		},
		{
			name:   "definition order follows reference order",
			define: [][2]string{{"1", "First."}, {"2", "Second."}, {"3", "Third."}},
			refs:   []string{"3", "1", "2"},
			inner: func(label string) string {
				switch label {
				case "3":
					return "<p>Third." + footnoteBackrefHTML("fnref:1") + "</p>\n"
				case "1":
					return "<p>First." + footnoteBackrefHTML("fnref:2") + "</p>\n"
				default:
					return "<p>Second." + footnoteBackrefHTML("fnref:3") + "</p>\n"
				}
			},
			want: footnoteWrapperHTML(
				footnoteItemHTML(1, "<p>Third."+footnoteBackrefHTML("fnref:1")+"</p>\n"),
				footnoteItemHTML(2, "<p>First."+footnoteBackrefHTML("fnref:2")+"</p>\n"),
				footnoteItemHTML(3, "<p>Second."+footnoteBackrefHTML("fnref:3")+"</p>\n"),
			),
		},
		{
			name:   "unreferenced definitions omitted",
			define: [][2]string{{"1", "Used."}, {"2", "Unused."}},
			refs:   []string{"1"},
			inner:  footnoteFixedHTML("<p>Used." + footnoteBackrefHTML("fnref:1") + "</p>\n"),
			want: footnoteWrapperHTML(
				footnoteItemHTML(
					1,
					"<p>Used."+footnoteBackrefHTML("fnref:1")+"</p>\n",
				),
			),
		},
		{
			name:   "no references renders nothing",
			define: [][2]string{{"1", "Unused."}},
			refs:   nil,
			inner:  footnoteFixedHTML("<p>Unused.</p>\n"),
			want:   "",
		},
		{
			name:   "code-only note keeps backref outside the code block",
			define: [][2]string{{"1", "text"}},
			refs:   []string{"1"},
			inner: footnoteFixedHTML(
				"<pre><code>code\n</code></pre>\n" +
					footnoteBackrefHTML("fnref:1"),
			),
			want: footnoteWrapperHTML(
				footnoteItemHTML(
					1,
					"<pre><code>code\n</code></pre>\n"+footnoteBackrefHTML("fnref:1"),
				),
			),
		},
		{
			name:   "empty note body gets bare backref",
			define: [][2]string{{"1", ""}},
			refs:   []string{"1"},
			inner:  footnoteFixedHTML(footnoteBackrefHTML("fnref:1")),
			want:   footnoteWrapperHTML(footnoteItemHTML(1, footnoteBackrefHTML("fnref:1"))),
		},
		{
			name:   "note ending in a list keeps backref outside it",
			define: [][2]string{{"1", "text"}},
			refs:   []string{"1"},
			inner: footnoteFixedHTML(
				"<p>text</p>\n<ul>\n<li>item</li>\n</ul>\n" +
					footnoteBackrefHTML("fnref:1"),
			),
			want: footnoteWrapperHTML(
				footnoteItemHTML(
					1,
					"<p>text</p>\n<ul>\n<li>item</li>\n</ul>\n"+footnoteBackrefHTML("fnref:1"),
				),
			),
		},
		{
			name:   "note ending in a blockquote keeps backref outside it",
			define: [][2]string{{"1", "text"}},
			refs:   []string{"1"},
			inner: footnoteFixedHTML(
				"<p>text</p>\n<blockquote>\n<p>quoted</p>\n</blockquote>\n" +
					footnoteBackrefHTML("fnref:1"),
			),
			want: footnoteWrapperHTML(
				footnoteItemHTML(
					1,
					"<p>text</p>\n<blockquote>\n<p>quoted</p>\n</blockquote>\n"+footnoteBackrefHTML("fnref:1"),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				var f footnotes
				for _, d := range tt.define {
					f.define(d[0], d[1])
				}
				for _, label := range tt.refs {
					_, ok := f.reference(label)
					require.True(t, ok)
				}
				assert.Equal(t, tt.want, f.render(tt.inner))
			},
		)
	}
}

// footnoteResolver stands in for the document inline pass: it renders one
// definition body, registering the footnote references it contains.
type footnoteResolver struct {
	notes  *footnotes
	counts map[string]int
}

func (r *footnoteResolver) resolve(body string) string {
	r.counts[body]++
	var b strings.Builder
	b.WriteString("<p>")
	for rest := body; ; {
		i := strings.Index(rest, "[^")
		j := -1
		if i >= 0 {
			if k := strings.IndexByte(rest[i:], ']'); k >= 0 {
				j = i + k
			} else {
				i = -1
			}
		}
		if i < 0 {
			b.WriteString(rest)
			break
		}
		if sup, ok := r.notes.reference(rest[i+2 : j]); ok {
			b.WriteString(rest[:i])
			b.WriteString(sup)
		} else {
			b.WriteString(rest[:j+1])
		}
		rest = rest[j+1:]
	}
	b.WriteString("</p>\n")
	return b.String()
}

func TestFootnoteRenderDiscoversReferences(t *testing.T) {
	var f footnotes
	f.define("1", "see [^2]")
	f.define("2", "two [^1]")
	_, ok := f.reference("1")
	require.True(t, ok)
	r := &footnoteResolver{notes: &f, counts: map[string]int{}}
	// The document pass resolves every body exactly once, in order, before
	// rendering attaches backlinks.
	in := map[string]string{}
	for _, label := range []string{"1", "2"} {
		in[label] = r.resolve(f.defs[label])
	}
	got := f.render(
		func(label string) string {
			return attachFootnoteBacklinks(in[label], f.backlinks(label, f.numbers[label]))
		},
	)
	want := footnoteWrapperHTML(
		footnoteItemHTML(
			1,
			"<p>see "+footnoteSupHTML("fnref:2", 2)+
				footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n",
		),
		footnoteItemHTML(
			2,
			"<p>two "+footnoteSupHTML("fnref1:1", 1)+
				footnoteBackrefHTML("fnref:2")+"</p>\n",
		),
	)
	assert.Equal(t, want, got)
	assert.Equal(t, map[string]int{"see [^2]": 1, "two [^1]": 1}, r.counts)
}

func TestFootnoteRenderFinalizesBackrefs(t *testing.T) {
	var f footnotes
	f.define("a", "note a")
	f.define("b", "b refs [^a]")
	_, ok := f.reference("a")
	require.True(t, ok)
	_, ok = f.reference("b")
	require.True(t, ok)
	r := &footnoteResolver{notes: &f, counts: map[string]int{}}
	in := map[string]string{}
	for _, label := range []string{"a", "b"} {
		in[label] = r.resolve(f.defs[label])
	}
	got := f.render(
		func(label string) string {
			return attachFootnoteBacklinks(in[label], f.backlinks(label, f.numbers[label]))
		},
	)
	want := footnoteWrapperHTML(
		footnoteItemHTML(
			1,
			"<p>note a"+
				footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n",
		),
		footnoteItemHTML(
			2,
			"<p>b refs "+footnoteSupHTML("fnref1:1", 1)+
				footnoteBackrefHTML("fnref:2")+"</p>\n",
		),
	)
	assert.Equal(t, want, got)
	assert.Equal(t, map[string]int{"note a": 1, "b refs [^a]": 1}, r.counts)
}

func TestFootnoteRenderSelfCycle(t *testing.T) {
	var f footnotes
	f.define("a", "self [^a]")
	_, ok := f.reference("a")
	require.True(t, ok)
	r := &footnoteResolver{notes: &f, counts: map[string]int{}}
	in := map[string]string{}
	for _, label := range []string{"a"} {
		in[label] = r.resolve(f.defs[label])
	}
	got := f.render(
		func(label string) string {
			return attachFootnoteBacklinks(in[label], f.backlinks(label, f.numbers[label]))
		},
	)
	want := footnoteWrapperHTML(
		footnoteItemHTML(
			1,
			"<p>self "+footnoteSupHTML("fnref1:1", 1)+
				footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n",
		),
	)
	assert.Equal(t, want, got)
	assert.Equal(t, map[string]int{"self [^a]": 1}, r.counts)
}

func TestFootnoteRenderNone(t *testing.T) {
	var f footnotes
	assert.Empty(t, f.render(func(string) string { return "" }))
}

func TestAttachFootnoteBacklinks(t *testing.T) {
	back := footnoteBackrefHTML("fnref:1")
	assert.Equal(
		t,
		"<p>note"+back+"</p>\n",
		attachFootnoteBacklinks("<p>note</p>\n", back),
	)
	assert.Equal(
		t,
		"<p>one</p>\n<p>two"+back+"</p>\n",
		attachFootnoteBacklinks("<p>one</p>\n<p>two</p>\n", back),
	)
	assert.Equal(
		t,
		"<pre><code>x\n</code></pre>\n"+back,
		attachFootnoteBacklinks("<pre><code>x\n</code></pre>\n", back),
	)
	assert.Equal(
		t,
		"<ul>\n<li>item</li>\n</ul>\n"+back,
		attachFootnoteBacklinks("<ul>\n<li>item</li>\n</ul>\n", back),
	)
	assert.Equal(
		t,
		""+back,
		attachFootnoteBacklinks("", back),
	)
	assert.Equal(
		t,
		"<!-- c --><p>note"+back+"</p>\n",
		attachFootnoteBacklinks("<!-- c --><p>note</p>\n", back),
	)
	assert.Equal(
		t,
		"<p>a <em>x</em>"+back+"</p>\n",
		attachFootnoteBacklinks("<p>a <em>x</em></p>\n", back),
	)
}
