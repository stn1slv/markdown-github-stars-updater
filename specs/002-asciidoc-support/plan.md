# Implementation Plan: AsciiDoc Support

**Branch**: `002-asciidoc-support` | **Date**: 2026-02-03 | **Spec**: `specs/002-asciidoc-support/spec.md`
**Input**: Feature specification from `/specs/002-asciidoc-support/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

This feature adds support for processing AsciiDoc (`.adoc`, `.asciidoc`) files to the existing CLI tool. It will reuse the core star-fetching logic but introduce a new parsing strategy to handle AsciiDoc-specific link syntax (`link:url[text]` and `url[text]`). Additionally, the project dependencies will be updated to their latest versions.

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

**Language/Version**: Go 1.25 (Ensure dependencies are updated)
**Primary Dependencies**: Standard library only (regexp, os, strings)
**Storage**: Local files (read/write)
**Testing**: Go standard `testing` package
**Target Platform**: CLI (macOS, Linux, Windows)
**Performance Goals**: Similar to Markdown processing (limited by network I/O)
**Constraints**: Single binary, no CGO
**Scale/Scope**: Small extension to existing utility

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

*   **Go Idiomatic**: Yes. Will use standard library `regexp` for parsing, consistent with existing codebase.
*   **CLI First**: Yes. Flag usage (`-dry-run`, `-out`) remains consistent. Input file type inferred from extension or potentially flagged (though inference is preferred per spec).
*   **Security & Configuration**: Yes. Uses `GITHUB_TOKEN` environment variable. No new secrets or config required.
*   **Simplicity & Focus**: Yes. Directly supports the core value proposition (updating stars) for a new format. No unnecessary features.

## Project Structure

### Documentation (this feature)

```text
specs/002-asciidoc-support/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)
<!--
  ACTION REQUIRED: Replace the placeholder tree below with the concrete layout
  for this feature. Delete unused options and expand the chosen structure with
  real paths.
-->

```text
.
├── main.go              # Main entry point (update to select parser)
├── markdown.go          # Extract Markdown logic (refactor)
├── asciidoc.go          # New AsciiDoc parsing logic
├── asciidoc_test.go     # Tests for AsciiDoc parsing
├── common.go            # Shared types/interfaces (if needed)
└── main_test.go         # Integration tests
```

**Structure Decision**: Moving to a slightly more modular flat structure. Extracting the existing Markdown logic from `main.go` into `markdown.go` (or similar) and adding `asciidoc.go` will keep `main.go` clean as a coordinator/CLI entry point.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A | | |