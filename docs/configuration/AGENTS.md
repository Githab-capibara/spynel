# Configuration Documentation DOX

## Purpose

- Own `.spynel/config.yaml` documentation and the live settings reference.

## Local Contracts

- Document the canonical private configuration, workspace-root-relative path resolution, the fixed non-configurable `.spynel` state directory, and validation of current settings with ignored unused keys removed on the next canonical save.
- Keep the settings matrix synchronized with `internal/config` catalog defaults, command examples, and executable behavior.
- Only the three extension controls are restart-bound; every other exposed setting applies live behind the shared save-and-reload boundary.
- Secrets are excluded from documentation and status; document how token references resolve without ever showing values.

## Child DOX Index

No child DOX files.