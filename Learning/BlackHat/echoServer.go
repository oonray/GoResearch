package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

func echo(con net.Conn){
	defer con.Close()
	var b []byte = make([]byte,512)
	var size int
	var err error

	for{
		size,err = con.Read(b[0:])
		if(err != nil){
			if(err == io.EOF){
				logrus.Errorf("The clent disconnected")
				break
			}
			logrus.Error("Could not read: %s",err)
			break
		}

		logrus.Infof("Recieved %d bytes: %s",size,string(b[:size]))
		_,err = con.Write(b[:size])
		if(err != nil){
			logrus.Errorf("Could Not Write Data")
		}
	}
}

func main(){
	listener, err := net.Listen("tcp",":20080")
	if(err != nil){
		logrus.Errorf("Could not start listener %s",err)
		return
	}

	logrus.Infof("Listening on :20080")
	for{
		con, err := listener.Accept()
		if(err != nil){
			logrus.Errorf("Could not Accept request")
			continue
		}
		go echo(con)
	}
}