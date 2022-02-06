package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGenerateScaffold(t *testing.T) {
	testCases := []struct {
		conf ProjectConfiguration
		expectedOutput string
	} {
		{
			conf:  ProjectConfiguration{
				Name: "TestProject",
				LocalPath: "/c/work/conf",
				RepoUrl: "https://github.com/repo-for-scaffold",
				StaticAssets: false,
			},
			expectedOutput: "Generating scaffold for project TestProject in /c/work/conf",
		},
	}

	bytBuf := new(bytes.Buffer)
	for _, tc := range testCases {
		generateScaffold(bytBuf, tc.conf)
		gotOutput := bytBuf.String()
		if strings.Index(gotOutput, tc.expectedOutput) == -1 {
			t.Errorf("Expected output: %s, Got: %s", tc.expectedOutput, gotOutput)
		}
	}
	bytBuf.Reset()
}