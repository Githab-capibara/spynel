# Extensions Template

Documents here describe trusted extensions and executable hooks. Start from the master [`../template.md`](../template.md).

## Extension entry skeleton

```markdown
## Installing an extension
## Hook contract (JSON over stdio)
## At-least-once delivery
## Version transitions
## Restart-bound controls
```

Rules:

- Extensions are explicitly installed Git repositories; only their declared executable hooks run, from the repository directory.
- Hook delivery is at least once with a stable `event_id`; consumers persistently deduplicate visible effects; never claim exactly-once arbitrary side effects.
- Only `extensions.enabled`, `extensions.directory`, and `extensions.hook_timeout` require a restart; every other exposed setting applies live.

## Verification

`scripts/dev.sh dox` validates structure; extension-loading smoke coverage and hook contract tests are the behavior evidence.