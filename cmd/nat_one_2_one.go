package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/cmd/nat_one_2_one"
)

var NatOneToOneCmd = &cobra.Command{
	Use:   "nat_one_to_one COMMAND",
	Short: "NAT 1-to-1 Mappings",
}


func init() {


}

func NatOneToOneAddCommands()  {
	NatOneToOneCmd.AddCommand(nat_one_2_one.CreateCmd)
	NatOneToOneCmd.AddCommand(nat_one_2_one.DeleteCmd)
	NatOneToOneCmd.AddCommand(nat_one_2_one.ListCmd)

}