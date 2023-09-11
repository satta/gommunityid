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
	const (
		maxParameterSizeIPv6 = 40
		maxParameterSizeIPv4 = 16
	)
	buffer := make([]byte, 2, maxParameterSizeIPv4)

	binary.BigEndian.PutUint16(buffer, cid.Seed)

	if v4SrcAddress := ft.Srcip.To4(); v4SrcAddress != nil {
		buffer = append(buffer, v4SrcAddress...)
	} else if v6SrcAddress := ft.Srcip.To16(); v6SrcAddress != nil {
		// As we are now dealing with IPv6, grow the buffer once to fit both addresses.
		buffer = append(make([]byte, 0, maxParameterSizeIPv6), buffer...)
		buffer = append(buffer, v6SrcAddress...)
	}
	if v4DstAddress := ft.Dstip.To4(); v4DstAddress != nil {
		buffer = append(buffer, v4DstAddress...)
	} else if v6DstAddress := ft.Dstip.To16(); v6DstAddress != nil {
		buffer = append(buffer, v6DstAddress...)
	}
	buffer = append(buffer, ft.Proto, 0)
	buffer = binary.BigEndian.AppendUint16(buffer, ft.Srcport)
	buffer = binary.BigEndian.AppendUint16(buffer, ft.Dstport)

	h := sha1.New()
	h.Write(buffer)
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
