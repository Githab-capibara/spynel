# ADR Template

Each Architecture Decision Record follows the strict Michael Nygard format. Start from the master [`../template.md`](../template.md).

```markdown
# NN. Title in present-tense imperative

## Status

Accepted | Proposed | Deprecated | Superseded by ADR-NN

## Context

What forces are at play?

## Decision

What are we doing? (present tense, affirmative)

## Consequences

- **Easier:** ...
- **Harder:** ...
- **Given up:** ...
- **Migration:** ...

## Alternatives considered

- **Option A:** rejected because ...
```

- **Context** and **Decision** are required; **Alternatives considered** must name at least the runner-up.
- Write in the present tense: the decision is what Spynel does now, not what it once did.
- Records describe only decisions already reflected in `docs/product/02-architecture.md` and the root DOX contract; they never invent capabilities, integrations, guarantees, or adoption claims.
- Keep entries in this folder's `README.md` index in sync with the records themselves.

## Verification

`scripts/dev.sh dox` guards the folder index and the owning `AGENTS.md`; content claims must trace to the architecture documentation or the root contract.