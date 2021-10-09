package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type ProjectConfiguration struct {
	projectName string
	locationOnDisk string
	remoteRepoUrl string
	staticAssets bool
}

func setupParseFlags(w io.Writer, args []string) ProjectConfiguration{
	config := ProjectConfiguration{}
	flag.StringVar(&config.projectName, "n", "", "Project name")
	flag.StringVar(&config.locationOnDisk, "d", "", "Project location on disk")
	flag.StringVar(&config.remoteRepoUrl, "r", "", "Project remote repository")
	flag.BoolVar(&config.staticAssets, "s", false, "Project will have static assets or not")
	flag.Parse()

	// fs := flag.NewFlagSet("scaffold-gen", flag.ContinueOnError)
	// err := fs.Parse(args)

	// if fs.NArg() != 0 {
	// 	return config, errors.New("No positional parameters expected")
	// }
	// return config, err
	return config
}

func validateConf(config *ProjectConfiguration) []error {
	var validationErrors []error

	if len(config.projectName) == 0 {
		validationErrors = append(validationErrors, errors.New("Project name cannot be empty"))
	}
	if len(config.locationOnDisk) ==  0 {
		validationErrors = append(validationErrors, errors.New("Project location cannot be empty"))
	}
	if len(config.remoteRepoUrl) == 0 {
		validationErrors = append(validationErrors, errors.New("Project repository URL cannot be empty"))
	}

	return validationErrors
}

func main() {
	config := setupParseFlags(os.Stdout, os.Args[1:])
	validationErrors := validateConf(&config)

	if (len(validationErrors) > 0) {
		for _ , error := range validationErrors {
			fmt.Printf("%s\n", error)
		}
		os.Exit(1)
	}
}