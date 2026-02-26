+++
title = "fsh"
date = 2025-12-09
description = "A stack-based shell bringing concatenative programming to Unix command-line tools"
category = "Forth"
tags = ["Forth"]
github = "https://github.com/alexanderbrevig/fsh"
visibility = "public"
languages = "OCaml"
interest_score = 8

[extra]
post_tags = ["Forth"]
+++




Why this matters: Unix pipes are powerful but unintuitive if you're not already thinking linearly left-to-right. What if you could use the same stack-based thinking that makes Forth elegant? Push values, manipulate them, pass them through commands. The cognitive model is different—not linear pipelines, but value transformation on a stack.

The problem is that shell composition has inherent ordering constraints. With pipes, you write `cat file | grep pattern | sort`. The order is fixed by the data flow. If you want to reorder operations, you rebuild the entire pipeline. More fundamentally, pipes are a specific answer to the problem of chaining commands. What if you tried a different answer?

My idea is to bring stack-based (RPN) semantics to the Unix shell. Instead of pipes, you use a stack. Commands pop arguments from the stack and push results back. This means you can manipulate arguments without nested subshells or complex quoting.

I built `fsh` as a complete stack-based shell. It's not meant to replace bash—it's an exploration of what Unix command chaining looks like with different semantics. Commands automatically print their output *and* keep it on the stack for the next command. You get control flow (if/then/else, loops), custom word definitions, string operations, glob expansion, and a configurable prompt.

The interaction model is different:

```forth
/etc/legal cat free grep    \ → cat /etc/legal | grep free
5 3 > if greater . else smaller . then   \ → conditionals
lib 1 ls each cat then      \ → iterate and cat each file

: ll -la 1 ls ;              \ → define reusable words
eval lib ll                  \ → use them
```

The stack-aware prompt shows you what's ready to use: `fsh[2:1]>` means two string arguments and one output waiting. You can compose custom utilities with word definitions. The documentation walks from basic piping to FizzBuzz scripts written entirely in stack notation. This is experimental and probably not your daily shell. But it's a genuine exploration of *different ways to think about command composition*. It shows that pipes aren't the only answer.

## Why RPN?

Reverse Polish Notation might feel backwards at first, but it provides:

1. **Zero ambiguity** - no operator precedence rules
2. **Natural composition** - `f g` means "do f, then do g"
3. **Stack manipulation** - powerful reordering without variables
4. **No parentheses** - structure emerges from order

Once you internalize the flow, it becomes second nature.