package InterfaceVLAN

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteInterfaceVlanCommand(t *testing.T) {
	cmd := DeleteCmd

	vlanif := cmd.Flag("vlanif").Name
	assert.Equal(t, "vlanif", vlanif, "Name flag for interface should be vlanif")
	v := cmd.Flag("vlanif").Shorthand
	assert.Equal(t, "v", v, "Shorthand flag for interface should be v")


	id := cmd.Flag("id").Name
	assert.Equal(t, "id", id, "Name flag for interface should be id")
	i := cmd.Flag("id").Shorthand
	assert.Equal(t, "i", i, "Shorthand flag for interface should be i")
}