package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
)

func scan(host string,port int,wg sync.WaitGroup) error {
	defer wg.Done()
	conn, err := net.Dial("tcp",fmt.Sprintf("%s:%d",host,port))
	if(err != nil){
		return err
	}
	defer conn.Close()

	logrus.Infof("Port %s Open",port)
	return nil
}

func main(){
	var wg sync.WaitGroup
	for i := 0; i<= 1024; i++ {
		wg.Add(1)
		go scan("scanme.nmap.org",i,wg)
	}
	wg.Wait()
}
