// This is example to Use SO_REUSEPORT by Golang.
//
// net.ListenConfig の Control フィールドに指定する func には
// 「OS側にバインドされる前」のコネクションが渡される。
//
// なので、バインド前に設定しないと行けないソケットオプションなどは
// このタイミングで指定することが出来る。
//
// # REFERENCES:
//	- https://christina04.hatenablog.com/entry/go-so-reuseport
//	- https://pkg.go.dev/net@go1.20#ListenConfig

package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		listenCtl = net.ListenConfig{
			Control: setSockOpt,
		}
		listener net.Listener
		sigCh    = make(chan os.Signal, 1)
		myNo     = os.Args[1]
		err      error
	)

	listener, err = listenCtl.Listen(context.Background(), "tcp4", ":9999")
	if err != nil {
		return err
	}

	signal.Notify(sigCh, os.Interrupt)
	go func() {
		defer listener.Close()
		<-sigCh
		fmt.Println("shutdown..." + myNo)
	}()

	for {
		var (
			conn    net.Conn
			connErr error
		)

		conn, connErr = listener.Accept()
		if connErr != nil {
			return err
		}

		go func(c net.Conn) {
			defer c.Close()
			c.Write([]byte(myNo))
		}(conn)
	}
}

func setSockOpt(network string, address string, rc syscall.RawConn) error {
	var (
		sockOpErr error
		sockOpFn  = func(fd uintptr) {
			sockOpErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		}
		err error
	)

	err = rc.Control(sockOpFn)
	if err != nil {
		return err
	}

	if sockOpErr != nil {
		return sockOpErr
	}

	return nil
}
