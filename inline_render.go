package cfm

import (
	"html"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"
)

func renderNodes(b *strings.Builder, n *inNode) {
	switch n.kind {
	case nText:
		writeText(b, n.seg)
		if n.hard {
			b.WriteString("<br>\n")
		} else if n.soft {
			b.WriteByte('\n')
		}
	case nRaw:
		b.WriteString(n.seg)
	case nLit:
		b.WriteString(n.str)
	case nCode:
		b.WriteString("<code>")
		rawEscape(b, strings.ReplaceAll(n.str, "\n", " "))
		b.WriteString("</code>")
	case nEmph, nStrike:
		writeSpan(b, n)
	case nLink:
		b.WriteString(`<a href="`)
		rawEscape(b, linkDestinationAttr(n.dest))
		b.WriteByte('"')
		if n.hasTitle {
			b.WriteString(` title="`)
			writeText(b, n.title)
			b.WriteByte('"')
		}
		b.WriteByte('>')
		renderChildren(b, n)
		b.WriteString("</a>")
	case nAuto:
		renderAutoLink(b, n)
	default:
		writeText(b, n.seg)
		renderChildren(b, n)
	}
}
func writeSpan(b *strings.Builder, n *inNode) {
	open, closing := "<em>", "</em>"
	switch {
	case n.kind == nEmph && n.level == 2:
		open, closing = "<strong>", "</strong>"
	case n.kind == nStrike:
		open, closing = "<del>", "</del>"
	}
	b.WriteString(open)
	renderChildren(b, n)
	b.WriteString(closing)
}
func renderChildren(b *strings.Builder, n *inNode) {
	for c := n.firstChild; c != nil; c = c.next {
		renderNodes(b, c)
	}
}
func renderAutoLink(b *strings.Builder, n *inNode) {
	href := urlEscapePass(n.seg)
	switch {
	case n.email && !mailtoPrefix(href):
		href = "mailto:" + href
	case n.www:
		href = "http://" + href
	}
	b.WriteString(`<a href="`)
	rawEscape(b, href)
	b.WriteString(`">`)
	rawEscape(b, n.seg)
	b.WriteString("</a>")
}

// mailtoPrefix reports whether s starts with a case-insensitive "mailto:",
// avoiding an allocation-prone lowercase copy of the whole address.
func mailtoPrefix(s string) bool {
	return len(s) >= 7 && strings.EqualFold(s[:7], "mailto:")
}

// linkDestinationAttr resolves escapes and character references in a link
// destination, then percent-escapes it for an href attribute.
func linkDestinationAttr(dest string) string {
	dest = unescapePunctuations(dest)
	dest = resolveNumericRefs(dest)
	dest = resolveEntityNames(dest)
	return urlEscapePass(dest)
}

// writeText renders plain text: it unescapes backslash-escaped punctuation,
// resolves character references, and HTML-escapes what remains. All three
// happen in a single pass: pending segments between special bytes are kept
// escape-free by stopping at every escapable byte, so they are written as-is.
func writeText(b *strings.Builder, s string) {
	n := 0
	i := 0
	escaped := false
	for i < len(s) {
		if !escaped {
			// Plain runs cannot contain anything the switch below reacts to.
			j := i
			for j < len(s) && !textBreakByte[s[j]] {
				j++
			}
			if j == len(s) {
				break
			}
			i = j
		}
		c := s[i]
		switch {
		case escaped && isPunctByte(c):
			b.WriteString(s[n : i-1])
			if esc := escapeHTMLByte(c); esc != "" {
				b.WriteString(esc)
			} else {
				b.WriteByte(c)
			}
			n = i + 1
			escaped = false
			i++
		case c == 0:
			b.WriteString(s[n:i])
			b.WriteString("\uFFFD")
			n = i + 1
			escaped = false
			i++
		case c == '&':
			if end, chars, preEscaped, ok := resolveCharReference(s, i); ok {
				b.WriteString(s[n:i])
				if preEscaped {
					b.WriteString(chars)
				} else {
					rawEscape(b, chars)
				}
				n = end
				i = end
				escaped = false
				continue
			}
			b.WriteString(s[n:i])
			b.WriteString("&amp;")
			n = i + 1
			i++
		case c == '<':
			b.WriteString(s[n:i])
			b.WriteString("&lt;")
			n = i + 1
			i++
		case c == '>':
			b.WriteString(s[n:i])
			b.WriteString("&gt;")
			n = i + 1
			i++
		case c == '"':
			b.WriteString(s[n:i])
			b.WriteString("&quot;")
			n = i + 1
			i++
		case c == '\\' && !escaped:
			escaped = true
			i++
		default:
			escaped = false
			i++
		}
	}
	b.WriteString(s[n:])
}

// textBreakByte marks bytes writeText must stop at: backslash escapes,
// character references and the bytes raw HTML escaping replaces.
var textBreakByte = func() (brk [256]bool) {
	for c := range 256 {
		brk[c] = c == '\\' || escapeHTMLByte(byte(c)) != ""
	}
	return brk
}()

// resolveCharReference resolves the character reference at s[i] == '&',
// reporting whether the replacement is already HTML-escaped.
func resolveCharReference(s string, i int) (int, string, bool, bool) {
	next := i + 1
	if next >= len(s) {
		return 0, "", false, false
	}
	if s[next] == '#' {
		v, end, ok := numericCharReference(s, next+1)
		if ok {
			return end, escapedRune(v), true, true
		}
		return 0, "", false, false
	}
	return namedCharReference(s, i, next)
}

// numericCharReference parses the digits of a numeric reference starting
// just after the '#'; up to six hexadecimal or seven decimal digits.
func numericCharReference(s string, start int) (rune, int, bool) {
	base, maxDigits := 10, 7
	digit := isDigitByte
	if start < len(s) {
		if c := s[start]; c == 'x' || c == 'X' {
			base, maxDigits, digit = 16, 6, isHexDigit
			start++
		}
	}
	j := start
	for j < len(s) && digit(s[j]) {
		j++
	}
	if j == start || j >= len(s) || s[j] != ';' || j-start > maxDigits {
		return 0, 0, false
	}
	v, _ := strconv.ParseUint(s[start:j], base, 32)
	return rune(v), j + 1, true
}

// namedCharReference resolves an HTML5 entity name beginning at s[i] == '&'.
func namedCharReference(s string, i, next int) (int, string, bool, bool) {
	j := next
	for j < len(s) && isAlphaNumeric(s[j]) {
		j++
	}
	if j == next || j >= len(s) || s[j] != ';' {
		return 0, "", false, false
	}
	name := s[i : j+1]
	chars, ok := decodeNamedEntity(name)
	if !ok {
		return 0, "", false, false
	}
	return j + 1, chars, false, true
}
func escapedRune(r rune) string {
	if r < 256 {
		if esc := escapeHTMLByte(byte(r)); esc != "" {
			return esc
		}
	}
	return string(toValidRune(r))
}
func toValidRune(r rune) rune {
	if r <= 0 || r >= 0xD800 && r < 0xE000 || r > utf8.MaxRune {
		return '\uFFFD'
	}
	return r
}
func rawEscape(b *strings.Builder, s string) {
	n := 0
	for i := range len(s) {
		if !htmlEscapeByte[s[i]] {
			continue
		}
		b.WriteString(s[n:i])
		b.WriteString(escapeHTMLByte(s[i]))
		n = i + 1
	}
	b.WriteString(s[n:])
}

// htmlEscapeByte marks the bytes rawEscape must replace with an entity.
var htmlEscapeByte = func() (esc [256]bool) {
	for c := range 256 {
		esc[c] = escapeHTMLByte(byte(c)) != ""
	}
	return esc
}()

func escapeHTMLByte(c byte) string {
	switch c {
	case '"':
		return "&quot;"
	case '&':
		return "&amp;"
	case '<':
		return "&lt;"
	case '>':
		return "&gt;"
	case 0:
		return "\uFFFD"
	}
	return ""
}
func unescapePunctuations(s string) string {
	var b strings.Builder
	n := 0
	for i := 0; i < len(s); {
		if i < len(s)-1 && s[i] == '\\' && isPunctByte(s[i+1]) {
			b.WriteString(s[n:i])
			b.WriteByte(s[i+1])
			i += 2
			n = i
			continue
		}
		i++
	}
	if n == 0 {
		return s
	}
	b.WriteString(s[n:])
	return b.String()
}

// resolveNumericRefs resolves numeric character references in a link
// destination.
func resolveNumericRefs(s string) string {
	var b strings.Builder
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '&' || i+1 >= len(s) || s[i+1] != '#' {
			continue
		}
		if v, end, ok := numericRefValue(s, i); ok {
			b.WriteString(s[n:i])
			b.WriteRune(toValidRune(v))
			i = end - 1
			n = end
		}
	}
	if n == 0 {
		return s
	}
	b.WriteString(s[n:])
	return b.String()
}
func numericRefValue(s string, i int) (rune, int, bool) {
	next := i + 1
	if next+1 >= len(s) || s[i] != '&' || s[next] != '#' {
		return 0, 0, false
	}
	if nc := s[next+1]; nc == 'x' || nc == 'X' {
		start := next + 2
		j := start
		for j < len(s) && isHexDigit(s[j]) {
			j++
		}
		if j > start && j-start <= 6 && j < len(s) && s[j] == ';' {
			v, _ := strconv.ParseUint(s[start:j], 16, 32)
			return rune(v), j + 1, true
		}
		return 0, 0, false
	}
	start := next + 1
	j := start
	for j < len(s) && s[j] >= '0' && s[j] <= '9' {
		j++
	}
	if j > start && j < len(s) && s[j] == ';' && j-start < 8 {
		v, _ := strconv.ParseUint(s[start:j], 10, 32)
		return rune(v), j + 1, true
	}
	return 0, 0, false
}

// resolveEntityNames resolves named character references (HTML5 names).
func resolveEntityNames(s string) string {
	var b strings.Builder
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '&' {
			continue
		}
		j := i + 1
		for j < len(s) && isAlphaNumeric(s[j]) {
			j++
		}
		if j == i+1 || j >= len(s) || s[j] != ';' {
			continue
		}
		name := s[i : j+1]
		chars, ok := decodeNamedEntity(name)
		if !ok {
			continue
		}
		b.WriteString(s[n:i])
		b.WriteString(chars)
		i = j
		n = j + 1
	}
	if n == 0 {
		return s
	}
	b.WriteString(s[n:])
	return b.String()
}

// Legacy-prefix decoding leaves a semicolon suffix. Only &semi; is a valid
// named reference whose decoded value contains a semicolon.
func decodeNamedEntity(name string) (string, bool) {
	chars := html.UnescapeString(name)
	if chars == name || name != "&semi;" && strings.HasSuffix(chars, ";") {
		return "", false
	}
	return chars, true
}

// urlEscapePass percent-escapes a URL: a fixed safe
// set passes through, existing %XX escapes survive, spaces become %20.
func urlEscapePass(s string) string {
	var b strings.Builder
	n := 0
	limit := len(s)
	for i := 0; i < limit; {
		c := s[i]
		if urlSafeByte(c) {
			i++
			continue
		}
		if c == '%' && i+2 < limit && isHexDigit(s[i+1]) && isHexDigit(s[i+2]) {
			i += 3
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size <= 1 {
			i++
			continue
		}
		b.WriteString(s[n:i])
		if r == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteString(url.QueryEscape(string(r)))
		}
		i += size
		n = i
	}
	if n < limit {
		b.WriteString(s[n:])
	}
	return b.String()
}
func urlSafeByte(c byte) bool {
	if isAlphaNumeric(c) {
		return true
	}
	switch c {
	case '!', '#', '$', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '=', '?', '@', '_', '~':
		return true
	}
	return false
}
