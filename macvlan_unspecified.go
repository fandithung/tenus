// +build !linux

package tenus

import (
	"net"
)

// Default MacVlan mode
const (
	default_mode = "bridge"
)

// Supported macvlan modes by tenus package
var MacVlanModes = map[string]bool{
	"private": true,
	"vepa":    true,
	"bridge":  true,
}

// MacVlanOptions allows you to specify some options for macvlan link.
type MacVlanOptions struct {
	// macvlan device name
	Dev string
	// macvlan mode
	Mode string
	// MAC address
	MacAddr string
}

// MacVlaner embeds Linker interface and adds few more functions.
type MacVlaner interface {
	// Linker interface
	Linker
	// MasterNetInterface returns macvlan master network device
	MasterNetInterface() *net.Interface
	// Mode returns macvlan link's network mode
	Mode() string
}

// MacVlanLink is Link which has a master network device and operates in
// a given network mode. It implements MacVlaner interface.
type MacVlanLink struct {
	Link
	// Master device logical network interface
	masterIfc *net.Interface
	// macvlan operatio nmode
	mode string
}

// NewMacVlanLink creates macvlan network link
//
// It is equivalent of running:
//		ip link add name mc${RANDOM STRING} link ${master interface} type macvlan
// NewMacVlanLink returns MacVlaner which is initialized to a pointer of type MacVlanLink if the
// macvlan link was created successfully on the Linux host. Newly created link is assigned
// a random name starting with "mc". It sets the macvlan mode to "bridge" mode which is a default.
// It returns error if the link could not be created.
func NewMacVlanLink(masterDev string) (MacVlaner, error) {
	return nil, errNotImplemented
}

// NewMacVlanLinkWithOptions creates macvlan network link and sets som of its network parameters
// passed in as MacVlanOptions.
//
// It is equivalent of running:
// 		ip link add name ${macvlan name} link ${master interface} address ${macaddress} type macvlan mode ${mode}
// NewMacVlanLinkWithOptions returns MacVlaner which is initialized to a pointer of type MacVlanLink if the
// macvlan link was created successfully on the Linux host. If particular option is empty, it sets default value if possible.
// It returns error if the macvlan link could not be created or if incorrect options have been passed.
func NewMacVlanLinkWithOptions(masterDev string, opts MacVlanOptions) (MacVlaner, error) {
	return nil, errNotImplemented
}

// NetInterface returns macvlan link's network interface
func (macvln *MacVlanLink) NetInterface() *net.Interface {
	return nil
}

// MasterNetInterface returns macvlan link's master network interface
func (macvln *MacVlanLink) MasterNetInterface() *net.Interface {
	return nil
}

// Mode returns macvlan link's network operation mode
func (macvln *MacVlanLink) Mode() string {
	return "not implemented"
}

func validateMacVlanOptions(opts *MacVlanOptions) error {
	return errNotImplemented
}
