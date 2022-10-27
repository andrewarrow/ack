package tcp

const MAX_SEGMENT_SIZE = 512       // MSS
const MAX_TRANSMISSION_UNIT = 1500 // MTU 1460 bytes, 40 bytes headers

type Segment struct {
	Header *Header
	Data   []byte
}

func NewSegment() *Segment {
	s := Segment{}
	s.Header = NewHeader()
	return &s
}

func BreakIntoSegments(text string) []string {
	segments := []string{}
	for {
		if len(text) < MAX_SEGMENT_SIZE {
			segments = append(segments, text)
			break
		}
		segment := text[0:MAX_SEGMENT_SIZE]
		segments = append(segments, segment)
		text = text[MAX_SEGMENT_SIZE:]
	}
	return segments
}
