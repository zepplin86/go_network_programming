package main

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestMyReadInfoBuffer(t *testing.T) {
	//비트 연산자 16 byte
	payload := make([]byte, 16)
	//난수 데이터 payload 에 채움
	_, err := rand.Read(payload)

	t.Logf("data: %x", payload)

	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:7384")

	if err != nil {
		t.Fatal(err)
	}

	//서버
	go func() {
		//연결 수립되면 payload 읽어들여 tcp 데이터씀
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)

		if err != nil {
			t.Error(err)
		}
	}()

	//클라이언트
	conn, err := net.Dial("tcp", listener.Addr().String())

	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 2) // 1byte buffer

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			//eof 들어오면 break;
			break;
		}
		t.Logf("read data: %x", buf[:n])
		//t.Logf("read %d bytes", n) // buf[:n] 읽은 데이터 바이트
	}

	conn.Close()
}