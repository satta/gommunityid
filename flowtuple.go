package gommunityid

import (
	"bytes"
	"net"
)

// FlowTuple is a collection of all values required for ID calculation.
type FlowTuple struct {
	Srcip    net.IP
	Dstip    net.IP
	Srcport  uint16
	Dstport  uint16
	Proto    uint8
	IsOneWay bool
}

// MakeFlowTuple returns a FlowTuple for the given set of communication
// details: protocol, IPs (source, destination) and ports (source,
// destination).
func MakeFlowTuple(srcip, dstip net.IP, srcport, dstport uint16, proto uint8) FlowTuple {
	var isOneWay bool
	if proto == ProtoICMP {
		srcport, dstport, isOneWay = GetICMPv4PortEquivalents(uint8(srcport), uint8(dstport))
	} else if proto == ProtoICMP6 {
		srcport, dstport, isOneWay = GetICMPv6PortEquivalents(uint8(srcport), uint8(dstport))
	}
	v := FlowTuple{
		Srcip:    srcip,
		Dstip:    dstip,
		Srcport:  srcport,
		Dstport:  dstport,
		Proto:    proto,
		IsOneWay: isOneWay,
	}
	return v
}

// MakeFlowTupleTCP returns a FlowTuple with the TCP protocol preconfigured.
func MakeFlowTupleTCP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoTCP)
}

// MakeFlowTupleUDP returns a FlowTuple with the UDP protocol preconfigured.
func MakeFlowTupleUDP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoUDP)
}

// MakeFlowTupleSCTP returns a FlowTuple with the SCTP protocol preconfigured.
func MakeFlowTupleSCTP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoSCTP)
}

// MakeFlowTupleICMP returns a FlowTuple with the ICMPv4 protocol preconfigured.
func MakeFlowTupleICMP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoICMP)
}

// MakeFlowTupleICMP6 returns a FlowTuple with the ICMPv6 protocol preconfigured.
func MakeFlowTupleICMP6(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoICMP6)
}

func flowTupleOrdered(addr1, addr2 []byte, port1, port2 uint16) bool {
	return bytes.Compare(addr1, addr2) == -1 || (bytes.Equal(addr1, addr2) && port1 < port2)
}

// IsOrdered returns true if the flow tuple direction is ordered.
func (ft FlowTuple) IsOrdered() bool {
	return ft.IsOneWay || flowTupleOrdered(ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport)
}

// InOrder returns a new copy of the flow tuple, with guaranteed IsOrdered()
// property.
func (ft FlowTuple) InOrder() FlowTuple {
	if ft.IsOrdered() {
		return FlowTuple{
			Srcip:    ft.Srcip,
			Dstip:    ft.Dstip,
			Srcport:  ft.Srcport,
			Dstport:  ft.Dstport,
			Proto:    ft.Proto,
			IsOneWay: ft.IsOneWay,
		}
	}
	return FlowTuple{
		Srcip:    ft.Dstip,
		Dstip:    ft.Srcip,
		Srcport:  ft.Dstport,
		Dstport:  ft.Srcport,
		Proto:    ft.Proto,
		IsOneWay: ft.IsOneWay,
	}
}
