package shodan

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/shadowscatcher/shodan"
	"github.com/sirupsen/logrus"
)


const BaseUrl string = "https://api.shodan.io"

type Client struct {
	apiKey string
}

type ApiInfo struct {
	Query_Credits int    `json:"query_credits,omitempty"`
	Scan_Credits  int    `json:"scan_credits,omitempty"`
	Telnet        bool   `json:"telnet,omitempty"`
	Plan          string `json:"plan,omitempty"`
	Https         bool   `json:"https,omitempty"`
	Unlocked      bool   `json:"unlocked,omitempty"`
}

type HostLocation struct {
	City         string  `json:"city,omitempty"`
	RegionCode   string  `json:"region_code,omitempty"`
	AereaCode    string  `json:"aerea_code,omitempty"`
	Logitude     float32 `json:"logitude,omitempty"`
	Latitude     float32 `json:"latitude,omitempty"`
	CountryCode3 string  `json:"country_code3,omitempty"`
	CountryName  string  `json:"country_name,omitempty"`
	PostalCode   string  `json:"postal_code,omitempty"`
	DMACode      string  `json:"dma_code,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
}
	
type Host struct {
	OS        string       `json:"os,omitempty"`
	Timestamp string       `json:"timestamp,omitempty"`
	ISP       string       `json:"isp,omitempty"`
	ASN       string       `json:"asn,omitempty"`
	Hostnames string       `json:"hostnames,omitempty"`
	Location  HostLocation `json:"location,omitempty"`
	IP        int64        `json:"ip,omitempty"`
	Domains   []string     `json:"domains,omitempty"`
	Org       string       `json:"org,omitempty"`
	Data      string       `json:"data,omitempty"`
	Port      int          `json:"port,omitempty"`
	IPString  string       `json:"ip_string,omitempty"`
}

type HostSearch struct {
	Matches []Host `json:"matches,omitempty"`
}

func New(apiKey string) *Client {
	return &Client{apiKey:apiKey}
}

func (c *Client) ApiInfo() (*ApiInfo,error) {
	var data ApiInfo 

	url := fmt.Sprintf("%s/api-info?key%s",BaseUrl,c.apiKey)
	resp, err := http.Get(url)
	if(err!=nil){
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if(err != nil){
		return nil, err	
	}
	return &data,nil
}

func (c *Client)HostSearch(query string) (*HostSearch,error){
	var ret HostSearch

	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s",BaseUrl,c.apiKey,query))

	if(err != nil){
		logrus.Errorf("Could not query shodan %s",err)
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ret)
	if(err != nil){
		logrus.Errorf("Could not parse shodan info %s",err)
		return nil, err
	}
	return &ret, err	
}

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

