# tmux-repo-runner

Repo-Runner is a CLI tool for managing Git repositories in combination with TMUX.

## Installation

(TODO Instructions on how to install and build your application.)

## Usage

### Adding a Repository

To add a new repository:

Add a repository to the list of repositories to be managed by Repo-Runner.

```sh
rr --add <git-repo-url>
```

Open the repository in a new TMUX session.

```sh
rr run . --open
```

This command clones the given Git repository to the specified directory.

## Contributing

(Guidelines for contributing to the project.)

## License

(Your chosen license or "Unlicensed" if not applicable.)

## TODOS

Basic

- [x] `add`
- [x] `open`
- [ ] `remove`
- [ ] `status`
- [ ] `help`

Advanced

- [ ] config: Make session layout configurable
- [ ] config: User `config` in `~/.config/reporunner/config.json/toml` & Default Config
- [ ] Package and global install

Extra

- [ ] Add TUI
  - [charmbracelet/bubbles: TUI components for Bubble Tea 🫧](https://github.com/charmbracelet/bubbles/tree/master)
