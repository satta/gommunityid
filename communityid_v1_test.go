package gommunityid

import (
	"fmt"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSet struct {
	Srcip       netip.Addr
	Dstip       netip.Addr
	Srcport     uint16
	Dstport     uint16
	Base64Seed0 string
	HexSeed0    string
	Base64Seed1 string
}

type makeFunc func(netip.Addr, netip.Addr, uint16, uint16) FlowTuple

func verifyTuples(t *testing.T, ts []testSet, mf makeFunc) {
	cid0 := CommunityIDv1{
		Seed: 0,
	}
	cid1 := CommunityIDv1{
		Seed: 1,
	}
	for _, tset := range ts {
		ft := mf(tset.Srcip, tset.Dstip, tset.Srcport, tset.Dstport)
		assert.Equal(t, tset.Base64Seed0, cid0.CalcBase64(ft))
		assert.Equal(t, tset.HexSeed0, cid0.CalcHex(ft))
		assert.Equal(t, tset.Base64Seed1, cid1.CalcBase64(ft))
		assert.Equal(t, tset.HexSeed0, fmt.Sprintf("1:%x", cid0.Calc(ft)))
	}
}

func TestCommunityIDv1ICMP(t *testing.T) {
	verifyTuples(t, []testSet{
		{
			Srcip:       netip.MustParseAddr("192.168.0.89"),
			Dstip:       netip.MustParseAddr("192.168.0.1"),
			Srcport:     8,
			Dstport:     0,
			Base64Seed0: "1:X0snYXpgwiv9TZtqg64sgzUn6Dk=",
			HexSeed0:    "1:5f4b27617a60c22bfd4d9b6a83ae2c833527e839",
			Base64Seed1: "1:03g6IloqVBdcZlPyX8r0hgoE7kA=",
		},
		{
			Srcip:       netip.MustParseAddr("192.168.0.1"),
			Dstip:       netip.MustParseAddr("192.168.0.89"),
			Srcport:     0,
			Dstport:     8,
			Base64Seed0: "1:X0snYXpgwiv9TZtqg64sgzUn6Dk=",
			HexSeed0:    "1:5f4b27617a60c22bfd4d9b6a83ae2c833527e839",
			Base64Seed1: "1:03g6IloqVBdcZlPyX8r0hgoE7kA=",
		},
		{
			Srcip:       netip.MustParseAddr("192.168.0.89"),
			Dstip:       netip.MustParseAddr("192.168.0.1"),
			Srcport:     20,
			Dstport:     0,
			Base64Seed0: "1:3o2RFccXzUgjl7zDpqmY7yJi8rI=",
			HexSeed0:    "1:de8d9115c717cd482397bcc3a6a998ef2262f2b2",
			Base64Seed1: "1:lCXHHxavE1Vq3oX9NH5ladQg02o=",
		},
		{
			Srcip:       netip.MustParseAddr("192.168.0.89"),
			Dstip:       netip.MustParseAddr("192.168.0.1"),
			Srcport:     20,
			Dstport:     1,
			Base64Seed0: "1:tz/fHIDUHs19NkixVVoOZywde+I=",
			HexSeed0:    "1:b73fdf1c80d41ecd7d3648b1555a0e672c1d7be2",
			Base64Seed1: "1:Ie3wmFyxiEyikbsbcO03d2nh+PM=",
		},
		{
			Srcip:       netip.MustParseAddr("192.168.0.1"),
			Dstip:       netip.MustParseAddr("192.168.0.89"),
			Srcport:     0,
			Dstport:     20,
			Base64Seed0: "1:X0snYXpgwiv9TZtqg64sgzUn6Dk=",
			HexSeed0:    "1:5f4b27617a60c22bfd4d9b6a83ae2c833527e839",
			Base64Seed1: "1:03g6IloqVBdcZlPyX8r0hgoE7kA=",
		},
	},
		MakeFlowTupleICMP)
}

func TestCommunityIDv1ICMP6(t *testing.T) {
	verifyTuples(t, []testSet{
		{
			Srcip:       netip.MustParseAddr("fe80::200:86ff:fe05:80da"),
			Dstip:       netip.MustParseAddr("fe80::260:97ff:fe07:69ea"),
			Srcport:     135,
			Dstport:     0,
			Base64Seed0: "1:dGHyGvjMfljg6Bppwm3bg0LO8TY=",
			HexSeed0:    "1:7461f21af8cc7e58e0e81a69c26ddb8342cef136",
			Base64Seed1: "1:kHa1FhMYIT6Ym2Vm2AOtoOARDzY=",
		},
		{
			Srcip:       netip.MustParseAddr("fe80::260:97ff:fe07:69ea"),
			Dstip:       netip.MustParseAddr("fe80::200:86ff:fe05:80da"),
			Srcport:     136,
			Dstport:     0,
			Base64Seed0: "1:dGHyGvjMfljg6Bppwm3bg0LO8TY=",
			HexSeed0:    "1:7461f21af8cc7e58e0e81a69c26ddb8342cef136",
			Base64Seed1: "1:kHa1FhMYIT6Ym2Vm2AOtoOARDzY=",
		},
		{
			Srcip:       netip.MustParseAddr("3ffe:507:0:1:260:97ff:fe07:69ea"),
			Dstip:       netip.MustParseAddr("3ffe:507:0:1:200:86ff:fe05:80da"),
			Srcport:     3,
			Dstport:     0,
			Base64Seed0: "1:NdobDX8PQNJbAyfkWxhtL2Pqp5w=",
			HexSeed0:    "1:35da1b0d7f0f40d25b0327e45b186d2f63eaa79c",
			Base64Seed1: "1:OlOWx9psIbBFi7lOCw/4MhlKR9M=",
		},
		{
			Srcip:       netip.MustParseAddr("3ffe:507:0:1:200:86ff:fe05:80da"),
			Dstip:       netip.MustParseAddr("3ffe:507:0:1:260:97ff:fe07:69ea"),
			Srcport:     3,
			Dstport:     0,
			Base64Seed0: "1:/OGBt9BN1ofenrmSPWYicpij2Vc=",
			HexSeed0:    "1:fce181b7d04dd687de9eb9923d66227298a3d957",
			Base64Seed1: "1:Ij4ZxnC87/MXzhOjvH2vHu7LRmE=",
		},
	},
		MakeFlowTupleICMP6)
}

func TestCommunityIDv1SCTP(t *testing.T) {
	verifyTuples(t, []testSet{
		{
			Srcip:       netip.MustParseAddr("192.168.170.8"),
			Dstip:       netip.MustParseAddr("192.168.170.56"),
			Srcport:     7,
			Dstport:     80,
			Base64Seed0: "1:jQgCxbku+pNGw8WPbEc/TS/uTpQ=",
			HexSeed0:    "1:8d0802c5b92efa9346c3c58f6c473f4d2fee4e94",
			Base64Seed1: "1:Y1/0jQg6e+I3ZwZZ9LP65DNbTXU=",
		},
		{
			Srcip:       netip.MustParseAddr("192.168.170.56"),
			Dstip:       netip.MustParseAddr("192.168.170.8"),
			Srcport:     80,
			Dstport:     7,
			Base64Seed0: "1:jQgCxbku+pNGw8WPbEc/TS/uTpQ=",
			HexSeed0:    "1:8d0802c5b92efa9346c3c58f6c473f4d2fee4e94",
			Base64Seed1: "1:Y1/0jQg6e+I3ZwZZ9LP65DNbTXU=",
		},
	},
		MakeFlowTupleSCTP)
}

func TestCommunityIDv1TCP(t *testing.T) {
	verifyTuples(t, []testSet{
		{
			Srcip:       netip.MustParseAddr("128.232.110.120"),
			Dstip:       netip.MustParseAddr("66.35.250.204"),
			Srcport:     34855,
			Dstport:     80,
			Base64Seed0: "1:LQU9qZlK+B5F3KDmev6m5PMibrg=",
			HexSeed0:    "1:2d053da9994af81e45dca0e67afea6e4f3226eb8",
			Base64Seed1: "1:3V71V58M3Ksw/yuFALMcW0LAHvc=",
		},
		{
			Srcip:       netip.MustParseAddr("66.35.250.204"),
			Dstip:       netip.MustParseAddr("128.232.110.120"),
			Srcport:     80,
			Dstport:     34855,
			Base64Seed0: "1:LQU9qZlK+B5F3KDmev6m5PMibrg=",
			HexSeed0:    "1:2d053da9994af81e45dca0e67afea6e4f3226eb8",
			Base64Seed1: "1:3V71V58M3Ksw/yuFALMcW0LAHvc=",
		},
	},
		MakeFlowTupleTCP)
}

func TestCommunityIDv1UDP(t *testing.T) {
	verifyTuples(t, []testSet{
		{
			Srcip:       netip.MustParseAddr("192.168.1.52"),
			Dstip:       netip.MustParseAddr("8.8.8.8"),
			Srcport:     54585,
			Dstport:     53,
			Base64Seed0: "1:d/FP5EW3wiY1vCndhwleRRKHowQ=",
			HexSeed0:    "1:77f14fe445b7c22635bc29dd87095e451287a304",
			Base64Seed1: "1:Q9We8WO3piVF8yEQBNJF4uiSVrI=",
		},
		{
			Srcip:       netip.MustParseAddr("8.8.8.8"),
			Dstip:       netip.MustParseAddr("192.168.1.52"),
			Srcport:     53,
			Dstport:     54585,
			Base64Seed0: "1:d/FP5EW3wiY1vCndhwleRRKHowQ=",
			HexSeed0:    "1:77f14fe445b7c22635bc29dd87095e451287a304",
			Base64Seed1: "1:Q9We8WO3piVF8yEQBNJF4uiSVrI=",
		},
	},
		MakeFlowTupleUDP)
}
