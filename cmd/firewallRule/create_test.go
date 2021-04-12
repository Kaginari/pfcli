package firewallRule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createFirewallRuleCommand(t *testing.T) {
	cmd := CreateCmd

	Interface:=cmd.Flag("interface").Name
	assert.Equal(t,"interface",Interface,"Name flag for interface should be interface" )
	i:=cmd.Flag("interface").Shorthand
	assert.Equal(t,"i",i,"Shorthand flag for interface should be i" )


	Type:=cmd.Flag("type").Name
	assert.Equal(t,"type",Type,"Name flag for interface should be type" )
	tt:=cmd.Flag("type").Shorthand
	assert.Equal(t,"t",tt,"Shorthand flag for interface should be t" )


	protocol:=cmd.Flag("protocol").Name
	assert.Equal(t,"protocol",protocol,"Name flag for interface should be protocol" )
	p:=cmd.Flag("protocol").Shorthand
	assert.Equal(t,"p",p,"Shorthand flag for interface should be p" )


	ipprotocol:=cmd.Flag("ipprotocol").Name
	assert.Equal(t,"ipprotocol",ipprotocol,"Name flag for interface should be ipprotocol" )

	icmptype:=cmd.Flag("icmptype").Name
	assert.Equal(t,"icmptype",icmptype,"Name flag for interface should be icmptype" )

	src:=cmd.Flag("src").Name
	assert.Equal(t,"src",src,"Name flag for interface should be src" )


	srcport:=cmd.Flag("srcport").Name
	assert.Equal(t,"srcport",srcport,"Name flag for interface should be srcport" )


	dstport:=cmd.Flag("dstport").Name
	assert.Equal(t,"dstport",dstport,"Name flag for interface should be dstport" )


	disabled:=cmd.Flag("disabled").Name
	assert.Equal(t,"disabled",disabled,"Name flag for interface should be disabled" )



	descr:=cmd.Flag("descr").Name
	assert.Equal(t,"descr",descr,"Name flag for interface should be descr" )

	log:=cmd.Flag("log").Name
	assert.Equal(t,"log",log,"Name flag for interface should be log" )



	top:=cmd.Flag("top").Name
	assert.Equal(t,"top",top,"Name flag for interface should be top" )


	apply:=cmd.Flag("apply").Name
	assert.Equal(t,"apply",apply,"Name flag for interface should be apply" )

}