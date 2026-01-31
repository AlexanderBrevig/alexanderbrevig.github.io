+++
title = "typst-forth"
date = 1769559994
category = "Forth"
tags = ["Forth"]
github = "https://github.com/alexanderbrevig/typst-forth"
visibility = "public"
languages = "Typst"
interest_score = 0
+++

## About

Forth syntax highlight for Typst
## Original README

# codeforth

Forth syntax highlighting for [Typst](https://typst.app/).

![](example.png)

## Installation

````typst
#import "@preview/codeforth:0.1.0": forth-syntax
#set raw(syntaxes: forth-syntax)

```forth
: fun 1 2 + ;
```
````

## Credits

- Syntax patterns informed by [mitranim/sublime-forth](https://github.com/mitranim/sublime-forth)

## License

MIT
