package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Kaginari/pfcli/functions"
	"github.com/Kaginari/pfcli/models"
	"io/ioutil"
	"log"
	"net/http"
)

const VIRTUAL_IPS_API_URI = "v1/firewall/virtual_ip"
type VirtualIpsResponse struct {

	Status  	string        `json:"status"`
	Code    	int64         `json:"code"`
	Return  	int64 `        json:"return"`
	Message 	string        `json:"message"`
	Date    	interface{}   `json:"data"`
}
type VirtualIpsResponseList struct {

	VirtualIpsResponse
	Date []interface{}  `json:"data"`
}

type IVirtualIpsService interface {
	Create(model models.Virtual_IPS) (VirtualIpsResponse , error)
	Delete(model models.DeleteVirtualIPS) (VirtualIpsResponse , error)
	List() (VirtualIpsResponseList , error)

}
type VirtualIpsServiceImp struct {
	client PfClient
	path string
}
func VirtualIpsServiceConstruct (client PfClient) *VirtualIpsServiceImp{
	return &VirtualIpsServiceImp{
		client: client ,
		path:   VIRTUAL_IPS_API_URI,
	}
}
var _ IVirtualIpsService = &VirtualIpsServiceImp{}
var VirtualIps VirtualIpsResponse
var VirtualIpsList VirtualIpsResponseList

func (v VirtualIpsServiceImp) Create(model models.Virtual_IPS) (VirtualIpsResponse, error) {
	body, _ := json.Marshal(model)
	response , err := v.client.Post(v.path ,body)
	if err != nil {
		return VirtualIps , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &VirtualIps)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return VirtualIps, nil
}

func (v VirtualIpsServiceImp) Delete(model models.DeleteVirtualIPS) (VirtualIpsResponse, error) {
	body, _ := json.Marshal(model)
	response , err := v.client.Delete(v.path ,body)
	if err != nil {
		return VirtualIps , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &VirtualIps)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return VirtualIps, nil
}

func (v VirtualIpsServiceImp) List() (VirtualIpsResponseList, error) {
	response , err := v.client.Get(v.path ,nil)
	if err != nil {
		return VirtualIpsList , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &VirtualIpsList)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return VirtualIpsList, nil
}
