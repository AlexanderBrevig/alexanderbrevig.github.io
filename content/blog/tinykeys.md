+++
title = "tinykeys"
date = 2023-08-07
description = "A superminimal choc keyboard - 10 keys with chord-based layout"
category = "Hardware"
tags = ["Hardware", "Keyboard"]
github = "https://github.com/alexanderbrevig/tinykeys"
visibility = "public"
languages = "Unknown"
interest_score = 8

[extra]
post_tags = ["Hardware", "Keyboard"]
+++

Why this matters: Keyboard design is a rabbit hole where enthusiasts obsess over mechanical switches, stabilizers, and layouts. Most keyboards are variations on a theme—QWERTY, Dvorak, Colemak. But the premise of a keyboard—one key per function—is actually arbitrary. What if you used chords? What if three keys could generate the entire character set through combination?

The problem is that keyboards are designed for historical reasons (typewriter key count, mechanical constraints). Modern keyboards inherit this inertia even though computers could handle any input scheme. Most people never question whether 100+ keys are necessary.

My idea is a radically minimal keyboard: 10 physical keys total. Three main keys (representing different regions—left, center, right), shift layers for case changes, and thumb keys for modifiers and common functions. Use chords (simultaneous key presses) to generate the full character set.

I built `tinykeys` as a 10-key keyboard design. The layout is:

```
↑ ◐ ● ◑ ↹ ⌫      ↩ ␣ ◐ ● ◑ ↓
```

The three main keys `◐●◑` combine to produce different characters. With timing or chords, you expand capacity without adding keys:

```
◐ = A
◐● = S  (both pressed simultaneously)
● = D
●◑ = F
◑ = G
```

Add the shift keys `↑` and `↓` to change rows, and you layer the character set. Numbers and symbols live on different layers. The thumb keys handle entry and spacing. Is this practical for daily typing? Probably not. Is it an interesting exploration of whether keyboards *must* have 100+ keys? Absolutely.

This pushes the boundaries of what "minimal keyboard" means. It's the kind of project that makes other keyboard enthusiasts say "that's wild, let me try it"—it demonstrates that core assumptions about input devices are actually design choices, not constraints. Novel ideas deserve credit.