package VirtualIps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createVirtualIpsCommand(t *testing.T) {
	cmd := CreateCmd

	mode := cmd.Flag("mode").Name
	assert.Equal(t, "mode", mode, "Name flag for interface should be mode")
	m := cmd.Flag("mode").Shorthand
	assert.Equal(t, "m", m, "Shorthand flag for interface should be m")


	Interface:=cmd.Flag("interface").Name
	assert.Equal(t,"interface",Interface,"Name flag for interface should be interface" )
	i:=cmd.Flag("interface").Shorthand
	assert.Equal(t,"i",i,"Shorthand flag for interface should be i" )



	subnet:=cmd.Flag("subnet").Name
	assert.Equal(t,"subnet",subnet,"Name flag for interface should be subnet" )
	s:=cmd.Flag("subnet").Shorthand
	assert.Equal(t,"s",s,"Shorthand flag for interface should be s" )

	descr:=cmd.Flag("descr").Name
	assert.Equal(t,"descr",descr,"Name flag for interface should be descr" )
	d:=cmd.Flag("descr").Shorthand
	assert.Equal(t,"d",d,"Shorthand flag for interface should be d" )


	noexpand:=cmd.Flag("noexpand").Name
	assert.Equal(t,"noexpand",noexpand,"Name flag for interface should be noexpand" )
	n:=cmd.Flag("noexpand").Shorthand
	assert.Equal(t,"n",n,"Shorthand flag for interface should be n" )


	vhid:=cmd.Flag("vhid").Name
	assert.Equal(t,"vhid",vhid,"Name flag for interface should be vhid" )
	v:=cmd.Flag("vhid").Shorthand
	assert.Equal(t,"v",v,"Shorthand flag for interface should be v" )



	advbase:=cmd.Flag("advbase").Name
	assert.Equal(t,"advbase",advbase,"Name flag for interface should be advbase" )



	advskew:=cmd.Flag("advskew").Name
	assert.Equal(t,"advskew",advskew,"Name flag for interface should be advskew" )


	password:=cmd.Flag("password").Name
	assert.Equal(t,"password",password,"Name flag for interface should be password" )
	p:=cmd.Flag("password").Shorthand
	assert.Equal(t,"p",p,"Shorthand flag for interface should be p" )

}
