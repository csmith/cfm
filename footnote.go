package cfm

import (
	"strconv"
	"strings"
)

// footnoteDefState assembles the body of one definition line by line: blank
// lines queue until further content commits them, indented lines dedent four
// columns, and unindented lines may lazily continue the open paragraph.
type footnoteDefState struct {
	label   string
	body    []string
	pending int
}

// newFootnoteDef parses the definition's first line. A false result means the
// line does not start a definition.
func newFootnoteDef(line string) (*footnoteDefState, bool) {
	off := 0
	for off < len(line) && line[off] == ' ' {
		off++
	}
	if off > 3 || off+1 >= len(line) || line[off] != '[' || line[off+1] != '^' {
		return nil, false
	}
	closes := footnoteLabelClosure(line, off+2)
	if closes < 0 || closes+1 >= len(line) || line[closes+1] != ':' {
		return nil, false
	}
	label := line[off+2 : closes]
	if isFootnoteBlank(label) {
		return nil, false
	}
	d := &footnoteDefState{label: label}
	if rest := footnoteTrimLead(line[closes+2:]); rest != "" {
		d.body = append(d.body, rest)
	}
	return d, true
}

// commit flushes queued blank lines into the body when content follows them,
// reporting how many were absorbed.
func (d *footnoteDefState) commit() int {
	n := d.pending
	if len(d.body) > 0 {
		for range d.pending {
			d.body = append(d.body, "")
		}
	}
	d.pending = 0
	return n
}

func (d *footnoteDefState) text() string {
	return strings.Join(d.body, "\n")
}

// lazyFootnoteBody continues the definition's open paragraph with an
// unindented line when the callback accepts it. Blank lines, a paragraph
// that has not started, and rejected lines end the definition.
func lazyFootnoteBody(d *footnoteDefState, line string, lazy func(string) bool) bool {
	if d.pending > 0 || len(d.body) == 0 || lazy == nil || !lazy(line) {
		return false
	}
	d.body = append(d.body, stripIndent(line))
	return true
}

// Call only at a block boundary, after stripping container indentation.
// A zero consumed count means no definition; body continuations are dedented.
// lazy reports whether an unindented line keeps the body paragraph open.
func parseFootnoteDefinition(lines []string, lazy func(string) bool) (label, body string, consumed int) {
	if len(lines) == 0 {
		return "", "", 0
	}
	d, ok := newFootnoteDef(lines[0])
	if !ok {
		return "", "", 0
	}
	taken := 0
	stopped := false
	for _, line := range lines[1:] {
		if isFootnoteBlank(line) {
			d.pending++
			continue
		}
		if s, ok := footnoteDedent(line); ok {
			taken += d.commit()
			taken++
			d.body = append(d.body, s)
			continue
		}
		if lazyFootnoteBody(d, line, lazy) {
			taken++
			continue
		}
		stopped = true
		break
	}
	consumed = 1 + taken
	if stopped {
		consumed += d.pending
	}
	return d.label, d.text(), consumed
}

type footnotes struct {
	defs    map[string]string
	numbers map[string]int
	order   []string
	refs    map[string]int
}

// define records the body of a label's first definition and reports whether
// the label was new.
func (f *footnotes) define(label, body string) bool {
	if f.defs == nil {
		f.defs = make(map[string]string)
	}
	if _, ok := f.defs[label]; ok {
		return false
	}
	f.defs[label] = body
	return true
}

func (f *footnotes) reference(label string) (string, bool) {
	if _, ok := f.defs[label]; !ok {
		return "", false
	}
	if f.numbers == nil {
		f.numbers = make(map[string]int)
		f.refs = make(map[string]int)
	}
	n, ok := f.numbers[label]
	if !ok {
		n = len(f.numbers) + 1
		f.numbers[label] = n
		f.order = append(f.order, label)
	}
	occ := f.refs[label]
	f.refs[label] = occ + 1
	var b strings.Builder
	b.WriteString(`<sup id="`)
	b.WriteString(footnoteRefID(n, occ))
	b.WriteString(`"><a href="#fn:`)
	b.WriteString(strconv.Itoa(n))
	b.WriteString(`" class="footnote-ref" role="doc-noteref">`)
	b.WriteString(strconv.Itoa(n))
	b.WriteString(`</a></sup>`)
	return b.String(), true
}

func (f *footnotes) backlinks(label string, n int) string {
	var b strings.Builder
	for occ := range f.refs[label] {
		b.WriteString(`&#160;<a href="#`)
		b.WriteString(footnoteRefID(n, occ))
		b.WriteString(`" class="footnote-backref" role="doc-backlink">&#x21a9;&#xfe0e;</a>`)
	}
	return b.String()
}

// render assembles the notes section. Every reference is registered by the
// time it runs, so renderBody returns each note's complete list item content,
// backlinks included, for its label.
func (f *footnotes) render(renderBody func(label string) string) string {
	if len(f.order) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("<div class=\"footnotes\" role=\"doc-endnotes\">\n<hr>\n<ol>\n")
	for _, label := range f.order {
		b.WriteString(`<li id="fn:`)
		b.WriteString(strconv.Itoa(f.numbers[label]))
		b.WriteString("\">\n")
		b.WriteString(renderBody(label))
		b.WriteString("</li>\n")
	}
	b.WriteString("</ol>\n</div>\n")
	return b.String()
}

func footnoteRefID(n, occ int) string {
	if occ == 0 {
		return "fnref:" + strconv.Itoa(n)
	}
	return "fnref" + strconv.Itoa(occ) + ":" + strconv.Itoa(n)
}

// Backlinks join the last top-level paragraph when it closes the body;
// otherwise they sit bare after the final block, matching Goldmark.
func attachFootnoteBacklinks(body, backlinks string) string {
	if p := lastTopFootnoteParagraph(body); p >= 0 && strings.TrimSpace(body[p+len("</p>"):]) == "" {
		return body[:p] + backlinks + body[p:]
	}
	return body + backlinks
}

func lastTopFootnoteParagraph(s string) int {
	depth, last := 0, -1
	for i := 0; i < len(s); i++ {
		if s[i] != '<' {
			continue
		}
		if i+1 >= len(s) {
			return last
		}
		if strings.HasPrefix(s[i:], "<!--") {
			end := strings.Index(s[i+4:], "-->")
			if end < 0 {
				return last
			}
			i += end + 6
			continue
		}
		closing := s[i+1] == '/'
		j := i + 1
		if closing {
			j++
		}
		nameStart := j
		for j < len(s) && isFootnoteNameByte(s[j]) {
			j++
		}
		name := s[nameStart:j]
		for j < len(s) && s[j] != '>' {
			j++
		}
		if j >= len(s) {
			return last
		}
		switch {
		case closing:
			if depth > 0 {
				depth--
			}
			if depth == 0 && strings.EqualFold(name, "p") {
				last = i
			}
		case name != "" && !footnoteVoidElement(name):
			depth++
		}
		i = j
	}
	return last
}

func footnoteVoidElement(name string) bool {
	switch strings.ToLower(name) {
	case "area", "base", "br", "col", "embed", "hr", "img", "input",
		"link", "meta", "param", "source", "track", "wbr":
		return true
	default:
		return false
	}
}

func isFootnoteNameByte(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-'
}

// The label closes at the first unescaped ]; an unescaped [ rejects the label.
func footnoteLabelClosure(line string, start int) int {
	for i := start; i < len(line); i++ {
		switch c := line[i]; {
		case c == '\\' && i+1 < len(line) && isFootnotePunct(line[i+1]):
			i++
		case c == '[':
			return -1
		case c == ']':
			return i
		}
	}
	return -1
}

func isFootnotePunct(c byte) bool {
	switch {
	case c >= '!' && c <= '/', c >= ':' && c <= '@',
		c >= '[' && c <= '`', c >= '{' && c <= '~':
		return true
	default:
		return false
	}
}

// footnoteDedent removes exactly four visual columns of leading indent,
// counting a tab as enough columns to reach the next multiple of four.
func footnoteDedent(line string) (string, bool) {
	width := 0
	for i := range len(line) {
		switch {
		case line[i] == ' ' && width < 4:
			width++
		case line[i] == '\t' && width < 4:
			width += 4 - width%4
		default:
			if width < 4 {
				return "", false
			}
			return line[i:], true
		}
	}
	return "", false
}

func isFootnoteBlank(s string) bool {
	for i := range len(s) {
		switch s[i] {
		case ' ', '\t', '\v', '\f', '\r':
		default:
			return false
		}
	}
	return true
}

// Preserve significant first-line indentation for the block parser.
func footnoteTrimLead(line string) string {
	width := 0
	for i := range len(line) {
		switch line[i] {
		case ' ':
			width++
		case '\t':
			width += 4 - width%4
		default:
			if width < 4 {
				return line[i:]
			}
			return line
		}
	}
	return ""
}
