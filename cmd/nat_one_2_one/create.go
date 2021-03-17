package nat_one_2_one

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/functions"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"io/ioutil"
	"log"
	"net/http"
)

var CreateCmd = &cobra.Command{
	Use:   "create",

	Run: func(cmd *cobra.Command, args []string) {
		jsonReq, _ := json.Marshal(natOneToOne)
		res := functions.JsonOutput(jsonReq)
		fmt.Println(res)

		//jsonReq, _ := json.Marshal(natOneToOne)
		fmt.Printf("%s \n" ,jsonReq)

		resp, err := http.Post("pfenseurl", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		// Convert response body to string
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)

	},
}

var natOneToOne models.NatOneToOne

func init() {
	createFlags()
}

func createFlags()  {

	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&natOneToOne.Interface, "interface", "i","",models.NAT121InterfaceDesc )
	pf.StringVarP(&natOneToOne.Interface, "source", "s","", models.NAT121SourceDesc)
	pf.StringVarP(&natOneToOne.Destination, "destination", "d","", models.NAT121DestinationDesc)
	pf.StringVarP(&natOneToOne.External, "external", "e","", models.NAT121ExternalDesc)
	pf.StringVarP(&natOneToOne.NatReflection, "nat-reflection", "n","", models.NAT121NatReflectionDesc)
	pf.StringVar(&natOneToOne.Description, "description", "", models.NAT121DescriptionDesc)
	pf.BoolVar(&natOneToOne.Disabled, "disabled", false, models.NAT121DisableDesc)
	pf.BoolVar(&natOneToOne.NoBinat, "no-binat", false, models.NAT121NoBinatDesc)
	pf.BoolVarP(&natOneToOne.Top, "top","t", false, models.NAT121TopDesc)
	pf.BoolVarP(&natOneToOne.Top, "apply","a", false, models.NAT121ApplyDesc)

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

