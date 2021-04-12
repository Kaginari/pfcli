package natOneToOne

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createNatOneToOneCommand(t *testing.T) {
	cmd := CreateCmd

    Interface:=cmd.Flag("interface").Name
	assert.Equal(t,"interface",Interface,"Name flag for interface should be interface" )
	i:=cmd.Flag("interface").Shorthand
	assert.Equal(t,"i",i,"Shorthand flag for interface should be i" )

	source:=cmd.Flag("source").Name
	assert.Equal(t,"source",source,"Name flag for interface should be source" )
	s:=cmd.Flag("source").Shorthand
	assert.Equal(t,"s",s,"Shorthand flag for source should be s" )

	destination:=cmd.Flag("destination").Name
	assert.Equal(t,"destination",destination,"Name flag for interface should be destination" )
	d:=cmd.Flag("destination").Shorthand
	assert.Equal(t,"d",d,"Shorthand flag for destination should be d" )

	external:=cmd.Flag("external").Name
	assert.Equal(t,"external",external,"Name flag for interface should be external" )
	e:=cmd.Flag("external").Shorthand
	assert.Equal(t,"e",e,"Shorthand flag for external should be e" )

	natReflection:=cmd.Flag("nat-reflection").Name
	assert.Equal(t,"nat-reflection",natReflection,"Name flag for interface should be nat-reflection" )
	n:=cmd.Flag("nat-reflection").Shorthand
	assert.Equal(t,"n",n,"Shorthand flag for nat-reflection should be n" )


	description:=cmd.Flag("description").Name
	assert.Equal(t,"description",description,"Name flag for nat-reflection should be nat-reflection " )


	disabled:=cmd.Flag("disabled").Name
	assert.Equal(t,"disabled",disabled,"Name flag for nat-reflection should be disabled" )


	noBbinat:=cmd.Flag("no-binat").Name
	assert.Equal(t,"no-binat",noBbinat,"Name flag for nat-reflection should be no-binat" )

	top:=cmd.Flag("top").Name
	assert.Equal(t,"top",top,"Name flag for nat-reflection should be top" )
	tt:=cmd.Flag("top").Shorthand
	assert.Equal(t,"t",tt,"Shorthand flag for nat-reflection should be t" )


	apply:=cmd.Flag("apply").Name
	assert.Equal(t,"apply",apply,"Name flag for nat-reflection should be apply" )
	a:=cmd.Flag("apply").Shorthand
	assert.Equal(t,"a",a,"Shorthand flag for nat-reflection should be a" )

}