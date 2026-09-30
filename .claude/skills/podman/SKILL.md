```markdown
# podman Development Patterns

> Auto-generated skill from repository analysis

## Overview
This skill provides guidance on contributing to the `podman` Go codebase. It covers core coding conventions, file organization, and the main workflows for updating dependencies and fixing bugs with associated tests. The patterns documented here are based on observed practices in the repository.

## Coding Conventions

### File Naming
- Use **camelCase** for file names.
  - Example: `containerEngine.go`, `imageStore.go`

### Import Style
- Use **relative imports** for internal packages.
  - Example:
    ```go
    import (
        "fmt"
        "os"
        "github.com/containers/podman/v4/libpod"
    )
    ```

### Export Style
- Use **named exports** for functions, types, and variables intended for use outside the package.
  - Example:
    ```go
    // Exported function
    func NewContainerEngine() *ContainerEngine {
        // ...
    }

    // Unexported function (internal use)
    func newHelper() {
        // ...
    }
    ```

### Commit Patterns
- Commits are typically freeform, sometimes prefixed with `libpod`.
- Average commit message length is about 60 characters.
  - Example:  
    ```
    libpod: fix race condition in container removal
    ```

## Workflows

### Update Go Module Dependency
**Trigger:** When you need to update a third-party Go module dependency to a newer version  
**Command:** `/update-dependency`

1. Update `go.mod` to specify the new version of the dependency.
    ```sh
    go get example.com/dependency@v1.2.3
    ```
2. Update `go.sum` to reflect new dependency checksums.
    ```sh
    go mod tidy
    ```
3. Update the `vendor/` directory to include new or changed files from the dependency.
    ```sh
    go mod vendor
    ```
4. Update `vendor/modules.txt` to reflect the new dependency state (automatically handled by `go mod vendor`).

**Files involved:**
- `go.mod`
- `go.sum`
- `vendor/<dependency>/*`
- `vendor/modules.txt`

### Bugfix With Associated Test
**Trigger:** When you need to fix a bug and ensure it is covered by a test  
**Command:** `/bugfix-with-test`

1. Modify the relevant Go source file to fix the bug.
    ```go
    // Before
    func calculateSum(a, b int) int {
        return a - b
    }

    // After
    func calculateSum(a, b int) int {
        return a + b
    }
    ```
2. Update or add a corresponding `*_test.go` file to test the fix.
    ```go
    // calculateSum_test.go
    func TestCalculateSum(t *testing.T) {
        result := calculateSum(2, 3)
        if result != 5 {
            t.Errorf("Expected 5, got %d", result)
        }
    }
    ```

**Files involved:**
- `libpod/*.go`
- `libpod/*_test.go`

## Testing Patterns

- Test files follow the pattern: `*_test.go`
- The testing framework is not explicitly specified, but standard Go testing is assumed.
- Example test file:
    ```go
    // containerEngine_test.go
    import "testing"

    func TestContainerStart(t *testing.T) {
        // test logic here
    }
    ```

## Commands

| Command               | Purpose                                                    |
|-----------------------|------------------------------------------------------------|
| /update-dependency    | Update a Go module dependency and all related files        |
| /bugfix-with-test     | Apply a bugfix and ensure it is covered by a test          |
```
