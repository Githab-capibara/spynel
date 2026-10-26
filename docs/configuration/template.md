# Configuration Template

Documents here describe `.spynel/config.yaml` and each setting's live effect. Start from the master [`../template.md`](../template.md).

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


## Settings entry skeleton

```markdown
| Setting | Type | Default | Restart-bound |
| `group.name` | `string` | `value` | no |
```

Rules:

- Every value must match the executable default from `internal/config` — no invented defaults.
- State whether the setting applies live (non-extension settings do) or is one of the three restart-bound extension controls.
- Secrets are referenced, never shown; document resolution via environment references without values.
- The retired `channels.tui.enabled` input, the fixed `.spynel` state directory, and workspace-root-relative path resolution are canonical; no configurable state-directory key exists.

## Verification

`scripts/dev.sh dox` validates structure; `spynel config` and `/config get` prove the default and live behavior claimed.
