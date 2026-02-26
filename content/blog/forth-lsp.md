+++
title = "forth-lsp"
date = 2023-06-06
description = "A Language Server Protocol implementation for Forth, bringing modern IDE features to Forth development."
category = "Forth"
tags = ["Forth"]
github = "https://github.com/alexanderbrevig/forth-lsp"
visibility = "public"
languages = "Rust"
interest_score = 8

[extra]
post_tags = ["Forth"]
+++




Why this matters: Forth is a powerful, minimal language from 1971 that never got proper IDE tooling. Most Forth programmers edit in whatever text editor they have at hand—no autocomplete, no hover documentation, no "go to definition". It's like doing serious programming with one hand tied behind your back. The Language Server Protocol standardizes how editors talk to language services, so if you implement LSP for Forth, you instantly get IDE support in VS Code, Helix, Neovim, Emacs—any editor that understands LSP.

The problem is that Forth remains niche partly because it's hard to work with. There's no ecosystem of development tools. You define words, but your editor has no idea where they're defined. You reference a built-in word, but you have to hunt through documentation to find its stack effect. Every Forth developer reinvents their own workflow.

My idea is to implement a complete Language Server Protocol for Forth. This means the editor can ask: "What's the type signature of this word? Where is it defined? What else uses it? Can I rename it safely?" The answer is the same across all editors, all workflows, all platforms.

I built `forth-lsp` as a full-featured LSP implementation in _Rust_. It's cargo-installable and published on crates.io. The feature set is comprehensive: hover documentation for both built-in words and user definitions, completion suggestions, go-to-definition (jump to where a word is defined), find-references (show all usages), rename (refactor symbols across your workspace), document symbols (outline view), diagnostics (real-time error detection for undefined words), signature help (see stack effects while typing), and formatting.

Install it with a single command and configure your editor to use it:

```bash
cargo install forth-lsp
```

Then point your editor to `forth-lsp`. The CI passes, tests run cleanly. If you're a Forth developer who's gotten used to working without IDE support, you might actually find this useful. Whether it's "finally" tooling worth having or just another attempt at making Forth less painful to work with depends on what you're trying to do, but it's there if you want it.
