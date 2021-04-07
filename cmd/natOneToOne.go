package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Kaginari/pfcli/cmd/natOneToOne"
)

var NatOneToOneCmd = &cobra.Command{
	Use:   "nat_one_to_one COMMAND",
	Short: "NAT 1-to-1 Mappings",
}


func init() {


}

func NatOneToOneAddCommands()  {
	NatOneToOneCmd.AddCommand(natOneToOne.CreateCmd)
	NatOneToOneCmd.AddCommand(natOneToOne.DeleteCmd)
	NatOneToOneCmd.AddCommand(natOneToOne.ListCmd)

}