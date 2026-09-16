# Channels Documentation DOX

## Purpose

- Own communication-channel documentation: Telegram and WhatsApp setup, authorization, pairing, and response-delivery behavior.

## Local Contracts

- Document fail-closed Telegram and WhatsApp sender allow-lists, required Telegram webhook verification, transport delivery and account modes, WhatsApp QR/session persistence, and `/start` acknowledgment behavior.
- Document that WhatsApp setup saves and enables the channel after access configuration, opens pairing automatically, uses a chrome-free full-terminal QR view, and retries expired sessions automatically.
- Document response-delivery differences: the TUI exposes live progress, CLI streaming is opt-in, and Telegram/WhatsApp deliver only the last terminal response or error.
- Remote routine confirmations never contain local-path Markdown links.
- Keep pairing credentials and recipient identifiers in private workspace state; never include them in static documentation.

## Child DOX Index

No child DOX files.