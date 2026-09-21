# Architecture Decision Records

Durable reasons behind Spynel's non-obvious architectural choices, recorded after the fact from architecture documentation and the root DOX contract.

| ADR | Decision | Status |
| --- | --- | --- |
| [01 — Single Primary Election](01-single-primary-election.md) | One elected owner commands the workspace application service. | Accepted |
| [02 — Harness-Neutral Composition Boundary](02-harness-neutral-composition-boundary.md) | Channels translate traffic; only the application and orchestrator speak to harnesses. | Accepted |
| [03 — Markdown as Durable Workflow State](03-markdown-durable-workflow-state.md) | Tasks and goals live in git-filesystem Markdown with leases. | Accepted |
| [04 — Serialized Save-and-Reload Configuration](04-save-and-reload-configuration.md) | Live settings commit atomically behind one admission boundary. | Accepted |
| [05 — Installation-Local Update Ownership](05-installation-local-update-ownership.md) | Each installation owns its runtime, updates, and restart path. | Accepted |