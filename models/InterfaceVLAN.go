package models

type InterfaceVLAN struct {
	If	  string    `json:"if"`
	Tag   string    `json:"tag"`
	Pcp	  string    `json:"pcp"`
	Descr string	`json:"descr"`
}
type DeleteIVlan struct {
	Vlanif	string	 `json:"vlanif"`
	Id      string   `json:"id"`
}


var IvlanIfDesc ="Set the parent interface to add the new VLAN to\n"
var IvlanTagDesc ="Set the VLAN tag to add to the parent interface\n"
var IvlanPcpDesc ="Set a 802.1Q VLAN priority. Must be an integer value between 0 and 7 (optional)\n"
var IvlanDescrDesc ="Set a description for the new VLAN interface\n"

var IvlanVlanifDesc= "Delete VLAN by full interface name (e.g. em0.2).\n"
var IvlanIdfDesc="Delete VLAN by pfSense ID. Required if vlanif was not specified previously. \n" +
	"If vlanif is specified, this value will be overwritten.\n"