package _interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createInterfaceCommand(t *testing.T) {
	cmd := CreateCmd

	If := cmd.Flag("if").Name
	assert.Equal(t, "if", If, "Name flag for interface should be If")
	i := cmd.Flag("if").Shorthand
	assert.Equal(t, "i", i, "Shorthand flag for interface should be i")



	enable:=cmd.Flag("enable").Name
	assert.Equal(t,"enable",enable,"Name flag for interface should be enable" )

	descr:=cmd.Flag("descr").Name
	assert.Equal(t,"descr",descr,"Name flag for interface should be descr" )
	d:=cmd.Flag("descr").Shorthand
	assert.Equal(t,"d",d,"Shorthand flag for interface should be d" )



	subnet:=cmd.Flag("subnet").Name
	assert.Equal(t,"subnet",subnet,"Name flag for interface should be subnet" )
	s:=cmd.Flag("subnet").Shorthand
	assert.Equal(t,"s",s,"Shorthand flag for interface should be s" )


	Type:=cmd.Flag("type").Name
	assert.Equal(t,"type",Type,"Name flag for interface should be type" )
	tt:=cmd.Flag("type").Shorthand
	assert.Equal(t,"t",tt,"Shorthand flag for interface should be t" )



	blockbogons:=cmd.Flag("blockbogons").Name
	assert.Equal(t,"blockbogons",blockbogons,"Name flag for interface should be blockbogons" )

}