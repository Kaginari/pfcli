package nat_one_2_one

import (
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"gitlab.com/nt-factory/2021/admin/pfcli/lib"
)

var CreateCmd = &cobra.Command{
	Use:   "create",

	Run: func(cmd *cobra.Command, args []string) {
		lib.CreateNatOneToOne(natOneToOne)
	},
}

var natOneToOne models.NatOneToOne

func init() {
	createFlags()
}

func createFlags()  {

	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&natOneToOne.Interface, "interface", "i","",models.NAT121InterfaceDesc )
	pf.StringVarP(&natOneToOne.Source, "source", "s","", models.NAT121SourceDesc)
	pf.StringVarP(&natOneToOne.Destination, "destination", "d","", models.NAT121DestinationDesc)
	pf.StringVarP(&natOneToOne.External, "external", "e","", models.NAT121ExternalDesc)
	pf.StringVarP(&natOneToOne.NatReflection, "nat-reflection", "n","", models.NAT121NatReflectionDesc)
	pf.StringVar(&natOneToOne.Description, "description", "", models.NAT121DescriptionDesc)
	pf.BoolVar(&natOneToOne.Disabled, "disabled", false, models.NAT121DisableDesc)
	pf.BoolVar(&natOneToOne.NoBinat, "no-binat", false, models.NAT121NoBinatDesc)
	pf.BoolVarP(&natOneToOne.Top, "top","t", false, models.NAT121TopDesc)
	pf.BoolVarP(&natOneToOne.Apply, "apply","a", true, models.NAT121ApplyDesc)

	//err := cobra.MarkFlagRequired(pf, "interface")
	//if err != nil {
	//	panic(err)
	//}
	//err = cobra.MarkFlagRequired(pf, "source")
	//if err != nil {
	//	fmt.Printf("dsdsdsd")
	//	panic(err)
	//}
	//err = cobra.MarkFlagRequired(pf, "destination")
	//if err != nil {
	//	panic(err)
	//}
	//err =cobra.MarkFlagRequired(pf, "nat-reflection")
	//if err != nil {
	//	panic(err)
	//}
	//err = cobra.MarkFlagRequired(pf, "description")
	//if err != nil {
	//	panic(err)
	//}
}



//cmd
// pfcli nat_one_to_one create  --description Test_1:1_NAT_entry -d 172.2.5.6 --external 1.2.3.5 --interface WAN --nat-reflection enable --source any