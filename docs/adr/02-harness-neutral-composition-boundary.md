# 02. Keep a Harness-Neutral Composition Boundary

- **Status:** Accepted
- **Date:** 2026-09-17
- **Deciders:** @Githab-capibara
- **Related:** [Harness compatibility](../harness/01-harness-compatibility.md) · [Architecture](../product/02-architecture.md)

## Context

Spynel must exploit improving external coding-harness intelligence without duplicating or competing with it. TUI, Telegram, and WhatsApp need identical behavior, and future harnesses must plug in without redrawing the transport layer. If channels knew which agent protocol was active, every new harness would fork the communication code.

## Decision

`cmd/spynel` composes only. Each channel translates external traffic into transport-neutral `core.Message` and calls the shared application; no channel ever knows which coding-agent protocol is active. The application and orchestrator depend only on the `harness.Harness` interface. Harness implementations, commands, arguments, and sandbox changes stay transactional and idle-only, while every surfaced choice uses one atomic active-safe dispatch snapshot of model, reasoning effort, and service mode.

Harness behavior lives in `internal/harness`: Codex, Claude Code, Agent Zero CLI, Pi, and ACP register metadata and factories there, implement the same interface, and explicitly declare native-steer or queued follow-up behavior.

## Consequences

- **Easier:** New harnesses register in `internal/harness` behind one contract; transports, configuration, history, and orchestration stay harness-agnostic.
- **Harder:** The shared application boundary must expose every cross-cutting capability (screens, status, jobs, branching, attachments) through provider-neutral models.
- **Given up:** No OpenAI-compatible HTTP façade and no per-harness transport shortcuts.
- **Migration:** New harnesses register metadata and factories in `internal/harness`; ACP-compatible CLIs reuse the shared v1 stdio adapter behind one concise command alias.

## Alternatives considered

- **Transport-speaking direct harness invocation:** rejected because it duplicates dispatch logic per channel and forfeits one global lifecycle.
- **An OpenAI-compatible HTTP façade:** rejected outright; harnesses extend through explicit interfaces only.