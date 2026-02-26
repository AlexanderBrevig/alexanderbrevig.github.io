+++
title = "Calc4th"
date = 2026-01-31
description = "Forth-based hardware lab multitool - RPN calculator interface for I2C/SPI/UART/GPIO on RP2040"
category = "Forth"
tags = ["Forth"]
github = "https://github.com/alexanderbrevig/Calc4th"
visibility = "private"
languages = "Typst"
interest_score = 8

[extra]
post_tags = ["Forth"]
+++




Why this matters: Embedded hardware debugging is one of those problems that looks simple until you're sitting in the lab with unknown chips and no way to quickly test them without rebuilding your entire test setup each time. Most people reach for a multimeter or oscilloscope, but those tools only show you voltage levels. You're building something that bridges the gap between general-purpose instruments and domain-specific hardware knowledge.

The problem is that hardware prototyping involves a lot of trial-and-error communication with unfamiliar chips. You write a test on your laptop, upload it, watch it fail, change the code, rebuild, re-upload. Each cycle takes minutes. Meanwhile you're juggling wires on a breadboard, squinting at datasheets, trying to remember if this device uses 7-bit or 8-bit addressing. It's tedious and error-prone.

My idea is to bring an interactive REPL directly into the hardware—a Forth interpreter running on an RP2040 with a button matrix and display. Instead of writing code on your computer, you probe devices *from the device itself*. Scan the I2C bus with a few button presses. Read a register. Toggle GPIO pins. Define custom Forth words and trigger them from hardware buttons.

I built `Calc4th` as a portable hardware lab multitool. It combines an RPN calculator interface with low-level hardware access. You get stack-based computation as your working memory, but more importantly, you get immediate feedback without the edit-compile-upload cycle. The buttons map to Forth words, and the OLED shows you results in real time.

The workflow becomes interactive. You can:
```forth
i2c-scan              \ discover I2C addresses at a button press
$50 i2c-addr!         \ set the address to probe
0 i2c-read .          \ read register 0 and print it
: read-temp 0 i2c-read 2/ ;   \ define a custom word for common operations
```

The hardware itself is straightforward—an RP2040, a 4×5 button matrix for input, a 128×64 OLED display for output. The real work is the Forth interpreter and the hardware abstraction layer that lets you talk to I2C, SPI, UART, and GPIO all from the stack. I've even included test routines so you can verify everything works before you start probing unknown devices.

This solves a real problem in the lab. Instead of swapping test rigs or rebuilding support circuitry for each chip, you have one portable tool that adapts to whatever you're working with. The PDF tutorial walks you through the entire workflow. This is the kind of tool that makes people in hardware development say "wait, why isn't this more common?"
