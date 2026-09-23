# 05. Keep Update and Uninstall Ownership Installation-Local

- **Status:** Accepted
- **Date:** 2026-09-17
- **Deciders:** @Githab-capibara
- **Related:** [Releasing](../releasing/01-releasing.md) · [Getting Started](../guides/01-getting-started.md)

## Context

Spynel installs through a root POSIX script or npm, and updates must restart every running instance of that installation across workspaces. A workspace-local or unmanaged record would lose track of processes after npm unlinks a running image, and a half-validated bundle must never replace a working installation.

## Decision

The installation owns its runtime records per user, never inside a workspace. The updater validates a bounded SHA-256 download, archive paths and types, required runtime files, and the executable version before publishing a complete immutable bundle behind an atomic `current` switch, under a crash-released installation lock. Existing processes retain their libraries; restart and generated services use the stable entry point.

Native process discovery retains visibility even after npm unlinks a running image. On macOS the retained executable path from `kern.procargs2` supplies the fallback when `proc_pidpath` loses the vnode path. A valid bundle is shared: `/update` acknowledges the action, requests a graceful owner shutdown and restart, asks every other registered instance of that installation to shut down and exec the updated binary in place, and succeeds only when new generations report the target version and application readiness. `spynel killall` and public uninstall use the same verified process inspection and bounded termination.

## Consequences

- **Easier:** Updates and uninstall restart or stop every registered server and TUI of one installation across workspaces; scripts and npm share one ownership model.
- **Harder:** The updater must verify native executable identity and process ownership before signaling, and must preserve unrelated user-bin entries and npm installations.
- **Given up:** Workspace-local or unmanaged process records; future startup registrations are still registered automatically.
- **Migration:** Older unregistered instances require an explicit one-time stop/relaunch instead of a coordinated restart.

## Alternatives considered

- **Per-workspace process registry:** rejected because instances span workspaces and npm unlink would strand them.
- **Trusting any binary at the recorded path:** rejected because native executable identity must be verified before signaling.