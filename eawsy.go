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
	app.Usage = "CLI for create project of lambda with golang from https://github.com/eawsy"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello friend!\nThis is Project for create aws-lambda-go from eawsy\nRead Doc at https://github.com/eawsy\n")
		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
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
					createProject(dir, projectName, projectType)
				} else {
					fmt.Println("Init fail pls use \"eawsy init [project_type] [project_name]")
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func createProject(dir string, projectName string, projectType string) {
	switch projectType {
	case "net":
		githubURL := "https://github.com/kingkong64/eawsy-net-template.git"
		cmd := exec.Command("git", "clone", githubURL, projectName)
		err := cmd.Run()
		if err == nil {
			removeUselessData(dir, projectName)
			fmt.Printf("cd %s\nglide install\n", projectName)
		}
	default:
		fmt.Printf("Project not found!!")
	}
}

func removeUselessData(dir string, projectName string) {
	gitFolderInProject := "./" + projectName + "/.git"
	rm := exec.Command("rm", "-rf", gitFolderInProject)
	rm.Run()
}
