package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ConfigSetCommand(t *testing.T) {
	cmd := SetCmd

	Interface:=cmd.Flag("host").Name
	assert.Equal(t,"host",Interface,"Name flag for interface should be host" )

	client_id:=cmd.Flag("client_id").Name
	assert.Equal(t,"client_id",client_id,"Name flag for interface should be client_id" )


	client_token:=cmd.Flag("client_token").Name
	assert.Equal(t,"client_token",client_token,"Name flag for interface should be client_token" )
}
