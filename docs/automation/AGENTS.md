# Automation Documentation DOX

## Purpose

- Own the non-visual plain CLI, programmatic v1 integration, and agent-readable documentation guidance.

## Local Contracts

- Document plain-CLI output contracts precisely: the last assistant-message item by default, all response deltas with `--stream`, NDJSON events with `--json`, structured JSON for status/history queries, and flags before positional arguments.
- `02-programmatic-integration.md` owns the supported v1 HTTP/NDJSON contract, request admission/retry limits, committed subscription/replay semantics, snapshot resynchronization, private Unix socket topology/limits, and runnable adapter/local-test examples. Keep it synchronized with the compiled `internal/agentdocs` integration topic and CLI help.
- Document headless operational stderr as content-free lifecycle metadata, with the optional private authenticated Unix socket supplementing loopback without changing foreign-environment fences.
- `spynel docs` is harness-free, server-free, offline, and never reads live or private workspace state; document the separation between static docs and live state.
- Keep `agent-readable-docs.md` synchronized with the curated `internal/agentdocs` catalog, concise `/help` metadata, and harness prompt guidance.

## Child DOX Index

No child DOX files.