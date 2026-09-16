# TUI Documentation DOX

## Purpose

- Own terminal/UI editing behavior and terminal-check guidance for the Bubble Tea TUI.

## Local Contracts

- Keep text-selection, editing, copy/paste, undo/redo, scrolling, and F6 static-copy guidance synchronized with `internal/channel/tui` behavior and its tests.
- Document terminal-handling contracts: bracketed paste, SGR mouse events, terminal-mode restore, and the alternatives for restoring the alternate screen.
- Keep `docs/tui/01-tui-editing.md` synchronized with `scripts/terminal-copy-screen.mjs` installation and replay commands.
- Never widen a viewport to a synthetic minimum beyond the physical terminal; document the cell-width contract honestly.

## Child DOX Index

No child DOX files.