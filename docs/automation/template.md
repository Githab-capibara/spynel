# Automation Template

Documents here describe the plain CLI and programmatic surfaces. Start from the master [`../template.md`](../template.md).

## CLI command reference skeleton

```markdown
### `spynel <command> [flags]`

- **Purpose:** ...
- **Output contract:** last assistant-message item by default; `--stream` for deltas; `--json` for NDJSON events.
- **Requires:** conversational identity; `--conversation NAME`.
- **Example:**
    ```sh
    spynel send --conversation NAME ...
    ```
```

## Programmatic contract skeleton

- Stable endpoint and version: `/v1/...`, `spynel.events/v1`.
- Admission/retry limits and honest completion semantics: provider admission is not completion.
- Committed conversation replay bounds, snapshot resynchronization, and outbox notifications.
- Optional private authenticated Unix socket topology and limits.
- Headless stderr stays content-free; private diagnostic bodies live in `spynel log`.

## Verification

Provider-free examples must be runnable without live services; `scripts/dev.sh dox` guards the index and `scripts/dev.sh test` keeps behavior claims honest.