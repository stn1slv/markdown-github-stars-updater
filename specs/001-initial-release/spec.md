# Feature Specification: Initial Release

**Feature Branch**: `main` (Retroactive)
**Created**: 2026-02-03
**Status**: Implemented
**Input**: Existing Codebase (`main.go`, `README.md`)

## User Scenarios & Testing

### User Story 1 - Update Star Counts in Markdown (Priority: P1)
As a developer maintaining a list of libraries (e.g., Awesome lists), I want to automatically update the star counts next to GitHub links so that the list remains accurate without manual checking.

**Why this priority**: Core value proposition.

**Independent Test**:
1. Create a `test.md` with `[Repo](https://github.com/google/go-github)`.
2. Run the tool: `export GITHUB_TOKEN=...; ./updater test.md`.
3. Verify `test.md` contains `[Repo (⭐12.3k)](...)` (or current count).

**Acceptance Scenarios**:
1. **Given** a Markdown file with `[Title](https://github.com/owner/repo)`, **When** the tool runs, **Then** the link text becomes `[Title (⭐123)](...)`.
2. **Given** a link already containing stars `[Title (⭐100)](...)`, **When** the tool runs, **Then** it updates to the new count `[Title (⭐105)](...)`.

### User Story 2 - Dry Run Mode (Priority: P2)
As a user, I want to preview changes without modifying the file to ensure the output looks correct before committing.

**Why this priority**: Safety mechanism for users.

**Independent Test**:
1. Run `./updater -dry-run test.md`.
2. Verify output is printed to console.
3. Verify `test.md` modification time has *not* changed.

**Acceptance Scenarios**:
1. **Given** a valid markdown file, **When** running with `-dry-run`, **Then** the updated content prints to stdout and the file remains unchanged.

### User Story 3 - Custom Output File (Priority: P3)
As a user, I want to save the updated list to a new file so I can compare it with the original.

**Why this priority**: Workflow flexibility.

**Acceptance Scenarios**:
1. **Given** `input.md`, **When** running with `-out output.md`, **Then** `input.md` is unchanged and `output.md` contains the updated links.

## Requirements

### Functional Requirements
- **FR-001**: The system MUST accept a GitHub Personal Access Token via the `GITHUB_TOKEN` environment variable.
- **FR-002**: The system MUST identify GitHub repository links in the format `[Title](https://github.com/owner/repo)`.
- **FR-003**: The system MUST fetch the stargazer count for each identified repository using the GitHub API.
- **FR-004**: The system MUST format star counts:
    - Exact count for < 1000 (e.g., `350`).
    - `X.Yk` for 1000-9999 (e.g., `1.2k`).
    - `Xk` for >= 10000 (e.g., `12k`).
- **FR-005**: The system MUST support a `-dry-run` flag.
- **FR-006**: The system MUST support an `-out` flag to specify a destination file.

## Success Criteria

### Measurable Outcomes
- **SC-001**: Accurately updates 100% of valid `github.com` links in a provided Markdown file.
- **SC-002**: Respects GitHub API rate limits (by nature of using the authenticated client).
