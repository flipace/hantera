package develop

import "github.com/urfave/cli"

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
				UsageText:   "hantera develop setup --config ./manifest.yml --target ./project",
				Description: "sets up a project for development",
				Action:      Setup,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "config, c, manifest",
						Usage: "load manifest/config from `FILE.yml`",
						Value: "./manifest.yml",
					},
					cli.StringFlag{
						Name:  "target, t",
						Usage: "setup project in `PATH`",
					},
					cli.StringFlag{
						Name:  "branch, b",
						Usage: "branch to checkout",
						Value: "refs/heads/develop",
					},
					cli.BoolFlag{
						Name:  "progress, p",
						Usage: "show clone progress",
					},
				},
			},
			{
				Name:        "install-dependencies",
				Aliases:     []string{"install-deps", "id"},
				Usage:       "hantera develop install-deps",
				UsageText:   "hantera develop install-deps --config ./manifest.yml --target ./project",
				Description: "installs dependencies for a project",
				Action:      InstallDependencies,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "config, c, manifest",
						Usage: "load manifest/config from `FILE.yml`",
						Value: "./manifest.yml",
					},
					cli.StringFlag{
						Name:  "target, t",
						Usage: "setup project in `PATH`",
					},
				},
			},
		},
	},
}
