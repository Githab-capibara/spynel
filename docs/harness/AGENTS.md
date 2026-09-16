# Harness Documentation DOX

## Purpose

- Own coding-harness compatibility documentation for Codex, Claude Code, Agent Zero CLI, Pi, ACP aliases, and custom ACP.

## Local Contracts

- Keep the compatibility matrix synchronized with the `internal/harness` catalog, capability detection, model/effort/service metadata, and tests.
- Document detected model reasoning/service capabilities and Custom model/effort entry when discovery is missing or incorrect: provider validation of manual identifiers, dependent selection, inherit/reset semantics, invalidation behavior, atomic dispatch snapshots, the unsupported ACP effort boundary, and capability-validated service modes.
- Document $harness.sandbox$ values, agent prefixes, task-review modes, and the one-line custom-ACP command grammar without shell expansion.
- A plan or CI definition is never evidence that a provider was executed; keep provider canary documentation gated by the reviewed threat model in `../security/01-provider-canary-threat-model.md`.

## Child DOX Index

No child DOX files.