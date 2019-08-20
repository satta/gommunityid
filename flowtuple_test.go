package gommunityid

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowTupleOrder(t *testing.T) {
	tpl := FlowTuple{
		Srcip:   net.IPv4(1, 2, 3, 4),
		Dstip:   net.IPv4(4, 5, 6, 7),
		Srcport: 1,
		Dstport: 2,
	}
	assert.True(t, tpl.IsOrdered())

	tpl2 := FlowTuple{
		Dstip:   net.IPv4(1, 2, 3, 4),
		Srcip:   net.IPv4(4, 5, 6, 7),
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
