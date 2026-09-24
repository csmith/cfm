# Tables and footnotes

Tables carry structured data; footnotes carry the sort of detail that
would ruin the flow of a sentence[^flow].

## Release matrix

| Version | Status      | Released   | Notes                    |
| ------- | ----------- | ---------- | ------------------------ |
| 1.0     | stable      | 2024-01-15 | Initial release          |
| 1.1     | stable      | 2024-06-02 | Performance fixes[^perf] |
| 2.0     | deprecated  | 2025-03-30 | Breaking API changes     |

## Alignment

| Left        | Centre | Right |
| :---------- | :----: | ----: |
| a           |   b    | 3     |
| longer cell |   x    | 12345 |

[^flow]: Which is not to say footnotes cannot have a flow of their own,
    links [included](/notes/footnote-style-guide).
[^perf]: Specifically, the allocator thrashing described in
    [the profiling post](/notes/profiling).

## Footnote gymnastics

References can repeat[^flow], and definitions can contain **emphasis**,
`inline code`, and further references to earlier notes[^flow]. The
reference in the table above behaves like any other, at least as far as
this fixture is concerned.

| Feature   | Since | Because                       |
| --------- | ----- | ----------------------------- |
| Footnotes | 1.0   | See [^flow]                   |
| Tables    | 1.0   | You are looking at one        |
| Both      | 1.1   | A table citing its own figure |

A final paragraph so the document does not end inside a footnote
definition, which would be a poor way to treat any parser.
