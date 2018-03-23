// +build !linux

package tenus

import (
	"net"
)

// VlanOptions allows you to specify options for vlan link.
type VlanOptions struct {
	// Name of the vlan device
	Dev string
	// VLAN tag id
	Id uint16
	// MAC address
	MacAddr string
}

// Vlaner is interface which embeds Linker interface and adds few more functions.
type Vlaner interface {
	// Linker interface
	Linker
	// MasterNetInterface returns vlan master network interface
	MasterNetInterface() *net.Interface
	// Id returns VLAN tag
	Id() uint16
}

// VlanLink is a Link which has a master network device.
// Each VlanLink has a VLAN tag id
type VlanLink struct {
	Link
	// Master device logical network interface
	masterIfc *net.Interface
	// VLAN tag
	id uint16
}

// NewVlanLink creates vlan network link.
//
// It is equivalent of running:
//		ip link add name vlan${RANDOM STRING} link ${master interface name} type vlan id ${tag}
// NewVlanLink returns Vlaner which is initialized to a pointer of type VlanLink if the
// vlan link was successfully created on the Linux host. Newly created link is assigned
// a random name starting with "vlan". It returns error if the link can not be created.
func NewVlanLink(masterDev string, id uint16) (Vlaner, error) {
	return nil, errNotImplemented
}

// NewVlanLinkWithOptions creates vlan network link and sets some of its network parameters
// to values passed in as VlanOptions
//
// It is equivalent of running:
//		ip link add name ${vlan name} link ${master interface} address ${macaddress} type vlan id ${tag}
// NewVlanLinkWithOptions returns Vlaner which is initialized to a pointer of type VlanLink if the
// vlan link was created successfully on the Linux host. It accepts VlanOptions which allow you to set
// link's options. It returns error if the link could not be created.
func NewVlanLinkWithOptions(masterDev string, opts VlanOptions) (Vlaner, error) {
	return nil, errNotImplemented
}

// NetInterface returns vlan link's network interface
func (vln *VlanLink) NetInterface() *net.Interface {
	return nil
}

// MasterNetInterface returns vlan link's master network interface
func (vln *VlanLink) MasterNetInterface() *net.Interface {
	return nil
}

// Id returns vlan link's vlan tag id
func (vln *VlanLink) Id() uint16 {
	return 0
}

func validateVlanOptions(opts *VlanOptions) error {
	return errNotImplemented
}
