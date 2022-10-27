package tcp

import (
	"fmt"
	"testing"
)

func TestSetFlag(t *testing.T) {
	h := NewHeader()
	h.SetFlag(SYN, 1)
	bits := fmt.Sprintf("%016b", h.OffsetReservedFlags)
	if bits != "0000000000000010" {
		t.Fatal("wrong bits", bits)
	}
}
