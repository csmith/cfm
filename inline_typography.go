package cfm

import (
	"unicode"
	"unicode/utf8"
)

// parseTypographer performs smart punctuation substitutions; quotes use the
// delimiter flanking rules plus goldmark's special cases and counters.
func (p *inlineParser) parseTypographer(line string) *inNode {
	c := line[0]
	switch {
	case len(line) > 2 && c == '-' && line[1] == '-' && line[2] == '-':
		p.pos += 3
		return p.newNode(inNode{kind: nLit, str: "&mdash;"})
	case len(line) > 2 && c == '.' && line[1] == '.' && line[2] == '.':
		p.pos += 3
		return p.newNode(inNode{kind: nLit, str: "&hellip;"})
	case len(line) > 2 && c == '.':
		return nil
	case len(line) > 1 && c == '<' && line[1] == '<':
		p.pos += 2
		return p.newNode(inNode{kind: nLit, str: "&laquo;"})
	case len(line) > 1 && c == '<':
		return nil
	case len(line) > 1 && c == '>' && line[1] == '>':
		p.pos += 2
		return p.newNode(inNode{kind: nLit, str: "&raquo;"})
	case len(line) > 1 && c == '>':
		return nil
	case len(line) > 1 && c == '-' && line[1] == '-':
		p.pos += 2
		return p.newNode(inNode{kind: nLit, str: "&ndash;"})
	case c == '\'':
		return p.typographerSingleQuote(line)
	case c == '"':
		return p.typographerDoubleQuote(line)
	}
	return nil
}
func (p *inlineParser) typographerSingleQuote(line string) *inNode {
	before := p.precedingRune()
	d := p.quoteDelimiter(line, before)
	if d == nil {
		return nil
	}
	if n := p.singleQuoteContraction(line, before, d); n != nil {
		return n
	}
	if d.canOpen && !d.canClose {
		return p.singleQuoteOpening(line)
	}
	return p.singleQuoteClosing(line, d)
}

// singleQuoteContraction handles apostrophes that can never open a
// quotation: decades, elisions and letters around the quote.
func (p *inlineParser) singleQuoteContraction(line string, before rune, d *delimiter) *inNode {
	if n := p.singleQuoteDecade(line, d); n != nil {
		return n
	}
	// Special cases: 'twas, 'em, 'net, 'll.
	if len(line) > 1 && (unicode.IsPunct(before) || unicode.IsSpace(before)) &&
		(line[1] == 't' || line[1] == 'e' || line[1] == 'n' || line[1] == 'l') {
		return p.quote("&rsquo;", 0)
	}
	// Apostrophes between alphanumerics.
	if len(line) > 1 && (unicode.IsDigit(before) || unicode.IsLetter(before)) {
		if r, _ := utf8.DecodeRuneInString(line[1:]); unicode.IsLetter(r) {
			return p.quote("&rsquo;", 0)
		}
	}
	return nil
}

// singleQuoteDecade matches decade abbreviations such as '90s.
func (p *inlineParser) singleQuoteDecade(line string, d *delimiter) *inNode {
	if !d.canOpen || d.canClose || len(line) <= 3 ||
		!isDigitByte(line[1]) || !isDigitByte(line[2]) || line[3] != 's' {
		return nil
	}
	after := rune(' ')
	if len(line) > 4 {
		after, _ = utf8.DecodeRuneInString(line[4:])
	}
	if isSpaceRune(after) || isPunctRune(after) {
		return p.quote("&rsquo;", 0)
	}
	return nil
}

// singleQuoteOpening renders a single quote that opens a quotation, using
// the apostrophe form for common contraction suffixes.
func (p *inlineParser) singleQuoteOpening(line string) *inNode {
	nt := "&lsquo;"
	// Special cases: Alice's, I'm, Don't, You'd.
	if len(line) > 1 && (line[1] == 's' || line[1] == 'm' || line[1] == 't' || line[1] == 'd') &&
		(len(line) < 3 || isPunctByte(line[2]) || isSpaceByte(line[2])) {
		nt = "&rsquo;"
	}
	// Special cases: I've, I'll, You're.
	if len(line) > 2 && ((line[1] == 'v' && line[2] == 'e') ||
		(line[1] == 'l' && line[2] == 'l') || (line[1] == 'r' && line[2] == 'e')) &&
		(len(line) < 4 || isPunctByte(line[3]) || isSpaceByte(line[3])) {
		nt = "&rsquo;"
	}
	return p.quote(nt, p.delta(nt))
}

// singleQuoteClosing renders a closing single quote: a plural possessive or
// abbreviation, or a quote that closes a still-open quotation.
func (p *inlineParser) singleQuoteClosing(line string, d *delimiter) *inNode {
	// Plural possessives and abbreviations: Smiths', doin'.
	if len(line) > 1 && unicode.IsSpace(firstRune(line)) ||
		unicode.IsPunct(firstRune(line)) && (len(line) > 2 && !unicode.IsDigit(firstRune(line[1:]))) {
		return p.quote("&rsquo;", 0)
	}
	if p.quoteSingle <= 0 {
		return nil
	}
	isClose := d.canClose && !d.canOpen
	maybeClose := d.canClose && d.canOpen && len(line) > 1 && unicode.IsPunct(firstRune(line[1:])) &&
		(len(line) == 2 || len(line) > 2 && (isPunctByte(line[2]) || isSpaceByte(line[2])))
	if isClose || maybeClose {
		return p.quote("&rsquo;", -1)
	}
	return nil
}
func (p *inlineParser) typographerDoubleQuote(line string) *inNode {
	before := p.precedingRune()
	d := p.quoteDelimiter(line, before)
	if d == nil {
		return nil
	}
	if d.canOpen && !d.canClose {
		return p.quote("&ldquo;", 1)
	}
	if p.quoteDouble > 0 {
		isClose := d.canClose && !d.canOpen
		maybeClose := d.canClose && d.canOpen && len(line) > 1 && unicode.IsPunct(firstRune(line[1:])) &&
			(len(line) == 2 || len(line) > 2 && (isPunctByte(line[2]) || isSpaceByte(line[2])))
		if isClose || maybeClose {
			// Special case: "Monitor 21""
			if len(line) > 1 && line[1] == '"' && unicode.IsDigit(before) {
				return nil
			}
			return p.quote("&rdquo;", -1)
		}
	}
	return nil
}
func (p *inlineParser) quote(entity string, delta int) *inNode {
	if entity == "&lsquo;" || entity == "&rsquo;" {
		p.quoteSingle += delta
	} else {
		p.quoteDouble += delta
	}
	p.pos++
	return p.newNode(inNode{kind: nLit, str: entity})
}
func (p *inlineParser) delta(nt string) int {
	if nt == "&lsquo;" {
		return 1
	}
	return 0
}
func (p *inlineParser) quoteDelimiter(line string, before rune) *delimiter {
	run := 1
	after := rune(' ')
	if run < len(line) {
		after, _ = utf8.DecodeRuneInString(line[1:])
	}
	canOpen, canClose := flank(before, after, line[0])
	return p.newDelimiter(delimiter{canOpen: canOpen, canClose: canClose})
}
func firstRune(s string) rune {
	r, _ := utf8.DecodeRuneInString(s)
	return r
}
