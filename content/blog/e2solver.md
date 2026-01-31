+++
title = "e2solver"
date = 1763386900
category = "Eternity"
tags = ["Eternity"]
github = "https://github.com/alexanderbrevig/e2solver"
visibility = "private"
languages = "Rust"
interest_score = 0
+++

## About

For fun experiment Eternity 2
## Original README

# Eternity 2 Solver Visualizer

An interactive egui-based visualizer for the Eternity 2 puzzle, built to work with Wayland through WSL2.

## Features

- **Interactive Board Visualization**: View the 16x16 puzzle board with all placed pieces
- **Zoom & Pan**: Zoom in/out with buttons and pan by dragging the board
- **Hover Tooltips**: Hover over any piece to see detailed information:
  - Position (row, col)
  - Piece ID
  - Rotation angle
  - Edge pattern IDs (North, East, South, West)
- **Pattern Rendering**: Edge patterns are rendered with distinct colors and can load SVG textures
- **Wayland Support**: Optimized for Wayland display server in WSL2

## Building & Running

### Quick Start

```bash
# Make the run script executable (first time only)
chmod +x run.sh

# Run the visualizer
./run.sh
```

### Manual Build

```bash
cargo build --release
./target/release/e2solver
```

## WSL2 Configuration

The `run.sh` script sets up the optimal environment for running egui on Wayland through WSL2:

- Sets `WAYLAND_DISPLAY` for Wayland protocol
- Configures `XDG_RUNTIME_DIR` for runtime files
- Enables software rendering fallback if GPU acceleration has issues
- Disables vsync for better performance in virtual environments

## Controls

- **Zoom In/Out**: Click the ➕/➖ buttons
- **Reset View**: Click the "Reset" button to restore default zoom and position
- **Pan**: Click and drag anywhere on the board
- **Inspect Piece**: Hover over any piece to see its details

## Data Files

The visualizer loads data from:

- `data/eternity2_256.csv` - Piece definitions (ID and edge patterns)
- `data/eternity2_256_all_hints.csv` - Known piece placements
- `data/patterns/*.svg` - Edge pattern graphics (optional)

## Technical Details

- **Framework**: eframe 0.29 with egui
- **Rendering**: OpenGL via glow backend
- **Display**: Native Wayland support with X11 fallback
- **Image Processing**: resvg/usvg for SVG pattern loading

## Project Structure

```
e2solver/
├── src/
│   ├── main.rs          # Application entry point and data loading
│   └── visualizer.rs    # Board rendering and UI logic
├── data/
│   ├── eternity2_256.csv
│   ├── eternity2_256_all_hints.csv
│   └── patterns/        # SVG pattern files
├── Cargo.toml
├── run.sh              # WSL2-optimized launcher script
└── README.md
```

## Dependencies

- **egui/eframe**: Immediate mode GUI framework
- **resvg/usvg**: SVG rendering
- **tiny-skia**: Software rasterization
- **csv**: CSV file parsing
- **serde**: Data serialization

## Development

The visualizer is designed to be extended with:

- Solver algorithms
- Step-by-step solution visualization
- Piece placement suggestions
- Constraint validation
- Search space exploration

Currently displays hint pieces loaded from the CSV files.
