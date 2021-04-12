package _interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteInterfaceCommand(t *testing.T) {
	cmd := DeleteCmd

	If := cmd.Flag("if").Name
	assert.Equal(t, "if", If, "Name flag for interface should be If")
	i := cmd.Flag("if").Shorthand
	assert.Equal(t, "i", i, "Shorthand flag for interface should be i")

}