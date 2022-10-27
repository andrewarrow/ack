package tcp

type Segment struct {
	Header Header
	Data   []byte
}

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
