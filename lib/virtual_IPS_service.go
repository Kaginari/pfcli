package lib

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

func CreateVirtualIPS(VirtualIPS models.Virtual_IPS)  {
	jsonReq, _ := json.Marshal(VirtualIPS)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("POST", ViperReadConfig().Host+"v1/firewall/virtual_ip", bytes.NewBuffer(jsonReq))
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


func DeleteVirtual_IPS(model models.DeleteVirtualIPS)  {
	jsonReq, _ := json.Marshal(model)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("DELETE", ViperReadConfig().Host+"v1/firewall/virtual_ip", bytes.NewBuffer(jsonReq))
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

func VrirtualIPSList()  {
	req, err := http.NewRequest("GET", ViperReadConfig().Host+"v1/firewall/virtual_ip", nil)
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
	fmt.Println("response Body : ", resp.Body)
	fmt.Println("response Headers : ", resp.Header)
}