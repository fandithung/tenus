// +build !linux

package tenus

import (
	"fmt"
	"net"

	"github.com/docker/libcontainer/netlink"
)

// Bridger embeds Linker interface and adds one extra function.
type Bridger interface {
	// Linker interface
	Linker
	// AddSlaveIfc adds network interface to the network bridge
	AddSlaveIfc(*net.Interface) error
	//RemoveSlaveIfc removes network interface from the network bridge
	RemoveSlaveIfc(*net.Interface) error
}

// Bridge is Link which has zero or more slave network interfaces.
// Bridge implements Bridger interface.
type Bridge struct {
	Link
	slaveIfcs []net.Interface
}

// NewBridge creates new network bridge on Linux host.
//
// It is equivalent of running: ip link add name br${RANDOM STRING} type bridge
// NewBridge returns Bridger which is initialized to a pointer of type Bridge if the
// bridge was created successfully on the Linux host. Newly created bridge is assigned
// a random name starting with "br".
// It returns error if the bridge could not be created.
func NewBridge() (Bridger, error) {
	return nil, errNotImplemented
}

// NewBridgeWithName creates new network bridge on Linux host with the name passed as a parameter.
// It is equivalent of running: ip link add name ${ifcName} type bridge
// It returns error if the bridge can not be created.
func NewBridgeWithName(ifcName string) (Bridger, error) {
	return nil, errNotImplemented
}

// BridgeFromName returns a tenus network bridge from an existing bridge of given name on the Linux host.
// It returns error if the bridge of the given name cannot be found.
func BridgeFromName(ifcName string) (Bridger, error) {
	if ok, err := NetInterfaceNameValid(ifcName); !ok {
		return nil, err
	}

	newIfc, err := net.InterfaceByName(ifcName)
	if err != nil {
		return nil, fmt.Errorf("Could not find the new interface: %s", err)
	}

	return &Bridge{
		Link: Link{
			ifc: newIfc,
		},
	}, nil
}

// AddToBridge adds network interfaces to network bridge.
// It is equivalent of running: ip link set ${netIfc name} master ${netBridge name}
// It returns error when it fails to add the network interface to bridge.
func AddToBridge(netIfc, netBridge *net.Interface) error {
	return netlink.NetworkSetMaster(netIfc, netBridge)
}

// RemoveFromBridge not yet implemented
func RemoveFromBridge(netIfc *net.Interface) error {
	return errNotImplemented
}

// AddSlaveIfc adds network interface to network bridge.
// It is equivalent of running: ip link set ${ifc name} master ${bridge name}
// It returns error if the network interface could not be added to the bridge.
func (br *Bridge) AddSlaveIfc(ifc *net.Interface) error {
	return errNotImplemented
}

// RemoveSlaveIfc removes network interface from the network bridge.
// It is equivalent of running: ip link set dev ${netIfc name} nomaster
// It returns error if the network interface is not in the bridge or
// it could not be removed from the bridge.
func (br *Bridge) RemoveSlaveIfc(ifc *net.Interface) error {
	return errNotImplemented
}
