# Go Paths and Environment - Quick Reference

## Key Environment Variables

| Variable   | Purpose                 | Default Location | Contains                    |
| ---------- | ----------------------- | ---------------- | --------------------------- |
| **GOROOT** | Go installation         | `/usr/local/go/` | Go compiler, standard tools |
| **GOPATH** | Go workspace            | `~/go/`          | Your binaries, module cache |
| **GOBIN**  | Binary install location | `~/go/bin/`      | Tools you install           |

## Directory Structure

```
GOROOT/                    # Go installation
├── bin/
│   ├── go                # Go compiler
│   ├── gofmt            # Code formatter
│   └── godoc            # Documentation tool

GOPATH/                   # Your workspace
├── bin/                 # Your installed tools (GOBIN)
│   ├── air              # Live reload tool
│   ├── golangci-lint    # Linter
│   └── myapp            # Your compiled programs
└── pkg/mod/             # Module cache
    └── github.com/gin-gonic/gin@v1.9.1/
```

## Essential Commands

### Module Management

```bash
go mod init myproject        # Create new module (creates go.mod)
go get package-name          # Add dependency to project
go mod download             # Download all deps to cache
go mod tidy                 # Clean up dependencies
```

### Binary Installation

```bash
go install package@latest   # Install CLI tool to GOBIN
go build                   # Compile current project
go run .                   # Compile and run current project
```

## PATH Setup (Add to ~/.bashrc or ~/.zshrc)

```bash
# Essential: Add GOBIN to PATH for global tool access
export PATH=$PATH:$(go env GOPATH)/bin

# Alternative explicit version:
export PATH=$PATH:~/go/bin
```

## How It All Works Together

### 1. **Standard Library** (net/http)

- Built into Go installation (GOROOT)
- Always available, no downloads needed
- Import: `import "net/http"`

### 2. **External Dependencies** (gin)

- Downloaded to module cache (GOPATH/pkg/mod/)
- Shared across all projects
- Import: `import "github.com/gin-gonic/gin"`

### 3. **Development Tools** (air, linters)

- Installed to GOBIN (GOPATH/bin/)
- Available globally if PATH is set
- Install: `go install github.com/air-verse/air@latest`

## Production Server Timeouts

### Critical Settings:

```go
server := &http.Server{
    Addr:    ":8080",
    Handler: ginRouter,

    // Security & Performance
    ReadTimeout:       15 * time.Second,  // Max time to read request
    WriteTimeout:      15 * time.Second,  // Max time to write response
    IdleTimeout:       60 * time.Second,  // Keep-alive idle time
    ReadHeaderTimeout: 5 * time.Second,   // Headers must arrive quickly
    MaxHeaderBytes:    1 << 20,           // 1MB header limit
}
```

### Timeout Purpose:

- **ReadTimeout**: Prevents slow request attacks, handles slow clients
- **WriteTimeout**: Prevents hanging responses, handles slow networks
- **IdleTimeout**: Manages keep-alive connections efficiently

## Quick Troubleshooting

### "Command not found" after go install:

```bash
# Check if binary exists
ls $(go env GOPATH)/bin/

# Fix: Add GOBIN to PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### Check current settings:

```bash
go env GOPATH    # Your workspace
go env GOBIN     # Where binaries install
go env GOROOT    # Go installation
```

### Verify PATH:

```bash
echo $PATH       # Should include ~/go/bin
which go         # Should show GOROOT/bin/go
which air        # Should show GOBIN/air (if installed)
```

## Development Workflow

### New Project:

```bash
mkdir myproject && cd myproject
go mod init github.com/username/myproject
go get github.com/gin-gonic/gin
```

### Install Tools (once):

```bash
go install github.com/air-verse/air@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Daily Development:

```bash
air                    # Live reload (from GOBIN)
golangci-lint run     # Linting (from GOBIN)
go build              # Uses cached deps (from module cache)
```

## Key Insights

1. **One cache, many projects**: All projects share the same module cache for efficiency
2. **Global tools**: Install development tools once, use everywhere
3. **Clean separation**: Project code vs dependencies vs tools
4. **Production ready**: Proper timeouts prevent attacks and resource exhaustion
5. **PATH is crucial**: Without it, you can't access your installed tools easily

This system makes Go development fast, secure, and scalable across multiple projects.
