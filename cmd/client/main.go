package main

import (
	"fmt"
	"net"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		conn net.Conn
		buf  = make([]byte, 6)
		err  error
	)

	conn, err = net.Dial("tcp4", ":9999")
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Read(buf)

	fmt.Printf("RESPONSE FROM: %s\n", string(buf))

	return nil
}
