package cfm

import "strings"

func (p *inlineParser) parseRawHTML() *inNode {
	line := p.src[p.pos:]
	switch {
	case len(line) > 1 && isAlpha(line[1]):
		return p.rawHTMLTag(false)
	case len(line) > 2 && line[1] == '/' && isAlpha(line[2]):
		return p.rawHTMLTag(true)
	case strings.HasPrefix(line, "<!--"):
		return p.parseHTMLComment()
	case strings.HasPrefix(line, "<?"):
		return p.parseHTMLUntil("?>")
	case len(line) > 2 && line[1] == '!' && line[2] >= 'A' && line[2] <= 'Z':
		return p.parseHTMLUntil(">")
	case strings.HasPrefix(line, "<![CDATA["):
		return p.parseHTMLUntil("]]>")
	}
	return nil
}

// rawHTMLTag consumes a raw HTML open or closing tag at p.pos and returns
// it as a verbatim node.
func (p *inlineParser) rawHTMLTag(closing bool) *inNode {
	n := rawTagLen(p.src[p.pos:], closing)
	if n == 0 {
		return nil
	}
	node := p.newNode(inNode{kind: nRaw, seg: p.src[p.pos : p.pos+n]})
	p.pos += n
	return node
}

// The dispatcher has checked the leading '<' and any closing '/'.
// Return the complete tag's byte length, or zero if invalid.
func rawTagLen(s string, closing bool) int {
	if closing {
		return rawCloseTagLen(s)
	}
	return rawOpenTagLen(s)
}
func rawOpenTagLen(s string) int {
	i := 1
	if i >= len(s) || !isAlpha(s[i]) {
		return 0
	}
	for i++; i < len(s) && (isAlphaNumeric(s[i]) || s[i] == '-'); i++ {
	}
	for {
		next, more, valid := scanTagAttr(s, i)
		if !valid {
			// No value follows an '='; no attribute or terminator can match.
			return 0
		}
		if !more {
			break
		}
		i = next
	}
	return tagCloserLen(s, i, true)
}
func rawCloseTagLen(s string) int {
	if len(s) < 3 || !isAlpha(s[2]) {
		return 0
	}
	i := 2
	for i++; i < len(s) && (isAlphaNumeric(s[i]) || s[i] == '-'); i++ {
	}
	// Closing tags carry only whitespace: no attributes, no slash.
	return tagCloserLen(s, i, false)
}

// scanTagAttr scans one whitespace-separated attribute with an optional value
// at s[i:]. It reports the offset after the attribute, whether an attribute
// was found, and whether the tag can still be valid.
func scanTagAttr(s string, i int) (next int, more, valid bool) {
	j := i
	for j < len(s) && isSpaceByte(s[j]) {
		j++
	}
	if j == i || j >= len(s) || !attrNameStart(s[j]) {
		// Attributes need at least one separating whitespace, and the run
		// left behind belongs to the tag terminator, which accepts a
		// narrower set than separators.
		return i, false, true
	}
	for i = j + 1; i < len(s) && attrNameByte(s[i]); i++ {
	}
	k := i
	for k < len(s) && isSpaceByte(s[k]) {
		k++
	}
	if k < len(s) && s[k] == '=' {
		for k++; k < len(s) && isSpaceByte(s[k]); k++ {
		}
		v, ok := attrValueLen(s, k)
		if !ok {
			return i, false, false
		}
		i = k + v
	}
	return i, true, true
}

// tagCloserLen completes a tag at s[i:]: whitespace (a bare carriage return
// is not tag whitespace), an optional slash on open tags, and '>'.
func tagCloserLen(s string, i int, allowSlash bool) int {
	for i < len(s) {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			i++
		} else if s[i] == '\r' && i+1 < len(s) && s[i+1] == '\n' {
			i += 2
		} else {
			break
		}
	}
	if allowSlash && i < len(s) && s[i] == '/' {
		i++
	}
	if i < len(s) && s[i] == '>' {
		return i + 1
	}
	return 0
}
func attrNameStart(c byte) bool {
	return isAlpha(c) || c == '_' || c == ':'
}
func attrNameByte(c byte) bool {
	return isAlphaNumeric(c) || c == ':' || c == '.' || c == '_' || c == '-'
}

// attrValueLen scans an attribute value at s[i]: 'single quoted', "double
// quoted", or an unquoted run excluding whitespace, quotes and =<>`.
func attrValueLen(s string, i int) (int, bool) {
	if i >= len(s) {
		return 0, false
	}
	if q := s[i]; q == '\'' || q == '"' {
		for j := i + 1; j < len(s); j++ {
			if s[j] == q {
				return j + 1 - i, true
			}
		}
		return 0, false
	}
	if attrValueStop(s[i]) {
		return 0, false
	}
	start := i
	for i++; i < len(s) && !attrValueStop(s[i]); i++ {
	}
	return i - start, true
}
func attrValueStop(c byte) bool {
	return c <= 0x20 || c == '"' || c == '\'' || c == '=' || c == '<' || c == '>' || c == '`'
}
func (p *inlineParser) parseHTMLComment() *inNode {
	src := p.src
	pos := p.pos
	for _, empty := range []string{"<!-->", "<!--->"} {
		if strings.HasPrefix(src[pos:], empty) {
			p.pos = pos + len(empty)
			return p.newNode(inNode{kind: nRaw, seg: src[pos:p.pos]})
		}
	}
	idx := strings.Index(src[pos+4:], "-->")
	if idx < 0 {
		return nil
	}
	end := pos + 4 + idx + 3
	p.pos = end
	return p.newNode(inNode{kind: nRaw, seg: src[pos:end]})
}
func (p *inlineParser) parseHTMLUntil(closer string) *inNode {
	idx := strings.Index(p.src[p.pos:], closer)
	if idx < 0 {
		return nil
	}
	end := p.pos + idx + len(closer)
	seg := p.src[p.pos:end]
	p.pos = end
	return p.newNode(inNode{kind: nRaw, seg: seg})
}
