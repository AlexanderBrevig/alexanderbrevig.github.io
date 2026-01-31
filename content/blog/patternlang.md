+++
title = "patternlang"
date = 1761718953
category = "Sewing"
tags = ["Sewing"]
github = "https://github.com/alexanderbrevig/patternlang"
visibility = "private"
languages = "OCaml"
interest_score = 0
+++

## About

Programmatic pattern generation
## Original README

# Pattern Language

A domain-specific language for creating digital, personalizable sewing patterns.

## Overview

Pattern Language is a DSL that allows pattern makers to create parametric sewing patterns using a clean, declarative syntax. The language supports:

- Parametric measurements that can be customized per individual
- Geometric primitives (points, lines, curves, shapes)
- Transformations (mirror, rotate, cut)
- Pattern piece composition and reuse through imports
- Export to PDF for printing and cutting

## Building

Requires OCaml >= 4.14 and dune >= 3.0.

```bash
# Install dependencies
opam install . --deps-only

# Build
dune build

# Run tests
dune test

# Install
dune install
```

## Usage

```bash
patternlang mypattern.pattern
```

## Example

```
import StandardMeasurements as Std

Collar = piece {
  // Measurements
  NeckCircumference = Std.NeckCircumference + EASE.x
  CollarHeight = 8cm
  CollarRise = 2cm

  // Points
  CenterBackNeck = point(0, 0)
  CenterFrontNeck = point(NeckCircumference / 2, CollarRise)
  MidNeck = point(NeckCircumference / 4, CollarRise / 2)

  // Lines
  NeckSeam = curve from CenterBackNeck to CenterFrontNeck through MidNeck, smooth

  // Shape
  Shape = closed [NeckSeam, FrontSeam, TopEdge, BackSeam]
}
```

## Documentation

- [Design Document](DESIGN.md) - Language design decisions and rationale
- [Grammar Specification](grammar.ebnf) - Formal EBNF grammar

## License

MIT
