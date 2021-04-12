package main

import (
	"flag"
	"os"

	"github.com/oonray/GoResearch/Learning/BlackHat/03.http/godan/shodan"
	"github.com/sirupsen/logrus"
)

func main(){
	search := flag.String("s","","The Search Term")
	flag.Parse()

	if(*search == ""){
		logrus.Errorf("Search term is manditory")
		flag.PrintDefaults()	
	}
	
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)

	info, err := s.ApiInfo()
	if(err != nil){
		logrus.Errorf("Could not get API info %s",err)
	}

	logrus.Infof("Query Credits: %d\nScan Credits: %d\n\n",info.Query_Credits,info.Scan_Credits)

	hostSearch, err := s.HostSearch(*search)
	if err != nil {
		logrus.Errorf("Problem making the search %s",err)
	}

	for _,host := range hostSearch.Matches{
		logrus.Infof("%18s%8d\n",host.IPString,host.Port)
	}
}

