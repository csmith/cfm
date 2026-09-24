# chameth.com-flavoured markdown 

Go library that converts Markdown to HTML in the manner required
for [chameth.com](https://chameth.com).

> [!WARNING]
> This code has been entirely written by LLMs, and I haven't reviewed
> it. The output seems to be byte-compatible with
> [Goldmark](https://github.com/yuin/goldmark) aside from deliberate
> deviations. It's primarily intended for use in chameth.com, not as
> a general purpose library. Feel free to report bugs or raise issues
> if you do use it, but buyer beware.

Supports most CommonMark syntax and standard extensions like tables
and footnotes, and applies smart typography and heading IDs.

Things it explicitly does NOT support:

- indents to create code blocks - causes headaches in HTML heavy uses
- setext headings
- markdown images
- reference links
- various kitchen-sink syntaxes like `*-*-` dividers, trailing `#` in headings

## Usage

```go
import "chameth.com/cfm"

html := cfm.Markdown("Hello *world*")

// or with syntax highlighting
renderer := cfm.Renderer{
	Highlight: func(code, language string) (string, error) {
		return myHighlighter(code, language)
	},
}
html = renderer.Markdown(source)
```
