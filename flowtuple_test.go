package gommunityid

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowTupleOrder(t *testing.T) {
	tpl := FlowTuple{
		Srcip:   netip.MustParseAddr("1.2.3.4"),
		Dstip:   netip.MustParseAddr("4.5.6.7"),
		Srcport: 1,
		Dstport: 2,
	}
	assert.True(t, tpl.IsOrdered())

	tpl2 := FlowTuple{
		Dstip:   netip.MustParseAddr("1.2.3.4"),
		Srcip:   netip.MustParseAddr("4.5.6.7"),
		Srcport: 2,
		Dstport: 1,
	}
	assert.True(t, !tpl2.IsOrdered())

	tplInOrder := tpl2.InOrder()
	assert.Equal(t, tpl, tplInOrder)
	assert.True(t, tplInOrder.IsOrdered())

	tplWasInOrder := tpl.InOrder()
	assert.Equal(t, tpl, tplWasInOrder)
}
