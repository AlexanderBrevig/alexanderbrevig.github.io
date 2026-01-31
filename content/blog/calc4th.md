+++
title = "Calc4th"
date = 1769873678
category = "Forth"
tags = ["Forth"]
github = "https://github.com/alexanderbrevig/Calc4th"
visibility = "private"
languages = "Typst"
interest_score = 0
+++

## About

Forth book and it's hardware and code
## Original README

# Calc4th: A Forth Lab Tool

A portable hardware lab multitool built around an RPN calculator interface. Write Forth interactively on an RP2040 to probe I2C/SPI buses, monitor UART, toggle GPIO, and test unknown devices in the field.

## Features

- **RPN Calculator** – stack-based computation as your working memory
- **I2C Scanner & Register Explorer** – probe unknown devices and read/write registers
- **SPI Communication** – talk to SPI peripherals
- **UART Monitor** – serial debugging and passthrough
- **GPIO Control** – digital I/O and PWM
- **Test Vector Injection** – define custom Forth words and trigger them from buttons

## Hardware

- Raspberry Pi Pico (RP2040)
- 4×5 button matrix (20 buttons)
- 128×64 OLED/LCD display
- I2C/SPI/UART/GPIO breakout headers

## Quick Start

Build and flash the Forth image:
```bash
./build.fish
```

Connect via USB for the Forth REPL, or use the button interface for portable operation.

## Usage Example

```forth
i2c-scan              \ discover I2C addresses
$50 i2c-addr!
0 i2c-read .          \ read register 0
: read-temp 0 i2c-read 2/ ;   \ define custom word
```

## Documentation

See `Calc4th.pdf` for the full tutorial.
