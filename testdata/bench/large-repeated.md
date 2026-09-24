## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

## Section: notes on throughput

Conversion cost is dominated by the number of inline constructs a
document carries, not by raw length. A [well-behaved](/notes/parsing)
document keeps its emphasis balanced and its links tidy, which is why
these sections repeat the *same* few shapes with `small variations`
between them.

## Section: a working list

- read the source file from disk;
- strip the front matter;
- convert the body to HTML;
- hand the result to the template engine;
- write the output file.

Each step is simple, and none of them is interesting on its own.

## Section: numbers

| Document | Bytes | Constructs | Notes               |
| -------- | ----- | ---------- | ------------------- |
| small    | 400   | 12         | paragraph only      |
| medium   | 4,200 | 180        | lists and a table   |
| large    | 38,000| 1,900      | everything at once  |

The large row is what this fixture is pretending to be, give or take
an order of magnitude.

## Section: a quoted aside

> Anyone can write a converter for the happy path.
>
> The work is in the *other* documents — the ones that arrive with
> unbalanced markers, pasted spreadsheets, and a `code block` that
> never seems to end.

Those documents are why the reference implementation has the extension
points it has.

## Section: a small program

```
type counter struct {
	paragraphs int
	words      int
}

func (c *counter) visit(n ast.Node) {
	switch n := n.(type) {
	case *ast.Paragraph:
		c.paragraphs++
	case *ast.Text:
		c.words += len(bytes.Fields(n.Segment.Value(n.Text())))
	}
}
```

No language tag on the fence, as with every block in these fixtures.

## Section: closing remarks

If you have read this far, you have now seen this exact block once,
plus or minus the repetition count. That is the point: volume without
novelty, so the converter spends its time converting rather than
deciding. See https://example.org/notes/repetition for why repeated
content is a fair stand-in for a long site-wide rebuild.

---

