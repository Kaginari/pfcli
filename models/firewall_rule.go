package models

type FirewallRule struct {
 Type string  				`json:"type"`
 Interface string  			`json:"interface"`
 IPProtocol string  		`json:"ipprotocol"`
 Protocol string  			`json:"protocol"`
 IMCPType string  			`json:"icmptype"`
 Source string  			`json:"src"`
 Destination string  		`json:"dst"`
 SourcePort int  			`json:"srcport"`
 DestinationPort  int	 	`json:"dstport"`
 Gateway  string 			`json:"gateway"`
 Disabled  bool 			`json:"disabled"`
 Description string  		`json:"descr"`
 Log bool  					`json:"log"`
 Top bool  					`json:"top"`
 Apply bool  				`json:"apply"`

}


var FWRTypeDesc = 		"Set a firewall rule type (pass, block, reject) \n"

var FWRInterfaceDesc = "Set which interface the rule will apply to.\n" +
						"You may specify either the interface's descriptive name, the pfSense ID\n" +
						"(wan, lan, optx), or the physical interface id (e.g. igb0). Floating rules are not supported\n"

var FWRIPProtocolDesc = "Set which IP protocol(s) the rule will apply to (inet, inet6, inet46)\n"

var FWRProtocolDesc = 	"Set which transfer protocol the rule will apply to.\n" +
						"If tcp, udp, tcp/udp, you must define a source and destination port\n"

var FWRIMCPTypeDesc =   "Set the ICMP subtype of the firewall rule. \n" +
						"Multiple values may be passed in as array, single values may be passed as string.\n" +
						"Only available when protocol is set to icmp. If icmptype is not specified all subtypes are assumed \n"

var FWRSourceDesc = 	"Set the source address of the firewall rule. This may be a single IP, network CIDR, alias name, or interface.\n" +
						"When specifying an interface, you may use the physical interface ID, the descriptive interfance name, or the pfSense ID.\n" +
						"To use only interface address, add ip to the end of the interface name otherwise the entire interface's subnet is implied.\n" +
						"To negate the context of the source address, you may prepend the address with !\n"

var FWRDestinationDesc =  "Set the destination address of the firewall rule. This may be a single IP, network CIDR, alias name, or interface.\n" +
						  "When specifying an interface, you may use the physical interface ID, the descriptive interface name, or the pfSense ID.\n" +
						  "To only use interface address, add ip to the end of the interface name otherwise the entire interface's subnet is implied.\n" +
	                      "To negate the context of the source address, you may prepend the address with !\n"

var FWRSourcePortDesc =   "Set the TCP and/or UDP source port or port alias of the firewall rule.\n" +
						  "This is only necessary if you have specified the protocol to tcp, udp, tcp/udp \n"

var FWRDestinationPortDesc = "Set the TCP and/or UDP destination port or port alias of the firewall rule.\n" +
							  "This is only necessary if you have specified the protocol to tcp, udp, tcp/udp\n"

var FWRGatewayDesc = 		"(optional) Set the routing gateway traffic will take upon match  \n"

var FWRDisabledDesc = 		"(optional) Disable the rule upon creation (optional) \n"

var FWRDescriptionDesc =  	"(optional) Set a description for the rule (optional) \n"

var FWRLogDesc = 			"(optional) Enabling rule matche logging\n"

var FWRTopDesc = 			"(optional) Add firewall rule to top of access control list \n"

var FWRApplyDesc = 			"Specify whether or not you would like this rule to be applied immediately,\n" +
							"or simply written to the configuration to be applied later. \n"

