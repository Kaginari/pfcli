package natOneToOne

import (
	"fmt"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",

	RunE: func(cmd *cobra.Command, args []string) error{
		config := lib.GetConfig()
		pfClient := config.Context()
		service  := lib.NatOneToOneServiceConstruct(pfClient)
		_ , err := service.Create(ModelnatOneToOne)
		if err != nil {
			fmt.Println("un error est occurred")
			return err
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		//jsonRes, _ := json.Marshal(res)
		//rest := functions.JsonOutput(jsonRes)

		return nil
	},
}

var ModelnatOneToOne models.NatOneToOne

func init() {
	createFlags()
}

func createFlags()  {

	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&ModelnatOneToOne.Interface, "interface", "i","LAN",models.NAT121InterfaceDesc )
	CreateCmd.MarkPersistentFlagRequired("interface")

	pf.StringVarP(&ModelnatOneToOne.Source, "source", "s","", models.NAT121SourceDesc)
	CreateCmd.MarkPersistentFlagRequired("source")

	pf.StringVarP(&ModelnatOneToOne.Destination, "destination", "d","", models.NAT121DestinationDesc)
	CreateCmd.MarkPersistentFlagRequired("destination")

	pf.StringVarP(&ModelnatOneToOne.External, "external", "e","", models.NAT121ExternalDesc)
	CreateCmd.MarkPersistentFlagRequired("external")

	pf.StringVarP(&ModelnatOneToOne.NatReflection, "nat-reflection", "n","", models.NAT121NatReflectionDesc)
	pf.StringVar(&ModelnatOneToOne.Description, "description", "", models.NAT121DescriptionDesc)
	pf.BoolVar(&ModelnatOneToOne.Disabled, "disabled", false, models.NAT121DisableDesc)
	pf.BoolVar(&ModelnatOneToOne.NoBinat, "no-binat", false, models.NAT121NoBinatDesc)
	pf.BoolVarP(&ModelnatOneToOne.Top, "top","t", false, models.NAT121TopDesc)
	pf.BoolVarP(&ModelnatOneToOne.Apply, "apply","a", true, models.NAT121ApplyDesc)
}



//cmd
// pfcli nat_one_to_one create  --description Test_1:1_NAT_entry -d 172.2.5.6 --external 1.2.3.5 --interface WAN --nat-reflection enable --source any