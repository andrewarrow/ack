package tcp

import "testing"

func TestHandshake(t *testing.T) {
	s := NewServer()
	s.Listen(80)
	if s.State != "LISTEN" {
		t.Fatal("wrong state", s.State)
	}
	c := NewClient()
	c.Connect(80)
	if c.State != "SYN-SENT" {
		t.Fatal("wrong state", c.State)
	}
}
