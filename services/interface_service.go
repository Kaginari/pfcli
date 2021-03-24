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

func ApplayInterface()  {
	
	req, err := http.NewRequest("POST", ViperReadConfig().Host+"v1/interface/apply", nil)
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
func CreateInterface(Interface_Model models.Interface )  {
	jsonReq, _ := json.Marshal(Interface_Model)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("POST", ViperReadConfig().Host+"v1/interface", bytes.NewBuffer(jsonReq))
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
func DeleteInterface(InterfaceDelete models.InterfaceDelete)  {
	jsonReq, _ := json.Marshal(InterfaceDelete)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("DELETE", ViperReadConfig().Host+"v1/interface", bytes.NewBuffer(jsonReq))
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

func InterfaceList()  {
	fmt.Println(ViperReadConfig().Host+"v1/interface")
	req, err := http.NewRequest("GET", ViperReadConfig().Host+"v1/interface", nil)
	req.Header.Add("Authorization", ViperReadConfig().ClientId+" "+ViperReadConfig().ClientToken)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response body :",resp.Body)
	fmt.Println("response Headers : ", resp.Header)
}
