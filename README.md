# npkg

`npkg` aims to provide some handy commands for managing packages from nixpkgs.
Under the hood, it use nix profile commands. The packages are managed with the
default profile of nix.

## Installation

```zsh
# install task with
#   brew install task
# or
#   nix profile install go-task
#
# make sure "echo $fpath" includes ~/.local/share/zsh/completions
# if not, add following line to .zshrc:
#   fpath+=("$HOME/.local/share/zsh/completions")

git clone https://github.com/zhengpd/npkg.git npkg
cd npkg && task build
```

## Usage

```text
$ npkg --help

NAME:
   npkg - A new cli application

USAGE:
   npkg [global options] command [command options] [arguments...]

COMMANDS:
   install, i  install package from nixpkgs to nix default profile
   list, l     list installed packages in nix default profile
   upgrade, u  upgrade installed package in nix default profile
   remove, r   remove installed package from nix default profile
   clean-up    clean up duplicate packages from nix default profile
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Credits

- [cli](https://github.com/urfave/cli): a useful package for building cli commands in Go
- [nixpkgs](https://github.com/NixOS/nixpkgs): a large software packages repository
- [nix](https://github.com/NixOS/nix): a powerful package manager
- [task](https://github.com/go-task/task): a task runner / build tool that aims to be simpler and easier to use
