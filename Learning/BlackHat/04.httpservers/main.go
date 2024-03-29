package main

import "net/http"

type badAuth struct {
	username string
	password string
}
  
func (b *badAuth) ServeHttp(w http.ResponseWriter, r *http.Request){
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username != b.Username || password != b.Password {
		http.Error(w,"Unauthorized",401)
		return
	}
	
	ctx := context.WithValue(r.Context(),"username",username)
	r = r.WithContext(ctx)
	next(w,r)
}

func hello(w http.ResponseWriter, r *http.Request){
	username := r.Context().Value("username")
	fmt.Fprintf(w,"Hi %s\n",username)
}

func main(){
	r := mux.NewRouter()	
	r.HandleFunc("/hello",hello).Methods("GET")

	n := negroni.Classic()
	n.Use(&badAuth{
		Username: "admin",
		Password: "password"
	})

	n.UseHndler(r)
	http.
}
