# Architecture Decision Records DOX

## Purpose

- Own the architecture decision records (ADRs) for Spynel: the durable reasons behind non-obvious architectural choices.

## Local Contracts

- ADRs use the strict Michael Nygard format documented in `template.md` and the master `../template.md` (section A).
- Records are written in present tense, affirmative voice, with a status of `Accepted` unless still under evaluation.
- Each record must state the concrete forces, the decision, the consequences (easier / harder / given up / migration), and the alternatives considered and rejected.
- ADRs describe decisions already reflected in `docs/product/02-architecture.md` and the root DOX contract; they never invent capabilities, guarantees, or adoption claims.
- Keep numbering sequential from `01` and never reuse a number; a superseded record links to its replacement instead of losing its history.
- Additions, removals, or renames update `README.md` (the living index) immediately.
- Do not duplicate live configuration, command, or release details; reference the owning topic folders instead.

## Child DOX Index

No child DOX files.