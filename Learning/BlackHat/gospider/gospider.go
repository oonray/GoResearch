package main

import (
	"flag"
	"net/url"
	"os"
	"time"

	"github.com/Willyham/gospider/spider"
	"github.com/sirupsen/logrus"
)

func main(){
	u := flag.String("url","","The url To Spider")
	flag.Parse()

	uri,_ := url.Parse(*u)
	spdr := spider.New(
		spider.WithRoot(uri),
		spider.WithConcurrency(5),
		spider.WithTimeout(time.Second*2),
		)

	err := spdr.Run()
	if(err != nil){
		logrus.Errorf("Spider Failed: %s",err)
	}

	err = spdr.Report(os.Stdout)
	if(err != nil){
		logrus.Errorf("Spider Failed: %s",err)
	}

	return
}
