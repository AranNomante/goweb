package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/"{
		fmt.Fprintf(w,"<h1>Welcome to site</h1>")
	}else if r.URL.Path == "/contact"{
		fmt.Fprintf(w,"To get in touch, please send an email to <a href=\"mailto:altug.ceylan@dt.net.tr\">Altug Ceylan</a>")
	}else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"Page not found, please send an email to <a href=\"mailto:altug.ceylan@dt.net.tr\">Altug Ceylan if you keep...</a>")
	}
}
func main(){
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServe(":3000",nil)
}