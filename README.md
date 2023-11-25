# repo-runner

`repo-runner` is a CLI tool for managing Git repositories in combination with TMUX.

## Installation

(TODO Instructions on how to install and build your application.)

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

## Contributing

(Guidelines for contributing to the project.)

## License

(Your chosen license or "Unlicensed" if not applicable.)

## TODOS

Basic

- [x] `add`
- [x] `open`
- [x] `remove`
- [ ] `status`
- [x] `help`

Advanced

- [x] config: Make session layout configurable
- [ ] config: User `config` in `~/.config/reporunner/config.json/toml` & Default Config
- [ ] Package and global install

Extra

- [ ] Add Interface when running `rr` without flag
  - [charmbracelet/bubbles: TUI components for Bubble Tea 🫧](https://github.com/charmbracelet/bubbles/tree/master)
