package tcp

import "testing"

func TestHandshake(t *testing.T) {
	s := NewServer()
	s.Listen(80)
}
