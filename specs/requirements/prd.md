# greeter — PRD

## Problem Statement

Developers building applications or integrations that need a simple, reliable way to greet a named user programmatically currently have to stand up their own trivial greeting logic, or skip the niceties of a proper greeting feature altogether because even a one-line text response needs a hosted, documented endpoint.

## Solution

Greeter is a small, public HTTP API: a single endpoint that takes a name and returns a JSON greeting, so any external developer can wire a greeting into their application without building or hosting the logic themselves.

## Actors

- **API Consumer** — an external developer (or their application/service) who calls the greeter API over HTTP to get a JSON greeting back. Has no account or sign-in; interacts only through the API.

## User Stories

1. As an API Consumer, I want to call GET /hello?name=X, so that I receive a JSON greeting addressed to that name.
2. As an API Consumer, I want GET /hello without a name (or with it empty), so that I still receive a generic JSON greeting rather than an error.

## Product Decisions

- The service exposes a single read-only endpoint, GET /hello, returning JSON.
- When the `name` query parameter is missing or empty, the response falls back to a generic greeting (e.g. "Hello, World!") rather than an error.
- Audience: external developers/public API consumers — the API is reachable outside the organization. *assumed*
- Authentication: the endpoint is open and requires no sign-in or API key, consistent with a simple public greeting utility. *assumed*
- The `name` value is accepted as free-form text with no length or character restrictions, and is echoed back in the greeting as-is. *assumed*
- Language/runtime: Go, per the organization's service default.

## Out of Scope

- User accounts, sign-in, or per-caller permissions.
- Persisting greetings, names, or request history.
- Localization/translation of the greeting text.
- Rate limiting, API keys, or usage quotas.

## Open Questions

(none)

## Further Notes

(none)