---

description: "Task list for AsciiDoc support feature"
---

# Tasks: AsciiDoc Support

**Input**: Design documents from `/specs/002-asciidoc-support/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, quickstart.md

**Tests**: Tests are included as per standard Go practices and to ensure regression safety during refactoring.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Path Conventions

- **Simple CLI**: `.` (root), `*_test.go` next to source

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and dependency updates

- [x] T001 Update go.mod to Go 1.25 and run `go get -u` to update dependencies
- [x] T002 Create new source files `markdown.go` and `asciidoc.go` (empty initially) per plan

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Refactoring to Strategy Pattern to support multiple file types.

**⚠️ CRITICAL**: Must complete before implementing AsciiDoc logic.

- [x] T003 Define `LinkUpdater` interface in `common.go` (or `main.go` if keeping it simple)
- [x] T004 [P] Move existing Markdown parsing logic from `main.go` to `MarkdownUpdater` struct in `markdown.go`
- [x] T005 Refactor `main.go` to use `LinkUpdater` interface instead of hardcoded logic
- [x] T006 Verify regression: Run existing `main_test.go` or manual test to ensure Markdown still works

**Checkpoint**: Foundation ready - code refactored, Markdown support preserved.

---

## Phase 3: User Story 1 - Update Star Counts in AsciiDoc (Priority: P1) 🎯 MVP

**Goal**: Enable updating of star counts in .adoc files.

**Independent Test**: Run `./updater test.adoc` and verify star counts appear.

### Tests for User Story 1
- [x] T007 [P] [US1] Create unit tests for AsciiDoc regex parsing in `asciidoc_test.go`
- [x] T008 [P] [US1] Implement `AsciiDocUpdater` struct and `FindRepos` logic in `asciidoc.go`
- [x] T009 [P] [US1] Implement `UpdateContent` logic for AsciiDoc in `asciidoc.go`
- [x] T010 [US1] Update `main.go` to select `AsciiDocUpdater` when file extension is `.adoc` or `.asciidoc`

**Checkpoint**: User Story 1 fully functional. AsciiDoc files can be updated.

---

## Phase 4: User Story 2 - Safety and Output Control (Priority: P2)

**Goal**: Ensure safety flags work for AsciiDoc.

**Independent Test**: Run with `-dry-run` and `-out` and verify behavior.

### Implementation for User Story 2
- [x] T011 [US2] Verify and ensure `main.go` passes correct flags/writer to `LinkUpdater` flow (likely already done in T005, this is a verification task)
- [x] T012 [P] [US2] Add integration test case in `main_test.go` for AsciiDoc with `-dry-run`

**Checkpoint**: User Story 2 verified.

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Documentation and final cleanup

- [x] T013 [P] Update `README.md` to mention AsciiDoc support and usage examples
- [x] T014 Run `go fmt` and `go vet` across all files
- [x] T015 Run `quickstart.md` validation steps manually

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies.
- **Foundational (Phase 2)**: Depends on Setup. Blocks US1.
- **User Stories (Phase 3+)**: US1 depends on Foundational. US2 depends on US1 (logically, for verification).

### Parallel Opportunities

- T004 (Markdown extraction) and T002 (File creation) can overlap slightly.
- T007 (Tests) and T008/T009 (Implementation) can be done in TDD fashion.
- T013 (Docs) can be done anytime after US1.

## Implementation Strategy

### MVP First (User Story 1)

1. **Refactor**: Extract Markdown logic to interface.
2. **Implement**: Add AsciiDoc logic.
3. **Verify**: Test both formats.
