# Guides Template

Documents here walk users through installations, launches, and recovery. Start from the master [`../template.md`](../template.md).

## Guide skeleton

```markdown
# 01. Guide title

## Prerequisites
## Steps
### Step 1 — ...
## Verify
## Troubleshooting
## Uninstall or rollback
```

Rules:

- Every command must run the way the shell actually executes it; step names start with verbs.
- Step "Verify" sections show the output that proves success — never assume success from silence.
- Troubleshooting offers offline checks before live diagnostics and never asks users to delete fresh leases.
- Installation/PATH mechanics belong here, not in the root README.

## Verification

Run each command shown in a clean shell before marking the guide done; `scripts/dev.sh dox` guards the index.