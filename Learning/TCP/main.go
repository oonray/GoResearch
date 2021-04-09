package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type con struct {
	in 	chan string;	
	out chan string;
	soc net.Conn;
}

func (c *con) Connect (host string, port string) error {
	var err error
	c.soc,err = net.Dial("tcp",fmt.Sprintf("%s:%s",host,port))	
	if(err != nil) {
		return err
	}
	return nil
}

func (c *con)Write(p []byte) (int,error) {
	return c.soc.Write(p)
}

func (c *con)Read(p []byte) (int,error) {
	return c.soc.Read(p)
}

type outin struct{}

func (outin)Write(p []byte) (int,error) {
	return os.Stdout.Write(p)
}

func (outin)Read(p []byte) (int,error) {
	return os.Stdin.Read(p)
}

func main(){
	host := flag.String("H","","The host to user")
	port := flag.String("p","","The host to user")
	flag.Parse()

	if *host == "" {flag.PrintDefaults();return}
	if *port == "" {flag.PrintDefaults();return}

	conn := &con{
		in: make(chan string),
		out: make(chan string),
	}

	stio := &outin{}


	conn.Connect(*host,*port)
	go io.Copy(conn,stio)
	io.Copy(stio,conn)

	
}
