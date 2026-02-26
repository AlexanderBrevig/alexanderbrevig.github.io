+++
title = "patternlang"
date = 2025-10-29
description = "OCaml DSL for parametric sewing pattern generation"
category = "Sewing"
tags = ["Sewing"]
github = "https://github.com/alexanderbrevig/patternlang"
visibility = "private"
languages = "OCaml"
interest_score = 8

[extra]
post_tags = ["Sewing"]
+++

Why this matters: Sewing patterns are geometry problems dressed up as making. Most patterns are static—printed on paper with fixed measurements. But bodies vary. A pattern that fits size 8 doesn't fit size 16. Either you photocopy and adjust by hand, or you buy size-specific versions. What if patterns were parametric? What if you could input your measurements and generate a custom pattern?

The problem is that traditional sewing patterns don't adapt to individual variation. The solution has always been to offer multiple sizes, which is wasteful, or to expect sewers to grade (resize) patterns manually using arcane rules. This is friction that discourages people from garment construction—it feels technical and unfriendly.

My idea is a domain-specific language for sewing patterns. Declare measurements as parameters, define geometry in terms of those measurements, export to PDF. A single pattern definition generates infinite custom variations. You describe the collar as "neck circumference plus ease, offset by rise amount," and the pattern automatically adjusts for whoever is making it.

I built `PatternLang` as an _OCaml_ DSL that does exactly this. The language supports parametric measurements (body circumferences, lengths, ease allowances), geometric primitives (points, lines, curves, shapes), transformations (mirror, rotate, cut), and PDF export. Define a pattern once in a declarative style, and it generates customized PDFs for any set of measurements:

```
import StandardMeasurements as Std

Collar = piece {
  NeckCircumference = Std.NeckCircumference + EASE.x
  CollarHeight = 8cm
  CollarRise = 2cm

  CenterBackNeck = point(0, 0)
  CenterFrontNeck = point(NeckCircumference / 2, CollarRise)
  MidNeck = point(NeckCircumference / 4, CollarRise / 2)

  NeckSeam = curve from CenterBackNeck to CenterFrontNeck through MidNeck, smooth
  Shape = closed [NeckSeam, FrontSeam, TopEdge, BackSeam]
}
```

This is domain-specific tooling at its best—creative intersection of programming and making. It solves a real problem in garment construction: friction between flexible design and practical manufacturing. Instead of managing dozens of pattern sizes, you manage one parametric definition.