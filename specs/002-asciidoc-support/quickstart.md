# Quickstart: AsciiDoc Support

## Usage

The tool now automatically detects AsciiDoc files based on the extension (`.adoc` or `.asciidoc`).

### Update an AsciiDoc file
```bash
export GITHUB_TOKEN=your_token
./markdown-github-stars-updater docs/manual.adoc
```

### Dry Run
Check what would change without modifying the file:
```bash
./markdown-github-stars-updater -dry-run docs/manual.adoc
```

### Output to new file
```bash
./markdown-github-stars-updater -out docs/manual-updated.adoc docs/manual.adoc
```

## Supported Syntax

The tool supports standard AsciiDoc links to GitHub:

1.  **Macro style**: `link:https://github.com/owner/repo[Title]`
    *   Becomes: `link:https://github.com/owner/repo[Title (⭐1.2k)]`
2.  **Inline style**: `https://github.com/owner/repo[Title]`
    *   Becomes: `https://github.com/owner/repo[Title (⭐1.2k)]`
