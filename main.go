package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build = "0"
)


func main() {
	app := cli.NewApp()
	app.Name = "pypi plugin"
	app.Usage = "pypi plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "repo",
			Usage: "pypi repo",
			EnvVar: "PLUGIN_REPOSITORY, PYPI_REPO, PYPI_URL",
			Value: "https://pypi.python.org/pypi/",
		},
		cli.StringFlag{
			Name: "username",
			Usage: "pypi username",
			EnvVar: "PLUGIN_USERNAME, PYPI_USERNAME, PYPI_KEY",
		},
		cli.StringFlag{
			Name: "password",
			Usage: "pypi password",
			EnvVar: "PLUGIN_PASSWORD, PYPI_PASSWORD, PYPI_SECRET",
		},
		cli.StringSliceFlag{
			Name: "distributions",
			Usage: "pypi distributions",
			EnvVar: "PLUGIN_DISTRIBUTIONS, PYPI_DISTRIBUTIONS",
		},
		cli.StringFlag{
			Name: "build.home",
			Usage: "build home directory",
			EnvVar: "HOME",
		},
		cli.StringFlag{
			Name: "build.workspace",
			Usage: "build workspace",
			EnvVar: "DRONE_WORKSPACE",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Build: Build{
			Workspace: c.String("build.workspace"),
			Home: c.String("build.home"),
		},
		Config: Config{
			Repo: c.String("repo"),
			Username: c.String("username"),
			Password: c.String("password"),
			Distributions: c.StringSlice("distributions"),
		},
	}

	return plugin.Exec()
}
