package gommunityid

import (
	"bytes"
	"net"
)

type FlowTuple struct {
	Srcip    net.IP
	Dstip    net.IP
	Srcport  uint16
	Dstport  uint16
	Proto    uint8
	IsOneWay bool
}

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

func MakeFlowTupleTCP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoTCP)
}

func MakeFlowTupleUDP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoUDP)
}

func MakeFlowTupleSCTP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoSCTP)
}

func MakeFlowTupleICMP(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoICMP)
}

func MakeFlowTupleICMP6(srcip, dstip net.IP, srcport, dstport uint16) FlowTuple {
	return MakeFlowTuple(srcip, dstip, srcport, dstport, ProtoICMP6)
}

func flowTupleOrdered(addr1, addr2 []byte, port1, port2 uint16) bool {
	return bytes.Compare(addr1, addr2) == -1 || (bytes.Equal(addr1, addr2) && port1 < port2)
}

func (ft FlowTuple) IsOrdered() bool {
	return ft.IsOneWay || flowTupleOrdered(ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport)
}

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
