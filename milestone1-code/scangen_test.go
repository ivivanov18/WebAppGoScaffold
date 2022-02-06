package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestSetupParseFlags(t *testing.T) {
	testCases := []struct {
		args []string
		err error
		expectedConf ProjectConfiguration
		expectedOutputContains string
	} {
		{
			args: []string{"-n", "MyProject", "-d", "/path/to/dir", "-r", "github.com/username/myproject"},
			err: nil,
			expectedConf: ProjectConfiguration{
				Name: "MyProject",
				LocalPath: "/path/to/dir",
				RepoUrl: "github.com/username/myproject",
				StaticAssets: false,
			},
			expectedOutputContains: "",
		},
		{
			args:         []string{"foo"},
			err:          errors.New("No positional parameters expected"),
			expectedConf: ProjectConfiguration{},
		},
	}

	bytesBuf := new(bytes.Buffer)
	for _, tc := range testCases {
		c, err := setupParseFlags(bytesBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Errorf("Expected non nil error, got %v", err)
		}
		if tc.err != nil {
			if err == nil || err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %v, Got: %v", tc.err, err)
			}
		}

		if c != tc.expectedConf {
			t.Errorf("Expected:%#v Got: %#v", c, tc.expectedConf)
		}

		if len(tc.expectedOutputContains) != 0 {
			gotOutput := bytesBuf.String()
			if strings.Index(gotOutput, tc.expectedOutputContains) == -1 {
				t.Errorf("Expected output: %s, Got: %s", tc.expectedOutputContains, gotOutput)
			}
		}
		bytesBuf.Reset()
	}
}