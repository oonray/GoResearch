package main

import (
	"bufio"
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

func better_echo(conn net.Conn) {
	defer conn.Close()

	_, err := io.Copy(conn, conn)
	if err != nil {
		logrus.Errorf("Could not read/write %s", err)

	}
}

func echo(con net.Conn) {
	defer con.Close()
	var err error

	reader := bufio.NewReader(con)
	s, err := reader.ReadString('\n')
	if err != nil {
		logrus.Errorf("Unable to Read data: %s", s)
	}
	logrus.Infof("Read %d Bytes %s", len(s), s)

	logrus.Infof("Writes Data")
	writer := bufio.NewWriter(con)
	_, err = writer.WriteString(s)
	if err != nil {
		logrus.Errorf("Could not Write: %s", err)
	}
	writer.Flush()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		logrus.Errorf("Could not start listener %s", err)
		return
	}

	logrus.Infof("Listening on :20080")
	for {
		con, err := listener.Accept()
		if err != nil {
			logrus.Errorf("Could not Accept request")
			continue
		}
		go echo(con)
	}
}
