package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"sort"
)

func scan(host string, port int) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	conn.Close()

	logrus.Infof("Port %d Open", port)
	return nil
}

func worker(ports chan int, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("scanme.nmap.org:%d", p)
		con, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		con.Close()
		results <- p
	}
}

func main() {
	var ports chan int = make(chan int, 100)
	var results chan int = make(chan int, 1024)
	var openports []int

	for i := 1; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 0; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	fmt.Print("\n")
	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		logrus.Infof("Port %d Open", port)
	}
}
