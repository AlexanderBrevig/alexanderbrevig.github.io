+++
title = "tinykeys"
date = 1691443437
category = "Hardware"
tags = ["Hardware", "Keyboard"]
github = "https://github.com/alexanderbrevig/tinykeys"
visibility = "public"
languages = "Unknown"
interest_score = 0
+++

## About

A superminimal choc keyboard _just_ for fun
## Original README

# tinykey

![tinykey design](img/tinykey.png)

A project just for fun, to see if I can type on it.

## The rough idea

The layout:

```
↑ ◐ ● ◑ ↹ ⌫      ↩ ␣ ◐ ● ◑ ↓
```

Imagine mapping the home row `ASDFG` to three keys, using chords.
Let's say we only have the `SDF` keys, the mapping will be like this:

```
S = A
SD = S
D = D
DF = F
F = G
```

So, if you wanted to type `SAD` then you would press `SD` `S` `D`. The grouped keys becomes a chord for your ring and middle finger.
Now then, what if you wanted to type `FEAST`?

Notice that `↑` on the pinky? It shifts the home row up.
Before things get crazy, let's rename `SDF` to `◐●◑`.

Now, then `FEAST` becomes:

```
F  E  A S  T
●◑ ↑● ● ◐● ↑◑ 
```

The same applies for the `↓` on the right pinky. This shifts the home row down.

So what about the thumbs then?

The left thumb has two keys, a TAB `↹` and a BACKSPACE `⌫`. 
The right thumb has RETURN `↩` and SPACE `␣`.
However, when held they enter other layers.

In my (still only slideware) config, I've mapped the numbers to the top row on the SPACE layer.
So, for me to type out `1337` it would be:

```
1   3   3   7
↑◐␣ ↑●␣ ↑●␣ ↑␣●◑
```

I will post my complete keymap when it's ready.

## Taking it further

We could also realize five keys by using timing and only two keys.

```
A S   D  F   G
◐ ◐-◑ ◑◐ ◑-◐ ◑
```

S and F has a roll, while D is _at the same time_ (within a confurable Δt ms).

Will try to play with this as well, and thus getting the layout down to 10 keys total.
