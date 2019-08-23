package gommunityid

import (
	"fmt"
	"hash"
)

// CommunityID is an interface defining the supported operations
// on a component calculating a specific community ID version.
type CommunityID interface {
	Calc(FlowTuple) []byte
	CalcHex(FlowTuple) string
	CalcBase64(FlowTuple) string
	Hash(FlowTuple) hash.Hash
	Render(hash.Hash) []byte
	RenderHex(hash.Hash) string
	RenderBase64(hash.Hash) string
}

// GetCommunityIDByVersion returns, for a given version number and seed, an
// object implementing the CommunityID interface for the specified version.
// This will be preconfigured with the given seed.
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
