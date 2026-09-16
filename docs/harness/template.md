# Harness Template

Documents here describe coding-harness compatibility. Start from the master [`../template.md`](../template.md).

## Harness entry skeleton

```markdown
### `<Harness name>`

- **Discovery:** from `PATH` and conventional per-user locations.
- **Follow-up mode:** native steering or queued.
- **Sandbox values:** `danger-full-access` | `workspace-write` | `read-only`.
- **Model/effort/service:** capability-validated; ACP effort unsupported.
- **Prerequisites:**
```

Rules:

- Every capability claim must match `internal/harness` catalogs and tests; use the compatibility matrix facts, never invent support boundaries.
- Model and supported-harness effort pickers end with Custom text entry even when discovery fails; explicit identifiers bypass membership checks and are validated by the harness.
- Agent Zero CLI must pass `a0 acp --check` before discovery reports it as available.
- Custom ACP uses a one-line command text with quoting/escaping and no shell expansion; never document shell-expanding argument lists.

## Verification

`scripts/dev.sh dox` validates structure; `internal/harness` lifecycle tests and the native-evidence suite are the behavior evidence.