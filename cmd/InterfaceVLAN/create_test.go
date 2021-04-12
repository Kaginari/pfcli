package InterfaceVLAN

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createInterfaceVlanCommand(t *testing.T) {
	cmd := CreateCmd

	If := cmd.Flag("if").Name
	assert.Equal(t, "if", If, "Name flag for interface should be if")
	i := cmd.Flag("if").Shorthand
	assert.Equal(t, "i", i, "Shorthand flag for interface should be i")


	tag := cmd.Flag("tag").Name
	assert.Equal(t, "tag", tag, "Name flag for interface should be tag")
	tt := cmd.Flag("tag").Shorthand
	assert.Equal(t, "t", tt, "Shorthand flag for interface should be t")


	pcp := cmd.Flag("pcp").Name
	assert.Equal(t, "pcp", pcp, "Name flag for interface should be pcp")
	p := cmd.Flag("pcp").Shorthand
	assert.Equal(t, "p", p, "Shorthand flag for interface should be p")



	descr := cmd.Flag("descr").Name
	assert.Equal(t, "descr", descr, "Name flag for interface should be descr")
	d := cmd.Flag("descr").Shorthand
	assert.Equal(t, "d", d, "Shorthand flag for interface should be d")

}

