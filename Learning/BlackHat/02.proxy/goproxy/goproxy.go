package main

import (
	"flag"
	"fmt"
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

func handle(src net.Conn, proto string, host string, port int) {
	dst, err := net.Dial(proto, fmt.Sprintf("%s:%d", host, port))

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
		http := flag.Bool("ht", false, "Use HTTP")
		udp := flag.Bool("u", false, "Use UDP")
		host := flag.String("H", "127.0.0.1", "Host to connect to")
		port := flag.Int("P", 80, "Port to bind to")
		host_b := flag.String("B", "0.0.0.0", "Host to bind to")
		port_b := flag.Int("p", 80, "Port to connect to")
		flag.Parse()

		proto := "tcp"

		if *http {
			proto = "http"
		}

		if *udp {
			proto = "udp"
		}

		url := fmt.Sprintf("%s:%d", *host_b, *port_b)
		logrus.Infof("Binding to %s", url)

		listener, err := net.Listen(proto, url)
		if err != nil {
			logrus.Errorf("Unable To Bind")
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				logrus.Errorf("Unable to Accept")
			}

			go handle(conn, proto, *host, *port)
		}
}
