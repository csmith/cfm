package cfm

import (
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

type blockKind uint8

const (
	blockDocument blockKind = iota
	blockParagraph
	blockHeading
	blockThematicBreak
	blockCode
	blockHTML
	blockQuote
	blockList
	blockItem
	blockTable
	blockFootnoteDef
)

// block is one node of the parsed document tree, which covers the whole
// input so later passes can inspect every block before rendering begins.
type block struct {
	kind     blockKind
	parent   *block
	children []*block

	lines []string

	level      int
	id         string
	fence      *fenceBlock
	html       *htmlBlock
	table      *markdownTable
	ordered    bool
	listStart  int
	marker     byte
	contentCol int
	indent     int

	blankBefore bool
	tight       bool

	// Cached inline output from the document-wide pass; tableCells holds the
	// rendered cells of a table block in render order.
	inlineHTML string
	tableCells *tableInline
}

// fenceBlock holds the fields of a fenced code block. Kind-specific fields
// live outside block to keep the common paragraph node small.
type fenceBlock struct {
	char   byte
	length int
	indent int
	lang   string
	code   []string
}

// htmlBlock holds the fields of a raw HTML block.
type htmlBlock struct {
	typ   int
	end   string
	lines []string
}

// tableInline caches the rendered cells of one table so the document pass
// can register footnote references in source order and rendering reuses the
// output in the same order.
type tableInline struct {
	cells []string
	pos   int
}

func (t *tableInline) inline(string) string {
	if t == nil || t.pos >= len(t.cells) {
		return ""
	}
	cell := t.cells[t.pos]
	t.pos++
	return cell
}

type docContext struct {
	ids       *headingIDs
	notes     *footnotes
	bodies    map[string]*blockParser
	defOrder  []*blockParser
	defAnchor *block
}

func newDocContext() *docContext {
	return &docContext{ids: &headingIDs{}, notes: &footnotes{}}
}

func (dc *docContext) inlineFunc() func(string, bool) string {
	body := &inlineContext{footnote: dc.notes.reference}
	cell := &inlineContext{footnote: body.footnote, table: true}
	return func(text string, table bool) string {
		if table {
			return renderInline(text, cell)
		}
		return renderInline(text, body)
	}
}

// Parse each definition once so heading IDs register in source order.
// Cache by label: distinct notes can contain identical source text.
func (dc *docContext) parseBody(label, body string, depth int) {
	if _, ok := dc.bodies[label]; ok {
		return
	}
	p := newBlockParser(dc)
	p.depth = depth + 1
	p.parse(body)
	if dc.bodies == nil {
		dc.bodies = make(map[string]*blockParser)
	}
	dc.bodies[label] = p
	dc.defOrder = append(dc.defOrder, p)
}

// Cache inline output in source order so rendering cannot register refs twice.
// Note bodies join the inline pass at the first definition's position.
func prepassInline(root *block, dc *docContext, inline func(string, bool) string) {
	ip := &inlinePrepass{dc: dc, inline: inline}
	ip.blocks(root.children)
}

type inlinePrepass struct {
	dc     *docContext
	inline func(string, bool) string
}

func (ip *inlinePrepass) blocks(children []*block) {
	for _, b := range children {
		switch b.kind {
		case blockParagraph:
			b.inlineHTML = ip.inline(blockText(b), false)
		case blockHeading:
			b.inlineHTML = ip.inline(b.lines[0], false)
		case blockTable:
			ip.table(b)
		case blockFootnoteDef:
			if b == ip.dc.defAnchor {
				for _, tree := range ip.dc.defOrder {
					ip.blocks(tree.root.children)
				}
			}
		default:
			ip.blocks(b.children)
		}
	}
}

func (ip *inlinePrepass) table(b *block) {
	cells := &tableInline{}
	for _, c := range b.table.header {
		cells.cells = append(cells.cells, ip.inline(c, true))
	}
	for _, row := range b.table.rows {
		for _, c := range row {
			cells.cells = append(cells.cells, ip.inline(c, true))
		}
	}
	b.tableCells = cells
}

func (dc *docContext) noteHTML(label string, r Renderer) string {
	p := dc.bodies[label]
	body := p.render(r)
	backlinks := dc.notes.backlinks(label, dc.notes.numbers[label])
	return attachFootnoteBacklinks(body, backlinks)
}

type blockParser struct {
	root       *block
	open       []*block
	doc        *docContext
	srcLines   []string
	lineNo     int
	emptyBlank bool
	rests      []string
	prevRests  []string
	raw        string
	depth      int
	srcLen     int
}

func newBlockParser(dc *docContext) *blockParser {
	doc := &block{kind: blockDocument}
	return &blockParser{root: doc, open: []*block{doc}, doc: dc}
}

func (p *blockParser) cur() *block { return p.open[len(p.open)-1] }

func (p *blockParser) addChild(b *block) {
	b.parent = p.cur()
	b.blankBefore = p.blankBeforeAt(b.parent)
	b.parent.children = append(b.parent.children, b)
}

func (p *blockParser) blankBeforeAt(parent *block) bool {
	i := p.indexOf(parent)
	return i >= 0 && i < len(p.prevRests) && isBlankLine(p.prevRests[i])
}

func (p *blockParser) openBlock(b *block) {
	p.addChild(b)
	p.open = append(p.open, b)
}

func (p *blockParser) addLeaf(b *block) {
	p.addChild(b)
}

func (p *blockParser) closeTo(n int) {
	if n < 0 {
		n = 0
	}
	for k := len(p.open) - 1; k > n; k-- {
		if b := p.open[k]; b.kind == blockList {
			p.setTight(b)
		}
	}
	p.open = p.open[:n+1]
}

func (p *blockParser) setTight(l *block) {
	tight := true
	for i, item := range l.children {
		if i > 0 && item.blankBefore && item.kind != blockFootnoteDef {
			tight = false
		}
		seen := false
		for _, c := range item.children {
			if c.kind == blockFootnoteDef {
				continue
			}
			if seen && c.blankBefore {
				tight = false
			}
			seen = true
		}
	}
	l.tight = tight
}

func (p *blockParser) parse(source string) {
	p.srcLen = len(source)
	p.srcLines = strings.Split(source, "\n")
	if last := len(p.srcLines) - 1; last >= 0 && p.srcLines[last] == "" && strings.HasSuffix(source, "\n") {
		p.srcLines = p.srcLines[:last]
	}
	for p.lineNo < len(p.srcLines) {
		p.line(p.srcLines[p.lineNo])
		p.lineNo++
	}
	p.closeTo(0)
	p.splitTables(p.root)
}

// Reverse traversal keeps sibling indices stable when a paragraph expands.
func (p *blockParser) splitTables(b *block) {
	for i, c := range slices.Backward(b.children) {
		p.splitTables(c)
		if c.kind != blockParagraph {
			continue
		}
		head, t := paragraphTable(c.lines)
		if t == nil {
			continue
		}
		parts := make([]*block, 0, 2)
		if head > 0 {
			parts = append(parts, &block{kind: blockParagraph, parent: b, blankBefore: c.blankBefore, lines: c.lines[:head]})
		}
		parts = append(parts, &block{kind: blockTable, parent: b, table: t})
		b.children = slices.Replace(b.children, i, i+1, parts...)
	}
}

// Only the first delimiter candidate can form a table; its preceding lines
// remain a paragraph. A run containing only dashes is not a table delimiter.
func paragraphTable(lines []string) (int, *markdownTable) {
	for i := 1; i < len(lines); i++ {
		if parseTableDelimiter(lines[i]) == nil || strings.Trim(lines[i], "-") == "" {
			continue
		}
		t, _ := parseTable(lines[i-1:])
		if t == nil {
			return 0, nil
		}
		return i - 1, t
	}
	return 0, nil
}

// walk consumes the container prefixes of raw, recording the remainder as
// seen at each container depth.
func (p *blockParser) walk(raw string) (rest string, baseCol, matched int, ok bool) {
	rest = raw
	p.rests = p.rests[:0]
	for i := range len(p.open) {
		switch c := p.open[i]; c.kind {
		case blockDocument, blockList:
		case blockQuote:
			s, w := splitIndent(rest)
			if w > 3 || !strings.HasPrefix(s, ">") {
				return rest, baseCol, i, false
			}
			rest = stripOneSpace(s[1:])
		case blockItem:
			if isBlankLine(rest) {
				p.rests = append(p.rests, rest)
				continue
			}
			w := indentWidth(rest)
			if w < c.indent {
				return rest, baseCol, i, false
			}
			rest = dedent(rest, c.indent)
			baseCol += c.indent
		}
		p.rests = append(p.rests, rest)
	}
	return rest, baseCol, len(p.open), true
}

func (p *blockParser) line(raw string) {
	p.raw = raw
	if p.emptyBlank && !p.isMarkerLine(raw) {
		p.emptyBlank = false
		if di := p.deepestItemIndex(); di >= 0 {
			p.closeTo(di - 2)
		}
	}
	rest, baseCol, matched, ok := p.walk(raw)
	if !ok {
		if p.lazyContinues(rest, baseCol, matched) {
			p.cur().lines = append(p.cur().lines, stripIndent(rest))
			return
		}
		p.closeTo(matched - 1)
		p.line(raw)
		return
	}
	// Marker lines keep an item-less list open so siblings can still join.
	if p.cur().kind == blockList && !p.isMarkerLine(rest) {
		p.closeTo(len(p.open) - 2)
	}
	switch cur := p.cur(); cur.kind {
	case blockCode:
		p.fenceLine(cur, rest)
		return
	case blockHTML:
		p.htmlLine(cur, rest)
		if isBlankLine(rest) && cur.html.typ >= 6 {
			p.swapRests()
		}
		return
	}
	p.lineAt(rest, baseCol)
	p.swapRests()
}

func (p *blockParser) swapRests() {
	p.rests, p.prevRests = p.prevRests, p.rests
}

func (p *blockParser) interrupts(rest string, baseCol int) bool {
	s, _ := splitIndent(rest)
	if s == "" {
		return false
	}
	if isHyphenBreak(s) {
		return true
	}
	if _, _, ok := parseATX(s); ok {
		return true
	}
	if _, ok := parseFenceStart(s); ok {
		return true
	}
	if t := htmlBlockType(s, true); t != 0 && t != 7 {
		return true
	}
	if strings.HasPrefix(s, ">") {
		return true
	}
	if m, ok := parseListMarker(s); ok {
		if indentWidth(rest) > 3 {
			return false
		}
		action, li := p.listCandidacy(baseCol+indentWidth(rest), m)
		if action == markerRefused {
			return false
		}
		if action == markerFresh && li < 0 && (m.blank || (m.ordered && m.num != 1)) {
			return false
		}
		return true
	}
	return false
}

func (p *blockParser) lazyContinues(rest string, baseCol, matched int) bool {
	cur := p.cur()
	_, isDef := newFootnoteDef(rest)
	return cur.kind == blockParagraph && !isBlankLine(rest) && !isDef &&
		!p.interrupts(rest, baseCol) && p.lazyAllowed(matched, baseCol)
}

func (p *blockParser) lazyAllowed(matched, baseCol int) bool {
	for i := matched; i < len(p.open); i++ {
		if p.open[i].kind == blockQuote {
			return true
		}
	}
	return baseCol == 0 && p.cur().parent == p.open[matched]
}

func (p *blockParser) isMarkerLine(raw string) bool {
	s, _ := splitIndent(raw)
	_, ok := parseListMarker(s)
	return ok
}

func (p *blockParser) deepestItemIndex() int {
	for i, v := range slices.Backward(p.open) {
		if v.kind == blockItem {
			return i
		}
	}
	return -1
}

func (p *blockParser) closeParagraph() {
	if p.cur().kind == blockParagraph {
		p.closeTo(len(p.open) - 2)
	}
}

func (p *blockParser) lineAt(rest string, baseCol int) {
	if isBlankLine(rest) {
		p.closeParagraph()
		if di := p.deepestItemIndex(); di >= 0 && len(p.open[di].children) == 0 {
			p.emptyBlank = true
		}
		return
	}
	s, w := splitIndent(rest)
	switch {
	case isHyphenBreak(s):
		p.closeParagraph()
		p.addLeaf(&block{kind: blockThematicBreak})
	case strings.HasPrefix(s, ">"):
		p.closeParagraph()
		p.openBlock(&block{kind: blockQuote})
		p.lineAt(stripOneSpace(s[1:]), 0)
	default:
		p.textLine(rest, s, w, baseCol)
	}
}

func (p *blockParser) textLine(rest, s string, w, baseCol int) {
	if label, body, consumed := p.footnoteDef(rest); consumed > 0 {
		p.closeParagraph()
		if p.doc.notes.define(label, body) {
			p.doc.parseBody(label, body, p.depth)
		}
		switch {
		case p.depth > 0:
			p.addLeaf(&block{kind: blockFootnoteDef})
		case p.doc.defAnchor == nil:
			p.doc.defAnchor = &block{kind: blockFootnoteDef}
			p.addLeaf(p.doc.defAnchor)
		}
		p.lineNo += consumed - 1
		return
	}
	if level, content, ok := parseATX(s); ok {
		p.closeParagraph()
		h := &block{kind: blockHeading, level: level, lines: []string{content}}
		h.id = p.doc.ids.generate(headingIDSource(content))
		p.addLeaf(h)
		return
	}
	if fs, ok := parseFenceStart(s); ok {
		p.closeParagraph()
		p.openBlock(
			&block{kind: blockCode, fence: &fenceBlock{
				char:   fs.char,
				length: fs.length,
				indent: w,
				lang:   fenceLanguage(fs.language),
			}},
		)
		return
	}
	if t := htmlBlockType(s, p.cur().kind == blockParagraph); t != 0 {
		p.closeParagraph()
		h := &htmlBlock{typ: t, end: htmlEndMarker(t, s)}
		h.lines = append(h.lines, rest)
		p.openBlock(&block{kind: blockHTML, html: h})
		if h.end != "" && strings.Contains(strings.ToLower(rest), h.end) {
			p.closeTo(len(p.open) - 2)
		}
		return
	}
	if m, ok := parseListMarker(s); ok {
		if p.paragraphKeepsMarker(w, baseCol, m) {
			p.cur().lines = append(p.cur().lines, s)
			return
		}
		p.handleMarker(s, baseCol+w, m)
		return
	}
	if p.cur().kind == blockParagraph {
		p.cur().lines = append(p.cur().lines, s)
		return
	}
	para := &block{kind: blockParagraph}
	p.openBlock(para)
	para.lines = append(para.lines, s)
}

// footnoteDef reads a footnote definition starting at the current line.
// Continuation lines are gathered one at a time through the container walk
// so their own indentation survives for the four-column body rule, and the
// scan stops as soon as the body cannot continue. The speculative walks run
// on a scratch slice to leave the parser's container bookkeeping alone.
func (p *blockParser) footnoteDef(rest string) (label, body string, consumed int) {
	d, ok := newFootnoteDef(rest)
	if !ok {
		return "", "", 0
	}
	saved := p.rests
	p.rests = nil
	taken := 0
	stopped := false
	for _, raw := range p.srcLines[p.lineNo+1:] {
		line, baseCol, _, ok := p.walk(raw)
		if !ok {
			// Containers closed: only lazy paragraph text continues, and a
			// blank line ends the body, leaving the line to close containers.
			if !isFootnoteBlank(raw) && p.lazyFootnoteLine(d, raw, 0) {
				taken++
				continue
			}
			stopped = true
			break
		}
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
		if p.lazyFootnoteLine(d, line, baseCol) {
			taken++
			continue
		}
		stopped = true
		break
	}
	p.rests = saved
	consumed = 1 + taken
	if stopped {
		consumed += d.pending
	}
	return d.label, d.text(), consumed
}

// lazyFootnoteLine continues the definition's open paragraph with an
// unindented line, matching lazy continuation through closed containers. A
// new definition start or an interrupting block ends it.
func (p *blockParser) lazyFootnoteLine(d *footnoteDefState, line string, baseCol int) bool {
	return lazyFootnoteBody(
		d,
		line,
		func(s string) bool {
			if _, ok := newFootnoteDef(s); ok {
				return false
			}
			return !p.interrupts(s, baseCol)
		},
	)
}

// paragraphKeepsMarker reports whether a list marker that cannot start a
// list at this position continues the open paragraph instead.
func (p *blockParser) paragraphKeepsMarker(w, baseCol int, m listMarker) bool {
	if p.cur().kind != blockParagraph {
		return false
	}
	if w > 3 {
		return true
	}
	action, li := p.listCandidacy(baseCol+w, m)
	return action == markerRefused ||
		action == markerFresh && li < 0 && (m.blank || m.ordered && m.num != 1)
}

type markerAction int

const (
	markerRefused markerAction = iota
	markerNested
	markerSibling
	markerFresh
)

// listCandidacy classifies a list marker at M visual columns into the current
// container level, returning the index of the target list for siblings.
func (p *blockParser) listCandidacy(M int, m listMarker) (markerAction, int) {
	if di := p.deepestItemIndex(); di >= 0 {
		item := p.open[di]
		if M >= item.contentCol {
			if M <= item.contentCol+3 {
				return markerNested, -1
			}
			return markerRefused, -1
		}
		return p.listCandidacyAt(item.parent, M, m)
	}
	if p.cur().kind == blockList {
		return p.listCandidacyAt(p.cur(), M, m)
	}
	return markerFresh, -1
}

func (p *blockParser) listCandidacyAt(list *block, M int, m listMarker) (markerAction, int) {
	li := p.indexOf(list)
	if !listSameType(list, m) {
		return markerFresh, li
	}
	up := 3
	if enc := list.parent; enc.kind == blockItem {
		up = enc.contentCol + 3
	}
	if M <= up {
		return markerSibling, li
	}
	return markerRefused, -1
}

func (p *blockParser) indexOf(b *block) int {
	for i := range p.open {
		if p.open[i] == b {
			return i
		}
	}
	return -1
}

func listSameType(l *block, m listMarker) bool {
	if l == nil || l.kind != blockList {
		return false
	}
	if m.ordered {
		return l.ordered && l.marker == m.marker
	}
	return !l.ordered && l.marker == m.marker
}

func (p *blockParser) handleMarker(s string, M int, m listMarker) {
	action, li := p.listCandidacy(M, m)
	switch action {
	case markerRefused:
		if p.cur().kind == blockParagraph {
			p.cur().lines = append(p.cur().lines, s)
			return
		}
		if di := p.deepestItemIndex(); di >= 0 {
			p.closeTo(di - 2)
			p.line(p.raw)
		}
		return
	case markerSibling:
		p.closeTo(li)
	default:
		p.closeParagraph()
		if li >= 0 {
			p.closeTo(li - 1)
		}
	}
	if p.cur().kind != blockList || !listSameType(p.cur(), m) {
		l := &block{kind: blockList, ordered: m.ordered, marker: m.marker, listStart: m.num}
		p.openBlock(l)
	}
	indent := m.markerW + m.spaces
	if m.blank {
		indent = m.markerW + 1
	}
	item := &block{kind: blockItem, contentCol: M + indent, indent: indent}
	p.openBlock(item)
	if m.blank {
		return
	}
	content := dedent(s[m.markerW:], m.spaces)
	p.lineAt(content, M+indent)
}

func (p *blockParser) fenceLine(b *block, rest string) {
	f := b.fence
	s, w := splitIndent(rest)
	n := 0
	for n < len(s) && s[n] == f.char {
		n++
	}
	if w-f.indent <= 3 && n >= f.length && isBlankLine(s[n:]) {
		p.closeTo(len(p.open) - 2)
		return
	}
	f.code = append(f.code, dedent(rest, f.indent))
}

func (p *blockParser) htmlLine(b *block, rest string) {
	h := b.html
	if h.typ <= 5 {
		h.lines = append(h.lines, rest)
		if h.end != "" && strings.Contains(strings.ToLower(rest), h.end) {
			p.closeTo(len(p.open) - 2)
		}
		return
	}
	if isBlankLine(rest) {
		p.closeTo(len(p.open) - 2)
		return
	}
	h.lines = append(h.lines, rest)
}

func (p *blockParser) render(r Renderer) string {
	br := blockRenderer{p: p, r: r}
	// Escapes and tags make the output somewhat longer than its source.
	br.out.Grow(p.srcLen + p.srcLen/2)
	br.stack = append(make([]renderFrame, 0, 16), renderFrame{b: p.root})
	br.run()
	return br.out.String()
}

type renderFrame struct {
	b         *block
	literal   string
	sep       bool
	tightPara bool
}

type blockRenderer struct {
	p     *blockParser
	out   strings.Builder
	r     Renderer
	stack []renderFrame
}

func (br *blockRenderer) run() {
	for len(br.stack) > 0 {
		f := br.stack[len(br.stack)-1]
		br.stack = br.stack[:len(br.stack)-1]
		if f.literal != "" {
			br.out.WriteString(f.literal)
			continue
		}
		br.frame(f)
	}
}

func (br *blockRenderer) frame(f renderFrame) {
	if f.sep && br.out.Len() > 0 && br.out.String()[br.out.Len()-1] != '\n' {
		br.out.WriteByte('\n')
	}
	switch b := f.b; b.kind {
	case blockDocument:
		br.children(b)
	case blockParagraph:
		br.paragraph(b, f.tightPara)
	case blockHeading:
		br.heading(b)
	case blockThematicBreak:
		br.out.WriteString("<hr>\n")
	case blockCode:
		br.out.WriteString(renderCode(b, br.r))
	case blockHTML:
		br.rawHTML(b)
	case blockQuote:
		br.wrap(b, "<blockquote>\n", "</blockquote>\n")
	case blockList:
		br.list(b)
	case blockItem:
		br.item(b)
	case blockTable:
		cells := b.tableCells
		br.out.WriteString(renderTable(b.table, cells.inline))
	}
}

func (br *blockRenderer) paragraph(b *block, tight bool) {
	text := b.inlineHTML
	if tight {
		br.out.WriteString(text)
		return
	}
	br.out.WriteString("<p>")
	br.out.WriteString(text)
	br.out.WriteString("</p>\n")
}

func (br *blockRenderer) heading(b *block) {
	level := byte('0' + b.level)
	br.out.WriteString("<h")
	br.out.WriteByte(level)
	br.out.WriteString(` id="`)
	br.out.WriteString(b.id)
	br.out.WriteString(`">`)
	br.out.WriteString(b.inlineHTML)
	br.out.WriteString("</h")
	br.out.WriteByte(level)
	br.out.WriteString(">\n")
}

func (br *blockRenderer) rawHTML(b *block) {
	for _, l := range b.html.lines {
		br.out.WriteString(l)
		br.out.WriteByte('\n')
	}
}

func (br *blockRenderer) wrap(b *block, openTag, closeTag string) {
	br.out.WriteString(openTag)
	br.stack = append(br.stack, renderFrame{literal: closeTag})
	br.children(b)
}

func (br *blockRenderer) list(b *block) {
	if !b.ordered {
		br.wrap(b, "<ul>\n", "</ul>\n")
		return
	}
	open := "<ol>\n"
	if b.listStart != 1 {
		open = `<ol start="` + strconv.Itoa(b.listStart) + "\">\n"
	}
	br.wrap(b, open, "</ol>\n")
}

func (br *blockRenderer) item(b *block) {
	br.out.WriteString("<li>")
	if c := firstRenderedChild(b); c != nil && (!b.parent.tight || c.kind != blockParagraph) {
		br.out.WriteByte('\n')
	}
	br.stack = append(br.stack, renderFrame{literal: "</li>\n"})
	br.children(b)
}

// firstRenderedChild skips definition anchors, which produce no output.
func firstRenderedChild(b *block) *block {
	for _, c := range b.children {
		if c.kind != blockFootnoteDef {
			return c
		}
	}
	return nil
}

func (br *blockRenderer) children(b *block) {
	br.stack = slices.Grow(br.stack, len(b.children))
	for i, v := range slices.Backward(b.children) {
		if v.kind == blockFootnoteDef {
			continue
		}
		br.stack = append(
			br.stack,
			renderFrame{
				b: v, sep: hasRenderedSibling(b.children[:i]),
				tightPara: b.kind == blockItem && b.parent.tight && v.kind == blockParagraph,
			},
		)
	}
}

func hasRenderedSibling(children []*block) bool {
	for _, c := range children {
		if c.kind != blockFootnoteDef {
			return true
		}
	}
	return false
}

func blockText(b *block) string {
	return strings.TrimRight(strings.Join(b.lines, "\n"), " \t")
}

func renderCode(b *block, r Renderer) string {
	f := b.fence
	code := strings.Join(f.code, "\n")
	if len(f.code) > 0 {
		code += "\n"
	}
	if r.Highlight != nil {
		if out, err := r.Highlight(code, f.lang); err == nil {
			return out
		}
	}
	attr := ""
	if f.lang != "" {
		attr = ` class="language-` + escapeHTML(f.lang) + `"`
	}
	return "<pre><code" + attr + ">" + escapeHTML(code) + "</code></pre>\n"
}

// splitIndent strips leading spaces and tabs, reporting their visual width
// (tabs advance to the next multiple of four).
func splitIndent(s string) (string, int) {
	i, w := 0, 0
	for i < len(s) {
		switch s[i] {
		case ' ':
			w++
		case '\t':
			w += 4 - w%4
		default:
			return s[i:], w
		}
		i++
	}
	return "", w
}

func stripIndent(s string) string {
	rest, _ := splitIndent(s)
	return rest
}

func indentWidth(s string) int {
	_, w := splitIndent(s)
	return w
}

// dedent removes up to cols leading whitespace columns. A tab straddling the
// boundary is replaced by the spaces that reach it.
func dedent(s string, cols int) string {
	i, w := 0, 0
	for i < len(s) && w < cols {
		switch s[i] {
		case ' ':
			w++
		case '\t':
			tw := 4 - w%4
			if w+tw > cols {
				return strings.Repeat(" ", w+tw-cols) + s[i+1:]
			}
			w += tw
		default:
			return s[i:]
		}
		i++
	}
	return s[i:]
}

func isBlankLine(s string) bool {
	return stripIndent(s) == ""
}

func stripOneSpace(s string) string {
	if len(s) > 0 && (s[0] == ' ' || s[0] == '\t') {
		return s[1:]
	}
	return s
}

func isHyphenBreak(s string) bool {
	if s == "" || s[0] != '-' {
		return false
	}
	s = strings.TrimRight(s, " \t")
	if len(s) < 3 {
		return false
	}
	for i := range len(s) {
		if s[i] != '-' {
			return false
		}
	}
	return true
}

type listMarker struct {
	ordered bool
	marker  byte // bullet character, or '.', ')' for ordered lists
	num     int
	markerW int
	spaces  int
	blank   bool // nothing but whitespace follows the marker
}

func parseListMarker(s string) (listMarker, bool) {
	var m listMarker
	if s == "" {
		return m, false
	}
	switch s[0] {
	case '-', '+', '*':
		m.marker = s[0]
		m.markerW = 1
	default:
		i := 0
		for i < len(s) && i < 9 && s[i] >= '0' && s[i] <= '9' {
			i++
		}
		if i == 0 || i >= len(s) || (s[i] != '.' && s[i] != ')') {
			return m, false
		}
		m.ordered = true
		m.marker = s[i]
		m.num, _ = strconv.Atoi(s[:i])
		m.markerW = i + 1
	}
	rest := s[m.markerW:]
	if rest == "" {
		m.blank = true
		return m, true
	}
	w := 0
	i := 0
	for i < len(rest) {
		switch rest[i] {
		case ' ':
			w++
		case '\t':
			w += 4 - w%4
		default:
			if w == 0 {
				return listMarker{}, false
			}
			m.spaces = w
			return m, true
		}
		i++
	}
	m.blank = true
	return m, true
}

// headingIDSource drops a closing hash sequence from the id source only; the
// sequence stays in the rendered heading text.
func headingIDSource(s string) string {
	i := len(s)
	for i > 0 && s[i-1] == '#' {
		i--
	}
	if i < len(s) && i > 1 && (s[i-1] == ' ' || s[i-1] == '\t') {
		return s[:i-1]
	}
	return s
}

func parseATX(s string) (level int, content string, ok bool) {
	i := 0
	for i < len(s) && s[i] == '#' {
		i++
	}
	if i == 0 || i > 6 || i == len(s) {
		return 0, "", false
	}
	if s[i] != ' ' && s[i] != '\t' {
		return 0, "", false
	}
	content = strings.Trim(s[i:], " \t")
	if content == "" {
		return 0, "", false
	}
	return i, content, true
}

type fenceStart struct {
	char     byte
	length   int
	language string
}

func parseFenceStart(s string) (fenceStart, bool) {
	if s == "" || (s[0] != '`' && s[0] != '~') {
		return fenceStart{}, false
	}
	n := 0
	for n < len(s) && s[n] == s[0] {
		n++
	}
	if n < 3 {
		return fenceStart{}, false
	}
	info := strings.Trim(s[n:], " \t")
	if s[0] == '`' && strings.ContainsRune(info, '`') {
		return fenceStart{}, false
	}
	return fenceStart{char: s[0], length: n, language: firstWord(info)}, true
}

func firstWord(s string) string {
	if i := strings.IndexAny(s, " \t"); i >= 0 {
		return s[:i]
	}
	return s
}

// headingIDs tracks the ids issued in one document. next records the first
// suffix worth trying for each base: ids are never released, so every
// earlier suffix is still taken and need not be probed again.
type headingIDs struct {
	used map[string]bool
	next map[string]int
}

// generate builds a goldmark-compatible heading id: surrounding whitespace
// is trimmed, ASCII letters and digits are lowercased, spaces, hyphens and
// underscores become hyphens, everything else is dropped; empty ids fall
// back to "heading" and collisions get numeric suffixes.
func (h *headingIDs) generate(raw string) string {
	raw = strings.Trim(raw, " \t\n\v\f\r")
	var b strings.Builder
	b.Grow(len(raw))
	for i := 0; i < len(raw); {
		c := raw[i]
		if c < utf8.RuneSelf {
			switch {
			case 'A' <= c && c <= 'Z':
				b.WriteByte(c + 'a' - 'A')
			case 'a' <= c && c <= 'z', '0' <= c && c <= '9':
				b.WriteByte(c)
			case c == ' ', c == '-', c == '_':
				b.WriteByte('-')
			}
			i++
			continue
		}
		_, size := utf8.DecodeRuneInString(raw[i:])
		i += size
	}
	base := b.String()
	if base == "" {
		base = "heading"
	}
	if h.used == nil {
		h.used = make(map[string]bool)
		h.next = make(map[string]int)
	}
	id := base
	if h.used[id] {
		n := max(h.next[base], 1)
		for {
			id = base + "-" + strconv.Itoa(n)
			if !h.used[id] {
				break
			}
			n++
		}
		h.next[base] = n + 1
	}
	h.used[id] = true
	return id
}

// fenceLanguage resolves the backslash escapes and character references
// that CommonMark recognizes in a fence info string.
func fenceLanguage(s string) string {
	if !strings.ContainsAny(s, "&\\") {
		return s
	}
	var b strings.Builder
	for i := 0; i < len(s); {
		if s[i] == '\\' && i+1 < len(s) && isPunctByte(s[i+1]) {
			b.WriteByte(s[i+1])
			i += 2
			continue
		}
		if s[i] == '&' {
			if r, end, ok := numericRefValue(s, i); ok {
				b.WriteRune(toValidRune(r))
				i = end
				continue
			}
			if chars, end, ok := namedEntityRef(s, i); ok {
				b.WriteString(chars)
				i = end
				continue
			}
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}

func namedEntityRef(s string, i int) (string, int, bool) {
	j := i + 1
	for j < len(s) && isAlphaNumeric(s[j]) {
		j++
	}
	if j == i+1 || j >= len(s) || s[j] != ';' {
		return "", 0, false
	}
	name := s[i : j+1]
	chars, ok := decodeNamedEntity(name)
	if !ok {
		return "", 0, false
	}
	return chars, j + 1, true
}

var htmlBlock1Names = []string{"script", "pre", "style", "textarea"}

var htmlBlock6Names = map[string]bool{
	"address": true, "article": true, "aside": true, "base": true,
	"basefont": true, "blockquote": true, "body": true, "caption": true,
	"center": true, "col": true, "colgroup": true, "dd": true, "details": true,
	"dialog": true, "dir": true, "div": true, "dl": true, "dt": true,
	"fieldset": true, "figcaption": true, "figure": true, "footer": true,
	"form": true, "frame": true, "frameset": true, "h1": true, "h2": true,
	"h3": true, "h4": true, "h5": true, "h6": true, "head": true,
	"header": true, "hr": true, "html": true, "iframe": true, "legend": true,
	"li": true, "link": true, "main": true, "menu": true, "menuitem": true,
	"nav": true, "noframes": true, "ol": true, "optgroup": true, "option": true,
	"p": true, "param": true, "search": true, "section": true, "summary": true,
	"table": true, "tbody": true, "td": true, "tfoot": true, "th": true,
	"thead": true, "title": true, "tr": true, "track": true, "ul": true,
}

// htmlBlockType classifies a raw HTML block start on the indentation-stripped
// line s. Type 7 requires that no paragraph is open.
func htmlBlockType(s string, paraOpen bool) int {
	if s == "" || s[0] != '<' {
		return 0
	}
	lower := strings.ToLower(s)
	for _, name := range htmlBlock1Names {
		if strings.HasPrefix(lower, "<"+name) {
			switch rest := s[1+len(name):]; {
			case rest == "", rest[0] == ' ', rest[0] == '\t', rest[0] == '>':
				return 1
			}
		}
	}
	switch {
	case strings.HasPrefix(s, "<!--"):
		return 2
	case strings.HasPrefix(s, "<?"):
		return 3
	case strings.HasPrefix(s, "<![CDATA["):
		return 5
	case len(s) >= 3 && s[1] == '!' && isASCIILetter(s[2]):
		return 4
	}
	if name, after, ok := htmlTagName(s); ok && htmlBlock6Names[name] {
		if after == "" || after[0] == ' ' || after[0] == '\t' ||
			strings.HasPrefix(after, ">") || strings.HasPrefix(after, "/>") {
			return 6
		}
	}
	if !paraOpen && isCompleteTagAlone(s) {
		return 7
	}
	return 0
}

func htmlEndMarker(t int, s string) string {
	switch t {
	case 2:
		return "-->"
	case 3:
		return "?>"
	case 4:
		return ">"
	case 5:
		return "]]>"
	case 1:
		lower := strings.ToLower(s)
		for _, name := range htmlBlock1Names {
			if strings.HasPrefix(lower, "<"+name) {
				return "</" + name + ">"
			}
		}
	}
	return ""
}

func htmlTagName(s string) (name, after string, ok bool) {
	i := 1
	if i < len(s) && s[i] == '/' {
		i++
	}
	start := i
	for i < len(s) && (isASCIILetter(s[i]) || isASCIIDigit(s[i]) || s[i] == '-') {
		i++
	}
	if i == start || !isASCIILetter(s[start]) {
		return "", "", false
	}
	return strings.ToLower(s[start:i]), s[i:], true
}

func isASCIILetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// isCompleteTagAlone reports whether s is a complete opening or closing tag
// followed only by whitespace.
func isCompleteTagAlone(s string) bool {
	if !strings.HasPrefix(s, "<") {
		return false
	}
	_, rest, ok := htmlTagName(s)
	if !ok {
		return false
	}
	if strings.HasPrefix(s, "</") {
		rest = strings.TrimLeft(rest, " \t")
		return strings.HasPrefix(rest, ">") && isBlankLine(rest[1:])
	}
	return isCompleteOpenTag(rest)
}

// isCompleteOpenTag consumes the attributes of an open tag body. Every
// attribute must be preceded by whitespace; a value may follow its name
// only through an equals sign.
func isCompleteOpenTag(s string) bool {
	for {
		space := strings.TrimLeft(s, " \t") != s
		s = strings.TrimLeft(s, " \t")
		switch {
		case strings.HasPrefix(s, ">"):
			return isBlankLine(s[1:])
		case strings.HasPrefix(s, "/>"):
			return isBlankLine(s[2:])
		case !space:
			return false
		}
		name, ok := cutAttributeName(s)
		if !ok {
			return false
		}
		s = strings.TrimLeft(name, " \t")
		if !strings.HasPrefix(s, "=") {
			continue
		}
		value, ok := cutAttributeValue(strings.TrimLeft(s[1:], " \t"))
		if !ok {
			return false
		}
		s = value
	}
}

func cutAttributeName(s string) (string, bool) {
	if s == "" {
		return s, false
	}
	c := s[0]
	if !isASCIILetter(c) && c != '_' && c != ':' {
		return s, false
	}
	i := 1
	for i < len(s) && isNameByte(s[i]) {
		i++
	}
	return s[i:], true
}

func isNameByte(c byte) bool {
	return isASCIILetter(c) || isASCIIDigit(c) || c == '-' || c == '_' || c == ':' || c == '.'
}

func cutAttributeValue(s string) (string, bool) {
	if s == "" {
		return s, false
	}
	switch s[0] {
	case '"', '\'':
		end := strings.IndexByte(s[1:], s[0])
		if end < 0 {
			return s, false
		}
		return s[end+2:], true
	}
	i := 0
	for i < len(s) && !strings.ContainsRune(" \t\"'=<>", rune(s[i])) && s[i] != '`' {
		i++
	}
	if i == 0 {
		return s, false
	}
	return s[i:], true
}
