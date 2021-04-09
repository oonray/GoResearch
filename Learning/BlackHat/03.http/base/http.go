package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main(){
	r1, _ := http.Get("https://google.com/robots.txt")
	defer r1.Body.Close()

	form := url.Values{}
	form.Add("foo","baar")
	
	r3, _ := http.Post("https://google.com/robots.txt","application/x-www-form-urlencoded",strings.NewReader(form.Encode()))
	defer r3.Body.Close()

	var client http.Client
	req, _ := http.NewRequest("PUT","https://google.com/robots.txt",strings.NewReader(form.Encode()))
	defer req.Body.Close()

	resp, _ := client.Do(req)
	defer resp.Body.Close()

}

