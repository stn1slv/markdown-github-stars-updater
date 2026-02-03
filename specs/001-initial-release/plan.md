# Implementation Plan: Initial Release

**Branch**: `main`
**Date**: 2026-02-03
**Spec**: `specs/001-initial-release/spec.md`
**Input**: Existing Codebase

## Summary
The project is a standalone CLI tool written in Go. It uses the `google/go-github` library to query the GitHub API and regex to parse/modify Markdown files.

## Technical Context
**Language**: Go 1.24
**Dependencies**:
- `github.com/google/go-github/v68` (GitHub API Client)
- `golang.org/x/oauth2` (Authentication)
**Platform**: Cross-platform CLI (macOS, Linux, Windows)
**Performance**: Dependent on GitHub API rate limits and network latency. Processing is sequential.

## Constitution Check
- **Go Idiomatic**: Yes, standard layout, `go.mod` used.
- **CLI First**: Yes, uses `flag` package.
- **Security**: Yes, uses `GITHUB_TOKEN` env var.

## Project Structure

### Documentation
```text
specs/001-initial-release/
├── plan.md              # This file
└── spec.md              # Feature specification
```

### Source Code
```text
.
├── main.go              # Entry point and core logic
├── go.mod               # Module definition
├── go.sum               # Dependency checksums
└── main_test.go         # (Existing tests if any, or to be added)
```

## Implementation Details

### Core Logic (`main.go`)
1.  **CLI Parsing**:
    -   `flag.String("out")`
    -   `flag.Bool("dry-run")`
    -   `os.Getenv("GITHUB_TOKEN")`
2.  **File Processing**:
    -   `os.ReadFile` to load content.
    -   `regexp.MustCompile` to find links.
3.  **API Interaction**:
    -   `oauth2.StaticTokenSource` for auth.
    -   `client.Repositories.Get` to fetch repo details.
4.  **Formatting**:
    -   Logic to handle `< 1k`, `1k-10k`, `> 10k` formatting logic.
    -   `strings.Replace` to inject new titles.

### Verification
-   Build: `go build`
-   Test: `go test ./...`
