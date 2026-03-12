# Project Changelog

## Merged Features Log

### Initial Release — 2026-02-03
**Branch:** `main`
**Spec:** `specs/001-initial-release`

**What was added:**
- Automatically update star counts next to GitHub links in Markdown files.
- Support for dry-run mode to preview changes.
- Support for custom output file destination.

**New Components:**
- `main.go`: Entry point and core logic.
- `go.mod`: Dependency management.
- `main_test.go`: Integration tests.

**Tasks Completed:** N/A (tasks.md not provided)

### AsciiDoc Support — 2026-02-03
**Branch:** `002-asciidoc-support`
**Spec:** `specs/002-asciidoc-support`

**What was added:**
- Support for AsciiDoc (`.adoc`, `.asciidoc`) files.
- Refactored parsing logic into a modular `LinkUpdater` interface.
- Detection of file type by extension.

**New Components:**
- `common.go`: Shared `LinkUpdater` interface and utilities.
- `asciidoc.go`: AsciiDoc-specific parsing and injection.
- `asciidoc_test.go`: Unit tests for AsciiDoc support.
- `markdown.go`: Extracted Markdown logic.

**Tasks Completed:** 15/15 tasks
