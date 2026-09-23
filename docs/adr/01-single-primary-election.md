# 01. Elect One Primary Server per Workspace

- **Status:** Accepted
- **Date:** 2026-09-17
- **Deciders:** @Githab-capibara
- **Related:** [Architecture](../product/02-architecture.md) · [Troubleshooting](../guides/02-troubleshooting.md)

## Context

Spynel's durable state, harness sessions, and histories must remain single-writer. Running a new owner freely alongside a live one would duplicate dispatch, split conversations, and corrupt Markdown leases. But a workspace must tolerate restarts, dead owners, and multiple live TUI and automation clients. The process boundary is a full application service, so the decision chooses a small numeric winner rather than replicating state machines.

## Decision

Every process rooted at the same workspace state directory participates in one OS-file-locked election for the `.spynel/runtime/primary.json` lease. The owner record carries a random instance ID, PID, loopback endpoint, non-secret environment ID, per-term bearer token, start time, and heartbeat. The owner renews every five seconds; other processes check once per second and may take over only after the heartbeat is 30 seconds old. Renewal and release verify both instance ID and term token so a stalled former owner cannot overwrite the winner.

The environment ID is a SHA-256 digest, domain-separated for this purpose, of a private random token in the OS per-user configuration directory. A known mismatch fails before HTTP dialing with host/container guidance and is never takeover-eligible. The first TUI in an ownerless workspace resumes the latest TUI conversation; additional live TUIs receive independent sessions.

## Consequences

- **Easier:** Telegram, WhatsApp, and continuous Markdown orchestration have one durable owner; every TUI and plain CLI client reaches the owner over authenticated loopback and sees one global job, task, and goal set.
- **Harder:** Every lease check must serialize on a short-held cross-platform lock, and takeover decisions must reconcile fresh-owner fencing with compatibility readiness for older records.
- **Given up:** No per-workspace leaderless consistency; a stale or mid-write owner is not consulted mid-transition.
- **Migration:** `/restart` persists and emits its final acknowledgment before publishing a coalesced process request that replaces the process in place.

## Alternatives considered

- **Per-process writer locks on every durable file:** rejected because it fragments single-writer ownership and leaks authorization decisions into file layout.
- **No owner, lock-free last-write-wins:** rejected because it would duplicate harness dispatch and corrupt session persistence.