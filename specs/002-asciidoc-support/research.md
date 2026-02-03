# Research: AsciiDoc Support

**Feature**: AsciiDoc Support
**Status**: Complete

## Decision: Regular Expression Parsing

**Decision**: Use Go's `regexp` package to identify AsciiDoc links, similar to the existing Markdown implementation.

**Rationale**:
*   **Consistency**: Matches the current approach for Markdown.
*   **Simplicity**: Avoids the overhead of a full AsciiDoc parser (like Asciidoctor) which would be overkill for just finding links.
*   **Dependencies**: Keeps the project dependency-free (standard library only).

**Alternatives Considered**:
*   **Full AST Parser**: Using a library to parse AsciiDoc into an AST, modifying the AST, and rendering back.
    *   *Rejected*: Overly complex, introduces heavy dependencies, and might alter formatting/comments unrelated to links.
*   **Line-by-line Scanning**: Manually scanning strings without regex.
    *   *Rejected*: More verbose and error-prone than `regexp` for this specific pattern matching.

## Decision: Strategy Pattern for File Types

**Decision**: Refactor `main.go` to use a simple interface or strategy pattern for handling different file types.

**Rationale**:
*   **Extensibility**: Makes it easy to add other formats (e.g., reStructuredText) in the future.
*   **Clean Code**: Separates parsing logic from CLI orchestration and API interaction.

**Alternatives Considered**:
*   **If/Else in Main**: Just adding `if strings.HasSuffix(file, ".adoc")` blocks inside the main loop.
    *   *Rejected*: Leads to spaghetti code and makes `main` hard to test.

## AsciiDoc Link Syntax Research

**Macro Form**: `link:url[text]`
*   Regex: `link:(https://github\.com/[^\[]+)\[([^\]]*)\]`

**Inline Form**: `url[text]`
*   Regex: `(https://github\.com/[^\[]+)\[([^\]]*)\]` (Note: Need to be careful not to double-match or match other things. The macro form is safer to prioritize or combine carefully).

**Combined Regex Strategy**:
Ideally, capture the URL and the Text.
Pattern: `(?:link:)?(https://github\.com/[^\[]+)\[([^\]]*)\]`
*   `(?:link:)?`: Non-capturing group for optional "link:" prefix.
*   Group 1: URL (`https://github.com/...`)
*   Group 2: Text (`Title`)
