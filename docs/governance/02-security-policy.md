# 02. Security Policy

- **Status:** Accepted
- **Date:** 2026-09-17
- **Deciders:** @Githab-capibara
- **Researcher:** document_specialist agent
- **Purpose:** How to report a vulnerability and what support the alpha project commits to.
- **Feeds into:** docs/README.md

Spynel is alpha software. It coordinates external coding harnesses, so its trust boundaries, channel authorization, and secret handling are documented risk surfaces; see [Security](../security/02-security.md) and the [Provider-Canary Threat Model](../security/01-provider-canary-threat-model.md).

## Reporting a vulnerability

Do **not** open a public issue that describes an exploitable vulnerability. Report it privately through the canonical repository's **Security** tab at <https://github.com/agent0ai/spynel/security>, which supports private vulnerability reporting and lets maintainers reply without public disclosure.

Include, when known:

- the affected version and platform,
- a minimal reproduction that does not expose live credentials,
- the observed and expected behavior,
- whether the issue discloses secrets or workspace state.

## Supported boundary

Spynel supports only its current configuration, state, manifest, session, and workflow schemas. Security fixes land in the current development line and the next release; older schema generations are not patched retroactively because pre-release data is not read back.

## Disclosure

Maintainers aim to confirm a report, develop a fix, and release it before public disclosure. Forking the repository for private reproduction of a reported issue is expected; publishing a proof-of-concept trade that exposes real credentials is not.

## Secret discipline

Credentials, tokens, session keys, and workspace identity never appear in static documentation, issues, commits, or example output. Telegram and WhatsApp credentials and session state live only in the private workspace `.spynel/` directory; installation identity lives in the OS per-user configuration directory. Installers and updaters never claim to have changed a caller's PATH or shells without verification.