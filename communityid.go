package gommunityid

import (
	"hash"
)

type CommunityID interface {
	Calc(FlowTuple) []byte
	CalcHex(FlowTuple) string
	CalcBase64(FlowTuple) string
	Hash(FlowTuple) hash.Hash
	Render(hash.Hash) []byte
	RenderHex(hash.Hash) string
	RenderBase64(hash.Hash) string
}
