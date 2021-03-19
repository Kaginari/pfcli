package models

type Virtual_IPS struct {
	Mode  	    string  `json:"mode"`
	Interface 	string  `json:"interface"`
	Subnet		string  `json:"subnet"`
	Descr 		string  `json:"descr"`
	Noexpand	bool    `json:"noexpand"`
	Vhid 		string  `json:"vhid"`
	Advbase 	string  `json:"advbase"`
	Advskew 	string  `json:"advskew"`
	Password	string  `json:"password"`
}
type DeleteVirtualIPS struct {
	Id string `json:"id"`
}

var VertualIPS_Mode_Desc =
	"Set our virtual IP mode (ipalias, carp, proxyarp, other)\n"

var VertualIPS_Interface_Desc =
	"Set the interface that will listen for the virtual IP. You may specify either the interface's\n" +
	" descriptive name, the pfSense ID (wan, lan, optx), or the physical interface id (e.g. igb0)\n"
var VertualIPS_Subnet_Desc =
	"Set the IP or subnet (CIDR) for the virtual IP(s)\n"
var VertualIPS_Descr_Desc =
	"Set a description for the new virtual IP (optional)\n"
var VertualIPS_Noexpand_Desc =
	"Disable IP expansion for the virtual IP address. This is only available on proxyarp and \n" +
		"other mode virtual IPs (optional)\n"
var VertualIPS_Vhid_Desc =
	"Set the VHID for the new CARP virtual IP. Only available for carp mode. Leave unspecified to \n" +
		" automatically detect the next available VHID. (optional)\n"
var VertualIPS_Advbase_Desc =
	"Set an advertisement base for the new CARP virtual IP. Only available for carp mode.\n" +
		" Leave unspecified to assume default value of 1. (optional)\n"
var VertualIPS_Advskew_Desc =
	"Set an advertisement skew for the new CARP virtual IP. Only available for carp mode.\n" +
		" Leave unspecified to assume default value of 0. (optional)\n"
var VertualIPS_Password_Desc =
	"Set a password for the new CARP virtual IP. Required for carp mode.\n"
