// +build !linux

package tenus

// MacVtaper embeds MacVlaner interface
type MacVtaper interface {
	MacVlaner
}

// MacVtapLink is MacVlanLink. It implements MacVtaper interface
type MacVtapLink struct {
	*MacVlanLink
}

// NewMacVtapLink creates macvtap network link
//
// It is equivalent of running:
//		ip link add name mvt${RANDOM STRING} link ${master interface} type macvtap
// NewMacVtapLink returns MacVtaper which is initialized to a pointer of type MacVtapLink if the
// macvtap link was created successfully on the Linux host. Newly created link is assigned
// a random name starting with "mvt". It sets the macvlan mode to "bridge" which is a default.
// It returns error if the link could not be created.
func NewMacVtapLink(masterDev string) (MacVtaper, error) {
	return nil, errNotImplemented
}

// NewMacVtapLinkWithOptions creates macvtap network link and can set some of its network parameters
// passed in as MacVlanOptions.
//
// It is equivalent of running:
// 		ip link add name ${macvlan name} link ${master interface} address ${macaddress} type macvtap mode ${mode}
// NewMacVtapLinkWithOptions returns MacVtaper which is initialized to a pointer of type MacVtapLink if the
// macvtap link was created successfully on the Linux host. It returns error if the macvtap link could not be created.
func NewMacVtapLinkWithOptions(masterDev string, opts MacVlanOptions) (MacVtaper, error) {
	return nil, errNotImplemented
}
