package cfm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableParse(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		want      *markdownTable
		wantLines int
	}{
		{
			name:  "no lines",
			lines: nil,
		},
		{
			name:  "header only",
			lines: []string{"| just | one | row |"},
		},
		{
			name:  "header count exceeds delimiter count",
			lines: []string{"no | table | here", ":--- | :---"},
		},
		{
			name:  "header count exceeds delimiter count via pipes",
			lines: []string{"| a | b | c |", "|---|---|"},
		},
		{
			name:  "unescaped pipe in code span splits header cells",
			lines: []string{"| `a|b` |", "|---|"},
		},
		{
			name:  "delimiter with junk column",
			lines: []string{"| a | b |", "| :-x | --- |"},
		},
		{
			name:  "delimiter with only colons",
			lines: []string{"| a |", "|:::|"},
		},
		{
			name:  "delimiter with backslash",
			lines: []string{"| a |", "|-\\-|"},
		},
		{
			name:  "delimiter without hyphen",
			lines: []string{"| a |", "|||"},
		},
		{
			name:  "delimiter indented four spaces",
			lines: []string{"| a |", "    ---"},
		},
		{
			name:  "delimiter indented with tab",
			lines: []string{"| a |", "\t---"},
		},
		{
			name: "basic alignment",
			lines: []string{
				"| Left | Center | Right |",
				"|:-----|:------:|------:|",
				"| a | b | c |",
				"| wider cell | m | z |",
			},
			want: &markdownTable{
				header: []string{"Left", "Center", "Right"},
				align:  []string{"left", "center", "right"},
				rows:   [][]string{{"a", "b", "c"}, {"wider cell", "m", "z"}},
			},
			wantLines: 4,
		},
		{
			name: "no alignment",
			lines: []string{
				"| One | Two |",
				"|-----|-----|",
				"| 1 | 2 |",
			},
			want: &markdownTable{
				header: []string{"One", "Two"},
				align:  []string{"", ""},
				rows:   [][]string{{"1", "2"}},
			},
			wantLines: 3,
		},
		{
			name: "inline content kept raw",
			lines: []string{
				"| Format | Example |",
				"|--------|---------|",
				"| bold | **bold** |",
				"| italic | *italic* |",
				"| code | `code()` |",
				"| struck | ~~gone~~ |",
				"| link | [site](/home) |",
			},
			want: &markdownTable{
				header: []string{"Format", "Example"},
				align:  []string{"", ""},
				rows: [][]string{
					{"bold", "**bold**"},
					{"italic", "*italic*"},
					{"code", "`code()`"},
					{"struck", "~~gone~~"},
					{"link", "[site](/home)"},
				},
			},
			wantLines: 7,
		},
		{
			name: "escaped pipes stay in cells",
			lines: []string{
				"| Expression | Meaning |",
				"|------------|---------|",
				"| a \\| b | literal pipe |",
				"| `x \\| y` | pipe in code span |",
			},
			want: &markdownTable{
				header: []string{"Expression", "Meaning"},
				align:  []string{"", ""},
				rows: [][]string{
					{"a \\| b", "literal pipe"},
					{"`x \\| y`", "pipe in code span"},
				},
			},
			wantLines: 4,
		},
		{
			name: "escaped pipe at end of line loses its pipe",
			lines: []string{
				"| a \\|",
				"---",
			},
			want: &markdownTable{
				header: []string{"a \\"},
				align:  []string{""},
			},
			wantLines: 2,
		},
		{
			name: "escaped backslash before pipe does not split",
			lines: []string{
				"| a \\\\| b |",
				"|---|",
			},
			want: &markdownTable{
				header: []string{"a \\\\| b"},
				align:  []string{""},
			},
			wantLines: 2,
		},
		{
			name: "irregular rows are padded and truncated",
			lines: []string{
				"| A | B | C |",
				"|---|---|---|",
				"| short row |",
				"| x | y | z | extra cell |",
			},
			want: &markdownTable{
				header: []string{"A", "B", "C"},
				align:  []string{"", "", ""},
				rows:   [][]string{{"short row", "", ""}, {"x", "y", "z"}},
			},
			wantLines: 4,
		},
		{
			name: "empty cells",
			lines: []string{
				"| | | |",
				"|---|---|---|",
				"| | filled | |",
			},
			want: &markdownTable{
				header: []string{"", "", ""},
				align:  []string{"", "", ""},
				rows:   [][]string{{"", "filled", ""}},
			},
			wantLines: 3,
		},
		{
			name: "empty body",
			lines: []string{
				"| a | b |",
				"|---|---|",
			},
			want: &markdownTable{
				header: []string{"a", "b"},
				align:  []string{"", ""},
			},
			wantLines: 2,
		},
		{
			name: "short header is padded",
			lines: []string{
				"| a |",
				"|---|---|",
			},
			want: &markdownTable{
				header: []string{"a", ""},
				align:  []string{"", ""},
			},
			wantLines: 2,
		},
		{
			name: "header without pipes",
			lines: []string{
				"a",
				"|---|",
			},
			want: &markdownTable{
				header: []string{"a"},
				align:  []string{""},
			},
			wantLines: 2,
		},
		{
			name: "minimal delimiter",
			lines: []string{
				"| a |",
				"-|-",
			},
			want: &markdownTable{
				header: []string{"a", ""},
				align:  []string{"", ""},
			},
			wantLines: 2,
		},
		{
			name: "mixed alignment without outer pipes",
			lines: []string{
				"| a | b | c |",
				":---|-|---:",
			},
			want: &markdownTable{
				header: []string{"a", "b", "c"},
				align:  []string{"left", "", "right"},
			},
			wantLines: 2,
		},
		{
			name: "center alignment",
			lines: []string{
				"| a |",
				"|:---:|",
			},
			want: &markdownTable{
				header: []string{"a"},
				align:  []string{"center"},
			},
			wantLines: 2,
		},
		{
			name: "delimiter surrounded by spaces",
			lines: []string{
				"| a | b |",
				"  | :--- | ---: |  ",
			},
			want: &markdownTable{
				header: []string{"a", "b"},
				align:  []string{"left", "right"},
			},
			wantLines: 2,
		},
		{
			name: "cell spaces are trimmed",
			lines: []string{
				"|   x   |   |",
				"|---|---|",
			},
			want: &markdownTable{
				header: []string{"x", ""},
				align:  []string{"", ""},
			},
			wantLines: 2,
		},
		{
			name: "body stops at blank line",
			lines: []string{
				"| a |",
				"|---|",
				"| x |",
				"",
				"| y |",
			},
			want: &markdownTable{
				header: []string{"a"},
				align:  []string{""},
				rows:   [][]string{{"x"}},
			},
			wantLines: 3,
		},
		{
			name: "plain body line becomes a single padded cell",
			lines: []string{
				"| a | b |",
				"|---|---|",
				"plain text",
			},
			want: &markdownTable{
				header: []string{"a", "b"},
				align:  []string{"", ""},
				rows:   [][]string{{"plain text", ""}},
			},
			wantLines: 3,
		},
		{
			name: "solo pipe body row becomes empty cells",
			lines: []string{
				"| a |",
				"|---|",
				"|",
			},
			want: &markdownTable{
				header: []string{"a"},
				align:  []string{""},
				rows:   [][]string{{""}},
			},
			wantLines: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, consumed := parseTable(tt.lines)
			if tt.want == nil {
				assert.Nil(t, got)
				assert.Zero(t, consumed)
				return
			}
			require.NotNil(t, got)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantLines, consumed)
		})
	}
}

func TestTableRender(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{
			name: "alignment styles",
			lines: []string{
				"| Left | Center | Right |",
				"|:-----|:------:|------:|",
				"| a | b | c |",
				"| wider cell | m | z |",
			},
			want: `<table>
<thead>
<tr>
<th style="text-align:left">Left</th>
<th style="text-align:center">Center</th>
<th style="text-align:right">Right</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align:left">a</td>
<td style="text-align:center">b</td>
<td style="text-align:right">c</td>
</tr>
<tr>
<td style="text-align:left">wider cell</td>
<td style="text-align:center">m</td>
<td style="text-align:right">z</td>
</tr>
</tbody>
</table>
`,
		},
		{
			name: "no alignment attributes",
			lines: []string{
				"| One | Two |",
				"|-----|-----|",
				"| 1 | 2 |",
			},
			want: `<table>
<thead>
<tr>
<th>One</th>
<th>Two</th>
</tr>
</thead>
<tbody>
<tr>
<td>1</td>
<td>2</td>
</tr>
</tbody>
</table>
`,
		},
		{
			name: "empty body has no tbody",
			lines: []string{
				"| a | b |",
				"|---|---|",
			},
			want: `<table>
<thead>
<tr>
<th>a</th>
<th>b</th>
</tr>
</thead>
</table>
`,
		},
		{
			name: "irregular rows",
			lines: []string{
				"| A | B | C |",
				"|---|---|---|",
				"| short row |",
				"| x | y | z | extra cell |",
			},
			want: `<table>
<thead>
<tr>
<th>A</th>
<th>B</th>
<th>C</th>
</tr>
</thead>
<tbody>
<tr>
<td>short row</td>
<td></td>
<td></td>
</tr>
<tr>
<td>x</td>
<td>y</td>
<td>z</td>
</tr>
</tbody>
</table>
`,
		},
		{
			name: "cells are escaped by the callback",
			lines: []string{
				"| a & b | <i>c</i> |",
				"|---|---|",
				"| x | y |",
			},
			want: `<table>
<thead>
<tr>
<th>a &amp; b</th>
<th>&lt;i&gt;c&lt;/i&gt;</th>
</tr>
</thead>
<tbody>
<tr>
<td>x</td>
<td>y</td>
</tr>
</tbody>
</table>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table, consumed := parseTable(tt.lines)
			require.NotNil(t, table)
			assert.Equal(t, len(tt.lines), consumed)
			assert.Equal(t, tt.want, renderTable(table, escapeHTML))
		})
	}
}

func TestTableRenderNil(t *testing.T) {
	assert.Empty(t, renderTable(nil, escapeHTML))
}

func TestTableRenderDelegatesInline(t *testing.T) {
	table, _ := parseTable([]string{
		"| Expression | Meaning |",
		"|------------|---------|",
		"| a \\| b | literal pipe |",
		"| `x \\| y` | pipe in code span |",
	})
	require.NotNil(t, table)

	var seen []string
	got := renderTable(table, func(s string) string {
		seen = append(seen, s)
		return "[" + s + "]"
	})
	assert.Equal(t, []string{
		"Expression", "Meaning",
		"a \\| b", "literal pipe",
		"`x \\| y`", "pipe in code span",
	}, seen)
	assert.Equal(t, `<table>
<thead>
<tr>
<th>[Expression]</th>
<th>[Meaning]</th>
</tr>
</thead>
<tbody>
<tr>
<td>[a \| b]</td>
<td>[literal pipe]</td>
</tr>
<tr>
<td>[`+"`x \\| y`"+`]</td>
<td>[pipe in code span]</td>
</tr>
</tbody>
</table>
`, got)
}
