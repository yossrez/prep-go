# go command

[Primary tool](https://pkg.go.dev/cmd/go) for managing Go source code with
unified interface for compiling, testing, formatting, and managing dependencies.
Includes subcommands like `build`, `run`, `get` | `install`, `test`, `fmt`,
`mod`, `env`, `fix`. Handles the entire development workflow automatically.

## go build

[go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies) |
Compile packages and dependencies

```bash
go build [-o ouput] [build flags] [packages]
```

## go run

[go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program) |
Compile and run Go program

```bash
go run  [build flags] [-exec xprog] package [arguments...]
```

## go get or go install

[go get](https://pkg.go.dev/cmd/go#hdr-Add_dependencies_to_current_module_and_install_them) |
[go install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) |
Add dependencies to current module and install them

```bash
# In earlier versions of Go `go get` was used to build and install pkgs.
# Now `go get` is dedicated to adjusting dependencies in go.mod.
go get [-t] [-u] [-tool] [build flags] [packages]

# 'go install' may be used to build and commands insteads.
go install [build flags] [packages]
```

## go test

[go test](https://pkg.go.dev/cmd/go#hdr-Testing_flags) |
Ending a file's name with `_test.go` tells the `go test` command that this file
containts test functions.

```bash
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

### function signature

Function prefix with `Test` and take a pointer to the
testing package's `testing.T` type.

```go
package test

import "testing"

// where Xxx does not start with lower case letter
// and should have the signature. 
func TestXxx(t *testing.T) { ... }
````

## go fmt

[go fmt](https://pkg.go.dev/cmd/go#hdr-Gofmt__reformat__package_sources) |
gofmt (reformat) package sources (code)

```bash
go fmt [-n] [-x] [packages]
```

## go mod

[go mod](https://pkg.go.dev/cmd/go#hdr-Module_maintenance) |
Module maintenance

```bash
go mod <command> [arguments]
```

The commands are:

```bash
download    download modules to local cache
edit        edit go.mod from tools or scripts
graph       print module requirement graph
init        initialize new module in current directory
tidy        add missing and remove unused modules
vendor      make vendored copy of dependencies
verify      verify dependencies have expected content
why         explain why packages or modules are needed
```

## go env

[go env](https://pkg.go.dev/cmd/go#hdr-Environment_variables) |
Env prints Go environment information

```bash
go env [-json] [-changed] [-u] [-w] [var ...]
```

## go fix

[go fix](https://pkg.go.dev/cmd/go#hdr-Update_packages_to_use_new_APIs) |
Fix runs the Go fix command on the packages
named by the import paths.

```bash
go fix [-fix list] [packages]
```
