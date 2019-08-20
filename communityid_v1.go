package gommunityid

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash"
)

type CommunityIDv1 struct {
	Seed uint16
}

func (cid CommunityIDv1) Calc(ft FlowTuple) []byte {
	ft = ft.InOrder()
	return cid.Render(cid.Hash(ft))
}

func (cid CommunityIDv1) CalcHex(ft FlowTuple) string {
	ft = ft.InOrder()
	return cid.RenderHex(cid.Hash(ft))
}

func (cid CommunityIDv1) CalcBase64(ft FlowTuple) string {
	ft = ft.InOrder()
	return cid.RenderBase64(cid.Hash(ft))
}

func (cid CommunityIDv1) Hash(ft FlowTuple) hash.Hash {
	h := sha1.New()
	binary.Write(h, binary.BigEndian, cid.Seed)
	if ft.Srcip.To4() != nil {
		binary.Write(h, binary.BigEndian, ft.Srcip.To4())
	} else if ft.Srcip.To16() != nil {
		binary.Write(h, binary.BigEndian, ft.Srcip.To16())
	}
	if ft.Dstip.To4() != nil {
		binary.Write(h, binary.BigEndian, ft.Dstip.To4())
	} else if ft.Dstip.To16() != nil {
		binary.Write(h, binary.BigEndian, ft.Dstip.To16())
	}
	h.Write([]byte{ft.Proto, 0})
	binary.Write(h, binary.BigEndian, ft.Srcport)
	binary.Write(h, binary.BigEndian, ft.Dstport)
	return h
}

func (cid CommunityIDv1) Render(h hash.Hash) []byte {
	return h.Sum(nil)
}

func (cid CommunityIDv1) RenderBase64(h hash.Hash) string {
	return "1:" + base64.StdEncoding.EncodeToString(cid.Render(h))
}

func (cid CommunityIDv1) RenderHex(h hash.Hash) string {
	return fmt.Sprintf("1:%x", cid.Render(h))
}
