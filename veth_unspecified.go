// +build !linux

package tenus

import (
	"net"
)

// VethOptions allows you to specify options for veth link.
type VethOptions struct {
	// Veth pair's peer interface name
	PeerName string
	// TX queue length
	TxQueueLen int
}

// Vether embeds Linker interface and adds few more functions mostly to handle peer link interface
type Vether interface {
	// Linker interface
	Linker
	// PeerNetInterface returns peer network interface
	PeerNetInterface() *net.Interface
	// SetPeerLinkUp sets peer link up - which also brings up the other peer in VethPair
	SetPeerLinkUp() error
	// DeletePeerLink deletes peer link - this also deletes the other peer in VethPair
	DeletePeerLink() error
	// SetPeerLinkIp configures peer link's IP address
	SetPeerLinkIp(net.IP, *net.IPNet) error
	// SetPeerLinkNsToDocker sends peer link into Docker
	SetPeerLinkNsToDocker(string, string) error
	// SetPeerLinkNsPid sends peer link into container specified by PID
	SetPeerLinkNsPid(int) error
	// SetPeerLinkNsFd sends peer link into container specified by path
	SetPeerLinkNsFd(string) error
	// SetPeerLinkNetInNs configures peer link's IP network in network namespace specified by PID
	SetPeerLinkNetInNs(int, net.IP, *net.IPNet, *net.IP) error
}

// VethPair is a Link. Veth links are created in pairs called peers.
type VethPair struct {
	Link
	// Peer network interface
	peerIfc *net.Interface
}

// NewVethPair creates a pair of veth network links.
//
// It is equivalent of running:
// 		ip link add name veth${RANDOM STRING} type veth peer name veth${RANDOM STRING}.
// NewVethPair returns Vether which is initialized to a pointer of type VethPair if the
// veth link was successfully created on Linux host. Newly created pair of veth links
// are assigned random names starting with "veth".
// NewVethPair returns error if the veth pair could not be created.
func NewVethPair() (Vether, error) {
	return nil, errNotImplemented
}

// NewVethPairWithOptions creates a pair of veth network links.
//
// It is equivalent of running:
// 		ip link add name ${first device name} type veth peer name ${second device name}
// NewVethPairWithOptions returns Vether which is initialized to a pointer of type VethPair if the
// veth link was successfully created on the Linux host. It accepts VethOptions which allow you to set
// peer interface name. It returns error if the veth pair could not be created.
func NewVethPairWithOptions(ifcName string, opts VethOptions) (Vether, error) {
	return nil, errNotImplemented
}

// NetInterface returns veth link's primary network interface
func (veth *VethPair) NetInterface() *net.Interface {
	return nil
}

// NetInterface returns veth link's peer network interface
func (veth *VethPair) PeerNetInterface() *net.Interface {
	return nil
}

// SetPeerLinkUp sets peer link up
func (veth *VethPair) SetPeerLinkUp() error {
	return errNotImplemented
}

// DeletePeerLink deletes peer link. It also deletes the other peer interface in VethPair
func (veth *VethPair) DeletePeerLink() error {
	return errNotImplemented
}

// SetPeerLinkIp configures peer link's IP address
func (veth *VethPair) SetPeerLinkIp(ip net.IP, nw *net.IPNet) error {
	return errNotImplemented
}

// SetPeerLinkNsToDocker sends peer link into Docker
func (veth *VethPair) SetPeerLinkNsToDocker(name string, dockerHost string) error {
	return errNotImplemented
}

// SetPeerLinkNsPid sends peer link into container specified by PID
func (veth *VethPair) SetPeerLinkNsPid(nspid int) error {
	return errNotImplemented
}

// SetPeerLinkNsFd sends peer link into container specified by path
func (veth *VethPair) SetPeerLinkNsFd(nspath string) error {
	return errNotImplemented
}

// SetPeerLinkNetInNs configures peer link's IP network in network namespace specified by PID
func (veth *VethPair) SetPeerLinkNetInNs(nspid int, ip net.IP, network *net.IPNet, gw *net.IP) error {
	return errNotImplemented
}
