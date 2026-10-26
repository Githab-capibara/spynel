# Releasing Template

Documents here describe release, packaging, and update mechanics. Start from the master [`../template.md`](../template.md).

## Document skeleton

Every document in this folder starts from the master design-document header:

```markdown
# NN. Title

- **Status:** Research note | Accepted
- **Date:** YYYY-MM-DD
- **Deciders:** @Githab-capibara
- **Researcher:** document_specialist agent
- **Purpose:** ...
- **Feeds into:** docs/...
- **Related:** links to sibling documents
```


## Release skeleton

```markdown
# 01. Releasing

## Versioning (v-prefixed semantic tags)
## Native archives (Linux/macOS amd64/arm64)
## npm triggering
## Installers: npm and POSIX script
## Updates and update check
## Uninstall
```

Rules:

- Windows is explicitly and temporarily unsupported by both native packaging and npm; no Windows promises.
- Document update-by-default, read-only `update check`, caller-installation versus channel-primary selection, cross-workspace restarts, and `killall` with preserved autostart registrations.
- Keep native asset coverage, first-publication credentials versus OIDC trusted publishing, and mirror/version overrides accurate.

## Verification

`scripts/dev.sh dox` validates structure; `scripts/package-native.sh` and `node npm/test.js` are the behavior evidence.
