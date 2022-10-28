package tcp

import (
	"sync"
	"testing"
)

func TestHandshake(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

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

	wg.Wait()
}
