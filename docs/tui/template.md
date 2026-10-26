# TUI Template

Documents here describe terminal/UI editing and terminal checks. Start from the master [`../template.md`](../template.md).

## Document skeleton

Every document in this folder starts from the master design-document header:

```markdown
# NN. Title

- **Status:** Research note | Accepted
- **Date:** YYYY-MM-DD
- **Deciders:** @Githab-capibara
- **Researcher:** document_specialist agent
- **Purpose:** ...
- **Feeds into:** docs/...
- **Related:** links to sibling documents
```


## TUI behavior entry skeleton

```markdown
## Selection
## Editing
## Copy / paste
## Undo / redo
## Scrolling
## Terminal restore (F6) and alternatives
```

Rules:

- Behavior claims must match `internal/channel/tui` and its tests: selection units, drag semantics, composer keys, undo/redo, pagination, and static-copy steps.
- Keep `scripts/terminal-copy-screen.mjs` installation and replay commands synchronized with this folder's documents.
- Never describe synthetic minimum widths; the viewport is never widened beyond the physical terminal.

## Verification

`scripts/dev.sh dox` validates structure; TUI PTY tests plus `capture-tui.sh` PNG evidence back visual/terminal claims.
