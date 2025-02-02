package ch03

import (
	"io"
	"net"
	"testing"
)

func TestMyDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:7384")
	if err != nil {
		t.Fatal(err)
	}

	done :=make(chan struct{})

	go func() {
		defer func() { done <- struct{}{} }()

		for {
			conn, err := listener.Accept()

			if err != nil {
				t.Log(err)
				return
			}

			if conn != nil {
				t.Logf("conn open")
			}

			go func(c net.Conn){
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				t.Logf("conn go rotuien")

				buf := make([]byte, 1024)

				for {
					n, err := c.Read(buf)
					if err != nil {
						t.Error(err)
						if err != io.EOF {
							t.Error(err)
						}
						return
					}

					t.Logf("received: %q", buf[:n])
				}

			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String());

	conn.Write([]byte("dial message"))
	conn.Write([]byte("dial message2"))
	conn.Write([]byte("dial message3"))

	if err != nil{
		t.Fatal(err)
	}

	conn.Close()
	<-done
	listener.Close()
	<-done
}