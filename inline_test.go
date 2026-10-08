package cfm

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func render(t *testing.T, source string) string {
	t.Helper()
	return renderInline(source, nil)
}

func renderTableCell(t *testing.T, source string) string {
	t.Helper()
	return renderInline(source, &inlineContext{table: true})
}

func TestEmphasis(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"simple star", "foo*bar*", "foo<em>bar</em>"},
		{"digits around", "5*6*78", "5<em>6</em>78"},
		{"simple underscore", "_foo bar_", "<em>foo bar</em>"},
		{"underscore with spaces", "_ foo bar_", "_ foo bar_"},
		{"intraword star", "in*tra*word works", "in<em>tra</em>word works"},
		{"intraword underscore", "in__ter__ underscore does not", "in__ter__ underscore does not"},
		{"triple star", "***all three***", "<em><strong>all three</strong></em>"},
		{"nested emphasis", "**bold *nested***", "<strong>bold <em>nested</em></strong>"},
		{"alternating runs", "*foo**bar**baz*", "<em>foo<strong>bar</strong>baz</em>"},
		{"parens inside", "*(**foo**)*", "<em>(<strong>foo</strong>)</em>"},
		{"rule of three", "*foo**bar*", "<em>foo**bar</em>"},
		{"unclosed", "**bold never closed and *neither is this", "**bold never closed and *neither is this"},
		{"left flank only", "*a", "*a"},
		{"mixed star underscore", "*a _b_ c*", "<em>a <em>b</em> c</em>"},
		{"underscore punct open", "a *foo bar*", "a <em>foo bar</em>"},
		{"unicode flanking", "foo漢字_bar_baz", "foo漢字_bar_baz"},
		{"unicode emphasis", "Grüße *wunderbar*", "Grüße <em>wunderbar</em>"},
		{"list-like run literal", "****", "****"},
		{"empty emphasis", "**()**", "<strong>()</strong>"},
		{"adjacent pair", "foo_bar_baz", "foo_bar_baz"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestStrikethrough(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"single tilde", "~single~ strike", "<del>single</del> strike"},
		{"double tilde", "~~double~~ strike", "<del>double</del> strike"},
		{"both", "~single~ and ~~double~~ both strike", "<del>single</del> and <del>double</del> both strike"},
		{"triple literal", "~~~triple tildes~~~ inside", "~~~triple tildes~~~ inside"},
		{"unclosed", "~~never closed here", "~~never closed here"},
		{"rich content", "~~struck **bold** and *emph*~~", "<del>struck <strong>bold</strong> and <em>emph</em></del>"},
		{"nested tilde", "~~a ~b~~", "~<del>a <del>b</del></del>"},
		{"run inside", "~a~~b~", "<del>a~~b</del>"},
		{"leading tilde run", "a~~b", "a~~b"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestCodeSpans(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"simple", "a `code` b", "a <code>code</code> b"},
		{"protects raw text", "Code keeps raw text: `\"quoted\" -- 'single' ...`",
			"Code keeps raw text: <code>&quot;quoted&quot; -- 'single' ...</code>"},
		{"space trimming", "a ` x ` b", "a <code>x</code> b"},
		{"double backtick", "a `` code `` b", "a <code>code</code> b"},
		{"backtick inside", "``a ` ``", "<code>a ` </code>"},
		{"unclosed", "a `code b", "a `code b"},
		{"run mismatch", "a ``code` b", "a ``code` b"},
		{"multiple spans", "`not `code` here`", "<code>not </code>code<code> here</code>"},
		{"backslash literal", "`a\\`b`", "<code>a\\</code>b`"},
		{"empty span", "a `` b", "a `` b"},
		{"multiline", "a `one\ntwo` b", "a <code>one two</code> b"},
		{"no typography inside", "`&copy;`", "<code>&amp;copy;</code>"},
		{"escaped pipe literal", "`x \\| y`", "<code>x \\| y</code>"},
		{"double backslash literal", "`a\\\\|b`", "<code>a\\\\|b</code>"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestTableCodeSpans(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"escaped pipe in code span", "`x \\| y`", "<code>x | y</code>"},
		{"multiple escaped pipes", "`a\\|b\\|c`", "<code>a|b|c</code>"},
		{"double backslash before pipe", "`a\\\\|b`", "<code>a\\|b</code>"},
		{"triple backslash before pipe", "`a\\\\\\|b`", "<code>a\\\\|b</code>"},
		{"double backslash not before pipe", "`a\\\\b`", "<code>a\\\\b</code>"},
		{"plain pipe untouched", "`a|b`", "<code>a|b</code>"},
		{"unmatched backtick", "`x \\| y", "`x | y"},
		{"unmatched backtick run", "``x \\| y", "``x | y"},
		{"escaped pipe outside code span", "x \\| y", "x | y"},
		{"code pipe with typographer", "`a\\|b` -- done", "<code>a|b</code> – done"},
		{"code pipe with quotes", "`a\\|b` it's here", "<code>a|b</code> it’s here"},
		{"code pipe with entity", "`\\|` &amp; done", "<code>|</code> &amp; done"},
		{"emphasis around escaped pipe", "*x \\| y*", "<em>x | y</em>"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, renderTableCell(t, tt.source))
			},
		)
	}
}

func TestEscapes(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"escaped star", "\\*not emphasised\\*", "*not emphasised*"},
		{"escaped underscore", "\\_under\\_", "_under_"},
		{"escaped bracket", "\\[bracket\\]", "[bracket]"},
		{"escaped backslash", "a \\\\ z", "a \\ z"},
		{"before letter kept", "a\\ab", "a\\ab"},
		{"escaped quote no typography", "\\\"hi\\\"", "&quot;hi&quot;"},
		{"escaped hash", "\\# not heading", "# not heading"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestEntities(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"mixed", "Known &copy; &amp; &lt;tag&gt; &#8212; unknown &fakeentity; raw & alone.",
			"Known © &amp; &lt;tag&gt; — unknown &amp;fakeentity; raw &amp; alone."},
		{"hex numeric", "&#x1F600;", "😀"},
		{"decimal numeric", "&#128512;", "😀"},
		{"too many digits", "&#9999999999;", "&amp;#9999999999;"},
		{"out of range", "&#x110000;", "\uFFFD"},
		{"no semicolon", "&copy", "&amp;copy"},
		{"legacy prefix is not an entity", "&notit; &copycat; &notx;", "&amp;notit; &amp;copycat; &amp;notx;"},
		{"valid names including semicolon", "&not; &semi; &fjlig;", "¬ ; fj"},
		{"bare amp", "1 < 2 > 3 & 4", "1 &lt; 2 &gt; 3 &amp; 4"},
		{"empty name", "&; amp", "&amp;; amp"},
		{"entity quote not typographic", "&quot;hi&quot; there", "&quot;hi&quot; there"},
		{"entity dash not typographic", "a &amp; b -- c", "a &amp; b – c"},
		{"entity in code span", "`&copy;` stays", "<code>&amp;copy;</code> stays"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestLineBreaks(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"soft break", "line one\nline two", "line one\nline two"},
		{"hard break two spaces", "trailing two spaces  \nnext line.", "trailing two spaces<br>\nnext line."},
		{"hard break backslash", "backslash break\\\nsecond line.", "backslash break<br>\nsecond line."},
		{"trailing space soft break", "a \nb", "a\nb"},
		{"trailing spaces last line", "paragraph ends with trailing spaces   ", "paragraph ends with trailing spaces"},
		{"trailing newline ignored", "hello\n", "hello"},
		{"trailing spaces on last line trimmed", "text  \n", "text"},
		{"double backslash soft", "a\\\\\nb", "a\\\nb"},
		{"empty lines", "a\n\nb", "a\n\nb"},
		{"multi paragraph shape", "one\ntwo\n\nthree", "one\ntwo\n\nthree"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestInlineLinks(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"simple", "[inline link](/dest)", `<a href="/dest">inline link</a>`},
		{"with title", `[inline link](/dest "the title")`, `<a href="/dest" title="the title">inline link</a>`},
		{"styled text", "[**styled** text](/other)", `<a href="/other"><strong>styled</strong> text</a>`},
		{"angle destination", "[a](</x y> 't')", `<a href="/x%20y" title="t">a</a>`},
		{"balanced parens", "[a](/x(y))", `<a href="/x(y)">a</a>`},
		{"empty destination", "[a]() empty", `<a href="">a</a> empty`},
		{"leading space", "[a]( /x)", `<a href="/x">a</a>`},
		{"title trailing space", "[a](/x 'title' )", `<a href="/x" title="title">a</a>`},
		{"relative url", "[a](./local/page.html)", `<a href="./local/page.html">a</a>`},
		{"query kept", "[a](/s?q=x&y=1)", `<a href="/s?q=x&amp;y=1">a</a>`},
		{"space in bare dest fails", "[a](/x y)", `[a](/x y)`},
		{"unclosed paren", "[text](missing-paren", "[text](missing-paren"},
		{"unclosed title", `[a](/x "unclosed`, `[a](/x “unclosed`},
		{"escaped quote in title", `[a](/x "ti\"tle")`, `<a href="/x" title="ti&quot;tle">a</a>`},
		{"multiline title", "[a](/x \"ti\n tle\")", "<a href=\"/x\" title=\"ti\ntle\">a</a>"},
		{"empty title", `[a](/x "")`, `<a href="/x" title="">a</a>`},
		{"percent kept", "[a](/x%20y)", `<a href="/x%20y">a</a>`},
		{"invalid percent pair", "[a](/x%2zy)", `<a href="/x%252zy">a</a>`},
		{"unknown entity in destination", "[a](/&notit;)", `<a href="/&amp;notit;">a</a>`},
		{"numeric zero in destination", "[a](/x&#0;y)", `<a href="/x%EF%BF%BDy">a</a>`},
		{"entity in url", `[a](/x?n=&copy;)`, `<a href="/x?n=%C2%A9">a</a>`},
		{"backslash escape in url", `[a](/x\(y)`, `<a href="/x(y">a</a>`},
		{"stray brackets literal", "]stray[ and a[rbitrary] text ]here[", "]stray[ and a[rbitrary] text ]here["},
		{"unclosed label literal", "[a b c", "[a b c"},
		{"failed then literal", "[a] b", "[a] b"},
		{"reference style literal", "See [logo] and the [full form][logo] for details.",
			"See [logo] and the [full form][logo] for details."},
		{"failed reference then inline link", "[a][b](/x)", `[a]<a href="/x">b</a>`},
		{"no nested links", "[a [b](/x) c](/y)", `[a <a href="/x">b</a> c](/y)`},
		{"emphasis in label", "[a *b* c](/x)", `<a href="/x">a <em>b</em> c</a>`},
		{"link in emphasis", "*em [link](/x) end*", `<em>em <a href="/x">link</a> end</em>`},
		{"linkify suppressed in label", "[see https://example.com now](/x)", `<a href="/x">see https://example.com now</a>`},
		{"dest with utf8", "[a](/ünï)", `<a href="/%C3%BCn%C3%AF">a</a>`},
		{"angle dest gt", "[a](<b)c>)", `<a href="b)c">a</a>`},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestImagesLiteral(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"block image", "![Alt text](/images/photo.png)", "![Alt text](/images/photo.png)"},
		{"inline image", "Inline ![small](/i.png) within a sentence.", "Inline ![small](/i.png) within a sentence."},
		{"image with title", `![a](/i.png "t")`, `![a](/i.png &quot;t&quot;)`},
		{"unclosed image", "![alt](/i.png", "![alt](/i.png"},
		{"image in link", "[![a](/i.png)](/l)", `<a href="/l">![a](/i.png)</a>`},
		{"bang alone", "a ! b", "a ! b"},
		{"bang bracket not image", "![unclosed", "![unclosed"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestAutolinks(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"url", "<https://example.com/auto>", `<a href="https://example.com/auto">https://example.com/auto</a>`},
		{"email", "<fred@example.com>", `<a href="mailto:fred@example.com">fred@example.com</a>`},
		{"mailto scheme", "<mailto:a@b.io>", `<a href="mailto:a@b.io">mailto:a@b.io</a>`},
		{"ftp", "<ftp://a.b/c>", `<a href="ftp://a.b/c">ftp://a.b/c</a>`},
		{"unterminated", "<https://example.com", "&lt;https://example.com"},
		{"spaces stop", "<https://example.com/a b>", "&lt;https://example.com/a b&gt;"},
		{"loose attributes match goldmark", "<not a link>", "<not a link>"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestLinkify(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"plain url", "See https://example.com/page for details.",
			`See <a href="https://example.com/page">https://example.com/page</a> for details.`},
		{"trailing dot", "Ends with dot https://example.com.",
			`Ends with dot <a href="https://example.com">https://example.com</a>.`},
		{"paren wrap", "Wrapped (https://example.com/paren) too.",
			`Wrapped (<a href="https://example.com/paren">https://example.com/paren</a>) too.`},
		{"query escape", "Query https://example.com/search?q=x&y=1 works.",
			`Query <a href="https://example.com/search?q=x&amp;y=1">https://example.com/search?q=x&amp;y=1</a> works.`},
		{"email", "Mail fred.bloggs+tag@example.co.uk now.",
			`Mail <a href="mailto:fred.bloggs+tag@example.co.uk">fred.bloggs+tag@example.co.uk</a> now.`},
		{"bare domain", "Bare example.com is not linkified.", "Bare example.com is not linkified."},
		{"url at line start", "https://example.com first", `<a href="https://example.com">https://example.com</a> first`},
		{"ftp scheme", "ftp://files.example.com/x", `<a href="ftp://files.example.com/x">ftp://files.example.com/x</a>`},
		{"www scheme", "try www.example.com/path", `try <a href="http://www.example.com/path">www.example.com/path</a>`},
		{"trailing question", "query https://example.com/s?q?",
			`query <a href="https://example.com/s?q">https://example.com/s?q</a>?`},
		{"punctuation after", "link https://example.com, next",
			`link <a href="https://example.com">https://example.com</a>, next`},
		{"colon then stop", "at https://example.com: ending",
			`at <a href="https://example.com">https://example.com</a>: ending`},
		{"no scheme http only", "xhttp://example.com does not", "xhttp://example.com does not"},
		{"email without tld dot fails", "a@b nolink", "a@b nolink"},
		{"email label includes inner dash", "mail a@b.io-x no", `mail <a href="mailto:a@b.io-x">a@b.io-x</a> no`},
		{"url with escaped paren", "u (https://e.io/a(b)) done",
			`u (<a href="https://e.io/a(b)">https://e.io/a(b)</a>) done`},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestRawHTML(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"inline tags", "Inline <b>bold</b> and a <br>break.", "Inline <b>bold</b> and a <br>break."},
		{"attributes", "a <span title=\"a&amp;b\">tag</span> after & marker.",
			`a <span title="a&amp;b">tag</span> after &amp; marker.`},
		{"malformed angle", "a < not a tag > b and </ dangling.", "a &lt; not a tag &gt; b and &lt;/ dangling."},
		{"comment", "before <!-- hidden --> after", "before <!-- hidden --> after"},
		{"empty comment", "<!----> x", "<!----> x"},
		{"processing instruction", "<?pi data?> ok", "<?pi data?> ok"},
		{"declaration", "<!A decl> ok", "<!A decl> ok"},
		{"cdata", "<![CDATA[x]]> ok", "<![CDATA[x]]> ok"},
		{"unclosed comment", "a <!-- x", "a &lt;!– x"},
		{"autolink precedence", "<http://a.b>", `<a href="http://a.b">http://a.b</a>`},
		{"typographer angle quotes", "foo...bar..x.--y---z<<w>>v",
			"foo…bar..x.–y—z«w»v"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestTypography(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"double quotes", `"Double quoted" and 'single quoted' and "nested 'inner' outer".`,
			"“Double quoted” and ‘single quoted’ and " +
				"“nested ‘inner’ outer”."},
		{"apostrophes", "Don't can't won't it's a party.", "Don’t can’t won’t it’s a party."},
		{"dashes ellipsis", "An en -- dash, an em --- dash and an elipsis... done.",
			"An en – dash, an em — dash and an elipsis… done."},
		{"quotes around link", `"[linked](/page)" and 'plain'.`,
			"“<a href=\"/page\">linked</a>” and ‘plain’."},
		{"decade", "e.g. '90s style", "e.g. ’90s style"},
		{"twas", "'twas the night", "’twas the night"},
		{"em contraction", "'em ok", "’em ok"},
		{"possessive plural", "dogs' bowls", "dogs’ bowls"},
		{"rock and roll", "rock'n'roll", "rock’n’roll"},
		{"open close counter", `"a" 'b' "c`, "“a” ‘b’ “c"},
		{"quote after digit", `"Monitor 21"" ok`, "“Monitor 21&quot;” ok"},
		{"plain quotes between letters", `a"b"c`, "a&quot;b&quot;c"},
		{"single between letters", "a'b'c", "a’b’c"},
		{"double ellipsis", "foo..bar", "foo..bar"},
		{"four dashes", "x----y", "x—-y"},
		{"caret unaffected", "a ^ b", "a ^ b"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestTextEscaping(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"specials", `a & b < c > d " e`, `a &amp; b &lt; c &gt; d &quot; e`},
		{"apostrophe kept", "it's", "it’s"},
		{"nul replaced", "a\x00b", "a\uFFFDb"},
		{"unicode preserved", "Grüße aus München, 東京 and emoji 🎉", "Grüße aus München, 東京 and emoji 🎉"},
		{"nbsp preserved", "a\u00a0b", "a\u00a0b"},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestFootnoteReferences(t *testing.T) {
	sup := func(n string) string {
		return `<sup id="fnref:` + n + `"><a href="#fn:` + n + `" class="footnote-ref">` + n + `</a></sup>`
	}
	ctx := &inlineContext{footnote: func(label string) (string, bool) {
		if label == "1" {
			return sup("1"), true
		}
		return "", false
	}}
	t.Run(
		"defined reference",
		func(t *testing.T) {
			assert.Equal(t, "A statement"+sup("1")+".", renderInline("A statement[^1].", ctx))
		},
	)
	t.Run(
		"undefined literal",
		func(t *testing.T) {
			assert.Equal(t, "Unknown note[^99] stays literal.", renderInline("Unknown note[^99] stays literal.", ctx))
		},
	)
	t.Run(
		"nil context literal",
		func(t *testing.T) {
			assert.Equal(t, "note[^1] here", render(t, "note[^1] here"))
		},
	)
	t.Run(
		"image bang stays",
		func(t *testing.T) {
			assert.Equal(t, "!"+sup("1")+" img", renderInline("![^1] img", ctx))
		},
	)
	t.Run(
		"image undefined literal",
		func(t *testing.T) {
			assert.Equal(t, "![^9] img", renderInline("![^9] img", ctx))
		},
	)
	t.Run(
		"reference wins over link",
		func(t *testing.T) {
			assert.Equal(t, sup("1")+"(/x)", renderInline("[^1](/x)", ctx))
		},
	)
	t.Run(
		"label with escape",
		func(t *testing.T) {
			assert.Equal(t, "note[^1] x", renderInline("note[^1\\] x", ctx))
		},
	)
	t.Run(
		"inside emphasis",
		func(t *testing.T) {
			assert.Equal(t, "<em>a "+sup("1")+" b</em>", renderInline("*a [^1] b*", ctx))
		},
	)
}

func TestPrecedence(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{"code beats emphasis", "*`not em`*", "<em><code>not em</code></em>"},
		{"html beats emphasis", "*<b>x</b>*", "<em><b>x</b></em>"},
		{"escape beats emphasis", "\\*not\\*", "*not*"},
		{"link beats linkify", "[https://example.com](/x)", `<a href="/x">https://example.com</a>`},
		{"backslash in url", "[a](/x\\)y)", `<a href="/x)y">a</a>`},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, render(t, tt.source))
			},
		)
	}
}

func TestRenderInlineDoesNotPanic(t *testing.T) {
	inputs := []string{
		"", "\n", "   ", "*", "**", "***", "****", "`", "``", "![", "](", "[", "]",
		"<", ">", "&", "\\", "~", "'", "\"", "![a", "[a", "[a]", "![a]", "`a",
		"strings.Repeat test filler", "*a*b*c*d*e*", "~~~~", "<<<<", "&amp",
		"[a](b", "![a](b", "___", "~~", "^", "-.-", strings.Repeat("*", 5000),
		strings.Repeat("[", 3000), strings.Repeat("*a", 3000),
	}
	for _, in := range inputs {
		assert.NotPanics(t, func() { _ = renderInline(in, nil) }, "input %q", in)
	}
}

func FuzzRenderInline(f *testing.F) {
	seeds := []string{
		"",
		"plain paragraph text\n",
		"a `code span` and *emphasis* and ~~strike~~",
		"cells with `x \\| y` and \\| escapes",
		"| a \\| b | `c \\| d` |\n| --- | ---: |\n| e | f |",
		"`unclosed span and \\* escapes",
		"**bold *nested* [link](/x \"title\")**",
		"<https://example.com> www.example.com a@b.io",
		"smart 'quotes' \"double\" -- --- ... and &copy; &#x1F600; &fake; &",
		"line one  \nline two\\\nline three\nline four",
		"\x00 invalid \xff utf8 \xe2\x82",
	}
	for _, seed := range seeds {
		f.Add(seed)
	}
	ctx := &inlineContext{table: true}
	f.Fuzz(
		func(t *testing.T, source string) {
			plain := renderInline(source, nil)
			assert.Equal(t, plain, renderInline(source, nil))
			assert.NotContains(t, plain, "\x00")
			assert.NotPanics(t, func() { _ = renderInline(source, ctx) })
		},
	)
}

var inlineBenchSink string

func BenchmarkRenderInlineParagraph(b *testing.B) {
	source := "Parsing *inline* content turns `code`, [links](/dest), and 'smart quotes' -- " +
		"plus https://example.com/docs and <b>raw HTML</b> -- into HTML."
	b.ReportAllocs()
	for b.Loop() {
		inlineBenchSink = renderInline(source, nil)
	}
}

var (
	refOpenTagRe = regexp.MustCompile(
		`^<[A-Za-z][A-Za-z0-9-]*` +
			`(?:[\r\n \t]+[a-zA-Z_:][a-zA-Z0-9:._-]*` +
			`(?:[\r\n \t]*=[\r\n \t]*(?:[^"'=<>` + "`" + `\x00-\x20]+|'[^']*'|"[^"]*"))?)*` +
			`(?:[ \t]|(?:\r\n|\n){0,1})*/?>`,
	)
	refCloseTagRe = regexp.MustCompile(`^</[A-Za-z][A-Za-z0-9-]*(?:[ \t]|(?:\r\n|\n){0,1})*>`)

	refLinkifyURLRe = regexp.MustCompile(
		`^(?:http|https|ftp)://[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]+` +
			`(?::\d+)?(?:[/#?][-a-zA-Z0-9@:%_+.~#$!?&/=\();,'">\^{}\[\]` + "`" + `]*)?`,
	)
	refLinkifyWWWRe = regexp.MustCompile(
		`^www\.[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]+` +
			`(?:[/#?][-a-zA-Z0-9@:%_\+.~#!?&/=\();,'">\^{}\[\]` + "`" + `]*)?`,
	)
	refEmailDomainRe = regexp.MustCompile(
		`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?` +
			`(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*`,
	)
)

func TestRawTagScannerMatchesPattern(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		closing bool
	}{{
		name: "plain open", source: `<a href="/x"> rest`, closing: false,
	}, {
		name: "open no attrs", source: `<br>`, closing: false,
	}, {
		name: "open slash close", source: `<hr/>`, closing: false,
	}, {
		name: "slash space gt fails", source: `<br / >`, closing: false,
	}, {
		name: "unquoted value eats slash", source: `<a href=x/>`, closing: false,
	}, {
		name: "quoted value with gt", source: `<a title="a>b"> t`, closing: false,
	}, {
		name: "value equals excluded", source: `<a x==>`, closing: false,
	}, {
		name: "attr newline separator", source: "<a\n\nb>x", closing: false,
	}, {
		name: "bare carriage return fails", source: "<a\rx>", closing: false,
	}, {
		name: "crlf before gt", source: "<a\r\n>", closing: false,
	}, {
		name: "dashes in name", source: `<x-data->`, closing: false,
	}, {
		name: "colon underscore attr", source: `<a _x:1.2-y=3>`, closing: false,
	}, {
		name: "unclosed quote fails", source: `<a x="y>`, closing: false,
	}, {
		name: "empty fails", source: `<`, closing: false,
	}, {
		name: "close tag", source: `</a> rest`, closing: true,
	}, {
		name: "close with attrs fails", source: `</a href="foo">`, closing: true,
	}, {
		name: "close with slash fails", source: `</a/>`, closing: true,
	}, {
		name: "close whitespace", source: "</a\n\n> x", closing: true,
	}}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				re := refOpenTagRe
				if tt.closing {
					re = refCloseTagRe
				}
				want := 0
				if loc := re.FindStringIndex(tt.source); loc != nil && loc[0] == 0 {
					want = loc[1]
				}
				assert.Equal(t, want, rawTagLen(tt.source, tt.closing))
			},
		)
	}
}

func TestLinkifyScannerMatchesPattern(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   int // expected linkifyURLEnd / linkifyWWWEnd value
		www    bool
	}{{
		name: "plain url", source: `https://example.com/a rest`, want: len("https://example.com/a"),
	}, {
		name: "port and path", source: `http://a.example.com:8080/x#frag`, want: len("http://a.example.com:8080/x#frag"),
	}, {
		name: "multiple dots last wins", source: `https://a.b.example.io/x`, want: len("https://a.b.example.io/x"),
	}, {
		name: "no tld fails", source: `https://example/`, want: -1,
	}, {
		name: "uppercase tld only", source: `https://example.COM`, want: -1,
	}, {
		name: "tld then digit stops", source: `https://example.c0m/`, want: len("https://example.c"),
	}, {
		name: "star not host byte", source: `https://ex**a.com`, want: -1,
	}, {
		name: "url with underscore path", source: `https://ab.cd/x_y?z=1&w=2`, want: len("https://ab.cd/x_y?z=1&w=2"),
	}, {
		name: "bare port no digits", source: `https://ab.cd:/x`, want: len("https://ab.cd"),
	}, {
		name: "hash path", source: `ftp://ab.cd#x y`, want: len("ftp://ab.cd#x"),
	}, {
		name: "www plain", source: `www.example.com/x`, want: len("www.example.com/x"), www: true,
	}, {
		name: "www keeps port in path", source: `www.example.com:8080/`, want: len("www.example.com"), www: true,
	}, {
		name: "www no tld", source: `www.example/x`, want: -1, www: true,
	}}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				var got int
				if tt.www {
					got = linkifyWWWEnd(tt.source)
				} else {
					got = linkifyURLEnd(tt.source)
				}
				assert.Equal(t, tt.want, got)
			},
		)
	}
}

func TestLinkifyDomainScannerMatchesPattern(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   int
	}{
		// Expected values are offsets after the '@', matching the domain
		// part of the reference pattern: full labels capped at 63 bytes,
		// where a label's trailing hyphen ends the domain unless a dot
		// follows it.
		{name: "simple", source: `a@b.io`, want: 6},
		{name: "hyphen label stops domain", source: `a@b-.io`, want: 3},
		{name: "hyphen inside label", source: `a@ab-cd.ef`, want: 10},
		{name: "label too long capped", source: "a@" + strings.Repeat("x", 70), want: 2 + 63},
		{name: "starts with dot", source: `a@.b`, want: -1},
		{name: "single label", source: `a@b`, want: 3},
		{name: "all hyphen label", source: `a@----`, want: -1},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				after := strings.IndexByte(tt.source, '@') + 1
				assert.Equal(t, tt.want, linkifyDomainEnd(tt.source, after))
			},
		)
	}
}

func FuzzRawTagScanner(f *testing.F) {
	seeds := []string{
		"<a>", "<a href=1>", "<a href='1'>", "<a href=\"1\">", "<a x>",
		"<a x=1 y='2' z=\"3\"/>", "<a\n\nx=1>", "<a\rx>", "<a\r\n>",
		"<a /=", "<a \"'>", "</a>", "</a x>", "</a/>", "</a\n>",
		"<1a>", "<>", "<a", "<a " + strings.Repeat("x", 300) + ">",
		"<a " + strings.Repeat("x=1 ", 100) + ">", "<\x00>", "<\xff>",
	}
	for _, s := range seeds {
		f.Add(s)
	}
	f.Fuzz(
		func(t *testing.T, s string) {
			// rawTagLen assumes the '<' the dispatcher already checked.
			if strings.HasPrefix(s, "<") {
				assert.Equal(t, refMatchLen(refOpenTagRe, s), rawTagLen(s, false))
				if strings.HasPrefix(s, "</") {
					assert.Equal(t, refMatchLen(refCloseTagRe, s), rawTagLen(s, true))
				}
			}
		},
	)
}

func FuzzLinkifyScanner(f *testing.F) {
	seeds := []string{
		"http://a.co", "https://a.b.co/p?q=1", "ftp://x.io#f", "http://a.co:8080/p",
		"https://a.co:", "https://a.co:://", "http://a...b.co", "http://.a.co",
		"http://a.co.", "https://EXAMPLE.com", "https://a.co/" + strings.Repeat("p", 400),
		"https://" + strings.Repeat("a.", 200) + "co", "www.a.co", "www.a.co:80/p",
		"www..a.co", "a@b.co", "a@b-.co", "a@" + strings.Repeat("x", 70),
		"a@b.co.", "a@-b.co", "a@b..co", "\xff\xfe", "a.co\n b.co",
	}
	for _, s := range seeds {
		f.Add(s)
	}
	f.Fuzz(
		func(t *testing.T, s string) {
			want := refMatchLen(refLinkifyURLRe, s)
			if want == 0 {
				want = -1
			}
			assert.Equal(t, want, linkifyURLEnd(s))
			want = refMatchLen(refLinkifyWWWRe, s)
			if want == 0 {
				want = -1
			}
			assert.Equal(t, want, linkifyWWWEnd(s))
			end := refMatchLen(refEmailDomainRe, s)
			want = -1
			if end > 0 {
				want = end
			}
			assert.Equal(t, want, linkifyDomainEnd(s, 0))
		},
	)
}

// refMatchLen reports the length of an anchored leftmost match of re at the
// start of s, or 0 when the pattern does not match at the start.
func refMatchLen(re *regexp.Regexp, s string) int {
	loc := re.FindStringIndex(s)
	if loc == nil || loc[0] != 0 {
		return 0
	}
	return loc[1]
}
