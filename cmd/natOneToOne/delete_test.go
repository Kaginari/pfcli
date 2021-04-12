package natOneToOne

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteNatOneToOneCommand(t *testing.T) {
	cmd := DeleteCmd

	id := cmd.Flag("id").Name
	assert.Equal(t, "id", id, "Name flag for interface should be id")
	i := cmd.Flag("id").Shorthand
	assert.Equal(t, "i", i, "Shorthand flag for interface should be i")


	apply := cmd.Flag("apply").Name
	assert.Equal(t, "apply", apply, "Name flag for interface should be apply")
}

