+++
title = "ChessML"
date = 2025-10-21
description = "A for fun project to write a chess engine that's better than me using OCaml"
category = "Chess"
tags = ["Chess"]
github = "https://github.com/alexanderbrevig/ChessML"
visibility = "public"
languages = "OCaml"
interest_score = 9

[extra]
post_tags = ["Chess"]
+++




Why this matters: Chess engines are a proving ground for computer science fundamentals—move generation, search algorithms, heuristic evaluation, transposition tables. Building one teaches you about algorithm design, memory efficiency, and how to think about games as search problems. Plus, you get to play against something you built.

Learning a new language while building something non-trivial is hard. Most people learn syntax and write toy programs. I'm building a serious chess engine in _OCaml_ while learning the language itself. The problem is balancing ambition with learning curve—do I build something toy and boring, or something that actually works?

My idea is to build a complete chess engine: efficient move generation with bitboards, alpha-beta search with transposition table caching, opening book support from real games, and full UCI protocol support so you can play it in any chess GUI. Make it fast enough to be interesting, test it thoroughly, document the design decisions.

I built `ChessML` using OCaml's functional paradigm and pattern matching for board representation. Bitboards let me generate moves efficiently. The alpha-beta search with transposition tables dramatically cuts the search space—you cache positions you've already evaluated so you don't recompute them. The opening book comes from parsing PGN files (1M+ position-move combinations from 9M+ real games), so the engine doesn't play randomly in the opening.

The test suite is comprehensive—quick tests run in ~11 seconds, deep validation takes ~2 minutes. Release builds get ~40% faster with optimizations (120K nodes per second vs 86K in debug mode). Performance matters when you're searching millions of positions. The UCI and XBoard protocol support means you can plug it into Arena, ChessBase, or any standard chess GUI and get thoroughly embarrassed by a program written by someone still learning the language:

```bash
# Build for performance
dune build --profile=release

# Run quick tests for development feedback
dune runtest

# Play against it and prepare to lose
cutechess-cli -engine name=ChessML cmd=./_build/default/bin/chessml_uci.exe proto=uci
```

It probably has bugs. There are definitely things I did that a real OCaml programmer would recoil at. But it plays chess better than I do, which was the whole point, and that's good enough for me.
