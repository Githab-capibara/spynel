# Channels Template

Documents here describe Telegram and WhatsApp setup and delivery. Start from the master [`../template.md`](../template.md).

## Transport section skeleton

```markdown
## Prerequisites (allow-list)
## Authorization
## Setup
## Pairing
## Delivery and response behavior
## Reconnect and failure
```

Rules:

- Emergency: never print tokens, pairing credentials, recipient identifiers, session databases, or public webhook URLs in static docs.
- State that Telegram/WhatsApp fail closed on missing, malformed, or revoked allow-lists without reconnect retry.
- State delivery behavior precisely: only the last non-continuing final response or terminal error is delivered remotely; do not claim "real-time" behavior that does not exist.
- Document WhatsApp's automatic save/enable/pairing flow and chrome-free QR view; do not describe an enable-choice step that does not exist.

## Verification

`scripts/dev.sh dox` validates structure; adapter lifecycle tests in `internal/channel` are the behavior evidence.