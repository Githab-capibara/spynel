# ADR Template

Each Architecture Decision Record follows the strict Michael Nygard format: a bullet metadata header per the master [`../template.md`](../template.md) section A, then four required content sections. Start from the example below.

```markdown
# NN. Title in present-tense imperative

- **Status:** Proposed | Accepted | Deprecated | Superseded by ADR-NN
- **Date:** YYYY-MM-DD
- **Deciders:** GitHub handles
- **Related:** links

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

- **Status** is a top-level bullet, never an `## Status` section, and the record carries no separate **Researcher**, **Purpose**, or **Feeds into** bullets — those belong to design documents.
- **Context** and **Decision** are required; **Alternatives considered** must name at least the runner-up.
- Write in the present tense: the decision is what Spynel does now, not what it once did.
- Records describe only decisions already reflected in `docs/product/02-architecture.md` and the root DOX contract; they never invent capabilities, integrations, guarantees, or adoption claims.
- Keep entries in this folder's `README.md` index in sync with the records themselves.

## Verification

`scripts/dev.sh dox` guards the folder index and the owning `AGENTS.md`; content claims must trace to the architecture documentation or the root contract.
