package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
		fmt.Println("expected 'pcap' or 'tuple' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "pcap":
		pcapCmd.Parse(os.Args[2:])
		if len(pcapCmd.Args()) == 0 {
			log.Println("No input file given")
			flag.PrintDefaults()
			os.Exit(0)
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
		fmt.Printf("not yet implemented")
		_, _ = tupleSeed, tupleVersion
	default:
		fmt.Println("expected 'pcap' or 'tuple' subcommands")
		os.Exit(1)
	}
}
