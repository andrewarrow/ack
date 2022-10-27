package tcp

const MAX_SEGMENT_SIZE = 512       // MSS
const MAX_TRANSMISSION_UNIT = 1500 // MTU 1460 bytes, 40 bytes headers

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

func BreakIntoSegments(text string) []string {
	if len(text) < MAX_SEGMENT_SIZE {
		return []string{text}
	}
	segments := []string{}
	for {
		segment := text[0:MAX_SEGMENT_SIZE]
		segments = append(segments, segment)
		if true {
			break
		}
	}
	return segments
}
