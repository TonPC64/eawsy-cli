package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "eawsy"
	app.Usage = "cli for create project of lambda with golang from https://github.com/eawsy"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"-i"},
			Usage:   "initial project",
			Action: func(c *cli.Context) error {
				args := c.Args()
				if len(args) == 2 {
					projectType := c.Args().First()
					projectName := args[1]
					dir, _ := os.Getwd()
					switch projectType {
					case "net":
						cmd := exec.Command("git", "clone", "https://github.com/kingkong64/eawsy-net-template.git", projectName)
						cmd.Run()
						removeUselessData(dir, projectName)
						fmt.Printf("cd %s\nglide install\n", projectName)
					}
				} else {
					fmt.Println("Init fail pls use \"eawsy init [project_type] [project_name]")
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func removeUselessData(dir string, projectName string) {
	gitFolderInProject := "./" + projectName + "/.git"
	rm := exec.Command("rm", "-rf", gitFolderInProject)
	rm.Run()
}
