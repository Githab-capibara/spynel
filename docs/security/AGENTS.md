# Security Documentation DOX

## Purpose

- Own the provider-canary threat model and general security documentation.

## Local Contracts

- Keep authenticated provider canaries gated by the reviewed threat model: synthetic repositories only, disposable identities and homes, verified artifacts, bounded egress/cost/time, sanitized evidence, and per-run authorization. A plan or CI definition is not evidence that a provider was executed.
- Document trust boundaries, secret handling, executable trust, and instruction precedence without exposing secrets, credentials, notification origins, or conversation identities.
- Keep security claims synchronized with `internal/agentdocs` security content and channel authorization behavior.

## Child DOX Index

No child DOX files.