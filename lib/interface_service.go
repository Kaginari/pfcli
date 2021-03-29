package lib

import (
	"encoding/json"
	"github.com/Kaginari/pfcli/models"
	"io/ioutil"
	"log"
)

const INTERFACE_API_URI = "v1/interface"

type InterfaceResponse struct {

	Status      string 			`json:"status"`
	Code 		int64 			`json:"code"`
	Return 		int64 			`json:"return"`
	Message 	string 			`json:"message"`
	Date 		interface{}  	`json:"data"`
}
type IinterfaceService interface {
	Apply()(InterfaceResponse,error)
	Create(model models.Interface )(InterfaceResponse,error)
	Delete(model models.InterfaceDelete )(InterfaceResponse,error)
	List()(InterfaceResponse,error)
}
type InterfaceServiceImp struct {
	client PfClient
	path string
}

func InterfaceServiceConstruct(client PfClient) *InterfaceServiceImp {
	return &InterfaceServiceImp{
		client : client,
		path : INTERFACE_API_URI,
	}
}
var _ IinterfaceService=&InterfaceServiceImp{}
var interfaceResponse InterfaceResponse

func (i InterfaceServiceImp) Apply() (InterfaceResponse, error) {
	response , err := i.client.Post(i.path+"/apply" ,nil)
	if err != nil {
		return interfaceResponse , err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &interfaceResponse)
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	return interfaceResponse, nil
}

func (i InterfaceServiceImp) Create(model models.Interface) (InterfaceResponse, error) {
	body,_:=json.Marshal(model)
	response ,err :=i.client.Post(i.path,body)
	if err!=nil{
		return interfaceResponse, err
	}
	bytes,err :=ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes,&interfaceResponse)
	return interfaceResponse,nil
}

func (i InterfaceServiceImp) Delete(model models.InterfaceDelete) (InterfaceResponse, error) {
	body,_:=json.Marshal(model)
	response,err:=i.client.Delete(i.path,body)
	if err!=nil{
		return interfaceResponse, err
	}
	bytes ,err :=ioutil.ReadAll(response.Body)
	if err!=nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes,&interfaceResponse)
	return interfaceResponse,nil
}

func (i InterfaceServiceImp) List() (InterfaceResponse, error) {
	response,err:=i.client.Get(i.path,nil)
	if err!=nil{
		return interfaceResponse, err
	}
	bytes ,err :=ioutil.ReadAll(response.Body)
	if err!=nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes, &interfaceResponse)
	return interfaceResponse,nil
}


