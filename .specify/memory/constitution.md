# Markdown GitHub Stars Updater Constitution

## Core Principles

### I. Go Idiomatic
Code must adhere to standard Go formatting (`gofmt`), conventions, and idioms. Prefer the standard library for core functionality (file I/O, string manipulation). Dependencies should be minimal and justified (e.g., official GitHub clients).

### II. CLI First
The primary interface is the Command Line.
- **Flags**: Use the standard `flag` package for configuration (e.g., `-dry-run`, `-out`).
- **IO**: Read from files or stdin; write to files or stdout.
- **Feedback**: Errors go to `stderr`; useful status info goes to `stdout`.

### III. Security & Configuration
- **Secrets**: **NEVER** accept secrets (tokens) as command-line arguments. Use environment variables (e.g., `GITHUB_TOKEN`).
- **Read-Only Default**: Operations should be safe by default. Non-destructive modes (like dry-run) are encouraged for testing.

### IV. Simplicity & Focus
The tool does one thing: updates Markdown files. Avoid feature creep into general Markdown linting or formatting unless directly related to the core value proposition (updating dynamic data).

## Development Workflow

1.  **Spec-First**: Define behavior in `specs/`.
2.  **Test**: Ensure changes are verifiable.
3.  **Build**: Use `go build` to verify compilation.
4.  **Lint**: Code should pass standard Go linters.

## Governance
This constitution guides the architectural and stylistic decisions of the `markdown-github-stars-updater` project.
**Version**: 1.0.0 | **Ratified**: 2026-02-03