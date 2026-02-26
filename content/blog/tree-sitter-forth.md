+++
title = "tree-sitter-forth"
date = 2023-06-06
description = "Forth grammar for tree-sitter - fast, incremental parsing for syntax highlighting and code analysis."
category = "Forth"
tags = ["Forth"]
github = "https://github.com/alexanderbrevig/tree-sitter-forth"
visibility = "public"
languages = "JavaScript"
interest_score = 7

[extra]
post_tags = ["Forth"]
+++

Why this matters: Tree-sitter is a parser library that generates fast, incremental parsers from grammar definitions. Unlike traditional parsers that re-parse the entire file on each edit, tree-sitter updates only changed regions. This makes it ideal for editor integration—instant syntax highlighting, real-time error detection, code navigation, all without lag.

The problem is that tree-sitter grammars exist for popular languages, but Forth was missing. Forth is niche, so nobody bothered. That means any editor using tree-sitter had no native Forth support. You'd either edit in a language-agnostic editor with no highlighting, or use old-school regex-based highlighting that's crude and slow.

My idea is a complete tree-sitter grammar for Forth that handles all the language's quirks. Forth numbers come in multiple forms (decimal, hex with `$` or `0x`, binary with `%`, octal with `&`, characters, floats, double-cell numbers). Comments have their own syntax (line comments with `\`, block comments, stack effect comments). Strings use different prefixes depending on storage (`s"`, `c"`, `."` for different behaviors). Words are categorized semantically—control flow, operators, I/O, core words.

I built `tree-sitter-forth` as a comprehensive grammar that captures all of this. The grammar distinguishes between word categories for better syntax highlighting:

- **Control Flow**: `IF/THEN/ELSE`, `BEGIN/UNTIL`, `DO/LOOP`, `CASE/OF`
- **Operators**: Stack manipulation (`DUP`, `SWAP`, `ROT`), arithmetic, logic
- **I/O**: Input/output words (`.`, `EMIT`, `KEY`, `ACCEPT`)
- **Core**: Defining words, memory operations, compilation

The implementation ships with Node.js and _Rust_ bindings, so it works across the ecosystem. Combined with `forth-lsp`, you get the pieces for IDE support in a tree-sitter/LSP-compatible editor: syntax highlighting, go-to-definition, completions, hover documentation. Whether this actually changes how people work with Forth or just sits quietly in a GitHub repository is hard to say, but at least the infrastructure is there if someone wants it.