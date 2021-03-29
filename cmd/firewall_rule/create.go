package firewall_rule

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/lib"
	"github.com/Kaginari/pfcli/models"
)

var firewallRule models.FirewallRule

var CreateCmd = &cobra.Command{
	Use:   "create",
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		pfClient := config.Context()
		service  := lib.FirewallRuleServiceConstruct(pfClient)
		res , err := service.Create(firewallRule)
		if err != nil {
			fmt.Println("un error est occurred")
			// TODO FIN BETTER WAY TO HANDLE ERRORS
		}
		jsonRes, _ := json.Marshal(res)
		rest := functions.JsonOutput(jsonRes)

		fmt.Println(rest)
		created, _ := functions.Find(res.Date, "created")
		time , _ := functions.Find(created , "time")
		fmt.Println(time)

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


////cmd
////pfcli firewallRule create --descr test_pfcli --dst 172.12.25.18 --dstport 443  --icmptype kgndfjk  --interface wan  --protocol tcp/udp --src 172.16.209.138 --type block --ipprotocol inet --srcport 443 --top