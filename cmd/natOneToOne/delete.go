
package natOneToOne

import (
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient := config.Context()
		service  := lib.NatOneToOneServiceConstruct(pfClient)
		res , err := service.Delete(DeleteModel)
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
	},
}

func init() {
	pf := DeleteCmd.PersistentFlags()
	pf.StringVarP(&DeleteModel.Id, "id", "i", "", "Specify the rule tracker ID to delete")
	DeleteCmd.MarkPersistentFlagRequired("id")
	pf.BoolVar(&DeleteModel.Apply, "apply",true, models.NAT121ApplyDesc)
}
var DeleteModel models.DeleteNatOneToOne
//pfcli nat_one_to_one delete --id 0 --apply