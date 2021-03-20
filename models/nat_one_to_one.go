package models

type NatOneToOne struct {

	Interface string 		`json:"interface"`
	Source string 			`json:"src"`
	Destination string 		`json:"dst"`
	External string 		`json:"external"`
	NatReflection string 	`json:"natreflection"`
	Description string 		`json:"descr"`
	Disabled bool 			`json:"disabled"`
	NoBinat bool 			`json:"nobinat"`
	Top bool 				`json:"top"`
	Apply bool 				`json:"apply"`
}
type DeleteNatOneToOne struct {
	Id string `json:"id"`
	Apply bool `json:"apply"`
}
var NAT121InterfaceDesc  =
 	"Set which interface the mapping will apply to. \n" +
	"You may specify either the interface's descriptive name, \n" +
	"the pfSense ID (wan, lan, optx), or the physical interface id (e.g. igb0).\n" +
	"Floating rules are not supported.\n"

var NAT121SourceDesc =
	"Set the source address of the mapping. This may be a single IP,\n" +
	"network CIDR, alias name, or interface. When specifying an interface,\n" +
	"you may use the physical interface ID, the descriptive interfance name,\n" +
	"or the pfSense ID. To use only interface address, add ip to the end of the\n" +
	"interface name otherwise the entire interface's subnet is implied.\n" +
	"To negate the context of the source address, you may prepend the address with !\n"

var NAT121DestinationDesc =
	"Set the destination address of the mapping.\n" +
	"This may be a single IP,  network CIDR, alias name, or interface.\n" +
	"When specifying an interface, you may use the physical interface ID,\n" +
	"the descriptive interface name, or the pfSense ID. To only use interface address,\n" +
	"add ip to the end of the interface name otherwise the entire interface's subnet is implied.\n" +
	"To negate the context of the source address, you may prepend the address with !\n"

var NAT121ExternalDesc =
	"Specify IPv4 or IPv6 external address to map Inside traffic to.\n" +
	"This Is typically an address on an uplink Interface."

var NAT121NatReflectionDesc =
	"[Optional] Set the NAT reflection mode explicitly. Options are enable or disable\n"


var NAT121DescriptionDesc =
	"[Optional] Set a description for the mapping\n"

var NAT121DisableDesc =
	"[Optional] Disable the mapping upon creation\n"

var NAT121NoBinatDesc =
	"[Optional] Disable binat. This excludes the address from a later, more general, rule.\n"

var NAT121TopDesc =
	"[Optional] Add this mapping to top of access control list.\n"
var NAT121ApplyDesc =
	"[Optional] Specify whether or not you would like this 1:1 mapping to be applied immediately, or simply written to the configuration to be applied later. .\n"

