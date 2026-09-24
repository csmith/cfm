# CommonMark golden corpus

Golden input/expected-output pairs for `commonmark_test.go`, generated from
the CommonMark 0.31.2 specification. The fixtures — `spec.json`,
`manifest.json`, `examples/`, `overrides/` — are data; `generator/` holds
the Go code that regenerates them. The test harness is `commonmark_test.go`
at the repository root.

## Source and license

The corpus is derived from the CommonMark 0.31.2 specification:

- Source: <https://spec.commonmark.org/0.31.2/spec.json>
  (SHA-256 `d431b29d97b6f73e69d547109cf5081578fac931e72afe95639ebe766c1b2a20`,
  verified on every generation)
- Copyright (C) John MacFarlane, CommonMark spec 0.31.2
- License: [CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/)

`spec.json` is vendored verbatim as the upstream source. The generated
corpus is a derivative work and is likewise licensed CC-BY-SA-4.0. The
licence covers the fixture data only; the Go code in `generator/` is
project source, not part of the licensed corpus.

## Layout

- `spec.json` — upstream source, unmodified
- `manifest.json` — generated; records the disposition (included/excluded
  with reasons, file names, expectation source) of every one of the 652
  examples exactly once
- `examples/NNN.md`, `examples/NNN.html` — input and expected HTML for each
  included example, numbered by the original spec example number
- `examples/NNN.adapted.html` — chameth.com expected output, present only
  where it differs from the spec HTML (recorded per example as
  `adapted_html` in the manifest)
- `generator/` — self-contained Go module (own `go.mod`, pins goldmark
  v1.8.5) that classifies, renders and writes everything above
- `overrides/NNN.html` — hand-authored expectations, keyed by spec example
  number, for included examples whose desired rendering is a deliberate
  decision the site goldmark variant does not produce on its own:
  indentation-hidden headings and thematic breaks that must render as
  themselves, and content the site variant would lose (recorded as
  `expected_source: hand-authored`). File names are exactly `NNN.html`;
  overrides that no example consumes are rejected.

## Selection policy

Whole examples are excluded when they involve:

- indented code blocks (indented code is disabled for chameth.com)
- setext headings
- empty ATX headings or syntactic ATX closing sequences (`#`, `## ##`,
  `# foo #`)
- link reference definitions or reference links; malformed definitions that
  the spec renders as literal text are kept
- images
- thematic breaks other than runs of at least three contiguous hyphens
  (`---` is in scope, `***`, `___` and spaced variants are not)

Classification parses every example with goldmark and combines AST node
analysis with source-level checks on lines outside code blocks; link
reference definitions are additionally taken from the parser context so
consumed definitions cannot escape detection.

## Expected HTML

The spec HTML in `examples/NNN.html` is the upstream expectation. The
chameth.com renderer uses goldmark v1.8.5 with typographer, tables,
strikethrough, linkify, footnotes, auto heading IDs and unsafe (raw) HTML,
with the indented code block parser disabled. Where that rendering differs
from the spec HTML (heading IDs, `<br>`/`<hr>` without XHTML self-closing,
smart quotes and dashes, and similar), the chameth.com expectation is
written to `examples/NNN.adapted.html` and the manifest points the test at
it. A hand-authored `overrides/NNN.html` takes precedence over the site
rendering. `commonmark_test.go` therefore always compares `Markdown`
output against the chameth.com expectation: `NNN.adapted.html` when
present, otherwise `NNN.html`.

## Regeneration

```
cd generator
go test ./...
go run .
```

`go run .` re-verifies the spec.json checksum, reclassifies all 652
examples, rewrites `examples/` and `manifest.json`, and prints per-reason
exclusion counts plus audit lines for manual review. Generation is
deterministic. If the site rendering loses text that default goldmark
keeps, generation fails unless a hand-authored `overrides/NNN.html` exists
for the example; the same override also silences the audit warning.
