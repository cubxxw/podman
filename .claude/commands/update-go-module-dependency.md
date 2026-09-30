---
name: update-go-module-dependency
description: Workflow command scaffold for update-go-module-dependency in podman.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /update-go-module-dependency

Use this workflow when working on **update-go-module-dependency** in `podman`.

## Goal

Updates a Go module dependency to a new version, including go.mod, go.sum, and vendored files.

## Common Files

- `go.mod`
- `go.sum`
- `vendor/<dependency>/*`
- `vendor/modules.txt`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Update go.mod to specify the new version of the dependency.
- Update go.sum to reflect new dependency checksums.
- Update vendor/ directory to include new or changed files from the dependency.
- Update vendor/modules.txt to reflect the new dependency state.

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.