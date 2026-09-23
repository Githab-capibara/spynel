# 01. Contributing

- **Status:** Accepted
- **Date:** 2026-09-17
- **Deciders:** @Githab-capibara
- **Researcher:** document_specialist agent
- **Purpose:** How to propose, verify, and land changes in the Spynel repository.
- **Feeds into:** docs/README.md

Spynel accepts contributions through normal GitHub pull requests. Before touching any file, read the applicable AGENTS.md chain: the root `AGENTS.md` owns repository-wide contracts, and the nearest child DOX file owns the subtree you plan to change.

## Contribution flow

1. **Check for existing work.** Search issues and open pull requests for the same problem before opening a new one.
2. **Create an issue or discussion** for a non-trivial change so the direction is agreed before implementation.
3. **Work in a fork or feature branch** named for the change, for example `fix/webhook-header-verification`.
4. **Make the smallest maintainable change** that satisfies the owning AGENTS.md contract, and update the owning documentation when the change affects behavior, scope, or structure.
5. **Verify with the repository gates** described below.
6. **Open a pull request** with `gh` and a description that states the problem, the change, and the verification you ran.

## Working contracts

- Follow the root `AGENTS.md` operating principles: reconnaissance before action, documentation updates after edits, new coverage for new behavior, and re-verification before finishing.
- Do not leave temporary binaries, profiles, caches, or scratch data in the repository. Disposable state belongs under ignored test-managed temporary directories or purpose-specific `.tmp*` directories.
- Do not commit secrets, tokens, credentials, or live installation identity. Commits do not change git configuration, are not forced, and add only intended files.
- Keep public copy aligned with the product contract: Spynel is a classic, non-AI orchestration program that coordinates external coding harnesses. Do not invent capabilities, guarantees, or adoption claims.

## Commit identity

Commits in this repository are authored by the repository owner:

- Name: `Githab-capibara`
- Email: `rrrarrr37r@gmail.com`

Use this identity for every commit (`git commit --author="Githab-capibara <rrrarrr37r@gmail.com>"`), and write concise commit messages that match the surrounding history style.

## Verification gates

Run the checks that match your change before opening the pull request:

| Change | Gate |
| --- | --- |
| Any documentation or repository-structure change | `scripts/dev.sh dox` |
| Any Go change | `go test ./...`, `go vet ./...`, `go build -o .tmp-bin/spynel ./cmd/spynel` |
| Initialization, configuration, harness dispatch, extension loading, channel lifecycle, orchestration | `scripts/smoke.sh` |
| npm launcher or release-layout change | `node npm/test.js` |
| Packaging or release-workflow change | `scripts/package-native.sh` for the host target, then execute the extracted archive |
| Changes to the local Bubble Tea terminal module | its race tests plus the real TUI PTY test |

Verification commands reuse Go's shared user cache; do not point `GOCACHE` at a project-local directory. A deliberately requested cold-cache diagnostic must use `scripts/cold-cache.sh`.

## Review

Pull requests are reviewed for correctness against the owning contracts, evidence-backed test coverage, documentation synchronization, and the repository verification gates. Address review feedback in follow-up commits; review commits are never amended in place.