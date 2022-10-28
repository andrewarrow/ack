package tcp

import "fmt"

var Transport = make(chan *Segment, 1024)

//var PortMap = map[uint16]*chan *Segment{}

func init() {
	go HandleTransport()
}

func HandleTransport() {
	for seg := range Transport {
		fmt.Println("transport!", seg)
	}
}
