package gommunityid

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPcapFlowTupleSource(t *testing.T) {
	testfiles, err := filepath.Glob("testdata/*.pcap")
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for _, testfile := range testfiles {
		ftChan, err := PcapFlowTupleSource(testfile)
		if err != nil {
			log.Fatal(err)
		}
		for range ftChan {
			i++
		}
		log.Printf("read %s with %d packets", testfile, i)
		assert.Greater(t, i, 0)
	}
}
