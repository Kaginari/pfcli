package firewallRule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteFirewallRuleCommand(t *testing.T) {
	cmd := DeleteCmd

	tracker := cmd.Flag("tracker").Name
	assert.Equal(t, "tracker", tracker, "Name flag for interface should be tracker")
	tt := cmd.Flag("tracker").Shorthand
	assert.Equal(t, "t", tt, "Shorthand flag for interface should be t")

}