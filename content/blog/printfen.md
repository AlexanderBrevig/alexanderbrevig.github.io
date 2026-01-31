+++
title = "printfen"
date = 1754657045
category = "Chess"
tags = ["Chess"]
github = "https://github.com/alexanderbrevig/printfen"
visibility = "public"
languages = "OCaml"
interest_score = 0
+++

## About

Simple print utility for Chess FEN notation
## Original README

# Print FEN

Compile and run

```sh
ocamlc -o printfen printfen.ml
./printfen "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
```

Observe output

```sh
r n b q k b n r
p p . p p p p p
. . . . . . . .
. . p . . . . .
. . . . P . . .
. . . . . . . .
P P P P . P P P
R N B Q K B N R

Turn:            w
Castling:        KQkq
En passant:      c6
Halfmove clock:  0
Fullmove number: 2
```
