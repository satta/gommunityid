package gommunityid

import (
	"testing"
)

func TestCommunityIDMake(t *testing.T) {
	cid, err := GetCommunityIDByVersion(1, 42)
	if err != nil {
		t.Fatal(err)
	}
	cidv1, ok := cid.(CommunityIDv1)
	if !ok {
		t.Fatal("can't assert expected type")
	}
	if cidv1.Seed != 42 {
		t.Fatalf("wrong seed: %d", cidv1.Seed)
	}
}

func TestCommunityIDMakeFail(t *testing.T) {
	_, err := GetCommunityIDByVersion(23, 42)
	if err == nil {
		t.Fatal(err)
	}
}
