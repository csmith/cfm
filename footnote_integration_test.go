package cfm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationFootnoteForwardReference(t *testing.T) {
	src := "Intro cites[^later] here.\n\nA middle paragraph.\n\n[^later]: Arrived last.\n"
	want := "<p>Intro cites" + footnoteSupHTML("fnref:1", 1) + " here.</p>\n" +
		"<p>A middle paragraph.</p>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<p>Arrived last."+footnoteBackrefHTML("fnref:1")+"</p>\n"))
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteRepeatedAndReordered(t *testing.T) {
	src := "Three[^c], then [^a] and [^b], twice [^a].\n\n" +
		"[^a]: Alpha note.\n\n[^c]: Gamma note.\n\n[^b]: Beta note.\n"
	want := "<p>Three" + footnoteSupHTML("fnref:1", 1) + ", then " + footnoteSupHTML("fnref:2", 2) +
		" and " + footnoteSupHTML("fnref:3", 3) + ", twice " + footnoteSupHTML("fnref1:2", 2) + ".</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>Gamma note."+footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>Alpha note."+footnoteBackrefHTML("fnref:2")+
				footnoteBackrefHTML("fnref1:2")+"</p>\n"),
			footnoteItemHTML(3, "<p>Beta note."+footnoteBackrefHTML("fnref:3")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteUndefinedAndUnreferenced(t *testing.T) {
	src := "Unknown[^99] stays and [^ghost] too.\n\n[^unused]: Nobody cites this.\n"
	assert.Equal(t, "<p>Unknown[^99] stays and [^ghost] too.</p>\n", Markdown(src))
}

func TestIntegrationFootnoteSourceOrderNumbering(t *testing.T) {
	src := "ref[^a]\n\n[^a]: inner [^b]\n\nOther [^c] then [^b]\n\n[^b]: B\n\n[^c]: C\n"
	want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
		"<p>Other " + footnoteSupHTML("fnref:3", 3) + " then " + footnoteSupHTML("fnref1:2", 2) + "</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>inner "+footnoteSupHTML("fnref:2", 2)+
				footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>B"+footnoteBackrefHTML("fnref:2")+
				footnoteBackrefHTML("fnref1:2")+"</p>\n"),
			footnoteItemHTML(3, "<p>C"+footnoteBackrefHTML("fnref:3")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteInterleavedDefinitionsSourceOrder(t *testing.T) {
	src := "one[^a]\n\nmid[^c]\n\n[^a]: to [^b]\n\ntwo[^b] and [^a]\n\n" +
		"[^b]: back [^c]\n\n[^c]: C\n\nend[^b]\n"
	want := "<p>one" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
		"<p>mid" + footnoteSupHTML("fnref:2", 2) + "</p>\n" +
		"<p>two" + footnoteSupHTML("fnref1:3", 3) + " and " + footnoteSupHTML("fnref1:1", 1) + "</p>\n" +
		"<p>end" + footnoteSupHTML("fnref2:3", 3) + "</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>to "+footnoteSupHTML("fnref:3", 3)+
				footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>C"+footnoteBackrefHTML("fnref:2")+
				footnoteBackrefHTML("fnref1:2")+"</p>\n"),
			footnoteItemHTML(3, "<p>back "+footnoteSupHTML("fnref1:2", 2)+
				footnoteBackrefHTML("fnref:3")+footnoteBackrefHTML("fnref1:3")+
				footnoteBackrefHTML("fnref2:3")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteBodiesParseAtFirstDefinition(t *testing.T) {
	src := "use [^a]\n\n[^a]: A [^z]\n\nmid [^m]\n\n[^m]: M [^q]\n\n[^z]: Z\n\n[^q]: Q\n"
	want := "<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
		"<p>mid " + footnoteSupHTML("fnref:4", 4) + "</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>A "+footnoteSupHTML("fnref:2", 2)+
				footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>Z"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
			footnoteItemHTML(3, "<p>Q"+footnoteBackrefHTML("fnref:3")+"</p>\n"),
			footnoteItemHTML(4, "<p>M "+footnoteSupHTML("fnref:3", 3)+
				footnoteBackrefHTML("fnref:4")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteNestedDefinitionListOrder(t *testing.T) {
	src := "use [^n] now\n\n[^n]: plain\n\nuse2 [^m]\n\n[^m]: outer [^k]\n\n" +
		"    [^k]: inner\n\nuse3 [^k]\n"
	want := "<p>use " + footnoteSupHTML("fnref:1", 1) + " now</p>\n" +
		"<p>use2 " + footnoteSupHTML("fnref:3", 3) + "</p>\n" +
		"<p>use3 " + footnoteSupHTML("fnref1:2", 2) + "</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>plain"+footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>inner"+footnoteBackrefHTML("fnref:2")+
				footnoteBackrefHTML("fnref1:2")+"</p>\n"),
			footnoteItemHTML(3, "<p>outer "+footnoteSupHTML("fnref:2", 2)+
				footnoteBackrefHTML("fnref:3")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteNestedDefinitionPreservesBodies(t *testing.T) {
	src := "A[^1].\n\n[^1]: outer [^2] text\n\n    [^2]: nested body\n\nSee [^2] too.\n"
	want := "<p>A" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		"<p>See " + footnoteSupHTML("fnref1:2", 2) + " too.</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>outer "+footnoteSupHTML("fnref:2", 2)+" text"+
				footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>nested body"+footnoteBackrefHTML("fnref:2")+
				footnoteBackrefHTML("fnref1:2")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteDuplicateLabelFirstWins(t *testing.T) {
	src := "ref[^a]\n\n[^a]: one\n\n[^a]: two\n"
	want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<p>one"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteDuplicateBodyDistinctHeadingIDs(t *testing.T) {
	src := "a[^a] b[^b]\n\n[^a]: # Same\n\n[^b]: # Same\n"
	want := "<p>a" + footnoteSupHTML("fnref:1", 1) + " b" + footnoteSupHTML("fnref:2", 2) + "</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<h1 id=\"same\">Same</h1>\n"+footnoteBackrefHTML("fnref:1")),
			footnoteItemHTML(2, "<h1 id=\"same-1\">Same</h1>\n"+footnoteBackrefHTML("fnref:2")),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteReferenceInUnusedDefinition(t *testing.T) {
	src := "text[^a]\n\n[^a]: A\n\n[^unused]: see [^x]\n\n[^x]: X\n"
	want := "<p>text" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>A"+footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>X"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteMultiblockBodies(t *testing.T) {
	src := "P[^p], L[^l], Q[^q].\n\n" +
		"[^p]: First paragraph.\n\n    Second paragraph.\n\n    Final with *emphasis*.\n\n" +
		"[^l]: intro\n\n    - one\n    - two\n\n" +
		"[^q]: intro\n\n    > deep quote\n"
	want := "<p>P" + footnoteSupHTML("fnref:1", 1) + ", L" + footnoteSupHTML("fnref:2", 2) +
		", Q" + footnoteSupHTML("fnref:3", 3) + ".</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>First paragraph.</p>\n<p>Second paragraph.</p>\n"+
				"<p>Final with <em>emphasis</em>."+footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>intro</p>\n<ul>\n<li>one</li>\n<li>two</li>\n</ul>\n"+
				footnoteBackrefHTML("fnref:2")),
			footnoteItemHTML(3, "<p>intro</p>\n<blockquote>\n<p>deep quote</p>\n</blockquote>\n"+
				footnoteBackrefHTML("fnref:3")),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteContainerDefinitions(t *testing.T) {
	src := "- [^i]: item note\n\nQuoted:\n\n> [^q]: quoted note\n\nSee [^i] and [^q].\n"
	want := "<ul>\n<li></li>\n</ul>\n<p>Quoted:</p>\n<blockquote>\n</blockquote>\n" +
		"<p>See " + footnoteSupHTML("fnref:1", 1) + " and " + footnoteSupHTML("fnref:2", 2) + ".</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>item note"+footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>quoted note"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteSkippedInFencesAndHTML(t *testing.T) {
	src := "```\n[^1]: fence note\n```\n\n<div>\n[^2]: html note\n</div>\n\nSee [^1] and [^2].\n"
	want := "<pre><code>[^1]: fence note\n</code></pre>\n" +
		"<div>\n[^2]: html note\n</div>\n" +
		"<p>See [^1] and [^2].</p>\n"
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteEscapedLabels(t *testing.T) {
	src := "An \\[^esc] literal, a real[^a\\]b] one.\n\n[^a\\]b]: escaped label body\n"
	want := "<p>An [^esc] literal, a real" + footnoteSupHTML("fnref:1", 1) + " one.</p>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<p>escaped label body"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteTableReference(t *testing.T) {
	src := "| Note | Value |\n|------|-------|\n| a [^t] | b |\n\n[^t]: table body\n"
	want := "<table>\n<thead>\n<tr>\n<th>Note</th>\n<th>Value</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n" +
		"<td>a " + footnoteSupHTML("fnref:1", 1) + "</td>\n<td>b</td>\n</tr>\n</tbody>\n</table>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<p>table body"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteTableRepeatedReferences(t *testing.T) {
	src := "| h [^t] | k |\n|---|---|\n| a [^t] | b [^u] |\n\n[^t]: T\n\n[^u]: U\n"
	want := "<table>\n<thead>\n<tr>\n<th>h " + footnoteSupHTML("fnref:1", 1) + "</th>\n<th>k</th>\n</tr>\n</thead>\n" +
		"<tbody>\n<tr>\n<td>a " + footnoteSupHTML("fnref1:1", 1) + "</td>\n<td>b " +
		footnoteSupHTML("fnref:2", 2) + "</td>\n</tr>\n</tbody>\n</table>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>T"+footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>U"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteCyclesTerminate(t *testing.T) {
	selfWant := "<p>A" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<p>self "+footnoteSupHTML("fnref1:1", 1)+" ref"+
				footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n"))
	assert.Equal(t, selfWant, Markdown("A[^s].\n\n[^s]: self [^s] ref\n"))

	mutualWant := "<p>A" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>to "+footnoteSupHTML("fnref:2", 2)+
				footnoteBackrefHTML("fnref:1")+footnoteBackrefHTML("fnref1:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>back "+footnoteSupHTML("fnref1:1", 1)+
				footnoteBackrefHTML("fnref:2")+"</p>\n"),
		)
	assert.Equal(t, mutualWant, Markdown("A[^1].\n\n[^1]: to [^2]\n\n[^2]: back [^1]\n"))
}

func TestIntegrationFootnoteHeadingIDSourceOrder(t *testing.T) {
	defFirst := "Text[^1].\n\n[^1]: ## Intro\n\n# Intro\n"
	defFirstWant := "<p>Text" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		"<h1 id=\"intro-1\">Intro</h1>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<h2 id=\"intro\">Intro</h2>\n"+footnoteBackrefHTML("fnref:1")))
	assert.Equal(t, defFirstWant, Markdown(defFirst))

	mainFirst := "# Intro\n\nText[^1].\n\n[^1]: ## Intro\n"
	mainFirstWant := "<h1 id=\"intro\">Intro</h1>\n<p>Text" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<h2 id=\"intro-1\">Intro</h2>\n"+footnoteBackrefHTML("fnref:1")))
	assert.Equal(t, mainFirstWant, Markdown(mainFirst))
}

func TestIntegrationFootnoteKeepsHeadingNamespace(t *testing.T) {
	src := "# One\n\nSee[^n].\n\n# Two\n\n[^n]: ## Two\n\n# Two\n"
	want := "<h1 id=\"one\">One</h1>\n<p>See" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		"<h1 id=\"two\">Two</h1>\n" +
		"<h1 id=\"two-2\">Two</h1>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1,
			"<h2 id=\"two-1\">Two</h2>\n"+footnoteBackrefHTML("fnref:1")))
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteRemovalAndBlankSeparation(t *testing.T) {
	removed := "Note [^1] here.\n\n[^1]: gone\n\nSecond paragraph.\n"
	removedWant := "<p>Note " + footnoteSupHTML("fnref:1", 1) + " here.</p>\n<p>Second paragraph.</p>\n" +
		footnoteWrapperHTML(footnoteItemHTML(1, "<p>gone"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
	assert.Equal(t, removedWant, Markdown(removed))

	interrupted := "Text without blank\n[^1]: note after\n"
	assert.Equal(t, "<p>Text without blank</p>\n", Markdown(interrupted))

	// An unindented line directly after the definition paragraph continues
	// the note; the unreferenced note disappears with it.
	adjacent := "Plain para.\n\n[^1]: note\nImmediately after.\n"
	assert.Equal(t, "<p>Plain para.</p>\n", Markdown(adjacent))
}

func TestIntegrationFootnoteReferencesInBodies(t *testing.T) {
	src := "A[^a].\n\n[^a]: see [^b]\n\n[^b]: the b note\n"
	want := "<p>A" + footnoteSupHTML("fnref:1", 1) + ".</p>\n" +
		footnoteWrapperHTML(
			footnoteItemHTML(1, "<p>see "+footnoteSupHTML("fnref:2", 2)+
				footnoteBackrefHTML("fnref:1")+"</p>\n"),
			footnoteItemHTML(2, "<p>the b note"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
		)
	assert.Equal(t, want, Markdown(src))
}

func TestIntegrationFootnoteCodeSpanStartInterrupted(t *testing.T) {
	src := "text `code\n[^1]: x`\n"
	assert.Equal(t, "<p>text `code</p>\n", Markdown(src))
}

func TestIntegrationFootnoteLazyContinuation(t *testing.T) {
	t.Run("same paragraph", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\ncontinuation\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first\ncontinuation"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("multiple lines", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\ncontinuation\nmore lazy\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first\ncontinuation\nmore lazy"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("blank line ends the note", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\n\ncontinuation\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n<p>continuation</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("blank line between lazy lines keeps later indented paragraph", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\nlazy\n\n    indented para\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first\nlazy</p>\n<p>indented para"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("lazy continues an indented paragraph", func(t *testing.T) {
		src := "use [^a]\n\n[^a]: one\n\n    two\nthree\n"
		want := "<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>one</p>\n<p>two\nthree"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("heading interrupts", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\n# Head\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n<h1 id=\"head\">Head</h1>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("list interrupts", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\n- item\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n<ul>\n<li>item</li>\n</ul>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("fence interrupts", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\n```\nfence\n```\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n<pre><code>fence\n</code></pre>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("thematic break interrupts", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\n---\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n<hr>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("indented lazy lines lose their indent", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]: first\n   cont3\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first\ncont3"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("setext underline is plain text", func(t *testing.T) {
		src := "use [^a]\n\n[^a]: first\n=====\n"
		want := "<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first\n====="+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("html block interrupts", func(t *testing.T) {
		src := "use [^a]\n\n[^a]: first\n<div>\nx\n</div>\n"
		want := "<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n<div>\nx\n</div>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("hard breaks preserved", func(t *testing.T) {
		src := "use [^a]\n\n[^a]: first  \ncont  \ntail\n"
		want := "<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>first<br>\ncont<br>\ntail"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("indented definition line after a paragraph is text", func(t *testing.T) {
		src := "para\n    [^a]: x\n"
		assert.Equal(t, "<p>para\n[^a]: x</p>\n", Markdown(src))
	})

	t.Run("empty body stays empty", func(t *testing.T) {
		src := "ref[^a]\n\n[^a]:\ncontinuation\n"
		want := "<p>ref" + footnoteSupHTML("fnref:1", 1) + "</p>\n<p>continuation</p>\n" +
			footnoteWrapperHTML("<li id=\"fn:1\">\n"+footnoteBackrefHTML("fnref:1")+"</li>\n")
		assert.Equal(t, want, Markdown(src))
	})
}

func TestIntegrationFootnoteLazyThroughContainers(t *testing.T) {
	t.Run("unindented after quote definition", func(t *testing.T) {
		src := "> [^q]: note\ncont\n"
		assert.Equal(t, "<blockquote>\n</blockquote>\n", Markdown(src))
	})

	t.Run("quoted continuation joins the note", func(t *testing.T) {
		src := "> [^q]: note\n> cont\n\nuse [^q]\n"
		want := "<blockquote>\n</blockquote>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note\ncont"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("consecutive lazy lines", func(t *testing.T) {
		src := "> [^q]: note\ncont\nafter\n\nuse [^q]\n"
		want := "<blockquote>\n</blockquote>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note\ncont\nafter"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("blank line returns following text to main", func(t *testing.T) {
		src := "> [^q]: note\ncont\n\nafter\n\nuse [^q]\n"
		want := "<blockquote>\n</blockquote>\n<p>after</p>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note\ncont"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("deeply indented lazy lines strip all indent", func(t *testing.T) {
		src := "> [^q]: note\n        deep\n\nuse [^q]\n"
		want := "<blockquote>\n</blockquote>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note\ndeep"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("unindented after list definition", func(t *testing.T) {
		src := "- [^i]: note\ncont\n\nuse [^i]\n"
		want := "<ul>\n<li></li>\n</ul>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note\ncont"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("item-indented continuation joins the note", func(t *testing.T) {
		src := "- [^i]: note\n  cont\n\nuse [^i]\n"
		want := "<ul>\n<li></li>\n</ul>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note\ncont"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})
}

func TestIntegrationFootnoteDefinitionInterruptsContainers(t *testing.T) {
	t.Run("after list item paragraph", func(t *testing.T) {
		src := "- one\n[^a]: note\n\nuse [^a]\n"
		want := "<ul>\n<li>one</li>\n</ul>\n<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("after quote paragraph", func(t *testing.T) {
		src := "> text\n[^a]: note\n\nuse [^a]\n"
		want := "<blockquote>\n<p>text</p>\n</blockquote>\n" +
			"<p>use " + footnoteSupHTML("fnref:1", 1) + "</p>\n" +
			footnoteWrapperHTML(footnoteItemHTML(1,
				"<p>note"+footnoteBackrefHTML("fnref:1")+"</p>\n"))
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("zero indent after quote definition", func(t *testing.T) {
		src := "> [^q]: note\n[^b]: bee\n\nuse [^q] and [^b]\n"
		want := "<blockquote>\n</blockquote>\n" +
			"<p>use " + footnoteSupHTML("fnref:1", 1) + " and " + footnoteSupHTML("fnref:2", 2) + "</p>\n" +
			footnoteWrapperHTML(
				footnoteItemHTML(1, "<p>note"+footnoteBackrefHTML("fnref:1")+"</p>\n"),
				footnoteItemHTML(2, "<p>bee"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
			)
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("zero indent after item definition", func(t *testing.T) {
		src := "- [^i]: note\n[^j]: other\n\nuse [^i] and [^j]\n"
		want := "<ul>\n<li></li>\n</ul>\n" +
			"<p>use " + footnoteSupHTML("fnref:1", 1) + " and " + footnoteSupHTML("fnref:2", 2) + "</p>\n" +
			footnoteWrapperHTML(
				footnoteItemHTML(1, "<p>note"+footnoteBackrefHTML("fnref:1")+"</p>\n"),
				footnoteItemHTML(2, "<p>other"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
			)
		assert.Equal(t, want, Markdown(src))
	})

	t.Run("nested definitions inside one quote", func(t *testing.T) {
		src := "> [^q]: note\n> [^b]: bee\n\nuse [^q] and [^b]\n"
		want := "<blockquote>\n</blockquote>\n" +
			"<p>use " + footnoteSupHTML("fnref:1", 1) + " and " + footnoteSupHTML("fnref:2", 2) + "</p>\n" +
			footnoteWrapperHTML(
				footnoteItemHTML(1, "<p>note"+footnoteBackrefHTML("fnref:1")+"</p>\n"),
				footnoteItemHTML(2, "<p>bee"+footnoteBackrefHTML("fnref:2")+"</p>\n"),
			)
		assert.Equal(t, want, Markdown(src))
	})
}
