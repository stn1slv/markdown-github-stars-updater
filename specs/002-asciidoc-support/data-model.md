# Data Model: AsciiDoc Support

**Feature**: AsciiDoc Support
**Status**: Draft

## Core Interfaces

### FileProcessor
Interface to abstract parsing and updating logic for different file formats.

```go
type FileProcessor interface {
    // ProcessContent takes the raw file content and a map of repo URLs to star counts.
    // It returns the updated content with star counts injected.
    // Note: The actual fetching of stars happens outside to allow batching/caching if needed,
    // or we can pass the client in. For now, sticking to the current flow where we parse first.
    
    // Better approach matching current main.go logic:
    // 1. Parse to find Repos.
    // 2. (Main) Fetch Stars.
    // 3. Update Content.
    
    FindRepos(content string) ([]string, error)
    UpdateContent(content string, stars map[string]int) (string, error)
}
```

*Refinement*: The current `main.go` does `updateStarCounts` which parses AND fetches. To keep it simple and testable, separating "Find" and "Update" is cleaner, or we can keep a unified method that takes a `StarFetcher` function.

**Proposed Interface**:

```go
type LinkUpdater interface {
    // Update takes the content and a function to get stars for a repo URL.
    Update(ctx context.Context, content string, getStars func(repoURL string) (int, error)) (string, error)
}
```

## Entities

### MarkdownUpdater
Implements `LinkUpdater` for `.md` files.
*   **Regex**: `\[([^\]]+)\]\((https:\/\/github\.com\/[^\/)]+\/[^\/)]+)\)`

### AsciiDocUpdater
Implements `LinkUpdater` for `.adoc` files.
*   **Regex**: `(?:link:)?(https://github\.com/[^\[]+)\[([^\]]*)\]`

## File Type Detection

Simple mapping in `main`:

```go
var updaters = map[string]LinkUpdater{
    ".md":       &MarkdownUpdater{},
    ".markdown": &MarkdownUpdater{},
    ".adoc":     &AsciiDocUpdater{},
    ".asciidoc": &AsciiDocUpdater{},
}
```
