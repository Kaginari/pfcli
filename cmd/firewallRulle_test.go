package cmd

import "testing"

func Test_firewallRuleCommand(t *testing.T) {
	cmd := FirewallRuleCmd
	cmd.Execute()
}
