package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

var platforms = map[string]string{
	"arm64": "aarch64",
}

var platform = fmt.Sprintf("%s-%s", platforms[runtime.GOARCH], runtime.GOOS)

func build_pkg_path(name string) string {
	return fmt.Sprintf("legacyPackages.%s.%s", platform, name)
}

func get_local_pkgs() []string {
	output_bytes, err := exec.Command("nix", "profile", "list").Output()
	if err != nil {
		log.Fatal(err)
	}

	output_lines := strings.Split(strings.TrimRight(string(output_bytes[:]), "\n"), "\n")
	var names []string
	for _, line := range output_lines {
		pkg_path := strings.Split(line, " ")[1]
		splits := strings.Split(pkg_path, ".")
		pkg_name := splits[len(splits)-1]
		names = append(names, pkg_name)
	}

	return names
}

func already_installed(name string) bool {
	for _, item := range get_local_pkgs() {
		if item == name {
			return true
		}
	}

	return false
}

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "install package from nixpkgs to nix default profile",
				Action: func(cCtx *cli.Context) error {
					name := cCtx.Args().First()

					if already_installed(name) {
						fmt.Println("[Skipped]", name, "is already installed.")
						return nil
					}

					fmt.Println("installing", name)
					pkg := "nixpkgs#" + name
					cmd := exec.Command("nix", "profile", "install", pkg)
					cmd.Stdout = os.Stdout
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list installed packages in nix default profile",
				Action: func(cCtx *cli.Context) error {
					names := get_local_pkgs()

					sort.Slice(names, func(x, y int) bool {
						return names[x] < names[y]
					})

					for _, name := range names {
						fmt.Println(name)
					}

					return nil
				},
			},
			{
				Name:    "upgrade",
				Aliases: []string{"u"},
				Usage:   "upgrade installed package in nix default profile",
				Action: func(cCtx *cli.Context) error {
					pkg_name := cCtx.Args().First()
					fmt.Println("upgrading", pkg_name)
					cmd := exec.Command("nix", "profile", "upgrade", build_pkg_path(pkg_name))
					cmd.Stdout = os.Stdout
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "remove installed package from nix default profile",
				Action: func(cCtx *cli.Context) error {
					pkg_name := cCtx.Args().First()
					fmt.Println("removing", pkg_name)
					cmd := exec.Command("nix", "profile", "remove", build_pkg_path(pkg_name))
					cmd.Stdout = os.Stdout
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
