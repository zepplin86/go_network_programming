package ch03

import (
	"net"
	"testing"
)

func TestListenConneter(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:7384")
	if err != nil {
		t.Fatal(err)
	}

	defer func () {
		_ = listener.Close()
	}()

	t.Logf("bound to %q", listener.Addr())

	go func() {

		conn, err := net.Dial("tcp", listener.Addr().String())

		if err != nil {
			t.Error(err)
			return
		}

		conn.Close()

	}()

	conn, err := listener.Accept()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("conn to %q", conn.LocalAddr().String())


	conn.Close()
}
