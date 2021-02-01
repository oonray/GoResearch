package main

import (
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

func echo(conn net.Conn) {
	defer conn.Close()
	_, err := io.Copy(conn, conn)
	if err != nil {
		logrus.Errorf("Could not Read/Write %s", err)
	}
}

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")

	if err != nil {
		logrus.Errorf("Unable to conenct")
	}

	defer dst.Close()

	go func() {
		_, err := io.Copy(dst, src)
		if err != nil {
			logrus.Errorf("Could Not read/Write: %s", err)
		}
	}()

	_, err = io.Copy(src, dst)
	if err != nil {
		logrus.Errorf("Could Not read/Write: %s", err)
	}
}

func main() {
	listener, err := net.Listen("http", "0.0.0.0")
	if err != nil {
		logrus.Errorf("Unable To Bind")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Errorf("Unable to Accept")
		}

		go handle(conn)
	}
}
