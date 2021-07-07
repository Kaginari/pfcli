package lib

import (
	"encoding/json"
	"github.com/Kaginari/pfcli/models"
	"io/ioutil"
	"log"
)

const NAT_ONE_2_ONE ="v1/firewall/nat/one_to_one"

type NatOneToOneResponse struct {
	Status  	string        `json:"status"`
	Code    	int64         `json:"code"`
	Return  	int64 `        json:"return"`
	Message 	string        `json:"message"`
	Date    	interface{}   `json:"data"`
}
type NatOneToOneResponseList struct {
	NatOneToOneResponse
	Date    	[]interface{}   `json:"data"`
}
type InatOneToOneService interface {
	Create(model models.NatOneToOne) (NatOneToOneResponse , error)
	Delete(model models.DeleteNatOneToOne) (NatOneToOneResponse , error)
	List() (NatOneToOneResponseList , error)
}
type NatOneToOneServiceImp struct {
	client PfClient
	path string
}
func NatOneToOneServiceConstruct(c PfClient)*NatOneToOneServiceImp  {
	return &NatOneToOneServiceImp{
		client: c,
		path: NAT_ONE_2_ONE,
	}
}
var _ InatOneToOneService=NatOneToOneServiceImp{}
var NatOneToOne NatOneToOneResponse
var NatOneToOneList NatOneToOneResponseList
func (n NatOneToOneServiceImp) Create(model models.NatOneToOne) (NatOneToOneResponse, error) {
	body, _ := json.Marshal(model)
	response , err := n.client.Post(n.path ,body)
	if err != nil {
		return NatOneToOne , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &NatOneToOne)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return NatOneToOne, nil
}

func (n NatOneToOneServiceImp) Delete(model models.DeleteNatOneToOne) (NatOneToOneResponse, error) {
	body, _ := json.Marshal(model)
	response , err := n.client.Delete(n.path ,body)
	if err != nil {
		return NatOneToOne , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &NatOneToOne)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return NatOneToOne, nil
}

func (n NatOneToOneServiceImp) List() (NatOneToOneResponseList, error) {
	response , err := n.client.Get(n.path ,nil)
	if err != nil {
		return NatOneToOneList , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &NatOneToOneList)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return NatOneToOneList, nil
}


