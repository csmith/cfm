package cfm

import "strings"

// Markdown converts Markdown source to HTML with the default renderer.
func Markdown(source string) string {
	return Renderer{}.Markdown(source)
}

// Highlighter renders the body of a fenced code block as complete, trusted
// HTML. It receives the verbatim code and the block's language. When it
// reports an error, the renderer falls back to an escaped plain code block.
type Highlighter func(code, language string) (string, error)

// Renderer converts Markdown source to HTML. The zero value is ready to use.
type Renderer struct {
	// Highlight renders fenced code blocks. When nil, or when it reports an
	// error, fenced code falls back to escaped plain code.
	Highlight Highlighter
}

// Markdown converts Markdown source to HTML.
func (r Renderer) Markdown(source string) string {
	source = strings.ReplaceAll(source, "\r\n", "\n")
	source = strings.ReplaceAll(source, "\r", "\n")
	source = strings.ReplaceAll(source, "\x00", "\uFFFD")
	dc := newDocContext()
	p := newBlockParser(dc)
	p.parse(source)
	prepassInline(p.root, dc, dc.inlineFunc())
	body := p.render(r)
	return body + dc.notes.render(func(label string) string { return dc.noteHTML(label, r) })
}

func escapeHTML(s string) string {
	if !strings.ContainsAny(s, `&<>"`) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))
	for i := range len(s) {
		switch s[i] {
		case '"':
			b.WriteString("&quot;")
		case '&':
			b.WriteString("&amp;")
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		default:
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
