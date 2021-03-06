package develop

import "github.com/urfave/cli"

func check(err error) {
	if err != nil {
		panic(&err)
	}
}

// default flags for various develop commands
var defaultFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "load manifest from `FILE.yml`",
		Value: "./hantera.yml",
	},
	cli.StringFlag{
		Name:  "target, t",
		Usage: "setup project in `PATH`",
		Value: "./",
	},
}

// Commands : describes all commands for "develop" environments
var Commands = []cli.Command{
	{
		Name:      "develop",
		Aliases:   []string{"d"},
		Usage:     "run a task for develop environments",
		ArgsUsage: " ",
		Subcommands: []cli.Command{
			{
				Name:        "setup",
				Usage:       "hantera develop setup",
				UsageText:   "hantera develop setup --config ./hantera.yml --target ./project",
				Description: "sets up a project for development",
				Action:      Setup,
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "branch, b",
							Usage: "branch to checkout",
							Value: "develop",
						},
						cli.BoolFlag{
							Name:  "progress, p",
							Usage: "show clone progress",
						},
						cli.BoolFlag{
							Name:  "no-deps, no-dep",
							Usage: "don't install dependencies (e.g. npm i)",
						},
					},
					defaultFlags...,
				),
			},
			{
				Name:        "install-dependencies",
				Aliases:     []string{"install-deps", "id"},
				Usage:       "hantera develop install-deps",
				UsageText:   "hantera develop install-deps --config ./hantera.yml --target ./project",
				Description: "installs dependencies for a project",
				Action:      CmdInstallDependencies,
				Flags: append(
					[]cli.Flag{},
					defaultFlags...,
				),
			},
			{
				Name:        "update",
				Usage:       "hantera develop update",
				UsageText:   "hantera develop update --config ./hantera.yml --target ./project",
				Description: "updated dependencies of a project",
				Action:      Update,
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "branch, b",
							Usage: "branch to checkout",
							Value: "develop",
						},
						cli.BoolFlag{
							Name:  "progress, p",
							Usage: "show progress",
						},
						cli.BoolFlag{
							Name:  "no-deps, no-dep",
							Usage: "don't install dependencies (e.g. npm i)",
						},
					},
					defaultFlags...,
				),
			},
			{
				Name:        "clean",
				Usage:       "hantera develop clean",
				UsageText:   "hantera develop clean --config ./hantera.yml --target ./project",
				Description: "clean/delete dependencies of a project",
				Action:      Clean,
				Flags: append(
					[]cli.Flag{},
					defaultFlags...,
				),
			},
		},
	},
}
