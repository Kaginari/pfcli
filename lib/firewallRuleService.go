package lib

import (
	"encoding/json"
	"github.com/Kaginari/pfcli/models"
	"io/ioutil"
	"log"
)

const FIREWALL_RULES_API_URI = "v1/firewall/rule"

type FirewallRulesResponse struct {

	Status  	string        `json:"status"`
	Code    	int64         `json:"code"`
	Return  	int64 `        json:"return"`
	Message 	string        `json:"message"`
	Date    	interface{}   `json:"data"`
}
type FirewallRulesResponseList struct {

	FirewallRulesResponse
	Date []interface{}  `json:"data"`
}

type IFirewallRuleService interface {
		Create(model models.FirewallRule) (FirewallRulesResponse , error)
		Delete(model models.DeleteFirewallRule) (FirewallRulesResponse , error)
		List() (FirewallRulesResponseList , error)

}
type FirewallRuleServiceImp struct {
	client PfClient
	path string
}


func FirewallRuleServiceConstruct (client PfClient) *FirewallRuleServiceImp{
	return &FirewallRuleServiceImp{
		client: client ,
		path:   FIREWALL_RULES_API_URI,
	}
}
var _ IFirewallRuleService = &FirewallRuleServiceImp{}
var firewallRule FirewallRulesResponse
var firewallRuleList FirewallRulesResponseList
func (f FirewallRuleServiceImp) Create(model models.FirewallRule) (FirewallRulesResponse , error) {

		body, _ := json.Marshal(model)
		response , err := f.client.Post(f.path ,body)
		if err != nil {
			return firewallRule , err
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
		log.Fatal(err)
		}
		_ = json.Unmarshal(bytes , &firewallRule)
	  // TODO PARSE RESPONSE AND RETURN THE STRUCT
	   return firewallRule, nil
}

func (f FirewallRuleServiceImp) Delete(model models.DeleteFirewallRule)  (FirewallRulesResponse , error){

	body, _ := json.Marshal(model)
	response , err := f.client.Delete(f.path ,body)
	if err != nil {
		return firewallRule , err
	}
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &firewallRule)
	return firewallRule, nil
}

func (f FirewallRuleServiceImp) List() (FirewallRulesResponseList , error) {

	response , err := f.client.Get(f.path ,nil)
	if err != nil {
		return firewallRuleList , err
	}
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &firewallRuleList)
	return firewallRuleList, nil
}
