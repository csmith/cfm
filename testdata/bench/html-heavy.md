<div class="notice notice--draft">

This post mixes raw HTML with Markdown. The converter must pass the
HTML through unmodified — no escaping — while still processing the
Markdown in between.

</div>

## Inline HTML

You can <em>emphasise</em>, <strong>strengthen</strong>, highlight
<mark>like this</mark>, abbreviate <abbr title="extensible markup
language">XML</abbr>, and drop a <br> where prose demands drama.
Entities work too: &copy; &amp; &lt;escaped&gt; &mdash; provided the
renderer is not in safe mode.

Attributes can be <span data-kind="meta" data-id=42 complicated>more
elaborate</span>, with single quotes, double quotes, unquoted values,
and comments <!-- like this one --> scattered mid-sentence.

## Block HTML

<figure class="chart">
  <img src="/img/benchmarks.png" alt="Benchmark results" width="640" height="480">
  <figcaption>Figure 1. Results, artist's impression.</figcaption>
</figure>

<details>
<summary>Raw table, not a Markdown table</summary>

<table>
  <thead>
    <tr><th scope="col">Converter</th><th scope="col">Role</th></tr>
  </thead>
  <tbody>
    <tr><td>cfm</td><td>under test</td></tr>
    <tr><td>goldmark</td><td>reference</td></tr>
  </tbody>
</table>

</details>

<aside class="margin-note">
  <p>Self-closing tags (<img src="/img/1x1.gif" alt="">, <hr />, and
  <code>&lt;br/&gt;</code>) must survive alongside void elements written
  the long way, and <em>must not</em> leak into surrounding prose.</p>
</aside>

<p>Finally, a paragraph that is <em>entirely</em> wrapped in HTML tags,
containing <a href="https://example.com/linked">a link</a>, <code>code
spans</code>, and enough prose that the inline parser has real work to
do while deciding where the tags end and the Markdown begins.</p>
