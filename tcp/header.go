package tcp

const FIN = 0 // finish
const SYN = 1 // synchronize
const RST = 2 // reset
const PSH = 3 // push, Asks to push the buffered data
const ACK = 4 // ack, Acknowledgment
const URG = 5 // urgent
const ECE = 6 // Echo of Congestion Encountered
const CWR = 7 // Congestion window reduced
const NS = 8  // Explicit Congestion Notification Nonce

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

func (h *Header) SetFlag(flag byte, value byte) {
}
