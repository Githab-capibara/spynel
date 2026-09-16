# Security Template

Documents here describe trust boundaries and validated provider evaluation. Start from the master [`../template.md`](../template.md).

## Security document skeleton

```markdown
# 01. ...

- **Status:** Accepted
- **Date:** YYYY-MM-DD
- **Deciders:** @Githab-capibara
- **Related:** links

## Threat model
## Trust boundaries
## Gating rules
## Evidence and authorization
## Residual risks
```

Rules:

- Provider canaries require synthetic repositories only, disposable identities and homes, verified artifacts, bounded egress/cost/time, sanitized evidence, and per-run authorization. A plan or CI definition is never evidence that a provider was executed.
- Never include secrets, credentials, notification origins, recipient identifiers, conversation histories, or environment values in static documentation.
- Security claims stay synchronized with `internal/agentdocs` security content and channel fail-closed behavior.

## Verification

`scripts/dev.sh dox` validates structure; authorization and delivery behavior in `internal/channel` and `internal/app` is the evidence baseline.