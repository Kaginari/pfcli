package natOneToOne

import (
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",

	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient := config.Context()
		service  := lib.NatOneToOneServiceConstruct(pfClient)
		res , err := service.Create(ModelnatOneToOne)
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
	},
}

var ModelnatOneToOne models.NatOneToOne

func init() {
	createFlags()
}

func createFlags()  {

	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&ModelnatOneToOne.Interface, "interface", "i","",models.NAT121InterfaceDesc )
	pf.StringVarP(&ModelnatOneToOne.Source, "source", "s","", models.NAT121SourceDesc)
	pf.StringVarP(&ModelnatOneToOne.Destination, "destination", "d","", models.NAT121DestinationDesc)
	pf.StringVarP(&ModelnatOneToOne.External, "external", "e","", models.NAT121ExternalDesc)
	pf.StringVarP(&ModelnatOneToOne.NatReflection, "nat-reflection", "n","", models.NAT121NatReflectionDesc)
	pf.StringVar(&ModelnatOneToOne.Description, "description", "", models.NAT121DescriptionDesc)
	pf.BoolVar(&ModelnatOneToOne.Disabled, "disabled", false, models.NAT121DisableDesc)
	pf.BoolVar(&ModelnatOneToOne.NoBinat, "no-binat", false, models.NAT121NoBinatDesc)
	pf.BoolVarP(&ModelnatOneToOne.Top, "top","t", false, models.NAT121TopDesc)
	pf.BoolVarP(&ModelnatOneToOne.Apply, "apply","a", true, models.NAT121ApplyDesc)
}



//cmd
// pfcli nat_one_to_one create  --description Test_1:1_NAT_entry -d 172.2.5.6 --external 1.2.3.5 --interface WAN --nat-reflection enable --source any