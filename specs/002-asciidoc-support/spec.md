# Feature Specification: AsciiDoc Support

**Feature Branch**: `002-asciidoc-support`
**Created**: 2026-02-03
**Status**: Completed
**Input**: User description: "I want to add support of AsciiDocs. It should work exactly as markdown support."

## User Scenarios & Testing

### User Story 1 - Update Star Counts in AsciiDoc (Priority: P1)

As a technical writer using AsciiDoc, I want to automatically update GitHub star counts in my `.adoc` documentation so that my references remain accurate without manual maintenance.

**Why this priority**: This is the core functionality requested.

**Independent Test**:
1. Create `test.adoc` with content: `https://github.com/google/go-github[Go Github Client]`.
2. Run command: `./updater test.adoc`.
3. Verify `test.adoc` content updates to: `https://github.com/google/go-github[Go Github Client (⭐12.3k)]` (or current count).

**Acceptance Scenarios**:

1. **Given** an AsciiDoc file with `link:https://github.com/owner/repo[Title]` syntax, **When** executing the tool, **Then** the link text is updated to `[Title (⭐NUM)]`.
2. **Given** an AsciiDoc file with implicit `https://github.com/owner/repo[Title]` syntax, **When** executing the tool, **Then** the link text is updated to `[Title (⭐NUM)]`.
3. **Given** a file with existing stars `[Title (⭐100)]`, **When** executing the tool, **Then** the count is updated to the current value.

### User Story 2 - Safety and Output Control (Priority: P2)

As a user, I want to use the same safety flags (`-dry-run`, `-out`) available for Markdown when processing AsciiDoc files to prevent accidental data loss.

**Why this priority**: Ensures consistent user experience and safety across file types.

**Independent Test**:
1. Run `./updater -dry-run test.adoc`.
2. Verify output prints to stdout and file is unchanged.

**Acceptance Scenarios**:

1. **Given** an AsciiDoc file, **When** executing with `-dry-run`, **Then** the updated content is printed to stdout and the file remains untouched.
2. **Given** an AsciiDoc file, **When** executing with `-out new.adoc`, **Then** the original file is untouched and `new.adoc` contains the updates.

## Requirements

### Functional Requirements

- **FR-001**: The CLI MUST accept files with AsciiDoc extensions (e.g., `.adoc`, `.asciidoc`) as input arguments.
- **FR-002**: The system MUST identify GitHub repository links in common AsciiDoc formats:
    - Macro form: `link:https://github.com/user/repo[Text]`
    - Inline form: `https://github.com/user/repo[Text]`
- **FR-003**: The system MUST inject the formatted star count (e.g., `(⭐1.2k)`) into the link text area (inside the `[]` brackets) for AsciiDoc links.
- **FR-004**: The system MUST apply the same star formatting logic (<1k, 1k-10k, >10k) used for Markdown.
- **FR-005**: The system MUST support existing flags (`-dry-run`, `-out`) for AsciiDoc processing identical to Markdown processing.

### Key Entities

- **Link Parser**: Logic to identify and extract URLs and Link Text, now needing to support multiple syntaxes (Markdown vs AsciiDoc) based on file context.

## Success Criteria

### Measurable Outcomes

- **SC-001**: Accurately updates 100% of valid GitHub links in a provided AsciiDoc file.
- **SC-002**: Zero regression in Markdown processing capabilities.
- **SC-003**: Tool automatically detects file type or applies appropriate parsing without requiring a specific `--format` flag (inferred from extension).