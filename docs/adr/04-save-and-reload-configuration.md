# 04. Commit Live Settings Through a Serialized Save-and-Reload Boundary

## Status

Accepted

## Context

Spynel exposes many live settings across communications, harnesses, and themes. If every owner cached its own copy and applied changes independently, commands and channels would diverge and in-flight work would see half-new configuration. A validated change must be visible to the whole process as one atomic snapshot.

## Decision

Live configuration uses one serialized save-and-reload boundary. The store validates a complete candidate against current settings, atomically replaces canonical `.spynel/config.yaml`, reloads that file into the shared process snapshot before returning, and publishes the refreshed snapshot to runtime owners. Subsequent operations read the refreshed snapshot; only owners with derived cached state receive minimal direct hooks. Active channels and in-flight orchestration may finish under their already-admitted state.

Only `extensions.enabled`, `extensions.directory`, and `extensions.hook_timeout` are restart-bound. Harness implementation, command, arguments, and sandbox changes remain transactional and idle-only. Model selection commits under the same fence as harness dispatch admission: admitted turns keep their captured model while every dispatch snapshot ordered after the commit uses the new one.

## Consequences

- **Easier:** TUI, CLI, Telegram, and WhatsApp converge on one process snapshot; scalar form controls and text commands share the same config keys, so validation cannot diverge.
- **Harder:** The save path must reload before returning and publish atomically, and derived-cache owners need small direct refresh hooks.
- **Given up:** No per-key streaming reconfiguration; one canonical YAML file with a strictly current schema.
- **Migration:** Unused configuration keys are ignored on load and removed on the next canonical save; migrations and fallback readers are not retained.

## Alternatives considered

- **Per-owner setting caches refreshed by polling:** rejected because it would split decision-making across components and delay convergence.
- **No live reload (all settings restart-bound):** rejected because most settings must apply without interrupting active channels.