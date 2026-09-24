# Benchmarking Markdown converters

Every static site lives and dies by its build pipeline, and the part of
that pipeline nobody thinks about until it hurts is Markdown conversion.
This post walks through [my setup](/notes/build-pipeline) and the
measurements behind it.

## Why conversion cost matters

Most posts are small, but a few grow like weeds: long-form essays,
changelogs, accumulated notes. At some point the converter's throughput
stops being noise, because you notice it every time you rebuild the
site to check a one-line CSS change.

A typical rebuild converts every page from scratch, so the total cost
is roughly proportional to the sum of all page sizes. The alternative
is described at https://example.org/notes/incremental-builds, along
with the usual caveats about stale fragments.

## The pipeline

1. Read every source file from disk.
2. Parse front matter and convert the body.
3. Render templates.
4. Write the output tree.

Each stage is embarrassingly parallel, but only step 2 is interesting
here.

## What a post looks like

A post is mostly prose, with the occasional structural flourish:

- short paragraphs, because screens are narrow;
- lists like this one;
- `inline code` for commands such as `serve --drafts`;
- the odd image, linked as [a diagram of the pipeline](/img/pipeline.svg).

Occasionally a quote sums things up better than I can:

> The fastest build is the one you don't have to repeat.
>
> — Every developer, eventually

## Code

Most of the code in these posts is short enough to read without
scrolling:

```
func total(pages []page) int {
	sum := 0
	for _, p := range pages {
		sum += len(p.body)
	}
	return sum
}
```

Note that the block above deliberately carries no language tag.

## Details worth forgetting

Headings like this one get automatic identifiers, which is convenient
when you want to link to ***Why conversion cost matters*** directly —
although usually you would link with plain text and normal [emphasis
instead](#why-conversion-cost-matters) of shouting.

---

That is the whole post. There will be a follow-up with numbers, and if
history is any guide, the follow-up will itself need a follow-up.
