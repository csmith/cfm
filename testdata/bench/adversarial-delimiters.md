# Delimiter stress

Delimiter runs:

Openers, none closed: ** *** * ** *** **** ***** ****** ******* ********
Underscores, none closed: __ ___ _ __ ___ ____ _____ ______ _______ ________
*a*b*c* **a**b**c** ***a***b***c*** _a_b_c_ __a__b__c__

Intraword underscores: snake_case_name, file_name_extension,
a_much_longer_identifier_with_many_parts, __dunder__, single_.

Unbalanced openers: **never closed, *also never closed, *** mixing **
and * with nothing to pair them with.

Balanced-ish: **bold *italic* bold** ~~strike **bold strike**~~
***all three at once***.

Escapes: \*not emphasised\* \_nope\_ \\ and **\*escaped openers\***
inside real bold.

Backticks: `one`, ``two``, ```three```, a ` b` c `` d `` e,
```code with ` one tick```, and `unclosed code **is literal** here.

Links: [text](https://example.com/a(b)) and [empty]().
Then [dangling][nope] and [nope2] again.
Also [nested [brackets] inside](https://example.com/x).

Autolink-ish: <https://example.com/auto>, <not@a.link>, and
<http://example.com/edge>.

Strikethrough mixes: ~~a~~ ~~b**c**d~~ ~~~triple~~~ ~~*em strike*~~
and x~~y~~z for good measure.

Rule candidates:

---

----------

------------

Emphasis across punctuation: ***a***,b,c***d*** plus *a."* plus
**a,b** and no-space*em*after*em*more and ***a**b*c***.

Footnote-shaped but undefined: [^missing1] and [^missing2], one inside
**bold [^missing3] bold**.

Closing runs: *****a***** and ******b****** and *******c*******.

Reference-definition-shaped text stays literal in code spans:
`[used]: https://example.com/used "A definition"` and
`[unused]: https://example.com/unused 'Another definition'`.

Text that uses [inline links](https://example.com/used) instead, then
[inline links](https://example.com/unused) again for good measure.

Fence soup:

````
```
~~~
text between stray fences
```
~~~
````

And a last line so nothing dangles.
