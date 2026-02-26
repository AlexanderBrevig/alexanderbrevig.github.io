+++
title = "nixos-config"
date = 2022-09-14
description = "A complete NixOS configuration with Home Manager & Hyprland"
category = "NixOS"
tags = ["NixOS"]
github = "https://github.com/alexanderbrevig/nixos-config"
visibility = "public"
languages = "Nix"
interest_score = 8

[extra]
post_tags = ["NixOS"]
+++

Why this matters: NixOS is powerful but daunting. It forces you to declare everything—system packages, user environment, services, dotfiles. Most people build their configurations incrementally, discovering patterns through trial and error. A well-documented, modular example that shows how to structure a real system can accelerate that learning by months. It becomes a reference point for best practices.

The problem is that NixOS documentation is good but abstract. The NixOS manual explains *how* to write configurations, but doesn't show *what a good one looks like*. You end up searching GitHub for examples, finding half-finished configs with cryptic comments. Nobody wants to build their system configuration from first principles.

My idea is to create a complete, professional NixOS configuration that I actually use, document it thoroughly, and open-source it as a template. Show how to structure it with flakes, how to manage multiple hosts (laptop and desktop), how to set up Home Manager for user-specific configuration, how to handle secrets, development environments, and services. Make it forkable and customizable.

I built a complete NixOS flake configuration with Home Manager and Hyprland (a modern Wayland compositor). It includes dual-host support (ThinkPad T14 laptop + custom workstation), comprehensive development environment setup for _Rust_, _Go_, _Java_, _Flutter_, _C++_, _Python_, _OCaml_. The terminal stack is modern: Fish shell, Starship prompt, WezTerm, Helix editor.

The documentation covers prerequisites, step-by-step installation (including LUKS encryption and partitioning), hardware configuration, SSH setup, post-installation configuration, troubleshooting, and customization. The structure is modular:

```nix
nixos/
├── flake.nix              # Main configuration
├── hosts/                 # Host-specific configs (laptop/desktop)
└── modules/               # Reusable system modules
    ├── shared.nix
    └── hyprland.nix
```

It's a real system someone (me) actually uses and maintains, not a textbook example. Whether it becomes a reference point for others or just an interesting artifact depends on whether anyone finds it useful. Either way, it works for what it's built for.
