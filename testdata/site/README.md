# Site fixtures

Golden corpus for `Markdown`. Every `NNN-name.md` is paired with the expected
HTML in `NNN-name.html`; `markdown_test.go` renders each pair as an individual
subtest and compares byte for byte.

Except for the hand-authored 9xx series, fixtures were generated with Goldmark
v1.8.5 configured like the site renderer: Typographer, Table, Strikethrough,
Linkify and Footnote extensions, automatic heading IDs, unsafe HTML, and the
indented code block parser removed (no Chroma). Regenerate with:

    go run ./testdata/site/generate.go

The 9xx fixtures are hand-authored and pin deliberate deviations from stock
Goldmark: empty ATX markers, closing hashes, reference definitions, images and
non-hyphen thematic break markers stay literal text. Setext headings are not
recognised: equals underlines stay literal text (902), while a hyphen run under
a paragraph line renders the paragraph plus a thematic break (903).
Indentation is ignored for block recognition, so indented prose stays prose,
while indented fences, headings and lists are still recognised — stock
Goldmark silently drops those indented lines once its indented code parser is
removed, which is why 906–908 exist as regression tests. The generator does
not rewrite the 9xx series.
