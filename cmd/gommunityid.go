package main

import (
	"flag"
	"fmt"
	"log"
	"net/netip"
	"os"
	"strconv"

	"github.com/satta/gommunityid"
)

func main() {
	pcapCmd := flag.NewFlagSet("pcap", flag.ExitOnError)
	pcapVersion := pcapCmd.Uint("version", 1, "Community ID version")
	pcapSeed := pcapCmd.Uint("seed", 0, "seed value (default 0)")

	tupleCmd := flag.NewFlagSet("tuple", flag.ExitOnError)
	tupleVersion := tupleCmd.Uint("version", 1, "Community ID version")
	tupleSeed := tupleCmd.Uint("seed", 0, "seed value (default 0)")

	if len(os.Args) < 2 {
		fmt.Println("Usage: gommunityid <pcap|tuple> ...")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "pcap":
		pcapCmd.Parse(os.Args[2:])
		if len(pcapCmd.Args()) == 0 {
			fmt.Println("Usage: gommunityid pcap [options] <pcap-file>")
			pcapCmd.PrintDefaults()
			os.Exit(1)
		}
		ftChan, err := gommunityid.PcapFlowTupleSource(pcapCmd.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
		cid, err := gommunityid.GetCommunityIDByVersion(*pcapVersion, uint16(*pcapSeed))
		if err != nil {
			log.Fatal(err)
		}
		for ft := range ftChan {
			communityid := cid.CalcBase64(ft.FlowTuple)
			fmt.Printf("%d%s | %s | %s %s %d %d %d\n",
				ft.Metadata.Timestamp.Unix(),
				ft.Metadata.Timestamp.Format(".000000"),
				communityid,
				ft.FlowTuple.Srcip, ft.FlowTuple.Dstip,
				ft.FlowTuple.Proto,
				ft.FlowTuple.Srcport, ft.FlowTuple.Dstport)
		}
	case "tuple":
		tupleCmd.Parse(os.Args[2:])
		if len(tupleCmd.Args()) != 5 {
			fmt.Println("Usage: gommunityid tuple [options] <proto> <srcip> <dstip> <srcport> <dstport>")
			tupleCmd.PrintDefaults()
			os.Exit(1)
		}
		cid, err := gommunityid.GetCommunityIDByVersion(*tupleVersion, uint16(*tupleSeed))
		if err != nil {
			log.Fatal(err)
		}
		srcip, err := netip.ParseAddr(tupleCmd.Args()[1])
		if err != nil {
			log.Fatalf("%s is not a valid IP address: %s", tupleCmd.Args()[1], err)
		}
		dstip, err := netip.ParseAddr(tupleCmd.Args()[2])
		if err != nil {
			log.Fatalf("%s is not a valid IP address: %s", tupleCmd.Args()[2], err)
		}
		srcport, err := strconv.ParseUint(tupleCmd.Args()[3], 10, 16)
		if err != nil {
			log.Fatal(err)
		}
		dstport, err := strconv.ParseUint(tupleCmd.Args()[4], 10, 16)
		if err != nil {
			log.Fatal(err)
		}
		proto, err := strconv.ParseUint(tupleCmd.Args()[0], 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		ft := gommunityid.MakeFlowTuple(srcip, dstip, uint16(srcport), uint16(dstport), uint8(proto))
		communityid := cid.CalcBase64(ft)
		fmt.Printf("%s\n", communityid)
	default:
		fmt.Println("expected 'pcap' or 'tuple' subcommands")
		os.Exit(1)
	}
}
