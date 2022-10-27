package tcp

import (
	"fmt"
	"testing"
)

func TestToggleFlag(t *testing.T) {
	h := NewHeader()
	h.SetFlag(SYN, 1)
	bits := fmt.Sprintf("%016b", h.OffsetReservedFlags)
	if bits != "1111000000000010" {
		t.Fatal("wrong bits", bits)
	}
	h.SetFlag(ECE, 1)
	bits = fmt.Sprintf("%016b", h.OffsetReservedFlags)
	if bits != "1111000001000010" {
		t.Fatal("wrong bits", bits)
	}
}
