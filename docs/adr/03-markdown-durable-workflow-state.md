# 03. Make Markdown Files the Durable Workflow State

- **Status:** Accepted
- **Date:** 2026-09-17
- **Deciders:** @Githab-capibara
- **Related:** [Tasks and goals](../workflows/01-tasks-and-goals.md)

## Context

Tasks and goals must survive restarts, be inspectable by any AI harness, and be auditable by humans. A database or binary store would hide state from harnesses and complicate review. The durable source of truth must be the same files any agent and any person can read and edit.

## Decision

Markdown task and goal files under the fixed `.spynel/` state directory are the durable source of workflow state. Front matter is machine-readable; bodies and agent logs are human-readable. Tasks are finite objectives carrying explicit boolean `review_required`; missing or malformed values fail safe to review. Every claimed phase has a persisted lease, and ownership, recovery, and stale takeover are serialized on one cross-process lock above the sibling status folders.

Transitions, prompt paths, and stale thresholds are fixed in `internal/orchestrator`; there are no runtime route snapshots. Goals run a separate `proposed -> planning -> active -> review -> reviewing` state machine whose settled task evidence informs—but never mechanically completes—an independently reviewed goal outcome.

## Consequences

- **Easier:** Any harness and any person can read the authoritative task and goal files; `/tasks` and `/goals` render bounded projections without starting a harness.
- **Harder:** Every claimed phase needs a persisted lease, and front matter and progress journaling must stay machine-readable and consistent.
- **Given up:** No database or binary task store; cross-process writer discipline is enforced by leases and locks instead.
- **Migration:** Existing task and goal documents must match the current schema; obsolete representations are rejected rather than migrated.

## Alternatives considered

- **Relational or embedded database:** rejected because it would hide state from harnesses and complicate human auditing.
- **Binary/owned state files:** rejected because durable state must remain inspectable and editable by external agents.