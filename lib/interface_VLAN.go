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

func AddVlan(InetrfaceVLAN models.InterfaceVLAN)  {
	jsonReq, _ := json.Marshal(InetrfaceVLAN)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("POST", ViperReadConfig().Host+"v1/interface/vlan", bytes.NewBuffer(jsonReq))
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

func DeleteVlan(DeleteInetrfaceVLAN models.DeleteIVlan){
	jsonReq, _ := json.Marshal(DeleteInetrfaceVLAN)
	res := functions.JsonOutput(jsonReq)
	fmt.Println(res)
	req, err := http.NewRequest("DELETE", ViperReadConfig().Host+"v1/interface/vlan", bytes.NewBuffer(jsonReq))
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

func VlanList()  {
	req, err := http.NewRequest("GET", ViperReadConfig().Host+"v1/interface/vlan", nil)
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