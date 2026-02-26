+++
title = "literate-lsp"
date = 2026-01-30
description = "[WIP] An LSP that allows literate programming experience in markdown/typst/more"
category = "LSP"
tags = ["LSP"]
github = "https://github.com/alexanderbrevig/literate-lsp"
visibility = "public"
languages = "Rust"
interest_score = 8

[extra]
post_tags = ["LSP"]
+++




## About

[WIP] An LSP that allows literate programming experience in markdown/typst/more

## Article

The idea: extract code from markdown/typst documentation, delegate to language servers (forth-lsp, gopls, rust-analyzer, etc.), map results back. This turns docs from read-only into something more interactive—hover over code blocks and get completions, definitions, error checking without leaving the document. Works with 100+ LSPs, production-ready for Helix. Solves a real problem if you get tired of copying code out of documentation to actually test it.

