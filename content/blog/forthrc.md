+++
title = "forthrc"
date = 2026-01-28
description = "Forth controlled programmatic resistor and capacitor, for experimentation and hot swap of values"
category = "Hardware"
tags = ["Hardware"]
github = "https://github.com/alexanderbrevig/forthrc"
visibility = "private"
languages = "Forth"
interest_score = 9

[extra]
post_tags = ["Hardware"]
+++




Why this matters: Circuit prototyping on a breadboard involves a lot of component swapping. You breadboard a filter, and it needs a different resistor value. You grab parts from your bin, measure, try again. Repeat fifty times. What if you could dial the value in from software instead?

The problem is that component values are fixed. If you breadboard a filter and the frequency is slightly off, you need to unsolder and swap resistors. This kills iteration speed. You end up building "component kits" with multiple resistor values pre-assembled, which wastes parts and bench space.

My idea is to use digital potentiometers (variable resistors controlled by SPI) and switched capacitor arrays so I can change component values electronically. Combine them on an RP2040 with a Forth REPL and I get interactive circuit tuning from software.

I built `forthrc` as a programmable resistor-capacitor tool. The hardware uses three MCP4131 digital potentiometers (100kΩ each) for resistance, and a binary-weighted capacitor array (1nF to 512nF) controlled by analog switches. An opamp multiplier extends the capacitance range. Everything is wired to an RP2040 running Zeptoforth.

The Forth interface is intuitive. Set resistance with human-readable units:

```forth
\ Set Port A to 47 kiloohms
47 KOHMS PORTA!

\ Set capacitance to 470nF
470 NF CAP-NF!

\ Read back current values
PORTA@ raw>kohms . ." kΩ" cr
```

The tool includes hardware test routines so you can verify everything works before you start prototyping. The documentation covers pin connections, troubleshooting, and a step-by-step path from minimal testing (one resistor channel) to the full system. This solves a real pain point in circuit prototyping—instead of hunting through bins for the right resistor, you dial it in from the REPL. It's a creative, pragmatic solution to a problem every hardware designer faces.

# Programmable Component Tool

An RP2040 + Zeptoforth based tool for interactive circuit prototyping with programmable resistors and capacitors.

## Quick Start

### Prerequisites
- Raspberry Pi Pico (RP2040)
- Zeptoforth installed on the Pico (from `../zeptoforth`)
- Serial terminal connected at 115200 baud

### Hardware Needed
- 3× MCP4131-104 digital potentiometers (100kΩ, 8-bit, SPI)
- 10× capacitors for binary array: 1nF, 2nF, 4nF, 8nF, 16nF, 32nF, 64nF, 128nF, 256nF, 512nF
- 2× ADG1414 (or CD4066) analog switches for capacitor array
- 1× MCP6002 op-amp for capacitance multiplier
- Breadboard and wires

### Loading the Software

1. **Connect to Zeptoforth REPL:**
   ```bash
   screen /dev/ttyACM0 115200
   # or
   minicom -D /dev/ttyACM0 -b 115200
   ```

2. **Load the component tool:**
   ```forth
   #include component-tool-zeptoforth.fs
   ```

3. **Initialize hardware:**
   ```forth
   init-tool
   ```

4. **Verify it works:**
   ```forth
   help-tool
   show-ports
   ```

## Usage Examples

### Setting Resistances

```forth
\ Set Port A to 47kΩ
47 KOHMS PORTA!

\ Set Port B to 1500Ω (1.5kΩ)
1500 OHMS PORTB!

\ Set Port A to raw value (0-255)
128 PORTA!

\ Read current value
PORTA@ .                 \ Raw value
PORTA@ raw>kohms . ." kΩ" cr
```

### Setting Capacitances

```forth
\ Set nF capacitor to 470nF
470 NF CAP-NF!

\ Set µF multiplier (raw value)
128 CAP-UF!

\ Read values
CAP-NF@ raw>nf . ." nF" cr
```

### Display Current Settings

```forth
show-ports
\ Output:
\ Port A: 128 raw = 50 kΩ
\ Port B: 64 raw = 25 kΩ
\ Cap nF: 470 raw = 470 nF
\ Cap µF: 128 raw
```

## Pin Connections

### SPI (Digital Pots)
```
RP2040 Pin    Function         MCP4131 Pin
==========    ========         ===========
GP19          MOSI             SI (3)
GP18          SCK              SCK (2)
GP16          MISO             SO (4) - optional
GP17          CS0              CS (1) - Port A
GP20          CS1              CS (1) - Port B
GP21          CS2              CS (1) - Cap µF
3.3V          VDD              VDD (8)
GND           GND              VSS (4), GND (7)

MCP4131 Connections:
- Pin 5 (PB0): Terminal B
- Pin 6 (PA0): Wiper (output)
- Pin 7 (GND): Terminal A (ground reference)
```

### Capacitor Switches (nF Range)
```
RP2040 Pin    Function         Analog Switch
==========    ========         =============
GP0           1nF switch       Control
GP1           2nF switch       Control
GP2           4nF switch       Control
GP3           8nF switch       Control
GP4           16nF switch      Control
GP5           32nF switch      Control
GP6           64nF switch      Control
GP7           128nF switch     Control
GP8           256nF switch     Control
GP9           512nF switch     Control
```

## Hardware Testing

Before using the full component tool, test your hardware:

```forth
\ Load test routines
#include hardware-test.fs

\ Test CS pins (watch with LED or scope)
test-cs

\ Test capacitor switches
test-caps

\ Test SPI (requires MOSI->MISO jumper)
test-spi

\ Test MCP4131 chips
test-pots
```

## Breadboard Test Setup

**Minimum viable test (1 resistor port):**

1. Wire up SPI pins and one MCP4131 on GP17 (CS0)
2. Connect multimeter between wiper (pin 6) and terminal A (pin 7)
3. Load software and run:
   ```forth
   init-tool
   0 PORTA!       \ Should read ~75Ω
   128 PORTA!     \ Should read ~50kΩ
   255 PORTA!     \ Should read ~100kΩ
   ```

## Resistor Values

For 100kΩ MCP4131 (model -104):
- Raw value 0 = ~75Ω (wiper resistance)
- Raw value 128 = ~50kΩ
- Raw value 255 = ~100kΩ

For 10kΩ MCP4131 (model -103):
- Raw value 0 = ~75Ω
- Raw value 128 = ~5kΩ
- Raw value 255 = ~10kΩ

**To use different values:** Edit `PORTA-MAX-OHMS` in the source code.

## Capacitor Multiplier Circuit

```
IN ──[1kΩ]──┬─── OP+ (buffer)
            │
            └──[MCP4131]─┬─── OP-
                         │
                    [10nF C0G]
                         │
OUT ─────────────────────┴─── OP_OUT
```

Effective capacitance = 10nF × (R_pot / 1kΩ)
- Min: ~750pF (R_pot = 75Ω)
- Max: ~1µF (R_pot = 100kΩ)

For full 10µF range, add a switchable 100kΩ multiplier resistor.

## Files

- `CLAUDE.md` - Complete design documentation
- `component-tool-zeptoforth.fs` - Main implementation (load this!)
- `hardware-test.fs` - Hardware verification tests
- `component-tool.fth` - Generic Forth template (for reference)
- `README.md` - This file

## Troubleshooting

**"No response from MCP4131"**
- Check wiring (especially CS, MOSI, SCK)
- Verify 3.3V power to MCP4131
- Try `test-pots` to diagnose
- Check SPI configuration (1MHz, Mode 0)

**"Resistance values incorrect"**
- Verify MCP4131 model (-103 = 10kΩ, -104 = 100kΩ)
- Update `PORTA-MAX-OHMS` constant if needed
- Measure wiper resistance (should be ~75Ω at value 0)

**"Capacitor switches not working"**
- Check GPIO configuration (GP0-GP9 as outputs)
- Verify analog switch wiring and power
- Test individual switches with `test-caps`

## Next Steps

1. Build and test one resistor channel
2. Add remaining resistor channels
3. Build and test capacitor nF array
4. Build capacitor multiplier circuit
5. Design PCB for permanent tool
6. Add calibration routine
7. Optional: Add OLED display and rotary encoder

## Resources

- [MCP4131 Datasheet](https://www.microchip.com/en-us/product/MCP4131)
- [Zeptoforth Documentation](https://github.com/tabemann/zeptoforth)
- [RP2040 Datasheet](https://datasheets.raspberrypi.com/rp2040/rp2040-datasheet.pdf)

## License

MIT License - feel free to use, modify, and share!
