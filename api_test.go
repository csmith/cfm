package cfm_test

import (
	"testing"

	"chameth.com/cfm"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"empty", "", ""},
		{"spaces only", "   ", ""},
		{"blank lines only", " \n\t\n\n  ", ""},
		{"single paragraph", "hello", "<p>hello</p>\n"},
		{"paragraph with trailing newline", "hello\n", "<p>hello</p>\n"},
		{"soft line break", "hello\nworld", "<p>hello\nworld</p>\n"},
		{"two paragraphs", "para one\n\npara two", "<p>para one</p>\n<p>para two</p>\n"},
		{"three paragraphs", "one\n\ntwo\n\nthree", "<p>one</p>\n<p>two</p>\n<p>three</p>\n"},
		{"excess blank lines collapse", "one\n\n\n\n\ntwo", "<p>one</p>\n<p>two</p>\n"},
		{"blank line containing whitespace", "one\n   \ntwo", "<p>one</p>\n<p>two</p>\n"},
		{"CRLF line break", "a\r\nb", "<p>a\nb</p>\n"},
		{"CR line break", "a\rb", "<p>a\nb</p>\n"},
		{"CRLF paragraph separator", "a\r\n\r\nb", "<p>a</p>\n<p>b</p>\n"},
		{"mixed line endings", "a\r\nb\rc\r\nd", "<p>a\nb\nc\nd</p>\n"},
		{"NUL replaced", "a\x00b", "<p>a\uFFFDb</p>\n"},
		{"NUL alone", "\x00", "<p>\uFFFD</p>\n"},
		{"HTML special characters escaped", "a & b < c > d", "<p>a &amp; b &lt; c &gt; d</p>\n"},
		{"indented paragraph", "    indented four", "<p>indented four</p>\n"},
		{"tab-indented paragraph", "\tTabbed", "<p>Tabbed</p>\n"},
		{"mixed indentation across lines", "  two spaces\n      six spaces\n\ttab", "<p>two spaces\nsix spaces\ntab</p>\n"},
		{"continuation line indented", "text\n   more", "<p>text\nmore</p>\n"},
		{"trailing spaces stripped", "trailing spaces   ", "<p>trailing spaces</p>\n"},
		{"indented paragraphs", "  leading\npara\n\n  leading two", "<p>leading\npara</p>\n<p>leading two</p>\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, cfm.Markdown(tt.source), "Markdown(%q)", tt.source)
			assert.Equal(t, tt.want, (cfm.Renderer{}).Markdown(tt.source), "Renderer{}.Markdown(%q)", tt.source)
		})
	}
}
