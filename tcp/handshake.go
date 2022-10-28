package tcp

import "fmt"

type Server struct {
	State    string
	Port     int
	Outgoing chan *Segment
	Incoming chan *Segment
}

func NewServer() *Server {
	s := Server{}
	s.State = "CLOSED"
	s.Outgoing = make(chan *Segment, 1024)
	s.Incoming = make(chan *Segment, 1024)
	return &s
}

func (s *Server) Listen(port int) {
	s.State = "LISTEN"
	s.Port = port
	go s.HandleIncoming()
}

func (s *Server) HandleIncoming() {
	for seg := range s.Incoming {
		fmt.Println("server seg!", seg)
	}
}

type Client struct {
	State    string
	Outgoing chan *Segment
	Incoming chan *Segment
}

func NewClient() *Client {
	c := Client{}
	c.State = "CLOSED"
	c.Outgoing = make(chan *Segment, 1024)
	c.Incoming = make(chan *Segment, 1024)
	go c.HandleOutgoing()
	return &c
}

func (c *Client) Connect(port int) {
	seg := NewSegment()
	seg.Header.Sequence = 100
	seg.Header.SetFlag(SYN, 1)
	fmt.Println(seg.Header.String())
	c.Outgoing <- seg
	c.State = "SYN-SENT"
}

func (c *Client) HandleOutgoing() {
	for seg := range c.Outgoing {
		fmt.Println("seg!", seg)
	}
}
