package tcp

type Server struct {
}

func NewServer() *Server {
	s := Server{}
	return &s
}

func (s *Server) Listen(port int) {

}
