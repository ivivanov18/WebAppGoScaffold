package main

import (
	"errors"
	"testing"
)

func TestValidateConf(t *testing.T) {
	testCases := []struct {
		conf ProjectConfiguration
		errs []error
	} {
		{
			conf: ProjectConfiguration{
				Name: "TestProject",
				LocalPath: "/c/work/project",
				RepoUrl: "https://github.com/project",
				StaticAssets: false,
			},
			errs: []error{},
		},
		{
			conf: ProjectConfiguration{
				StaticAssets: false,
			},
			errs: []error{
				errors.New("Project name cannot be empty\n"),
				errors.New("Project location cannot be empty\n"),
				errors.New("Project repository URL cannot be empty\n"),
			},
		},
	}

	for _, tc := range testCases {
		errs := validateConf(&tc.conf)

		if (len(tc.errs) == 0 && len(errs) != 0) {
			t.Errorf("Expected no errors, got %v", errs)
		}

		if (len(tc.errs) != 0) {
			for i, e := range tc.errs {
				if (errs[i] == nil || errs[i].Error() != e.Error()) {
					t.Errorf("Expected error: %v, Got: %v", e, errs[i])
				}
			}
		}
	}
}