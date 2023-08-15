package gommunityid

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash"
)

// CommunityIDv1 encapsulates the calculation code for version 1 of the
// Community ID flow hashing algorithm.
type CommunityIDv1 struct {
	Seed uint16
}

// Calc returns the community id value for a given FlowTuple, as an
// unformatted byte slice.
func (cid CommunityIDv1) Calc(ft FlowTuple) []byte {
	ft = ft.InOrder()
	return cid.Render(cid.Hash(ft))
}

// CalcHex returns the community id value for a given FlowTuple, as an
// hex-encoded string.
func (cid CommunityIDv1) CalcHex(ft FlowTuple) string {
	ft = ft.InOrder()
	return cid.RenderHex(cid.Hash(ft))
}

// CalcBase64 returns the community id value for a given FlowTuple, as an
// Base64-encoded string.
func (cid CommunityIDv1) CalcBase64(ft FlowTuple) string {
	ft = ft.InOrder()
	return cid.RenderBase64(cid.Hash(ft))
}

// Hash returns a hash.Hash instance (SHA1) in a state corresponding to all
// input value already dealt with in the hash.
func (cid CommunityIDv1) Hash(ft FlowTuple) hash.Hash {
	h := sha1.New()
	binary.Write(h, binary.BigEndian, cid.Seed)
	if ft.Srcip.Is4() {
		binary.Write(h, binary.BigEndian, ft.Srcip.As4())
	} else if ft.Srcip.Is6() {
		binary.Write(h, binary.BigEndian, ft.Srcip.As16())
	}
	if ft.Dstip.Is4() {
		binary.Write(h, binary.BigEndian, ft.Dstip.As4())
	} else if ft.Dstip.Is6() {
		binary.Write(h, binary.BigEndian, ft.Dstip.As16())
	}
	h.Write([]byte{ft.Proto, 0})
	binary.Write(h, binary.BigEndian, ft.Srcport)
	binary.Write(h, binary.BigEndian, ft.Dstport)
	return h
}

// Render returns the value of the given hash, as an unformatted byte slice.
func (cid CommunityIDv1) Render(h hash.Hash) []byte {
	return h.Sum(nil)
}

// RenderBase64 returns the value of the given hash, as Base64-encoded string.
func (cid CommunityIDv1) RenderBase64(h hash.Hash) string {
	return "1:" + base64.StdEncoding.EncodeToString(cid.Render(h))
}

// RenderHex returns the value of the given hash, as hex-encoded string.
func (cid CommunityIDv1) RenderHex(h hash.Hash) string {
	return fmt.Sprintf("1:%x", cid.Render(h))
}
