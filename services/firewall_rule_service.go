package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gitlab.com/nt-factory/2021/admin/pfcli/functions"
	"gitlab.com/nt-factory/2021/admin/pfcli/models"
	"log"
	"net/http"
)

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