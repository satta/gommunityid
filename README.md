# gommunityid

[![Status](https://github.com/satta/gommunityid/actions/workflows/go.yml/badge.svg)](https://github.com/satta/gommunityid/actions)
[![Coverage Status](https://coveralls.io/repos/github/satta/gommunityid/badge.svg?branch=master)](https://coveralls.io/github/satta/gommunityid?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/satta/gommunityid)](https://goreportcard.com/report/github.com/satta/gommunityid)
[![Documentation](https://godoc.org/github.com/satta/gommunityid?status.svg)](http://godoc.org/github.com/satta/gommunityid)

gommunityid is a Golang implementation of the [Community ID flow hashing algorithm](https://github.com/corelight/community-id-spec). Its API design was clearly and obviously inspired by the [Python reference implementation](https://github.com/corelight/pycommunityid).

## Usage

```Go
package main

import (
	"fmt"
	"net"

	"github.com/satta/gommunityid"
)

func main() {
	// Get instance for version 1, seed 0
	cid, _ := gommunityid.GetCommunityIDByVersion(1, 0)

	// Obtain flow tuple. This can be done any way you like.
	ft := gommunityid.MakeFlowTuple(net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8), 9, 10, 1)

	// Calculate Base64-encoded value
	communityid := cid.CalcBase64(ft)
	fmt.Printf("%s\n", communityid)

	// Calculate hex-encoded value
	communityid = cid.CalcHex(ft)
	fmt.Printf("%s\n", communityid)

	// Calculate byte slice
	communityidByte := cid.Calc(ft)
	fmt.Printf("%v\n", communityidByte)
}
```
There is also a [convenience function](https://godoc.org/github.com/satta/gommunityid#PcapFlowTupleSource) for parsing pcap files and automated FlowTuple generation for all supported protocols.

## Command line interface

This package builds a simple [command line tool](cmd/gommunityid.go) to calculate IDs for pcaps:
```
$ ./gommunityid pcap
Usage: gommunityid pcap [options] <pcap-file>
  -seed uint
    	seed value (default 0)
  -version uint
    	Community ID version (default 1)
$ gommunityid pcap testdata/tcp.pcap
1071580904.891921 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 128.232.110.120 66.35.250.204 6 34855 80
1071580905.035577 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 66.35.250.204 128.232.110.120 6 80 34855
1071580905.035724 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 128.232.110.120 66.35.250.204 6 34855 80
1071580905.037333 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 128.232.110.120 66.35.250.204 6 34855 80
1071580905.181581 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 66.35.250.204 128.232.110.120 6 80 34855
1071580905.184528 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 66.35.250.204 128.232.110.120 6 80 34855
1071580905.184844 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 128.232.110.120 66.35.250.204 6 34855 80
1071580905.184698 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 66.35.250.204 128.232.110.120 6 80 34855
1071580905.184920 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 128.232.110.120 66.35.250.204 6 34855 80
1071580905.184736 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 66.35.250.204 128.232.110.120 6 80 34855
1071580905.203025 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 128.232.110.120 66.35.250.204 6 34855 80
1071580905.346457 | 1:LQU9qZlK+B5F3KDmev6m5PMibrg= | 66.35.250.204 128.232.110.120 6 80 34855
```
and explicit tuples:
```
$ gommunityid tuple
Usage: gommunityid tuple [options] <proto> <srcip> <dstip> <srcport> <dstport>
  -seed uint
    	seed value (default 0)
  -version uint
    	Community ID version (default 1)
$ gommunityid tuple 6 66.35.250.204 128.232.110.120 80 34855
1:LQU9qZlK+B5F3KDmev6m5PMibrg=
```

## Author/Contact

Sascha Steinbiss

## License

MIT
