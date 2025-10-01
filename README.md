# AeroManager

A Go-based utility for robust multi-monitor workspace management in Aerospace, inspired by Hyprland's window management on Linux.

## Overview

AeroManager enhances Aerospace's multi-monitor workspace handling by addressing common issues with workspace arrangement when monitors are plugged/unplugged and providing intelligent workspace switching based on cursor position.

## Features

### Workspace Rearrangement
Automatically reorganizes Aerospace workspaces based on your current multi-monitor setup. This fixes the common issue where workspaces become mixed up after connecting or disconnecting external displays.

### Smart Workspace Switching
Switches workspaces intelligently based on:
- Current mouse cursor position
- Active monitor configuration
- Display topology

## Usage

The tool operates through command-line flags that determine the type of operation:

```bash
# Rearrange workspaces based on current monitor setup
aeromanager --rearrange

# Switch workspace based on cursor position
aeromanager --switch <workspace>
```

## How It Works

1. **Gathers system information** - Queries monitor configuration and cursor position using terminal commands
2. **Determines operation** - Based on provided flags, selects the appropriate action
3. **Executes commands** - Runs Aerospace CLI commands to perform workspace management

## Motivation

Aerospace's default behavior can be unpredictable in multi-monitor scenarios. AeroManager brings the polish and reliability of Hyprland's multi-monitor workspace management to macOS, making workspace handling more intuitive and consistent.

## Requirements

- macOS
- [Aerospace](https://github.com/nikitabobko/AeroSpace) window manager

## Installation

```bash
go build -o aeromanager
```

## License

MIT
