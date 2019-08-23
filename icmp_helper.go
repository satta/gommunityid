package gommunityid

import (
	"github.com/google/gopacket/layers"
)

var icmpv4PortEquivalents = map[uint8]uint8{
	layers.ICMPv4TypeEchoRequest:         layers.ICMPv4TypeEchoReply,
	layers.ICMPv4TypeEchoReply:           layers.ICMPv4TypeEchoRequest,
	layers.ICMPv4TypeTimestampRequest:    layers.ICMPv4TypeTimestampReply,
	layers.ICMPv4TypeTimestampReply:      layers.ICMPv4TypeTimestampRequest,
	layers.ICMPv4TypeInfoRequest:         layers.ICMPv4TypeInfoReply,
	layers.ICMPv4TypeInfoReply:           layers.ICMPv4TypeInfoRequest,
	layers.ICMPv4TypeRouterSolicitation:  layers.ICMPv4TypeRouterAdvertisement,
	layers.ICMPv4TypeRouterAdvertisement: layers.ICMPv4TypeRouterSolicitation,
	layers.ICMPv4TypeAddressMaskRequest:  layers.ICMPv4TypeAddressMaskReply,
	layers.ICMPv4TypeAddressMaskReply:    layers.ICMPv4TypeAddressMaskRequest,
}

var icmpv6PortEquivalents = map[uint8]uint8{
	layers.ICMPv6TypeEchoRequest:                         layers.ICMPv6TypeEchoReply,
	layers.ICMPv6TypeEchoReply:                           layers.ICMPv6TypeEchoRequest,
	layers.ICMPv6TypeRouterSolicitation:                  layers.ICMPv6TypeRouterAdvertisement,
	layers.ICMPv6TypeRouterAdvertisement:                 layers.ICMPv6TypeRouterSolicitation,
	layers.ICMPv6TypeNeighborSolicitation:                layers.ICMPv6TypeNeighborAdvertisement,
	layers.ICMPv6TypeNeighborAdvertisement:               layers.ICMPv6TypeNeighborSolicitation,
	layers.ICMPv6TypeMLDv1MulticastListenerQueryMessage:  layers.ICMPv6TypeMLDv1MulticastListenerReportMessage,
	layers.ICMPv6TypeMLDv1MulticastListenerReportMessage: layers.ICMPv6TypeMLDv1MulticastListenerQueryMessage,
	144: 145,
	145: 144,
}

// GetICMPv4PortEquivalents returns ICMPv4 codes mapped back to pseudo port
// numbers, as well as a bool indicating whether a communication is one-way.
func GetICMPv4PortEquivalents(p1, p2 uint8) (uint16, uint16, bool) {
	if val, ok := icmpv4PortEquivalents[p1]; ok {
		return uint16(p1), uint16(val), false
	}
	return uint16(p1), uint16(p2), true
}

// GetICMPv6PortEquivalents returns ICMPv6 codes mapped back to pseudo port
// numbers, as well as a bool indicating whether a communication is one-way.
func GetICMPv6PortEquivalents(p1, p2 uint8) (uint16, uint16, bool) {
	if val, ok := icmpv6PortEquivalents[p1]; ok {
		return uint16(p1), uint16(val), false
	}
	return uint16(p1), uint16(p2), true
}
