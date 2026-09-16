# Releasing Template

Documents here describe release, packaging, and update mechanics. Start from the master [`../template.md`](../template.md).

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