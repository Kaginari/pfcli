package models

type Interface struct {
	If    			string   `json:"if"`
	Descr 			string   `json:"descr"`
	Enable 			bool	 `json:"enable"`
	Type 			string	 `json:"type"`
	Ipaddr 			string	 `json:"ipaddr"`
	Subnet 			string 	 `json:"subnet"`
	Blockbogons 	bool	 `json:"blockbogons"`
}
type InterfaceDelete struct {
	If string `json:"if"`
}

var InerfaceIfDescr="Specify the physical interface to configure"
var Inerface_descr_Descr="Set a descriptive name for the new interface"
var InerfaceEnableDescr="Enable the interface upon creation. Do not specify this value to leave disabled."
var InerfaceTypeDescr="Set the interface IPv4 configuration type (optional)"
var InerfaceIpaddrDescr="Set the interface's static IPv4 address (required if type is set to staticv4)"
var InerfaceSubnetDescr="Set the interface's static IPv4 address's subnet bitmask (required if type is set to staticv4)"
var InerfaceBlockbogonsDescr=""
var Inerface_ifDelete_Descr="Specify the interface to delete. You may specify either the interface's descriptive name, \n"+
								"the pfSense ID (wan, lan, optx), or the physical interface id (e.g. igb0)\n"