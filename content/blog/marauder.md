+++
title = "marauder"
date = 2021-02-05
description = "(fake command line screenshot for writing pentest reports) I solemnly swear I am up to no good"
category = "Misc"
tags = ["Misc"]
github = "https://github.com/alexanderbrevig/marauder"
visibility = "public"
languages = "Go"
interest_score = 7

[extra]
post_tags = ["Misc"]
+++

Why this matters: Security professionals and pentesters need to document their work clearly. Reports need to show what you tested, what you found, how you found it. Terminal screenshots are standard documentation. The problem is that real engagement screenshots contain sensitive information—live URLs, IP addresses, usernames, paths. You need clean examples that illustrate your technique without exposing your client.

The problem is creating convincing terminal screenshots for documentation without either (1) using real engagement data and leaking information, or (2) spending hours building fake examples manually. You need something that looks authentic, is configurable, and works fast.

My idea is a tool that runs any command, captures its output, and generates a fake terminal screenshot that looks like a real screenshot but is entirely clean. No sensitive data, authentic appearance, customizable styling.

I built `marauder` as a _Go_ tool that does exactly this. Run a command with `marauder`, and you get two outputs: a text file with the command and output (for reference), and a PNG image that looks like a terminal screenshot:

```bash
marauder ls -la
```

You get `$DATE ls -la.txt` (the actual output) and `$DATE ls -la.png` (a fake terminal rendering). The configuration system is flexible—YAML config file or environment variables. Customize colors, fonts, margins, and output directory. Colors can be overridden per-run:

```bash
MARAUDER_COLOR_COMMAND="00ff00" marauder sensitive-command
```

The tool understands the security space. The rendering looks authentic because it mimics real terminal aesthetics. No developer time wasted on Photoshop. No risk of accidentally including real data. This is the kind of tool that makes practitioners in security say "oh, I'm building something like this"—it solves a real workflow problem efficiently.
