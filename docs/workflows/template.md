# Workflows Template

Documents here describe durable task/goal documents and standing per-role instructions. Start from the master [`../template.md`](../template.md).

## Workflow section skeleton

```markdown
## Lifecycle
## Durable document (front matter and body)
## Leases and recovery
## Review semantics
## Notifications
```

Rules:

- Tasks are single finite objectives; goals are measurable multi-round outcomes. The distinction is never blurred.
- Every claimed phase has a persisted lease; recovery journals condition, evidence, exception, and next action and never bypasses required review.
- Persistent instruction files are exactly the five `.spynel/instructions/agent-*.md` roles with fresh loading, safe-file constraints, and content-free inspection.
- Notification guidance uses concrete `--workdir`, `--origin`, and `--message` forms; placeholders and stdin composition are not task-agent guidance.

## Verification

`scripts/dev.sh dox` validates structure; typed status-folder and review-inspection smoke coverage is the behavior evidence.