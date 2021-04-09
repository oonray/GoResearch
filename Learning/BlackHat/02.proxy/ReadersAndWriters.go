package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type FooReader struct {}

func (foo *FooReader) Read(b []byte) (int,error){
	fmt.Printf("$> ")
	return os.Stdin.Read(b)
}


type FooWriter struct {}

func (foo *FooWriter) Write(b []byte) (int, error){
	return os.Stdout.Write(b)
}

func main(){
	var reader FooReader
	var writer FooWriter

	var input []byte = make([]byte, 4096)

	s, err := reader.Read(input)
	if(err != nil){
		logrus.Errorf("Could not read data: %s",err)
	}
	logrus.Infof("Read %d bytes",s)

	s, err = writer.Write(input)
	if(err != nil){
		logrus.Errorf("Could not write data: %s",err)
	}

	_, err = io.Copy(&writer,&reader)
	if  err != nil {
		logrus.Errorf("Unable to Read/Write")
	}

	logrus.Infof("Wrote %d bytes",s)
}

