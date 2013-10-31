package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	)


var port = flag.String("port", "8080", "server port")
var addr = flag.String("addr", "localhost", "http service address") // Q=17, R=18
func main() {
	flag.Parse()

	var localUrl = fmt.Sprintf("%s:%s",*addr,*port)
	log.Print(localUrl)


	http.Handle("/", http.HandlerFunc(initRoot))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/github.com/Tok3n/OpenTok3n/webResources/static/"))))
	
	err := http.ListenAndServe(localUrl, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func initRoot(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w,"Hello Open Tok3n, Just for the lols at port: %s",*port)
}