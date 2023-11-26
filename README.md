# repo-runner

`repo-runner` is a CLI tool for managing Git repositories in combination with TMUX.

## Installation

To install the CLI application globally, run the following script:

### Automatic installation

> [!NOTE]
>
> The install script will install the CLI application to `/usr/local/bin/rr`. So make sure this is in your `$PATH`.

> [!IMPORTANT]
>
> Make sure you have `go` installed on your machine.

```bash
# Clone the repository (SSH)
git clone git@github.com:nikbrunner/repo-runner.git

# `cd` into that folder
cd repo-runner

# Give the install script execution permissions
chmod +x ./install.sh

# Run the install script
./install.sh

# Check if the app is in your PATH
which rr
```

### Manual installation

> [!IMPORTANT]
>
> Make sure you have `go` installed on your machine.

```bash
# Clone the repository (SSH)
git clone git@github.com:nikbrunner/repo-runner.git

# `cd` into that folder
cd repo-runner

# Build the app
go build -o rr

# Move the app to your PATH
mv rr /usr/local/bin/rr
```

## Usage

### Adding a Repository

To add a new repository:

Add a repository to the list of repositories to be managed by `repo-runner`.
This command clones the given Git repository to the specified directory.

```sh
rr --add <git-repo-url>
```

### Opening a Repository

Open the repository in a new TMUX session.
This will let you pick a repository from the list of repositories to be managed by `repo-runner` to be opened in a new TMUX session.

```sh
rr run . --open
```

### Removing a Repository

Remove a repository from the list of repositories to be managed by `repo-runner`.
This will let you pick a repository from the list of repositories to be managed by `repo-runner` to be removed.

```sh
rr --remove
```

### Status

Show the status of all repositories managed by `repo-runner`.

```sh
rr --status
```

### Show Help

Lists all available commands and flags.

```sh
rr --help
```

## Testing

```sh
go test ./...
```

## Contributing

(Guidelines for contributing to the project.)

## License

(Your chosen license or "Unlicensed" if not applicable.)

## ROADMAP

### Done

- [x] `add`
- [x] `open`
- [x] `remove`
- [x] `status`
- [x] `help`
- [x] Default session layout should come from config
- [x] Improve `fzf` styling
- [x] Global install & build script

### Next

- [ ] `RepoRunnerGPT`
- [ ] Improve `--status` performance
- [ ] Sketch for TUI
- [ ] Improve test coverage
- [ ] `--reset` - `git fullreset` for picked repository
- [ ] `--reset-all` - `git fullreset` for all repositories
- [ ] config: Make session layout configurable
- [ ] config: Enable user config in `~/.config/reporunner/config.toml/.yml`

### Future

- [ ] Add TUI when running `rr` without flag
