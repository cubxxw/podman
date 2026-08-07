---
name: bugfix-with-associated-test
description: Workflow command scaffold for bugfix-with-associated-test in podman.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /bugfix-with-associated-test

Use this workflow when working on **bugfix-with-associated-test** in `podman`.

## Goal

Implements a bugfix in a Go source file and updates or adds a corresponding test file to verify the fix.

## Common Files

- `libpod/*.go`
- `libpod/*_test.go`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Modify the relevant Go source file to fix the bug.
- Update or add a corresponding *_test.go file to test the fix.

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.