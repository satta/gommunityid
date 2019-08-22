package gommunityid

import (
	"fmt"
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

func GetCommunityIDByVersion(version uint, seed uint16) (CommunityID, error) {
	switch version {
	case 1:
		return CommunityIDv1{
			Seed: seed,
		}, nil
	default:
		return nil, fmt.Errorf("invalid version: %d", version)
	}
}
