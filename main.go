package main

import (
	"flag"
	"fmt"
	"os"
)

type ProjectConfiguration struct {
	projectName string
	locationOnDisk string
	remoteRepoUrl string
	staticAssets bool
}

var errors []string


func main() {
	config := ProjectConfiguration{}
	flag.StringVar(&config.projectName, "n", "", "Project name")
	flag.StringVar(&config.locationOnDisk, "d", "", "Project location on disk")
	flag.StringVar(&config.remoteRepoUrl, "r", "", "Project remote repository")
	flag.BoolVar(&config.staticAssets, "s", false, "Project will have static assets or not")
	flag.Parse()

	if config.projectName == "" {
		errors = append(errors, "Project name cannot be empty")
	}
	if config.locationOnDisk == "" {
		errors = append(errors, "Project location cannot be empty")
	}
	if config.remoteRepoUrl == "" {
		errors = append(errors, "Project repository URL cannot be empty")
	}

	if (len(errors) > 0) {
		for _ , error := range errors {
			fmt.Printf("%s\n", error)
		}
		os.Exit(1)
	}
}