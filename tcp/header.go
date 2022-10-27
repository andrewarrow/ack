package tcp

import (
	"fmt"
	"strconv"
	"strings"
)

const FIN = 15 // finish
const SYN = 14 // synchronize
const RST = 13 // reset
const PSH = 12 // push, Asks to push the buffered data
const ACK = 11 // ack, Acknowledgment
const URG = 10 // urgent
const ECE = 9  // Echo of Congestion Encountered
const CWR = 8  // Congestion window reduced
const NS = 7   // Explicit Congestion Notification Nonce

type Header struct {
	Source              uint16
	Destination         uint16
	Sequence            uint32
	Acknowledgment      uint32
	OffsetReservedFlags uint16 // 4 + 3 + 9 = 16
	Window              uint16
	Checksum            uint16
	Urgent              uint16
	Options             []uint32 // up to 10
}

func NewHeader() *Header {
	h := Header{}
	return &h
}

func (h *Header) String() string {
	bitString := fmt.Sprintf("%016b", h.OffsetReservedFlags)
	buffer := []string{}
	for i, c := range bitString {
		char := fmt.Sprintf("%c", c)
		if char == "1" && i == SYN {
			buffer = append(buffer, "SYN")
		}
	}
	flags := strings.Join(buffer, ",")
	return fmt.Sprintf("%d %s", h.Sequence, flags)
}

// TODO replace string use with golang bit logic if performance needed
func (h *Header) SetFlag(flag int, value byte) {
	bitString := fmt.Sprintf("%016b", h.OffsetReservedFlags)
	buffer := []string{}
	for i, c := range bitString {
		char := fmt.Sprintf("%c", c)
		if i == flag {
			char = fmt.Sprintf("%d", value)
		}
		buffer = append(buffer, char)
	}

	newValue, _ := strconv.ParseInt(strings.Join(buffer, ""), 2, 64)
	h.OffsetReservedFlags = uint16(newValue)
}
