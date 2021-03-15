package gommunityid

import (
	"net"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

// PcapFlowTuple represents a pair of the FlowTuple for a packet as
// well as its packet metadata (e.g. timestamp).
type PcapFlowTuple struct {
	FlowTuple FlowTuple
	Metadata  *gopacket.PacketMetadata
}

// PcapFlowTupleSource returns, for a given pcap file name, a channel
// delivering PcapFlowTuples for each packet in the file. If the file
// cannot be read for some reason, an error is returned as well
// accordingly.
func PcapFlowTupleSource(file string) (<-chan PcapFlowTuple, error) {
	outChan := make(chan PcapFlowTuple)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	handle, err := pcapgo.NewReader(f)
	if err != nil {
		return nil, err
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	go func() {
		for packet := range packetSource.Packets() {
			var ft FlowTuple

			var src, dst net.IP
			ip4Layer := packet.Layer(layers.LayerTypeIPv4)
			if ip4Layer != nil {
				ip, _ := ip4Layer.(*layers.IPv4)
				src, dst = ip.SrcIP, ip.DstIP
			} else {
				ip6Layer := packet.Layer(layers.LayerTypeIPv6)
				if ip6Layer != nil {
					ip, _ := ip6Layer.(*layers.IPv6)
					src, dst = ip.SrcIP, ip.DstIP
				} else {
					// no IP layer found
					continue
				}
			}

			var srcP, dstP uint16
			var proto uint8

			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)
				srcP, dstP = uint16(tcp.SrcPort), uint16(tcp.DstPort)
				proto = ProtoTCP
			}
			udpLayer := packet.Layer(layers.LayerTypeUDP)
			if udpLayer != nil {
				udp, _ := udpLayer.(*layers.UDP)
				srcP, dstP = uint16(udp.SrcPort), uint16(udp.DstPort)
				proto = ProtoUDP
			}
			icmp4Layer := packet.Layer(layers.LayerTypeICMPv4)
			if icmp4Layer != nil {
				icmp4, _ := icmp4Layer.(*layers.ICMPv4)
				srcP, dstP = uint16(icmp4.TypeCode.Type()), uint16(icmp4.TypeCode.Code())
				proto = ProtoICMP
			}
			icmp6Layer := packet.Layer(layers.LayerTypeICMPv6)
			if icmp6Layer != nil {
				icmp6, _ := icmp6Layer.(*layers.ICMPv6)
				srcP, dstP = uint16(icmp6.TypeCode.Type()), uint16(icmp6.TypeCode.Code())
				proto = ProtoICMP6
			}
			sctpLayer := packet.Layer(layers.LayerTypeSCTP)
			if sctpLayer != nil {
				sctp, _ := sctpLayer.(*layers.SCTP)
				srcP, dstP = uint16(sctp.SrcPort), uint16(sctp.DstPort)
				proto = ProtoSCTP
			}
			if proto == 0 {
				continue
			}

			ft = MakeFlowTuple(src, dst, srcP, dstP, proto)
			outChan <- PcapFlowTuple{
				FlowTuple: ft,
				Metadata:  packet.Metadata(),
			}
		}
		close(outChan)
	}()
	return outChan, nil
}
