package lib

import (
	"encoding/json"
	"github.com/Kaginari/pfcli/models"
	"io/ioutil"
	"log"
)

const INTERFACE_VLAN_API_URI = "v1/interface/vlan"
type InterfaceVlanResponse struct {
	Status  	string        `json:"status"`
	Code    	int64         `json:"code"`
	Return  	int64 `        json:"return"`
	Message 	string        `json:"message"`
	Date    	interface{}   `json:"data"`
}
type IinterfaceVlanService interface {
	Create(model models.InterfaceVLAN) (InterfaceVlanResponse , error)
	Delete(model models.InterfaceVlanDelete) (InterfaceVlanResponse , error)
	List() (InterfaceVlanResponse , error)
}
type InterfaceVlanServiceImp struct {
	client PfClient
	path string
}
func InterfaceVlanConstruct(c PfClient) *InterfaceVlanServiceImp   {
	return &InterfaceVlanServiceImp{
		client: c,
		path: INTERFACE_VLAN_API_URI,
	}
}
var _  IinterfaceVlanService=InterfaceVlanServiceImp{}
var InterfaceVlan InterfaceVlanResponse
func (i InterfaceVlanServiceImp) Create(model models.InterfaceVLAN) (InterfaceVlanResponse, error) {
	body, _ := json.Marshal(model)
	response , err := i.client.Post(i.path ,body)
	if err != nil {
		return InterfaceVlan , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &InterfaceVlan)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return InterfaceVlan, nil
}

func (i InterfaceVlanServiceImp) Delete(model models.InterfaceVlanDelete) (InterfaceVlanResponse, error) {
	body, _ := json.Marshal(model)
	response , err := i.client.Delete(i.path ,body)
	if err != nil {
		return InterfaceVlan , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &InterfaceVlan)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return InterfaceVlan, nil
}

func (i InterfaceVlanServiceImp) List() (InterfaceVlanResponse, error) {
	response , err := i.client.Get(i.path ,nil)
	if err != nil {
		return InterfaceVlan , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &InterfaceVlan)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return InterfaceVlan, nil
}
