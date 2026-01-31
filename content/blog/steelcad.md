+++
title = "steelcad"
date = 1769609202
category = "Rust"
tags = ["Rust"]
github = "https://github.com/alexanderbrevig/steelcad"
visibility = "private"
languages = "Rust"
interest_score = 0
+++

## About

Using steel language and vcad package to be a programmable cad
## Original README

# SteelCAD

A parametric CAD system combining [Steel](https://github.com/mattwparas/steel) (Scheme) and [vcad](https://github.com/ecto/vcad) (parametric CAD in Rust).

## Overview

SteelCAD allows you to define 3D models using Scheme, with support for:
- Parametric modeling with CSG operations
- Module system with git-based library loading
- Interactive REPL and batch compilation
- Export to STL, STEP, and other formats

## Project Structure

This is a Cargo workspace with the following crates:

- **steelcad-core**: Core library integrating Steel and vcad
- **steelcad-cli**: Command-line interface (REPL and compiler)

## Quick Start

```bash
# Start the REPL
cargo run --bin steelcad repl

# Compile a .scm file to STL
cargo run --bin steelcad compile examples/plate.scm --output plate.stl
```

## Example

```scheme
;; Load the core library
(require "steelcad/core.scm")

;; Create a simple cube
(define my-cube
  (-> (cube 10 10 10)
      (translate 5 0 0)))

;; Export to STL
(export-stl my-cube "output.stl")
```

## Development Status

🚧 **Early Development** - API is subject to change

## License

MIT OR Apache-2.0
