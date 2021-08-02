package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
    _ = mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello World")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/",DefaultHandler)

	srv := &http.Server{
			Handler: r,
			Addr:    "0.0.0.0:5801",
	}

	logrus.Fatal(srv.ListenAndServe())
}
