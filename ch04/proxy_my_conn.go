package main

import (
	"io"
	"net"
)

func proxyMyConn(source, destination string) error {
	connSource, err := net.Dial("tcp", source)
	if err != nil {
		return err
	}

	defer connSource.Close()

	connDestination, err := net.Dial("tcp", destination)

	if err != nil {
		return err
	}
	defer connDestination.Close()

	go func() { _, _ = io.Copy(connSource, connDestination)} ()

	_, err = io.Copy(connDestination, connSource)

	return err
}

var _ = proxyConn