# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

SCMP (Secure Client Messaging Protocol) is the client-to-server protocol for the infodancer messaging stack. It handles all communication between a user's client and their own domain server. A client has exactly one server relationship — its own domain.

This is a new protocol designed from scratch. See the parent `infodancer/infodancer` repo for the full requirements document (`docs/next-gen-messaging-protocol.md`) and the protocol outline (`docs/protocol-outlines.md`).

### Responsibilities

- User authentication
- Message composition and submission
- Message retrieval from store
- Folder and mailbox management
- Key management (submit, rotate, revoke, discover, monitor)
- Push notification to client
- Policy configuration (pull preferences, reputation thresholds, retention policies, hash address management)
- Multi-device synchronization

### Key Design Principles

- **Client encrypts before handoff**: server receives ciphertext only, never holds plaintext (except on escrow-mandatory domains)
- **Explicit acceptance**: pulling a message is an explicit protocol action, not automatic
- **Hashed recipient addressing**: server never sees which user a hash belongs to (blind store model)
- **One server relationship**: client speaks only to its own domain

## Technology

- **Go** for the server implementation
- **Protocol Buffers** for message serialization
- **gRPC** for all client-server operations (bidirectional streaming for push notifications)
- Future consideration: REST/HTTP gateway for web clients

## Development Commands

This module uses [Task](https://taskfile.dev/) as the build tool:

```bash
task build          # Build the binary
task test           # Run tests
task lint           # Run golangci-lint
task vulncheck      # Check for vulnerabilities
task all            # Run all checks (build, lint, vulncheck, test)
task test:coverage  # Run tests with coverage report
task install:deps   # Install golangci-lint and govulncheck
task hooks:install  # Configure git to use .githooks directory
```

## Module Structure

```
/cmd/scmp/main.go       # Entrypoint only, minimal logic
/internal/scmp/          # Domain-specific implementation
/proto/                  # Protocol buffer definitions
/errors/                 # Centralized error definitions
```

## Development Workflow

### Branch and Issue Protocol

**This workflow is MANDATORY.** All significant work must follow this process with no exceptions:

1. **Create a GitHub issue first** - Before creating a branch, draft an issue describing the purpose and design based on your understanding of the user's request. Assign the issue to the user who requested it. Ask the user to approve the issue before proceeding.

2. **Create a feature or content branch** - Only after issue approval, create the branch. Use descriptive names that include the issue id like `feature/UUID` or `bug/UUID`.

3. **Reference the issue in all commits** - Every commit message and pull request must include the issue URL.

4. **Stay focused on the issue** - Make only changes directly related to the approved issue. Do not refactor unrelated code, fix unrelated bugs, or make "improvements" outside the scope.

5. **Handle unrelated problems separately** - If you notice bugs, technical debt, or potential issues unrelated to your current work, ask the user to approve creating a separate GitHub issue. Do not address them in the current branch.

### Pull Request Workflow

- All branches merge to main via PR
- PRs should reference the originating issue
- **NEVER ask users to merge or approve a PR** - PR approval and merging must always be manual actions taken by the user
- After creating a PR, checkout the main branch before starting any further work

### Security Best Practices

- Never commit secrets, API keys, credentials, or tokens
- Use `crypto/rand` for random number generation in security contexts
- Validate all external input at system boundaries
- Use parameterized queries to prevent injection
- Regularly audit dependencies with `govulncheck`

Read CONVENTIONS.md for Go coding standards.
