# Governance Template

Documents here describe how people contribute and behave. Start from the master [`../template.md`](../template.md).

## Governance document skeleton

```markdown
# 01. Document title

- **Status:** Accepted
- **Date:** YYYY-MM-DD
- **Deciders:** @Githab-capibara
- **Researcher:** document_specialist agent
- **Purpose:** ...
- **Feeds into:** docs/README.md

## Scope

## Section — how it works

## Verification
```

Rules:

- Every command must run the way the shell actually executes it; GitHub operations use the configured `gh` CLI.
- Commit identity and DOX-compliance rules must match the repository contract in the root `AGENTS.md`.
- Never print secrets, tokens, credentials, or workspace identities in examples or guides.

## Verification

`scripts/dev.sh dox` guards the folder index and AGENTS.md coverage; the master `../template.md` shape guards structure.