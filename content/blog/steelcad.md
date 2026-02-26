+++
title = "steelcad"
date = 2026-01-28
description = "Parametric CAD system combining Steel (Scheme) and vcad (parametric CAD in Rust)"
category = "Rust"
tags = ["Rust"]
github = "https://github.com/alexanderbrevig/steelcad"
visibility = "private"
languages = "Rust"
interest_score = 7
draft = true

[extra]
post_tags = ["Rust"]
+++





## About

Parametric CAD system combining Steel (Scheme) and vcad (parametric CAD in Rust)

## Article

SteelCAD integrates a Scheme interpreter (Steel) with the vcad parametric CAD library. You define 3D models using Scheme, with CSG operations, module system with git-based library loading, and export to STL/STEP. It's early-stage ("API subject to change"), but the architecture shows real thinking about how to make CAD programmable in a high-level language. Most CAD is GUI-driven; making it script-based opens different possibilities. Interesting exploration.

