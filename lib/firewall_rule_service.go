package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gitlab.com/nt-factory/2021/admin/pfcli/functions"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"io/ioutil"
	"log"
	"net/http"
)

const FIREWALL_RULES_API_URI = "v1/firewall/rule"

type FirewallRulesResponse struct {

	Status string `json:"status"`
	Code int64 `json:"code"`
	Return int64 `json:"return"`
	Message string `json:"message"`
	Date interface{}  `json:"data"`
}

type IFirewallRuleService interface {
	Create(model models.FirewallRule) (FirewallRulesResponse , error)
	Delete(model models.DeleteFirewallRule) (*FirewallRulesResponse , error)
	List() (*FirewallRulesResponse , error)

}
type FirewallRuleService struct {
	client PfClient
	path string
}
var _ IFirewallRuleService = &FirewallRuleService{}

func FirewallRuleServiceConstruct (client PfClient) *FirewallRuleService{
	return &FirewallRuleService{
		client: client ,
		path:   FIREWALL_RULES_API_URI,
	}
}

func (f FirewallRuleService) Create(model models.FirewallRule) (FirewallRulesResponse , error) {
		body, _ := json.Marshal(model)
		response , err := f.client.Post(f.path ,body)
		var firewallRule FirewallRulesResponse
		if err != nil {
			return firewallRule , err
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
		log.Fatal(err)
		}

		json.Unmarshal(bytes , &firewallRule)
	  // TODO PARSE RESPONSE AND RETURN THE STRUCT
	   return firewallRule, nil
}

func (f FirewallRuleService) Delete(model models.DeleteFirewallRule)  (*FirewallRulesResponse , error){
	body, _ := json.Marshal(model)
	_ , err := f.client.Delete(f.path ,body)
	if err != nil {
		return nil , err
	}
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return nil, nil
}

func (f FirewallRuleService) List() (*FirewallRulesResponse , error) {
	_ , err := f.client.Get(f.path ,nil)
	if err != nil {
		return nil , err
	}
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return nil, nil
}



func AddRule(firewallRule models.FirewallRule)  {
	jsonReq, _ := json.Marshal(firewallRule)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("POST", ViperReadConfig().Host+"v1/firewall/rule", bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", ViperReadConfig().ClientId + " "+ViperReadConfig().ClientToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response body : ", resp.Body)
	fmt.Println("response Headers : ", resp.Header)
}


func RuleList()  {
	req, err := http.NewRequest("GET", ViperReadConfig().Host+"v1/firewall/rule", nil)
	req.Header.Add("Authorization", ViperReadConfig().ClientId + " "+ViperReadConfig().ClientToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)


	}



	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
}


func DeletRule(DeleteModel models.DeleteFirewallRule)  {
	jsonReq, _ := json.Marshal(DeleteModel)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("DELETE", ViperReadConfig().Host+"v1/firewall/rule", bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", ViperReadConfig().ClientId + " "+ViperReadConfig().ClientToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)


	}

	fmt.Println("response Status : ", resp.Status)
	jsonReq2, _ := json.Marshal(resp.Body)
	res2 := functions.JsonOutput(jsonReq2)
	fmt.Println(res2)
	fmt.Println("response Headers : ", resp.Header)

}