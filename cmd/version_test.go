package cmd

import "testing"

func Test_VersionCommand(t *testing.T) {
	cmd := versionCmd()
	cmd.Execute()
}
