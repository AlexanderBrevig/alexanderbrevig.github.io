+++
title = "voltara"
date = 1762562702
category = "Hardware"
tags = ["Hardware", "Audio / live performance"]
github = "https://github.com/alexanderbrevig/voltara"
visibility = "private"
languages = "Rust"
interest_score = 0
+++

## About

Wearable midi live show idea
## Original README

# Voltara - Wearable Drum Trigger System

A wearable gesture-based MIDI drum trigger system for live electronic music performance.

## Overview

This system enables 3 performers to trigger drum sounds through wrist and ankle strikes. Wearable accelerometer modules detect impacts and transmit MIDI over a wired RS-485 network to a central stage box, which outputs USB MIDI to a DAW.

**Status:** Phase 1 - Proof of Concept

## Project Structure

```
voltara/
├── hardware/          # PCB designs (KiCad) and enclosures
├── firmware/          # Embedded Rust code for RP2040
├── tests/             # Integration and hardware tests
├── tools/             # Development utilities
├── docs/              # Assembly, testing, calibration guides
├── DESIGN.md          # Complete technical specification
└── CLAUDE.md          # AI assistant context file
```

## Quick Start

### Firmware Development

```bash
cd firmware

# Install Rust toolchain (automatically configured via rust-toolchain.toml)
rustup show

# Build all firmware
cargo build --release

# Build specific module
cargo build -p extremity-module --release

# Run tests
cargo test -p voltara-common
```

### Hardware

See `hardware/README.md` for KiCad project details and PCB fabrication.

## Key Features

- **Wearable sensors** on wrists and ankles (4 per performer)
- **Wired connection** via CAT5 cable harness
- **Low latency** MIDI output (<10ms target)
- **Future-proof** architecture for gesture recognition
- **Theatrical integration** with plug-in ritual as part of performance

## Documentation

- `DESIGN.md` - Full system architecture and specifications
- `CLAUDE.md` - Quick context for AI-assisted development
- `docs/` - Assembly guides, testing procedures, calibration

## System Architecture

```
[Extremity Modules] (12 total: 4 per performer × 3)
      ↓ RS-485 + RTS over CAT5 EtherCon
[Belt Pack] (3 total: 1 per performer)
      ↓ RS-485 + Monitors + Instrument over CAT5 EtherCon
[Stage Box] (1 unit)
      ↓ USB MIDI
[DAW]
```

## Hardware Components

- **MCU:** RP2040 (Raspberry Pi Pico)
- **Sensor:** LIS3DH or ADXL345 (±16G accelerometer)
- **Communication:** RS-485 differential signaling
- **Power:** 12V power over CAT5 cables

## License

MIT
