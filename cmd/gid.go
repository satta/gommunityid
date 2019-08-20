package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"github.com/satta/gommunityid"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func tupleLessThan(addr1, addr2 []byte, port1, port2 uint16) bool {
	return bytes.Compare(addr1, addr2) == -1 || (bytes.Equal(addr1, addr2) && port1 < port2)
}

func main() {
	cid := gommunityid.CommunityIDv1{
		Seed: 0,
	}
	if handle, err := pcap.OpenOffline(os.Args[1]); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			var netLayer gopacket.NetworkLayer
			if netLayer = packet.NetworkLayer(); netLayer == nil {
				continue
			}
			var proto uint8
			switch netLayer.LayerType() {
			case layers.LayerTypeIPv4:
				proto = uint8(netLayer.(*layers.IPv4).Protocol)
			case layers.LayerTypeIPv6:
				proto = uint8(netLayer.(*layers.IPv6).NextLayerType())
			case layers.LayerTypeICMPv4:
				proto = 1
			case layers.LayerTypeICMPv6:
				proto = 58
			}
			src, dst := netLayer.NetworkFlow().Endpoints()

			ft := gommunityid.FlowTuple{
				Proto: proto,
				Srcip: net.IP(src.Raw()),
				Dstip: net.IP(dst.Raw()),
			}

			var transLayer gopacket.TransportLayer
			var srcP, dstP uint16
			if transLayer = packet.TransportLayer(); transLayer == nil {
				if got, ok := packet.Layer(layers.LayerTypeICMPv6).(*layers.ICMPv6); ok {
					var oneway bool
					srcP, dstP, oneway = gommunityid.GetICMPv6PortEquivalents(got.TypeCode.Type(), got.TypeCode.Code())
					ft.Srcport, ft.Dstport = srcP, dstP
					if !oneway && !tupleLessThan(ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport) {
						ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport = ft.Dstip, ft.Srcip, ft.Dstport, ft.Srcport
					}

				}
				if got, ok := packet.Layer(layers.LayerTypeICMPv4).(*layers.ICMPv4); ok {
					var oneway bool
					srcP, dstP, oneway = gommunityid.GetICMPv4PortEquivalents(got.TypeCode.Type(), got.TypeCode.Code())
					ft.Srcport, ft.Dstport = srcP, dstP
					if !oneway && !tupleLessThan(ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport) {
						ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport = ft.Dstip, ft.Srcip, ft.Dstport, ft.Srcport
					}
				}
			} else {
				_srcP, _dstP := transLayer.TransportFlow().Endpoints()
				srcP, dstP = binary.BigEndian.Uint16(_srcP.Raw()), binary.BigEndian.Uint16(_dstP.Raw())
				ft.Srcport, ft.Dstport = srcP, dstP
				if !tupleLessThan(ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport) {
					ft.Srcip, ft.Dstip, ft.Srcport, ft.Dstport = ft.Dstip, ft.Srcip, ft.Dstport, ft.Srcport
				}
			}

			communityid := cid.CalcBase64(ft)
			fmt.Printf("%d%s | %s | %s %s %d %d %d\n",
				packet.Metadata().Timestamp.Unix(),
				packet.Metadata().Timestamp.Format(".000000"),
				communityid,
				src, dst,
				proto,
				srcP, dstP)
		}
	}
}
