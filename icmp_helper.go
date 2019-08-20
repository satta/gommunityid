package gommunityid

import (
	"github.com/google/gopacket/layers"
)

var ICMPv4PortEquivalents = map[uint8]uint8{
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

var ICMPv6PortEquivalents = map[uint8]uint8{
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

func GetICMPv4PortEquivalents(p1, p2 uint8) (uint16, uint16, bool) {
	if val, ok := ICMPv4PortEquivalents[p1]; ok {
		return uint16(p1), uint16(val), false
	}
	return uint16(p1), uint16(p2), true
}

func GetICMPv6PortEquivalents(p1, p2 uint8) (uint16, uint16, bool) {
	if val, ok := ICMPv6PortEquivalents[p1]; ok {
		return uint16(p1), uint16(val), false
	}
	return uint16(p1), uint16(p2), true
}
