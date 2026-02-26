+++
title = "FXDev"
date = 2021-03-22
description = "breadboard psu and in/out with buffers, for prototyping guitar pedals on a breadboard"
category = "Hardware"
tags = ["Hardware", "Audio"]
github = "https://github.com/alexanderbrevig/FXDev"
visibility = "private"
languages = "Unknown"
interest_score = 5

[extra]
post_tags = ["Hardware", "Audio"]
+++




Why this matters: Guitar effects are where hardware engineering meets creativity. Breadboarding lets you experiment with circuits quickly, but the infrastructure overhead gets tedious. Every prototype needs the same power distribution, input buffering, output buffering, and bypass switching. You end up rebuilding these support circuits for every pedal you test.

The problem is that breadboard prototyping repeats the same boring infrastructure. You wire up a power supply rail, input/output buffers, a bypass switch, and *then* you finally get to the interesting part—the actual effect circuit. If you're testing multiple ideas, you do this setup five times. It's mechanical and error-prone.

My idea is to build a reusable module that handles all the infrastructure so you can focus on the effect circuit itself. Power distribution, buffered I/O, bypass switching—all on one plug-in board that sits on your breadboard. You add your effect circuit, test it, modify it, and move to the next idea without rewiring support circuitry.

I designed `FXDev` (also called `StompLab`) as a breadboard plugin module. It provides clean power rails (ground, input voltage 4-24V, and buffered Vin/2 for virtual ground), stereo input with optional buffers, stereo output with optional buffers, and bypass switching. Everything you need to prototype an effect without building the infrastructure each time.

The future plans show the trajectory: first the core module, then kits bundled with breadboard and 9V power, then pre-designed circuit blocks for common effects (tube screamer clone, big muff clone, metal zone tone stack). Eventually KiCAD files and box designs so people can take breadboard prototypes and move them to permanent enclosures.

This solves a real workflow problem for guitar effect designers and hardware hobbyists. Instead of wasting time on boilerplate circuitry, you spend time on what matters—exploring effect designs. It's a pragmatic tool for the builder's bench.
