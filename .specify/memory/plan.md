# Implementation Plan: GitHub Stars Updater

**Status**: Consolidated
**Last Updated**: 2026-03-12
**Spec**: `.specify/memory/spec.md`

## Summary
The project is a standalone CLI tool written in Go that updates GitHub star counts in Markdown and AsciiDoc files. It uses a modular architecture where format-specific logic is encapsulated in implementations of a `LinkUpdater` interface.

## Technical Context
**Language**: Go 1.25
**Primary Dependencies**:
- `github.com/google/go-github/v68` (GitHub API Client)
- `golang.org/x/oauth2` (Authentication)
**Platform**: Cross-platform CLI (macOS, Linux, Windows)
**Architecture**: Modular flat structure with interface-based polymorphism for file parsing.

## Constitution Check
- **Go Idiomatic**: Yes, uses interfaces for polymorphism and standard library for core I/O and parsing.
- **CLI First**: Yes, uses the `flag` package and supports piped-like output via `-dry-run`.
- **Security**: Yes, authentication handled via `GITHUB_TOKEN` environment variable.
- **Simplicity**: Yes, focuses on a single task (updating stars) with minimal external dependencies.

## Project Structure

### Documentation
```text
.specify/memory/
├── constitution.md      # Project principles
├── plan.md              # This file (Consolidated Plan)
└── spec.md              # Consolidated Specification
specs/                   # Archive of individual feature specifications
```

### Source Code
```text
.
├── main.go              # CLI Entry point and coordination
├── common.go            # LinkUpdater interface and shared utilities
├── markdown.go          # Markdown-specific parsing and injection
├── asciidoc.go          # AsciiDoc-specific parsing and injection
├── main_test.go         # Integration tests
├── asciidoc_test.go     # AsciiDoc unit tests
├── go.mod               # Module definition
└── go.sum               # Dependency checksums
```

## Implementation Details

### Core Components
1. **CLI Coordination (`main.go`)**: 
   - Handles flag parsing.
   - Detects file extension to select the appropriate `LinkUpdater`.
   - Manages the overall flow: Read -> Parse -> Fetch -> Update -> Write.
2. **Interface (`common.go`)**:
   - `LinkUpdater`: `FindRepos(content string) ([]string, error)` and `UpdateContent(content string, stars map[string]int) (string, error)`.
   - Shared helpers like `formatStarCount` and `removeStarsInfo`.
3. **Formatters**:
   - `MarkdownUpdater`: Uses regex tailored for `[text](url)`.
   - `AsciiDocUpdater`: Uses regex tailored for `link:url[text]` and `url[text]`.

### Verification Procedures
- **Unit Tests**: `go test ./...`
- **Manual Verification**: Running the tool against sample `.md` and `.adoc` files with various link formats.
