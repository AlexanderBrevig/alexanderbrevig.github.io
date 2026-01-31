+++
title = "LEB-Synth"
date = 1618182193
category = "Hardware"
tags = ["Hardware", "Audio / live performance"]
github = "https://github.com/alexanderbrevig/LEB-Synth"
visibility = "private"
languages = "Typst"
interest_score = 0
+++

## About

11 Eurorack modules
## Original README

exec(open('\\\\wsl.localhost\\\\Ubuntu\\\\home\\\\ab\\\\github.com\\\\leb-synth\\\\scripts\\\\eurorack_panel.py').read())

git clone --recurse-submodules git@github.com:AlexanderBrevig/ABKiCADLibs.git

# Rules

```
---
PANEL: 128.5mm
17mm to center from top
| 7 [jack] 13 [jack]
|16 [pot] 26 [pot] 20 [jack]
```

# Ideas

Probe plug, SP-3541 with resistor and bipolar led to show output

# Maintaining ABKiCADLibs

## Update

    git submodule update --remote --merge

# Modules & Synth

## Voice 36HP

- LEB-VCO 10HP
- LEB-VCA 8HP
- LEB-FIL 10HP
- LEB-ADSR 8HP

## Voice extra 18HP

- LEB-SUB 4HP
- LEB-LFO 10HP
- LEB-MULT 4HP

## Mix & Out 16HP

- LEB-OUT 8HP
- LEB-MIX 8HP

## Extras 18HP

- LEB-NOISE 10HP
- LEB-FOLD 4HP
- LEB-DIST 4HP

## Synth

4HP: 4
8HP: 4
10HP: 4

Big voice: 54
All modules: 88
Two voices: 124

Standard widths: 84 and 104 HP

# Notes

- 3U height
- 84 HP width = (5.08 \* 84) = 426.72
- Power
  - 1 A +/-12V + 0.5A 5V

POWER ?
MIDI (usb, MIDI, 2 outs (mono/poly), velocity, gate)

OSC
NOISE (blue, white, pink, red)
FIL

MULT
GLIDE
ENV (attack, release, trigger gate)

LFO
AMP (gain, cvamount - in cv out)
MIX
REV dual?
DLY

# Parts

https://www.thonk.co.uk/ (alpha pots)
https://oddvolt.com/ (jacks)
https://www.exploding-shed.com/
https://shop.befaco.org/167-misc

## Chosen

Jack: PJ301BM, new name WQP518MA-BM
StereoJack: PJ301CM
Pot: Alpha 9mm vertical D shaft

## Switches

SPDT on-on A101SYCQ04 or A101SYCB04
SPDT on-off-on A103SYCQ04 or A103SYCB04
DPDT on-on A201SYCQ04 or A201SYCB04

SPDT on-on 1MS1T2B4M7RE
SPDT on-off-on 1MS3T2B4M7RE
DPDT on-on 1MD1T2B4M7RE

## Trim

Multi Bourns 3296W
Horiz Bourns 3386G
Top 3/8" trim Bourns 3386F

## Pots

Alpha 9mm series

## Diodes

1N4148
1N4148WS - SOD-323
BAT42
BAT42WS
1N4004 - power THT
BZX384-C(5V1) Zener - SOD-323
3mm LEDs

## Capacitors

> For small values of capacitors, for example, 1uF, 2.2uF & 4.7uF, then go for 50V or 63V devices. While for bigger values like 10uF, 22uF, 47uF and 100uF go for 35V.

### Ceramic

Decade box

C0G / low-K / NP0
25V
+/- 5% min (pref 1%)
0805
1206

X7R good enough for powersupply
08055C104JAT2A - power 100nF

### Polyester

1nF and 2.2uF. 5mm (0.2'').
50V, 63V or 100V

## Resistors

CRCW0805(33K)0FKEA - get a kit of values

Decade box

## IC

### Opamp

TL072C - TL072ACD
4558 - MC4558CN - RC4558P (dual bipolar)
LM13700 (transconductance)

### That (farnell)

THAT2180LC
THAT300P
THAT320P
THAT340P

### Switch

DG403DJ

### Special

SSM2164 - V2164D (VCA)
MN3207 (bucket brigade delay)
AS3310 (envelope generator) electronic druid, banzai, thonk

## Transistor

BC549
BC559
BC850 (use B or C, eg BC850CLT1G)
BC860

MMBFJ201
J201
J112

## Misc

Ferrite beads
FK237 SA220 O - TO-220 heatsink (farnell/distrelec)

### Pin Headers

10 way socket 2212S-10SG-85
10 way header 2211S-10G
6 way socket 2212S-06SG-85
6 way header 2211S-06G
4 way socket 2212S-04SG-85
4 way header 2211S-04G

### Connectors

Molex KK
Molex KK crimper

112APCX - jack
PJ398SM - jack (3.5mm) + nut flat washer

WQP-PJ301M-12 jack standard
