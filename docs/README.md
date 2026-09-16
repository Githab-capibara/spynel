# Documentation

The repository documentation index for Spynel. The root README explains what Spynel is; operational and implementation detail lives in these folders.

## Start here

| Guide | Purpose |
| --- | --- |
| [Getting Started](guides/01-getting-started.md) | Install, launch, and first engagement in minutes. |
| [Product Vision](product/01-product-vision.md) | What Spynel is, the three pillars, and the relationship slogan. |
| [Configuration](configuration/01-configuration.md) | The canonical `.spynel/config.yaml` and path rules. |
| [Troubleshooting](guides/02-troubleshooting.md) | Offline checks and bounded live diagnostics. |

## Directory map

| Folder | Scope | Index |
| --- | --- | --- |
| [product](product/README.md) | Vision and architecture. | [Directory index](product/README.md) |
| [guides](guides/README.md) | Getting started and troubleshooting. | [Directory index](guides/README.md) |
| [configuration](configuration/README.md) | Settings and live change matrix. | [Directory index](configuration/README.md) |
| [channels](channels/README.md) | Telegram and WhatsApp behavior. | [Directory index](channels/README.md) |
| [automation](automation/README.md) | Plain CLI, programmatic integration, and agent-readable docs. | [Directory index](automation/README.md) |
| [workflows](workflows/README.md) | Tasks, goals, and persistent instructions. | [Directory index](workflows/README.md) |
| [harness](harness/README.md) | Coding-harness compatibility. | [Directory index](harness/README.md) |
| [security](security/README.md) | Threat model and provider-canary gating. | [Directory index](security/README.md) |
| [extensions](extensions/README.md) | Trusted executable extensions and hooks. | [Directory index](extensions/README.md) |
| [releasing](releasing/README.md) | Release, packaging, and update mechanics. | [Directory index](releasing/README.md) |
| [tui](tui/README.md) | TUI editing and terminal checks. | [Directory index](tui/README.md) |

## Key entry points

- [Plain CLI and Automation](automation/01-cli-and-automation.md) — messages, follow-ups, conversations, status, framework commands, and output contracts.
- [Communication Integrations](channels/01-integrations.md) — TUI, Telegram, WhatsApp, voice, histories, and delivery behavior.
- [Tasks and Goals](workflows/01-tasks-and-goals.md) — durable Markdown workflows, review, recovery, waiting, and notifications.
- [Programmatic Integration](automation/02-programmatic-integration.md) — the v1 HTTP/NDJSON contract, replay, and private Unix sockets.
- [Persistent Instructions](workflows/02-persistent-instructions.md) — workspace-level preferences, role mapping, and precedence.
- [Coding Harness Compatibility](harness/01-harness-compatibility.md) — evidence-backed lifecycle coverage and known gaps.
- [Extensions and Hooks](extensions/01-extensions-and-hooks.md) — trusted executable extensions and delivery guarantees.
- [Agent-Readable Documentation](automation/03-agent-readable-docs.md) — the offline `spynel docs` interface and its versioned JSON schema.
- [Releasing](releasing/01-releasing.md) — native packaging, npm publication, credentials, and release verification.

## Governance

| Contract | Purpose |
| --- | --- |
| [AGENTS.md](AGENTS.md) | Owning documentation contracts and the child-DOX index. |
| [README.md](README.md) | This master directory index (Start here / Directory map / Key entry points / Governance). |
| [template.md](template.md) | The documentation style reference (ADR, design, folder, and project shapes). |

Each folder owns `AGENTS.md`, `README.md`, and `template.md`; add, remove, or rename entries through those files and keep the mapping above current.

Static documentation is offline, bounded, and free of live or private workspace state; current state lives in `spynel status`, `jobs`, `tasks`, `goals`, and `log`.