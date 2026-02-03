# Implementation Plan: [FEATURE]

**Branch**: `[###-feature-name]` | **Date**: [DATE] | **Spec**: [link]
**Input**: Feature specification from `/specs/[###-feature-name]/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

[Extract from feature spec: primary requirement + technical approach from research]

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

**Language/Version**: Go 1.24
**Primary Dependencies**: [e.g., spf13/cobra, spf13/viper, or standard library only]
**Storage**: [e.g., local files, SQLite, or N/A]
**Testing**: Go standard `testing` package
**Target Platform**: CLI (macOS, Linux, Windows)
**Performance Goals**: [e.g., <100ms startup, minimal memory footprint]
**Constraints**: [e.g., single binary, no CGO]
**Scale/Scope**: [e.g., small utility, complex tool]

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

[Gates determined based on constitution file]

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
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
# [REMOVE IF UNUSED] Option 1: Simple CLI (Flat)
.
├── main.go
├── [feature].go
└── [feature]_test.go

# [REMOVE IF UNUSED] Option 2: Structured CLI (cmd/internal)
cmd/
└── [appname]/
    └── main.go
internal/
└── [feature]/
    ├── [feature].go
    └── [feature]_test.go
pkg/
└── [public_lib]/
```

**Structure Decision**: [Document the selected structure and reference the real
directories captured above]

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |
