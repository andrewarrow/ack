package tcp

type Server struct {
	State string
	Port  int
}

type Client struct {
	State string
}

func NewServer() *Server {
	s := Server{}
	s.State = "CLOSED"
	return &s
}

func NewClient() *Client {
	c := Client{}
	c.State = "CLOSED"
	return &c
}

func (s *Server) Listen(port int) {
	s.State = "LISTEN"
	s.Port = port
	go s.ListenForMessages()
}

func (c *Client) Connect(port int) {
	seg := Segment{}
	seg.Header.Sequence = 100
	c.State = "SYN-SENT"
}

func (s *Server) ListenForMessages() {
	for {
	}
}
