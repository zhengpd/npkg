# npkg

`npkg` aims to provide some handy commands for managing packages from nixpkgs.
Under the hood, it use nix profile commands. The packages are managed with the
default profile of nix.

## TODO

- [ ] zsh completion

## Installation

```zsh
git clone https://github.com/zhengpd/npkg.git npkg
cd npkg && go install
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
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Credits

- [cli](https://github.com/urfave/cli): a useful package for building cli commands in Go
- [nixpkgs](https://github.com/NixOS/nixpkgs): a large software packages repository
- [nix](https://github.com/NixOS/nix): a powerful package manager
