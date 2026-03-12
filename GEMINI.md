# markdown-github-stars-updater Development Guidelines

Auto-generated from all feature plans. Last updated: 2026-03-12

## Active Technologies

- Go 1.25 + Standard library (regexp, os, strings, flag, net/http) (001-initial-release, 002-asciidoc-support)

## Project Structure

```text
.
├── main.go              # CLI entry point and coordination
├── common.go            # LinkUpdater interface and shared utilities
├── markdown.go          # Markdown parsing logic
├── asciidoc.go          # AsciiDoc parsing logic
├── main_test.go         # Integration tests
└── asciidoc_test.go     # AsciiDoc unit tests
```

## Commands

- **Setup**: `make setup`
- **Build**: `make build`
- **Test**: `make test`
- **Lint**: `make lint` (runs `golangci-lint`, `govulncheck`, and `nilaway`)
- **Format**: `make format` (runs `gofumpt`)
- **Run**: `make run` (example usage)
- **Upgrade Deps**: `make upgrade-deps` (upgrades all dependencies)

## Code Style

- Go 1.25: Follow standard conventions.
- All dependencies must be listed in `go.mod`.
- Prefer standard library over external dependencies.

## Recent Changes

- 001-initial-release: Core Markdown update logic, dry-run mode, and custom output support.
- 002-asciidoc-support: Added support for AsciiDoc (.adoc) files via the `LinkUpdater` interface.

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
