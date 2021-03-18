package firewall_rule

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/nt-factory/2021/admin/pfcli/functions"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"log"
	"net/http"
)

var firewallRule models.FirewallRule

var CreateCmd = &cobra.Command{
	Use:   "create",
	//PreRunE: func(cmd *cobra.Command, args []string) error {
	//	return functions.CheckRequiredFlags(cmd.Flags())
	//},
	Run: func(cmd *cobra.Command, args []string) {
		AddRule()

		//jsonReq, _ := json.Marshal(natOneToOne)
		//fmt.Printf("%s \n" ,jsonReq)

		//resp, err := http.Post("pfenseurl", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//
		//defer resp.Body.Close()
		//bodyBytes, _ := ioutil.ReadAll(resp.Body)
		//
		//// Convert response body to string
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)

	},
}

func init()  {
	createFlags()
}

func createFlags() {
	pf := CreateCmd.PersistentFlags()
	pf.StringVarP(&firewallRule.Interface, "interface", "i", "", models.FWRInterfaceDesc)
	pf.StringVarP(&firewallRule.Type, "type", "t", "", models.FWRTypeDesc)
	pf.StringVarP(&firewallRule.Protocol, "protocol", "p", "", models.FWRProtocolDesc)
	pf.StringVar(&firewallRule.IPProtocol, "ipprotocol", "", models.FWRIPProtocolDesc)
	pf.StringVar(&firewallRule.IMCPType, "icmptype", "", models.FWRIMCPTypeDesc)
	pf.StringVar(&firewallRule.Source, "src", "", models.FWRSourceDesc)
	pf.StringVar(&firewallRule.Destination, "dst", "", models.FWRDestinationDesc)
	pf.IntVar(&firewallRule.SourcePort, "srcport", 0 , models.FWRSourcePortDesc)
	pf.IntVar(&firewallRule.DestinationPort, "dstport", 0, models.FWRDestinationPortDesc)
	//pf.StringVarP(&firewallRule.Gateway, "gateway","g","WANGW_2", models.FWRGatewayDesc)
	pf.BoolVar(&firewallRule.Disabled, "disabled",false, models.FWRDisabledDesc)
	pf.StringVar(&firewallRule.Description, "descr","", models.FWRDescriptionDesc)

	pf.BoolVar(&firewallRule.Log, "log",false, models.FWRLogDesc)
	pf.BoolVar(&firewallRule.Top, "top",false, models.FWRTopDesc)
	pf.BoolVar(&firewallRule.Disabled, "apply",true, models.FWRApplyDesc)

}
func AddRule()  {
	jsonReq, _ := json.Marshal(firewallRule)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("POST", functions.ViperReadConfig().UrlPfsense+"v1/firewall/rule", bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", functions.ViperReadConfig().ClientId + " "+functions.ViperReadConfig().ClientToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("response Status : ", resp.Body)
	fmt.Println("response Headers : ", resp.Header)
}