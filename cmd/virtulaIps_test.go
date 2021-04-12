package cmd

import "testing"

func Test_virtualIpsCommand(t *testing.T) {
	cmd := interfaceCmd
	cmd.Execute()
}

