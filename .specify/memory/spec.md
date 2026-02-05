# Project Specification: GitHub Stars Updater

**Status**: Consolidated
**Last Updated**: 2026-02-05
**Input**: Combined from `001-initial-release` and `002-asciidoc-support`

## User Scenarios

### User Story 1 - Update Star Counts in Markdown (Priority: P1)
As a developer maintaining a list of libraries in Markdown (e.g., Awesome lists), I want to automatically update the star counts next to GitHub links so that the list remains accurate without manual checking.

### User Story 2 - Update Star Counts in AsciiDoc (Priority: P1)
As a technical writer using AsciiDoc, I want to automatically update GitHub star counts in my `.adoc` documentation so that my references remain accurate without manual maintenance.

### User Story 3 - Safety and Output Control (Priority: P2)
As a user, I want to preview changes without modifying the file (Dry Run) or save the updated list to a new file so I can ensure the output is correct before overwriting originals.

## Requirements

### Functional Requirements
- **FR-001 (Multi-Format Support)**: The system MUST support both Markdown (`.md`, `.markdown`) and AsciiDoc (`.adoc`, `.asciidoc`) files.
- **FR-002 (Link Detection)**: 
    - **Markdown**: Identify links in `[Title](https://github.com/owner/repo)` format.
    - **AsciiDoc**: Identify links in `link:https://github.com/owner/repo[Title]` and `https://github.com/owner/repo[Title]` formats.
- **FR-003 (Star Fetching)**: The system MUST fetch stargazer counts for identified repositories using the GitHub API via an authenticated client (`GITHUB_TOKEN`).
- **FR-004 (Star Formatting)**: The system MUST format star counts consistently:
    - Exact count for < 1,000 (e.g., `350`).
    - `X.Yk` for 1,000-9,999 (e.g., `1.2k`).
    - `Xk` for >= 10,000 (e.g., `12k`).
- **FR-005 (In-Place Injection)**: The system MUST inject the formatted star count (e.g., `(⭐1.2k)`) into the existing link text.
- **FR-006 (Safety Flags)**: 
    - `-dry-run`: Print updated content to stdout without modifying files.
    - `-out <path>`: Write output to a specific file instead of updating in-place.

### Key Entities
- **LinkUpdater (Interface)**: Defines the contract for finding repos and updating content for different file formats.
- **GitHub Client**: Authenticated wrapper for the GitHub API.
- **Formatter**: Shared logic for star count string representation.

## Success Criteria
- **SC-001**: Accurately updates 100% of valid GitHub links in supported file formats.
- **SC-002**: Correctly detects file type by extension and applies appropriate parsing logic.
- **SC-003**: Respects GitHub API rate limits.
- **SC-004**: Ensures zero regressions in Markdown support when adding new formats.
