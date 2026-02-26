+++
title = "PicoFSM"
date = 2012-10-12
description = "Extremely lightweight FSM library, supports edge definitions in the form of TransitionToWhen(state,condition)"
category = "Old"
tags = ["Old"]
github = "https://github.com/alexanderbrevig/PicoFSM"
visibility = "public"
languages = "C#"
interest_score = 2
draft = true

[extra]
post_tags = ["Old"]
+++





## About

Extremely lightweight FSM library, supports edge definitions in the form of TransitionToWhen(state,condition)

## Article

Back in 2012, I made a minimal finite state machine library for C#. Nothing groundbreaking—FSMs are a solved problem—but this was a clean, no-nonsense implementation. You define states with Enter/Update/Exit callbacks, wire up transitions, and let it run. The `TransitionToWhen` method is a neat touch for conditional transitions, but ultimately this is the kind of utility that everyone eventually writes for themselves when they need one.

It got packaged up on NuGet and used by a few people, which was cool. If you're doing state management in C# and want something minimal without dragging in a heavy framework, it does the job. But realistically? There's nothing here that would surprise anyone familiar with FSM patterns.

