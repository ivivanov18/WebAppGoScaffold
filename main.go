package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type ProjectConfiguration struct {
	Name string
	LocalPath string
	RepoUrl string
	StaticAssets bool
}

func setupParseFlags(w io.Writer, args []string) (ProjectConfiguration, error){
	config := ProjectConfiguration{}
	fs := flag.NewFlagSet("scaffold-gen", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&config.Name, "n", "", "Project name")
	fs.StringVar(&config.LocalPath, "d", "", "Project location on disk")
	fs.StringVar(&config.RepoUrl, "r", "", "Project remote repository")
	fs.BoolVar(&config.StaticAssets, "s", false, "Project will have static assets or not")

	err := fs.Parse(args)

	if err != nil {
		return ProjectConfiguration{}, err
	}

	if fs.NArg() != 0 {
		return config, errors.New("No positional parameters expected")
	}

	return config, err
}

func validateConf(config *ProjectConfiguration) []error {
	var validationErrors []error

	if len(config.Name) == 0 {
		validationErrors = append(validationErrors, errors.New("Project name cannot be empty\n"))
	}
	if len(config.LocalPath) ==  0 {
		validationErrors = append(validationErrors, errors.New("Project location cannot be empty\n"))
	}
	if len(config.RepoUrl) == 0 {
		validationErrors = append(validationErrors, errors.New("Project repository URL cannot be empty\n"))
	}

	return validationErrors
}

func generateScaffold(w io.Writer, conf ProjectConfiguration) {
	fmt.Fprintf(w, "Generating scaffold for project %s in %s\n", conf.Name, conf.LocalPath)
}

func main() {
	config, err := setupParseFlags(os.Stdout, os.Args[1:])

	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	validationErrors := validateConf(&config)

	if (len(validationErrors) > 0) {
		for _ , error := range validationErrors {
			fmt.Printf("%s", error)
		}
		os.Exit(1)
	}
	generateScaffold(os.Stdout, config)
}