package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const BACKUP_API_URI = "v1/system/config"

type BackupResponse struct {

	Status  	string        `json:"status"`
	Code    	int64         `json:"code"`
	Return  	int64 `        json:"return"`
	Message 	string        `json:"message"`
	Date    	interface{}   `json:"data"`
}

type IBackupervice interface {
	Save() (BackupResponse , error)

}
type BackuperviceImp struct {
	client PfClient
	path string
}


func BackupServiceConstruct (client PfClient) *BackuperviceImp{
	return &BackuperviceImp{
		client: client ,
		path:   BACKUP_API_URI,
	}
}
var _ IBackupervice = &BackuperviceImp{}
var backupRes BackupResponse


func (b BackuperviceImp) Save() (BackupResponse, error) {

	response , err := b.client.Get(b.path ,nil)
	if err != nil {
		return backupRes , err
	}
	// TODO PARSE RESPONSE AND RETURN THE STRUCT
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(bytes , &backupRes)
	return backupRes, nil
}