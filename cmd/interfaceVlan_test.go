package cmd

import "testing"

func Test_interfaceVlanCommand(t *testing.T) {
	cmd := InterfaceVLANCmd
	cmd.Execute()
}

