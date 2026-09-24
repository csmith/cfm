package cfm

import "strings"

type markdownTable struct {
	header []string
	align  []string
	rows   [][]string
}

// Callers supply a single paragraph's lines, with interrupting blocks removed.
// The returned count includes the header, delimiter, and body rows.
func parseTable(lines []string) (*markdownTable, int) {
	if len(lines) < 2 {
		return nil, 0
	}
	align := parseTableDelimiter(lines[1])
	if align == nil {
		return nil, 0
	}
	header := splitTableRow(lines[0])
	if len(header) > len(align) {
		return nil, 0
	}
	t := &markdownTable{header: fitTableRow(header, len(align)), align: align}
	for _, line := range lines[2:] {
		if strings.Trim(line, tableSpace) == "" {
			break
		}
		t.rows = append(t.rows, fitTableRow(splitTableRow(line), len(align)))
	}
	return t, 2 + len(t.rows)
}

func renderTable(t *markdownTable, inline func(string) string) string {
	if t == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString("<table>\n<thead>\n<tr>\n")
	for i, cell := range t.header {
		writeTableCell(&b, "th", cell, tableAlignAttr(t.align, i), inline)
	}
	b.WriteString("</tr>\n</thead>\n")
	if len(t.rows) > 0 {
		b.WriteString("<tbody>\n")
		for _, row := range t.rows {
			b.WriteString("<tr>\n")
			for i, cell := range row {
				writeTableCell(&b, "td", cell, tableAlignAttr(t.align, i), inline)
			}
			b.WriteString("</tr>\n")
		}
		b.WriteString("</tbody>\n")
	}
	b.WriteString("</table>\n")
	return b.String()
}

const tableSpace = " \t\r"

func writeTableCell(b *strings.Builder, tag, cell, attr string, inline func(string) string) {
	b.WriteByte('<')
	b.WriteString(tag)
	b.WriteString(attr)
	b.WriteByte('>')
	b.WriteString(inline(cell))
	b.WriteString("</")
	b.WriteString(tag)
	b.WriteString(">\n")
}

func tableAlignAttr(align []string, i int) string {
	if i >= len(align) || align[i] == "" {
		return ""
	}
	return ` style="text-align:` + align[i] + `"`
}

func fitTableRow(cells []string, n int) []string {
	out := make([]string, n)
	copy(out, cells)
	return out
}

// Goldmark suppresses pipe splitting after any backslash, even an escaped one.
func splitTableRow(line string) []string {
	line = strings.Trim(line, tableSpace)
	pos, limit := 0, len(line)
	if pos < limit && line[pos] == '|' {
		pos++
	}
	if limit > 0 && line[limit-1] == '|' {
		limit--
	}
	var cells []string
	for pos < limit {
		closure := pos
		for closure < limit && (line[closure] != '|' || closure > 0 && line[closure-1] == '\\') {
			closure++
		}
		cells = append(cells, strings.Trim(line[pos:closure], tableSpace))
		pos = closure + 1
	}
	return cells
}

func parseTableDelimiter(line string) []string {
	if !isTableDelimiter(line) {
		return nil
	}
	cols := strings.Split(line, "|")
	if strings.Trim(cols[0], tableSpace) == "" {
		cols = cols[1:]
	}
	if len(cols) > 0 && strings.Trim(cols[len(cols)-1], tableSpace) == "" {
		cols = cols[:len(cols)-1]
	}
	align := make([]string, 0, len(cols))
	for _, col := range cols {
		a, ok := tableColumnAlign(strings.Trim(col, tableSpace))
		if !ok {
			return nil
		}
		align = append(align, a)
	}
	return align
}

func isTableDelimiter(line string) bool {
	if tableIndentWidth(line) > 3 {
		return false
	}
	hasDash := false
	for i := range len(line) {
		switch line[i] {
		case '-':
			hasDash = true
		case ' ', '\t', '\r', '|', ':':
		default:
			return false
		}
	}
	return hasDash
}

func tableIndentWidth(line string) int {
	width := 0
	for i := range len(line) {
		switch line[i] {
		case ' ':
			width++
		case '\t':
			width += 4 - width%4
		default:
			return width
		}
	}
	return width
}

func tableColumnAlign(col string) (string, bool) {
	if col == "" {
		return "", false
	}
	left := col[0] == ':'
	right := col[len(col)-1] == ':'
	body := col
	if left {
		body = body[1:]
	}
	if right && len(body) > 0 {
		body = body[:len(body)-1]
	}
	if body == "" || strings.Trim(body, "-") != "" {
		return "", false
	}
	switch {
	case left && right:
		return "center", true
	case left:
		return "left", true
	case right:
		return "right", true
	default:
		return "", true
	}
}
