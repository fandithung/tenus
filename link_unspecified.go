// +build !linux

package tenus

import (
	"net"
)

// LinkOptions allows you to specify network link options.
type LinkOptions struct {
	// MAC address
	MacAddr string
	// Maximum Transmission Unit
	MTU int
	// Link network flags i.e. FlagUp, FlagLoopback, FlagMulticast
	Flags net.Flags
	// Network namespace in which the network link should be created
	Ns int
}

// Linker is a generic Linux network link
type Linker interface {
	// NetInterface returns the link's logical network interface
	NetInterface() *net.Interface
	// DeleteLink deletes the link from Linux host
	DeleteLink() error
	// SetLinkMTU sets the link's MTU.
	SetLinkMTU(int) error
	// SetLinkMacAddress sets the link's MAC address.
	SetLinkMacAddress(string) error
	// SetLinkUp brings the link up
	SetLinkUp() error
	// SetLinkDown brings the link down
	SetLinkDown() error
	// SetLinkIp configures the link's IP address
	SetLinkIp(net.IP, *net.IPNet) error
	// UnsetLinkIp remove and IP address from the link
	UnsetLinkIp(net.IP, *net.IPNet) error
	// SetLinkDefaultGw configures the link's default gateway
	SetLinkDefaultGw(*net.IP) error
	// SetLinkNetNsPid moves the link to network namespace specified by PID
	SetLinkNetNsPid(int) error
	// SetLinkNetInNs configures network settings of the link in network namespace
	SetLinkNetInNs(int, net.IP, *net.IPNet, *net.IP) error
}

// Link has a logical network interface
type Link struct {
	ifc *net.Interface
}

// NewLink creates new network link on Linux host.
//
// It is equivalent of running: ip link add name ${ifcName} type dummy
// NewLink returns Linker which is initialized to a pointer of type Link if the
// link was created successfully on the Linux host.
// It returns error if the network link could not be created on Linux host.
func NewLink(ifcName string) (Linker, error) {
	return nil, errNotImplemented
}

// NewLinkFrom creates new tenus link on Linux host from an existing interface of given name
func NewLinkFrom(ifcName string) (Linker, error) {
	return nil, errNotImplemented
}

// NewLinkWithOptions creates new network link on Linux host and sets some of its network
// parameters passed in as LinkOptions
//
// Calling NewLinkWithOptions is equivalent of running following commands one after another if
// particular option is passed in as a parameter:
// 		ip link add name ${ifcName} type dummy
// 		ip link set dev ${ifcName} address ${MAC address}
//		ip link set dev ${ifcName} mtu ${MTU value}
//		ip link set dev ${ifcName} up
// NewLinkWithOptions returns Linker which is initialized to a pointer of type Link if the network
// link with given LinkOptions was created successfully on the Linux host.
// It attempts to delete the link if any of the LinkOptions are incorrect or if setting the options
// failed and returns error.
func NewLinkWithOptions(ifcName string, opts LinkOptions) (Linker, error) {
	return nil, errNotImplemented
}

// DeleteLink deletes netowrk link from Linux Host
// It is equivalent of running: ip link delete dev ${name}
func DeleteLink(name string) error {
	return errNotImplemented
}

// NetInterface returns link's logical network interface.
func (l *Link) NetInterface() *net.Interface {
	return nil
}

// DeleteLink deletes link interface on Linux host.
// It is equivalent of running: ip link delete dev ${interface name}
func (l *Link) DeleteLink() error {
	return errNotImplemented
}

// SetLinkMTU sets link's MTU.
// It is equivalent of running: ip link set dev ${interface name} mtu ${MTU value}
func (l *Link) SetLinkMTU(mtu int) error {
	return errNotImplemented
}

// SetLinkMacAddress sets link's MAC address.
// It is equivalent of running: ip link set dev ${interface name} address ${address}
func (l *Link) SetLinkMacAddress(macaddr string) error {
	return errNotImplemented
}

// SetLinkUp brings the link up.
// It is equivalent of running: ip link set dev ${interface name} up
func (l *Link) SetLinkUp() error {
	return errNotImplemented
}

// SetLinkDown brings the link down.
// It is equivalent of running: ip link set dev ${interface name} down
func (l *Link) SetLinkDown() error {
	return errNotImplemented
}

// SetLinkIp configures the link's IP address.
// It is equivalent of running: ip address add ${address}/${mask} dev ${interface name}
func (l *Link) SetLinkIp(ip net.IP, network *net.IPNet) error {
	return errNotImplemented
}

// UnsetLinkIp configures the link's IP address.
// It is equivalent of running: ip address del ${address}/${mask} dev ${interface name}
func (l *Link) UnsetLinkIp(ip net.IP, network *net.IPNet) error {
	return errNotImplemented
}

// SetLinkDefaultGw configures the link's default Gateway.
// It is equivalent of running: ip route add default via ${ip address}
func (l *Link) SetLinkDefaultGw(gw *net.IP) error {
	return errNotImplemented
}

// SetLinkNetNsPid moves the link to Network namespace specified by PID.
func (l *Link) SetLinkNetNsPid(nspid int) error {
	return errNotImplemented
}

// SetLinkNetInNs configures network settings of the link in network namespace specified by PID.
func (l *Link) SetLinkNetInNs(nspid int, ip net.IP, network *net.IPNet, gw *net.IP) error {
	return errNotImplemented
}

// SetLinkNsFd sets the link's Linux namespace to the one specified by filesystem path.
func (l *Link) SetLinkNsFd(nspath string) error {
	return errNotImplemented
}

// SetLinkNsToDocker sets the link's Linux namespace to a running Docker one specified by Docker name.
func (l *Link) SetLinkNsToDocker(name string, dockerHost string) error {
	return errNotImplemented
}

// RenameInterfaceByName renames an interface of given name.
func RenameInterfaceByName(old string, newName string) error {
	return errNotImplemented
}

// setLinkOptions validates and sets link's various options passed in as LinkOptions.
func setLinkOptions(ifc *net.Interface, opts LinkOptions) error {
	return errNotImplemented
}
