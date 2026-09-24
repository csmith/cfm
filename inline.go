package cfm

import (
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// inlineContext carries block-level state for inline parsing. The footnote
// callback resolves a reference label to complete, trusted HTML.
type inlineContext struct {
	footnote func(label string) (string, bool)
	table    bool // table cell source: \| inside code spans becomes |
}

// renderInline renders the inline content of a paragraph, heading or table
// cell. Newlines in the source become hard or soft line breaks.
func renderInline(source string, ctx *inlineContext) string {
	p := &inlineParser{src: source, ctx: ctx}
	p.slabs = slabsPool.Get().(*parserSlabs)
	defer p.slabs.release()
	p.root = p.newNode(inNode{})
	p.parse()
	var b strings.Builder
	// Rendered HTML is rarely shorter than its source: escapes and tags only
	// grow it, so the source length is a good single-allocation estimate.
	b.Grow(len(source) + len(source)/8)
	renderChildren(&b, p.root)
	return b.String()
}

// parserSlabs pools one parser's node storage so consecutive calls reuse
// chunks instead of allocating and zeroing them. The pool never shares a set
// between concurrent users, so calls stay independent.
type parserSlabs struct {
	node    slab[inNode]
	delim   slab[delimiter]
	bracket slab[bracket]
}

func (s *parserSlabs) release() {
	s.node.clear()
	s.delim.clear()
	s.bracket.clear()
	slabsPool.Put(s)
}

var slabsPool = sync.Pool{New: func() any { return new(parserSlabs) }}

type inKind uint8

const (
	nText inKind = iota
	nRaw         // verbatim source segment (raw inline HTML)
	nLit         // trusted HTML string (typographer output, footnote refs)
	nCode
	nEmph
	nStrike
	nLink
	nAuto
	nBracket // an opening "[" or "![" awaiting its "]"
)

type inNode struct {
	kind                                      inKind
	seg                                       string // source text for nText/nRaw/nBracket, label for nAuto
	segStart, segStop                         int    // source offsets of seg, for text-node merging
	str                                       string // content for nCode/nLit
	dest                                      string
	title                                     string
	level                                     int // 1 = em, 2 = strong
	hasTitle                                  bool
	hard, soft                                bool
	delim                                     bool // nText is an emphasis delimiter run
	email                                     bool
	www                                       bool
	image                                     bool
	parent, prev, next, firstChild, lastChild *inNode
}
type delimiter struct {
	node              *inNode
	char              byte
	canOpen, canClose bool
	length, orig      int
	prev, next        *delimiter
}
type bracket struct {
	node    *inNode
	start   int
	bottom  *delimiter
	isImage bool
}
type inlineParser struct {
	src         string
	pos         int
	escaped     bool
	ctx         *inlineContext
	root        *inNode
	firstDelim  *delimiter
	lastDelim   *delimiter
	brackets    []*bracket
	quoteSingle int
	quoteDouble int
	slabs       *parserSlabs
}

// slab hands out parser nodes from geometric chunks. Chunks are never moved,
// so returned pointers stay valid for the parser's lifetime and nothing is
// shared between concurrent users of a set.
type slab[T any] struct {
	chunk []T
	used  int
	next  int
}

func (s *slab[T]) alloc() *T {
	if s.used == len(s.chunk) {
		size := s.next
		if size == 0 {
			size = 16
		} else if size < 512 {
			size *= 2
		}
		s.chunk = make([]T, size)
		s.next = size
		s.used = 0
	}
	n := &s.chunk[s.used]
	s.used++
	return n
}

// clear drops the parsed nodes so pooled chunks do not retain their strings.
func (s *slab[T]) clear() {
	clear(s.chunk[:s.used])
	s.used = 0
}
func (p *inlineParser) newNode(n inNode) *inNode {
	node := p.slabs.node.alloc()
	*node = n
	return node
}
func (p *inlineParser) newDelimiter(n delimiter) *delimiter {
	d := p.slabs.delim.alloc()
	*d = n
	return d
}
func (p *inlineParser) newBracket(n bracket) *bracket {
	b := p.slabs.bracket.alloc()
	*b = n
	return b
}

// parse scans line by line, mirroring goldmark's inline parse loop.
func (p *inlineParser) parse() {
	p.src = strings.TrimSuffix(p.src, "\n")
	for p.pos < len(p.src) {
		for p.parseSegment() {
		}
	}
	p.processDelimiters(nil)
}

// parseSegment scans the source from p.pos to the end of the current line.
// It reports whether a parser consumed input and scanning must resume from
// the new position.
func (p *inlineParser) parseSegment() bool {
	lineStart := p.pos
	lineEnd := p.lineEnd()
	line := p.src[lineStart:lineEnd]
	hasNL := lineEnd < len(p.src)
	scanEnd, hard, hardVisible := breakSuffix(line, hasNL)
	runStart := p.pos
	for i := 0; i < scanEnd; {
		if i > 0 {
			skip := inertRunEnd(line, i, scanEnd)
			if skip != i {
				i = skip
				p.escaped = false
				if i >= scanEnd {
					break
				}
			}
		}
		c := line[i]
		isSp := isSpaceByte(c) && c != '\r' && c != '\n'
		isPu := isPunctByte(c)
		// Backslash escaping suppresses punctuation triggers but, like goldmark
		// with escaped spaces disabled, not whitespace ones.
		if pc := triggerChar(c, i, isSp, isPu); (isSp || (isPu && !p.escaped) || i == 0) && isParserChar(pc) {
			var consumed bool
			runStart, consumed = p.tryTrigger(lineStart, lineEnd, i, pc, runStart)
			if consumed {
				return true
			}
		}
		switch {
		case p.escaped:
			p.escaped = false
		case c == '\\':
			p.escaped = true
		default:
			p.escaped = false
		}
		i++
	}
	p.appendBreakText(runStart, lineStart+scanEnd, hard, hardVisible, hasNL)
	p.pos = lineEnd
	if hasNL {
		p.pos++
	}
	return false
}

// inertRunEnd returns the offset of the next byte in line[from:end] that can
// start an inline construct or change backslash-escape state, or end.
func inertRunEnd(line string, from, end int) int {
	i := from
	for i < end && inlineScanSkip[line[i]] {
		i++
	}
	return i
}

// tryTrigger flushes pending text before lineStart+i and runs the matching
// inline parser there. It reports the pending-run start and whether a parser
// consumed input, in which case parsing resumes from the new p.pos.
func (p *inlineParser) tryTrigger(lineStart, lineEnd, i int, pc byte, runStart int) (int, bool) {
	if runStart < lineStart+i {
		p.flushText(runStart, lineStart+i)
	}
	savedPos := p.pos
	p.pos = lineStart + i
	if n := p.runParsers(pc, p.src[p.pos:lineEnd]); n != nil {
		if n.parent == nil {
			p.appendNode(n) // parsers that build on existing tree nodes attach themselves
		}
		return runStart, true
	}
	p.pos = savedPos
	return lineStart + i, false
}
func isParserChar(c byte) bool {
	switch c {
	case '`', '[', ']', '!', '<', '*', '_', '~', '(', '\'', '"', '-', '.', ',', '>', ' ':
		return true
	}
	return false
}

// inlineScanSkip marks bytes that never trigger an inline parser (parser
// characters, whitespace) and never participate in backslash escaping.
var inlineScanSkip = func() (skip [256]bool) {
	for c := range 256 {
		skip[c] = !isParserChar(byte(c)) && byte(c) != ' ' && byte(c) != '\t' && byte(c) != '\\'
	}
	return skip
}()

func triggerChar(c byte, i int, isSp, isPu bool) byte {
	if isSp || (i == 0 && !isPu) {
		return ' '
	}
	return c
}
func (p *inlineParser) runParsers(c byte, line string) *inNode {
	switch c {
	case '`':
		return p.parseCodeSpan()
	case '[':
		if n := p.parseFootnote(false); n != nil {
			return n
		}
		return p.openBracket(false)
	case '!':
		if n := p.parseFootnote(true); n != nil {
			return n
		}
		return p.openBracket(true)
	case ']':
		return p.closeBracket()
	case '<':
		if n := p.parseAutoLink(line); n != nil {
			return n
		}
		if n := p.parseRawHTML(); n != nil {
			return n
		}
		return p.parseTypographer(line)
	case '*', '_':
		return p.parseDelimiterRun(c, line)
	case '~':
		if n := p.parseDelimiterRun(c, line); n != nil {
			return n
		}
		return p.parseLinkify(line)
	case ' ', '(':
		return p.parseLinkify(line)
	case '\'', '"', '-', '.', ',', '>':
		return p.parseTypographer(line)
	}
	return nil
}

// breakSuffix reports how much of a line the scanner may examine,
// classifying its trailing break the way goldmark does.
func breakSuffix(line string, hasNL bool) (scanEnd int, hard, hardVisible bool) {
	scanEnd = len(line)
	if !hasNL {
		return len(trimRightSpace(line)), false, false
	}
	switch {
	case scanEnd >= 2 && line[scanEnd-1] == '\\' && (scanEnd == 2 || line[scanEnd-2] != '\\'):
		return scanEnd - 1, true, true
	case scanEnd >= 2 && line[scanEnd-1] == '\r' && line[scanEnd-2] == '\\' &&
		(scanEnd == 2 || line[scanEnd-3] != '\\'):
		return scanEnd - 2, true, true
	case scanEnd >= 2 && line[scanEnd-2] == ' ' && line[scanEnd-1] == ' ':
		return scanEnd - 2, true, false
	case scanEnd >= 3 && line[scanEnd-3] == ' ' && line[scanEnd-2] == ' ' && line[scanEnd-1] == '\r':
		return scanEnd - 3, true, false
	}
	return scanEnd, false, false
}
func (p *inlineParser) flushText(start, stop int) {
	if start >= stop {
		return
	}
	// Adjacent plain text merges into one node so character references split
	// by trigger boundaries resolve as a whole, like goldmark's text merging.
	if last := p.root.lastChild; last != nil && last.kind == nText && !last.delim &&
		!last.hard && !last.soft && last.segStop == start {
		last.seg = p.src[last.segStart:stop]
		last.segStop = stop
		return
	}
	p.appendNode(p.newNode(inNode{kind: nText, seg: p.src[start:stop], segStart: start, segStop: stop}))
}

// appendBreakText emits a line's trailing text with its break flags; the
// backslash of a hard break was already removed from the scanned range.
func (p *inlineParser) appendBreakText(start, stop int, hard, hardVisible, hasNL bool) {
	seg := ""
	if start < stop {
		seg = trimRightSpace(p.src[start:stop])
	}
	if hardVisible {
		if start < stop {
			seg = p.src[start:stop]
		}
	}
	node := p.newNode(inNode{kind: nText, seg: seg, hard: hard, soft: hasNL && !hard})
	if seg == "" && !hard && !node.soft {
		return
	}
	p.appendNode(node)
}
func (p *inlineParser) appendNode(n *inNode) {
	parent := p.root
	n.parent = parent
	if parent.lastChild != nil {
		n.prev = parent.lastChild
		parent.lastChild.next = n
	} else {
		parent.firstChild = n
	}
	parent.lastChild = n
}
func (p *inlineParser) parseCodeSpan() *inNode {
	src := p.src
	start := p.pos
	op := 0
	for start+op < len(src) && src[start+op] == '`' {
		op++
	}
	i := start + op
	for i < len(src) {
		if src[i] == '`' {
			j := i
			for j < len(src) && src[j] == '`' {
				j++
			}
			if j-i == op {
				p.pos = j
				content := trimCodeSpanEdges(src[start+op : i])
				if p.ctx != nil && p.ctx.table {
					// Table cells keep a backslash in front of every pipe; code
					// spans are the one place escapes never apply, so the table
					// extension unescapes \| here.
					content = strings.ReplaceAll(content, "\\|", "|")
				}
				return p.newNode(inNode{kind: nCode, str: content})
			}
			i = j
			continue
		}
		i++
	}
	p.pos = start + op
	return p.newNode(inNode{kind: nText, seg: src[start:p.pos]})
}

// trimCodeSpanEdges strips one space or newline from both ends when the
// content is not entirely blank, per CommonMark.
func trimCodeSpanEdges(content string) string {
	if len(content) < 2 || isBlank(content) {
		return content
	}
	if isSpaceOrNewline(content[0]) && isSpaceOrNewline(content[len(content)-1]) {
		return content[1 : len(content)-1]
	}
	return content
}
func (p *inlineParser) parseFootnote(bang bool) *inNode {
	if p.ctx == nil || p.ctx.footnote == nil {
		return nil
	}
	src := p.src
	pos := p.pos
	if bang {
		if pos+1 >= len(src) || src[pos+1] != '[' {
			return nil
		}
		pos += 2
	} else {
		pos++
	}
	if pos >= len(src) || src[pos] != '^' {
		return nil
	}
	pos++
	end := p.labelClosure(pos)
	if end < 0 {
		return nil
	}
	htmlStr, ok := p.ctx.footnote(src[pos:end])
	if !ok {
		return nil
	}
	if bang {
		p.appendNode(p.newNode(inNode{kind: nText, seg: "!"}))
	}
	p.pos = end + 1
	return p.newNode(inNode{kind: nLit, str: htmlStr})
}

// labelClosure finds the closing ']' of a footnote label starting at pos,
// without nesting and honouring backslash escapes. It cannot cross lines.
func (p *inlineParser) labelClosure(pos int) int {
	limit := len(p.src)
	if e := strings.IndexByte(p.src[pos:], '\n'); e >= 0 {
		limit = pos + e
	}
	for i := pos; i < limit; {
		switch c := p.src[i]; {
		case c == '\\' && i < limit-1 && isPunctByte(p.src[i+1]):
			i += 2
		case c == ']':
			return i
		case c == '[':
			return -1
		default:
			i++
		}
	}
	return -1
}
func (p *inlineParser) lineEnd() int {
	if p.pos >= len(p.src) {
		return p.pos
	}
	if e := strings.IndexByte(p.src[p.pos:], '\n'); e >= 0 {
		return p.pos + e
	}
	return len(p.src)
}

// parseDelimiterRun scans a '*', '_' or '~' run onto the delimiter stack;
// '~' runs longer than two characters or after a tilde stay literal.
func (p *inlineParser) parseDelimiterRun(c byte, line string) *inNode {
	before := p.precedingRune()
	run := 0
	for run < len(line) && line[run] == c {
		run++
	}
	after := rune(' ')
	if run < len(line) {
		after, _ = utf8.DecodeRuneInString(line[run:])
	}
	if c == '~' && (run > 2 || before == '~') {
		return nil
	}
	canOpen, canClose := flank(before, after, c)
	node := p.newNode(inNode{kind: nText, seg: line[:run], delim: true})
	d := p.newDelimiter(delimiter{node: node, char: c, canOpen: canOpen, canClose: canClose, length: run, orig: run})
	if p.lastDelim != nil {
		p.lastDelim.next = d
		d.prev = p.lastDelim
	} else {
		p.firstDelim = d
	}
	p.lastDelim = d
	p.pos += run
	return node
}

// flank applies the CommonMark left/right-flanking rules; '_' adds the
// stricter intraword conditions.
func flank(before, after rune, c byte) (canOpen, canClose bool) {
	beforeIsPunct := isPunctRune(before)
	beforeIsWS := isSpaceRune(before)
	afterIsPunct := isPunctRune(after)
	afterIsWS := isSpaceRune(after)
	isLeft := !afterIsWS && (!afterIsPunct || beforeIsWS || beforeIsPunct)
	isRight := !beforeIsWS && (!beforeIsPunct || afterIsWS || afterIsPunct)
	if c == '_' {
		return isLeft && (!isRight || beforeIsPunct), isRight && (!isLeft || afterIsPunct)
	}
	return isLeft, isRight
}
func (p *inlineParser) precedingRune() rune {
	if p.pos == 0 {
		return '\n'
	}
	r, _ := utf8.DecodeLastRuneInString(p.src[:p.pos])
	return r
}
func (d *delimiter) consume(n int) {
	d.length -= n
	d.node.seg = d.node.seg[:d.length]
}

// calcConsumption applies the CommonMark "rule of three".
func calcConsumption(opener, closer *delimiter) int {
	if (opener.canClose || closer.canOpen) && (opener.orig+closer.orig)%3 == 0 && closer.orig%3 != 0 {
		return 0
	}
	if opener.length >= 2 && closer.length >= 2 {
		return 2
	}
	return 1
}

// processDelimiters matches emphasis and strikethrough delimiters above
// bottom, converting leftovers back to literal text.
func (p *inlineParser) processDelimiters(bottom *delimiter) {
	if p.lastDelim == nil {
		return
	}
	var closer *delimiter
	if bottom == nil {
		closer = p.firstDelim
	} else {
		for c := p.lastDelim; c != nil && c != bottom; c = c.prev {
			closer = c
		}
	}
	if closer == nil {
		p.clearDelimiters(bottom)
		return
	}
	for closer != nil {
		if !closer.canClose {
			closer = closer.next
			continue
		}
		consume, found, maybeOpener := 0, false, false
		var opener *delimiter
		for opener = closer.prev; opener != nil && opener != bottom; opener = opener.prev {
			if opener.canOpen && opener.char == closer.char {
				maybeOpener = true
				consume = calcConsumption(opener, closer)
				if consume > 0 {
					found = true
					break
				}
			}
		}
		if !found {
			next := closer.next
			if !maybeOpener && !closer.canOpen {
				p.removeDelimiter(closer)
			}
			closer = next
			continue
		}
		closer = p.matchDelimiters(opener, closer, consume)
	}
	p.clearDelimiters(bottom)
}

// matchDelimiters builds one emphasis or strikethrough span and returns the
// next closer to examine.
func (p *inlineParser) matchDelimiters(opener, closer *delimiter, consume int) *delimiter {
	opener.consume(consume)
	closer.consume(consume)
	kind := nEmph
	if opener.char == '~' {
		kind = nStrike
	}
	span := p.newNode(inNode{kind: kind, level: consume})
	for child := opener.node.next; child != nil && child != closer.node; {
		next := child.next
		appendTo(span, child)
		child = next
	}
	p.insertAfter(opener.node, span)
	for c := opener.next; c != nil && c != closer; {
		next := c.next
		p.unlinkDelimiter(c)
		c = next
	}
	if opener.length == 0 {
		p.removeDelimiter(opener)
	}
	if closer.length == 0 {
		next := closer.next
		p.removeDelimiter(closer)
		return next
	}
	return closer
}

// removeDelimiter unlinks a delimiter; a fully consumed delimiter's node is
// dropped from the tree, a partly consumed one keeps its leftover text.
func (p *inlineParser) removeDelimiter(d *delimiter) {
	p.unlinkDelimiter(d)
	if d.length == 0 {
		detach(d.node)
	}
}
func (p *inlineParser) unlinkDelimiter(d *delimiter) {
	if d.prev != nil {
		d.prev.next = d.next
	} else {
		p.firstDelim = d.next
	}
	if d.next != nil {
		d.next.prev = d.prev
	} else {
		p.lastDelim = d.prev
	}
	d.prev = nil
	d.next = nil
}
func (p *inlineParser) clearDelimiters(bottom *delimiter) {
	for c := p.lastDelim; c != nil && c != bottom; {
		prev := c.prev
		p.removeDelimiter(c)
		c = prev
	}
}

// pruneDelimiters drops every delimiter above bottom without touching the
// tree; used when an image span discards the nodes they belong to.
func (p *inlineParser) pruneDelimiters(bottom *delimiter) {
	for c := p.lastDelim; c != nil && c != bottom; {
		prev := c.prev
		p.unlinkDelimiter(c)
		c = prev
	}
}
func (p *inlineParser) replaceNode(old, n *inNode) {
	n.prev = old.prev
	n.next = old.next
	n.parent = old.parent
	if old.prev != nil {
		old.prev.next = n
	} else {
		old.parent.firstChild = n
	}
	if old.next != nil {
		old.next.prev = n
	} else {
		old.parent.lastChild = n
	}
	old.prev = nil
	old.next = nil
}
func (p *inlineParser) insertAfter(ref, n *inNode) {
	n.parent = ref.parent
	n.prev = ref
	n.next = ref.next
	if ref.next != nil {
		ref.next.prev = n
	} else {
		ref.parent.lastChild = n
	}
	ref.next = n
}
func (p *inlineParser) adoptTrailingSiblings(n *inNode) {
	child := n.next
	for child != nil {
		next := child.next
		appendTo(n, child)
		child = next
	}
}
func (p *inlineParser) removeTrailingSiblings(n *inNode) {
	child := n.next
	for child != nil {
		next := child.next
		detach(child)
		child = next
	}
}
func appendTo(parent, n *inNode) {
	detach(n)
	n.parent = parent
	if parent.lastChild != nil {
		n.prev = parent.lastChild
		parent.lastChild.next = n
	} else {
		parent.firstChild = n
	}
	parent.lastChild = n
}
func detach(n *inNode) {
	if n.prev != nil {
		n.prev.next = n.next
	} else if n.parent != nil {
		n.parent.firstChild = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else if n.parent != nil {
		n.parent.lastChild = n.prev
	}
	n.prev = nil
	n.next = nil
	n.parent = nil
}

// isPunctByte reports whether c is ASCII punctuation, using the same byte
// classes as goldmark.
func isPunctByte(c byte) bool {
	switch {
	case c >= '!' && c <= '/':
		return true
	case c >= ':' && c <= '@':
		return true
	case c >= '[' && c <= '`':
		return true
	case c >= '{' && c <= '~':
		return true
	}
	return false
}
func isSpaceByte(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}
func isSpaceOrNewline(c byte) bool {
	return c == ' ' || c == '\n'
}
func isAlphaNumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}
func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}
func isDigitByte(c byte) bool {
	return c >= '0' && c <= '9'
}
func isHexDigit(c byte) bool {
	return isDigitByte(c) || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}
func isBlank(s string) bool {
	for i := range len(s) {
		switch s[i] {
		case ' ', '\t', '\n', '\r':
		default:
			return false
		}
	}
	return true
}
func trimRightSpace(s string) string {
	for len(s) > 0 && isSpaceByte(s[len(s)-1]) {
		s = s[:len(s)-1]
	}
	return s
}
func isPunctRune(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSymbol(r)
}
func isSpaceRune(r rune) bool {
	if r >= 0 && r <= 256 {
		switch byte(r) {
		case ' ', '\t', '\n', '\r':
			return true
		}
	}
	return unicode.IsSpace(r)
}
