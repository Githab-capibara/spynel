# Documentation Template

Spynel documentation follows four fixed shapes. This file is the master reference; each folder carries a tailored copy describing the shapes used there.

## File and folder conventions

- Filenames are lowercase kebab-case with a two-digit numeric prefix: `01-getting-started.md`.
- The prefix numbers documents in reading order within a folder. Reuse never happens; when a document is removed, later documents keep their numbers.
- Audience and `spynel docs` topic classification stay localized to the owning folder.
- Folder names describe the topic (product, guides, configuration, channels, automation, workflows, harness, security, extensions, releasing, tui). Never create generic "report", "misc", or "other" namespaces.
- Every folder keeps its own `AGENTS.md`, `README.md` (the table index), and `template.md`.
- Additions, removals, or renames update the folder README index and owning `AGENTS.md` (DOX) immediately.
- Content is English. "Report" never appears as a folder or document scope word.
- Dates follow ISO-8601 (`YYYY-MM-DD`) and timestamps follow UTC.

## A. Architecture Decision Record (ADR)

Michael Nygard format, strict:

```markdown
# 01. Title in present-tense imperative

- **Status:** Proposed | Accepted | Deprecated | Superseded by ADR-NN
- **Date:** YYYY-MM-DD
- **Deciders:** @Githab-capibara
- **Related:** links

## Context

What forces are at play?

## Decision

What are we doing? (present tense, affirmative)

## Consequences

- **Easier:** ...
- **Harder:** ...
- **Given up:** ...
- **Migration:** ...

## Alternatives considered

- **Option A:** rejected because ...
```

## B. Design document

```markdown
# 01. Title

- **Status:** Research note
- **Date:** YYYY-MM-DD
- **Deciders:** @Githab-capibara
- **Researcher:** ...
- **Purpose:** ...
- **Feeds into:** docs/product/...
```

## C. Folder README

Tabular index:

```markdown
# Folder Name

One-line scope statement.

| Document | Purpose |
| --- | --- |
| [Getting Started](01-example-topic.md) | First engagement in ~5 minutes |
```

## D. Project README (root only)

Badges and shields (license, stars, Discord, website), a hero image or video, benchmark results with donut charts, architectural diagrams (SVG), and tables linking into `docs/`.

## Verification

- `scripts/dev.sh dox` validates that every new folder owns an `AGENTS.md` and that indexes stay current.
- Run `scripts/dev.sh test` and `scripts/dev.sh build` after any documentation change that touches code-adjacent contracts.
- A documentation change that alters command, configuration, or behavior claims must keep the compiled `spynel docs` catalog, CLI help, and repository docs synchronized.