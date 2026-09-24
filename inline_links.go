package cfm

import "strings"

func (p *inlineParser) openBracket(bang bool) *inNode {
	n := 1
	if bang {
		if p.pos+1 >= len(p.src) || p.src[p.pos+1] != '[' {
			return nil
		}
		n = 2
	}
	node := p.newNode(inNode{kind: nBracket, seg: p.src[p.pos : p.pos+n], image: bang})
	p.appendNode(node)
	p.brackets = append(p.brackets, p.newBracket(bracket{
		node:    node,
		start:   p.pos,
		bottom:  p.lastDelim,
		isImage: bang,
	}))
	p.pos += n
	return node
}

// closeBracket makes a single attempt to close the innermost bracket: a
// failed link leaves it literal so a later ']' can close an outer one.
func (p *inlineParser) closeBracket() *inNode {
	if len(p.brackets) == 0 {
		return nil
	}
	b := p.brackets[len(p.brackets)-1]
	if b.start+len(b.node.seg)-p.brackets[0].start > 998 {
		p.brackets = p.brackets[:len(p.brackets)-1]
		return nil
	}
	closePos := p.pos
	p.pos++
	dest, title, hasTitle, ok := "", "", false, false
	if p.pos < len(p.src) && p.src[p.pos] == '(' {
		dest, title, hasTitle, ok = p.parseInlineLink()
	}
	if !ok || (!b.isImage && hasLinkAfter(b.node)) {
		// No link, or a link inside the label would nest anchors: the
		// bracket marker becomes literal text.
		p.brackets = p.brackets[:len(p.brackets)-1]
		p.pos = closePos
		return nil
	}
	if b.isImage {
		return p.finishImageBracket(b, closePos)
	}
	return p.finishLinkBracket(b, dest, title, hasTitle)
}
func hasLinkAfter(n *inNode) bool {
	for c := n.next; c != nil; c = c.next {
		if c.kind == nLink {
			return true
		}
	}
	return false
}
func (p *inlineParser) finishLinkBracket(b *bracket, dest, title string, hasTitle bool) *inNode {
	link := p.newNode(inNode{kind: nLink, dest: dest, title: title, hasTitle: hasTitle})
	p.processDelimiters(b.bottom)
	p.replaceNode(b.node, link)
	p.adoptTrailingSiblings(link)
	p.brackets = p.brackets[:len(p.brackets)-1]
	return link
}

// finishImageBracket renders an image as its literal source text; it must
// not swallow an already-formed link, which would nest anchors.
func (p *inlineParser) finishImageBracket(b *bracket, closePos int) *inNode {
	p.brackets = p.brackets[:len(p.brackets)-1]
	if hasLinkAfter(b.node) {
		p.pos = closePos
		p.removeTrailingSiblings(b.node)
		p.pruneDelimiters(b.bottom)
		return nil
	}
	lit := p.newNode(inNode{kind: nText, seg: p.src[b.start:p.pos]})
	p.replaceNode(b.node, lit)
	p.removeTrailingSiblings(lit)
	p.pruneDelimiters(b.bottom)
	return lit
}
func (p *inlineParser) parseInlineLink() (dest, title string, hasTitle, ok bool) {
	src := p.src
	p.pos++ // consume '('
	p.skipSpaces()
	if p.pos < len(src) && src[p.pos] == ')' {
		p.pos++
		return "", "", false, true
	}
	dest, ok = p.parseLinkDestination()
	if !ok {
		return "", "", false, false
	}
	p.skipSpaces()
	if p.pos < len(src) && src[p.pos] == ')' {
		p.pos++
		return dest, "", false, true
	}
	title, hasTitle, ok = p.parseLinkTitle()
	if !ok {
		return "", "", false, false
	}
	p.skipSpaces()
	if p.pos < len(src) && src[p.pos] == ')' {
		p.pos++
		return dest, title, hasTitle, true
	}
	return "", "", false, false
}
func (p *inlineParser) skipSpaces() {
	for p.pos < len(p.src) && isSpaceByte(p.src[p.pos]) {
		p.pos++
	}
}

// parseLinkDestination parses a <...> enclosed or bare destination. Bare
// destinations stop at whitespace or an unbalanced ')'.
func (p *inlineParser) parseLinkDestination() (string, bool) {
	src := p.src
	p.skipSpaces()
	lineEnd := p.lineEnd()
	if p.pos < lineEnd && src[p.pos] == '<' {
		for i := p.pos + 1; i < lineEnd; {
			switch {
			case src[i] == '\\' && i < lineEnd-1 && isPunctByte(src[i+1]):
				i += 2
			case src[i] == '>':
				dest := src[p.pos+1 : i]
				p.pos = i + 1
				return dest, true
			default:
				i++
			}
		}
		return "", false
	}
	opened := 0
	i := p.pos
loop:
	for i < lineEnd {
		switch c := src[i]; {
		case c == '\\' && i < lineEnd-1 && isPunctByte(src[i+1]):
			i++
		case c == '(':
			opened++
		case c == ')':
			opened--
			if opened < 0 {
				break loop
			}
		case isSpaceByte(c):
			break loop
		}
		i++
	}
	dest := src[p.pos:i]
	p.pos = i
	return dest, len(dest) != 0
}

// parseLinkTitle parses an optional "..." '...' or (...) title. Titles may
// span lines; a repeated opener aborts the scan because titles do not nest.
func (p *inlineParser) parseLinkTitle() (string, bool, bool) {
	src := p.src
	p.skipSpaces()
	if p.pos >= len(src) {
		return "", false, false
	}
	opener := src[p.pos]
	var closer byte
	switch opener {
	case '"', '\'':
		closer = opener
	case '(':
		closer = ')'
	default:
		return "", false, false
	}
	for i := p.pos + 1; i < len(src); i++ {
		switch c := src[i]; {
		case c == '\\' && i < len(src)-1 && isPunctByte(src[i+1]):
			i++
		case c == closer:
			title := trimTitleLines(src[p.pos+1 : i])
			p.pos = i + 1
			return title, true, true
		case c == opener:
			return "", false, false
		}
	}
	return "", false, false
}

// trimTitleLines drops leading whitespace from every title line after the
// first, matching how goldmark assembles multi-line titles from segments.
func trimTitleLines(title string) string {
	lines := strings.Split(title, "\n")
	for i := 1; i < len(lines); i++ {
		lines[i] = strings.TrimLeft(lines[i], " \t\r")
	}
	return strings.Join(lines, "\n")
}
func (p *inlineParser) parseAutoLink(line string) *inNode {
	stop := findEmailIndex(line[1:])
	email := stop >= 0
	if !email {
		stop = findURLIndex(line[1:])
	}
	if stop < 0 {
		return nil
	}
	stop++
	if stop >= len(line) || line[stop] != '>' {
		return nil
	}
	seg := p.src[p.pos+1 : p.pos+stop]
	p.pos += stop + 1
	return p.newNode(inNode{kind: nAuto, seg: seg, email: email})
}

// parseLinkify turns bare URLs and email addresses after whitespace, a line
// start or one of *_~( into links, like goldmark's Linkify extension.
func (p *inlineParser) parseLinkify(line string) *inNode {
	if len(p.brackets) > 0 {
		return nil
	}
	consumes := 0
	if line != "" {
		switch line[0] {
		case ' ', '\t', '*', '_', '~', '(':
			consumes = 1
			line = line[1:]
		}
	}
	end, email, www, ok := linkifyTarget(line)
	if !ok {
		return nil
	}
	if consumes > 0 {
		p.appendNode(p.newNode(inNode{kind: nText, seg: p.src[p.pos : p.pos+1]}))
	}
	p.pos += consumes + end
	return p.newNode(inNode{kind: nAuto, seg: line[:end], email: email, www: www})
}

// linkifyTarget finds the end of the linkable text at the start of line and
// reports whether it is an email and whether a www. match implies http.
func linkifyTarget(line string) (end int, email, www, ok bool) {
	n := -1
	if strings.HasPrefix(line, "http:") || strings.HasPrefix(line, "https:") ||
		strings.HasPrefix(line, "ftp:") {
		n = linkifyURLEnd(line)
	} else if strings.HasPrefix(line, "www.") {
		n = linkifyWWWEnd(line)
		www = n >= 0
	}
	if n < 0 {
		end, ok := linkifyEmail(line)
		return end, true, false, ok
	}
	return trimURLTail(line, n), false, www, true
}

// linkifyURLEnd scans an http, https or ftp URL at the start of s: scheme, a
// host run ending in a dot and lowercase TLD, optional :port and optional
// /#?-introduced path. It returns the match end, or -1 when absent.
func linkifyURLEnd(s string) int {
	switch {
	case strings.HasPrefix(s, "http://"):
		return linkifyHostEnd(s, 7, true)
	case strings.HasPrefix(s, "https://"):
		return linkifyHostEnd(s, 8, true)
	case strings.HasPrefix(s, "ftp://"):
		return linkifyHostEnd(s, 6, true)
	}
	return -1
}
func linkifyWWWEnd(s string) int {
	if !strings.HasPrefix(s, "www.") {
		return -1
	}
	return linkifyHostEnd(s, 4, false)
}

// linkifyHostEnd scans the host portion shared by URL and www matches from
// s[i:]: a host run, its last dot followed by a lowercase TLD, then an
// optional :port (URLs only) and an optional /#? path.
func linkifyHostEnd(s string, i int, url bool) int {
	j := i
	for j < len(s) && urlHostByte(s[j]) {
		j++
	}
	// The run ends at the dot, whose position may sit right at the 256-byte
	// cap: the dot itself is consumed after the run, not as part of it.
	limit := min(i+256, j-1)
	dot := -1
	for k := limit; k > i; k-- {
		if s[k] == '.' && k+1 < len(s) && s[k+1] >= 'a' && s[k+1] <= 'z' {
			dot = k
			break
		}
	}
	if dot < 0 {
		return -1
	}
	i = dot + 1
	for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
		i++
	}
	if url && i < len(s) && s[i] == ':' {
		k := i + 1
		for k < len(s) && isDigitByte(s[k]) {
			k++
		}
		if k > i+1 {
			i = k
		}
	}
	if i < len(s) && (s[i] == '/' || s[i] == '#' || s[i] == '?') {
		for i++; i < len(s) && linkifyTailByte(s[i], url); i++ {
		}
	}
	return i
}

// urlHostByte reports the bytes allowed in the host run before the TLD.
func urlHostByte(c byte) bool {
	switch {
	case isAlphaNumeric(c):
		return true
	case c == '-' || c == '@' || c == ':' || c == '.' || c == '_' ||
		c == '+' || c == '~' || c == '#' || c == '%' || c == '=':
		return true
	}
	return false
}

// linkifyTailByte reports the bytes allowed inside a URL or www path. The
// two patterns differ in exactly one byte: '$' continues only URL paths.
func linkifyTailByte(c byte, url bool) bool {
	switch {
	case isAlphaNumeric(c):
		return true
	case c == '$':
		return url
	case c == '-' || c == '@' || c == ':' || c == '%' || c == '_' || c == '+' || c == '.' ||
		c == '~' || c == '#' || c == '!' || c == '?' || c == '&' || c == '/' ||
		c == '=' || c == '(' || c == ')' || c == ';' || c == ',' || c == '\'' || c == '"' ||
		c == '>' || c == '^' || c == '{' || c == '}' || c == '[' || c == ']' || c == '`':
		return true
	}
	return false
}

// linkifyEmail scans an email address at the start of line, rejecting
// addresses whose domain would end in '-' or '_'.
func linkifyEmail(line string) (end int, ok bool) {
	if line != "" && isPunctByte(line[0]) {
		return 0, false
	}
	stop := findEmailIndex(line)
	if stop < 0 {
		return 0, false
	}
	stop, ok = linkifyEmailDomain(line, stop)
	if !ok {
		return 0, false
	}
	return trimTrailingPunct(line, stop), true
}

// linkifyEmailDomain checks the '@' and the dots around it, and trims one
// trailing dot from the domain.
func linkifyEmailDomain(line string, stop int) (int, bool) {
	at := strings.IndexByte(line, '@')
	if at < 0 || at >= stop || strings.IndexByte(line[at:stop-1], '.') < 0 {
		return 0, false
	}
	if line[stop-1] == '.' {
		stop--
	}
	if stop < len(line) && (line[stop] == '-' || line[stop] == '_') {
		return 0, false
	}
	return stop, true
}
func trimURLTail(line string, end int) int {
	switch line[end-1] {
	case '.':
		end--
	case ')':
		closing := 0
		for i := end - 1; i >= 0; i-- {
			switch line[i] {
			case ')':
				closing++
			case '(':
				closing--
			}
		}
		if closing > 0 {
			end -= closing
		}
	case ';':
		i := end - 2
		for i >= 0 && isAlphaNumeric(line[i]) {
			i--
		}
		if i != end-2 && i >= 0 && line[i] == '&' {
			end -= end - i
		}
	}
	return trimTrailingPunct(line, end)
}
func trimTrailingPunct(line string, end int) int {
	for end > 1 {
		switch line[end-1] {
		case '?', '!', '.', ',', ':', '*', '_', '~':
			end--
			continue
		}
		break
	}
	return end
}

// findURLIndex reports whether b starts with something shaped like
// [A-Za-z][A-Za-z0-9.+-]{1,31}:[^ \t\n\v\f\r<>]* and returns the stop index.
func findURLIndex(b string) int {
	if len(b) == 0 || !isAlpha(b[0]) {
		return -1
	}
	i := 1
	for ; i < len(b) && urlSchemeByte(b[i]); i++ {
	}
	if i == 1 || i > 33 || i >= len(b) || b[i] != ':' {
		return -1
	}
	i++
	for ; i < len(b) && urlRestByte(b[i]); i++ {
	}
	return i
}
func urlSchemeByte(c byte) bool {
	return isAlphaNumeric(c) || c == '+' || c == '-' || c == '.'
}
func urlRestByte(c byte) bool {
	return c > ' ' && c != '<' && c != '>'
}
func findEmailIndex(b string) int {
	i := 0
	for ; i < len(b) && emailLocalByte(b[i]); i++ {
	}
	if i == 0 || i >= len(b) || b[i] != '@' {
		return -1
	}
	return linkifyDomainEnd(b, i+1)
}

// linkifyDomainEnd scans a dot-separated domain at s[i:]: labels of one to 63
// characters that neither start nor end with a hyphen. It returns the offset
// after the last complete label, or -1 when the first byte is not a label.
func linkifyDomainEnd(s string, i int) int {
	end := -1
	for i < len(s) && isAlphaNumeric(s[i]) {
		run := i + 1
		for run < len(s) && (isAlphaNumeric(s[run]) || s[run] == '-') {
			run++
		}
		// The label may extend up to 61 further characters, but its last
		// character must be alphanumeric, so back off over trailing hyphens.
		k := min(run-i-2, 61)
		for k >= 0 && !isAlphaNumeric(s[i+1+k]) {
			k--
		}
		i++
		if k >= 0 {
			i += k + 1
		}
		end = i
		if i >= len(s) || s[i] != '.' {
			return end
		}
		i++
	}
	return end
}
func emailLocalByte(c byte) bool {
	return emailLocalBytes[c]
}

// emailLocalBytes marks the bytes allowed in an email local part; linkify
// tests every word against it, so the class is a table lookup.
var emailLocalBytes = func() (set [256]bool) {
	for c := range 256 {
		set[c] = isAlphaNumeric(byte(c)) || strings.IndexByte("!#$%&'*+-/=?^_`{|}~.", byte(c)) >= 0
	}
	return set
}()
