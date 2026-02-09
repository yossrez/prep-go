# Setting up the Environment

Instal Go from [official website](https://go.dev), configure PATH, and set up workspace.
Configure editor with Go support (VS Code, GoLand, Vim/Nvim, Zed). Use modules for
dependency management. Verify installation with `go version` and test with simple program.

## Prefered way setting up environment

Use `.dotfiles` to manage cofigs files

### Linux

My prefered linux distro is [Fedora Workstation](https://fedoraproject.org/).

Steps:

- Setup [homebrew](https://brew.sh/)
- Install go with homebrew
- Manage multi version go the [go way](https://go.dev/doc/manage-install)
- Install Code Editor/IDE (Zed, VSCode), the rest will automatically
setup by the editor when opening go project.

### Windows

I use [WSL2 and Fedora](https://learn.microsoft.com/en-us/windows/wsl/install)
for the distro, the rest set is the same as Linux setup above.
