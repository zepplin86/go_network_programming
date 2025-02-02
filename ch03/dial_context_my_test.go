package ch03

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestMyDialContext(t *testing.T) {
	// 1
	dl := time.Now().Add(5 * time.Second)
	// 2
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	// 3
	defer func (){
		t.Logf("cancel call")
		cancel()
	} ()

	var d net.Dialer

	d.Control = /* 4 */func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(5*time.Second + time.Microsecond)
		return nil
	}

	conn, err := d.DialContext(/* 5 */ctx, "tcp", "8.0.0.0:80")

	if err == nil {
		conn.Close()
		t.Fatal("connnection did not time out")
	} else {
		t.Logf("%s", err.Error())
	}

	nErr, ok := err.(net.Error)

	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error is not a timeout: %v", err)
		}
	}

	/* 6 */if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded; actual: %v", ctx.Err())
	}
}