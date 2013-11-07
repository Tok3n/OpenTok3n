package main

import (
	"fmt"
	"log"
	"net/http"
	)

func main() {
	

	err := configure();
	if err != nil{
		log.Fatal(err)
	}

	var localUrl = fmt.Sprintf("0.0.0.0:%s",configData.Address,configData.Port)
	log.Print(localUrl)

	registerHandlers()
	
	err = http.ListenAndServe(localUrl, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func registerHandlers(){
	http.Handle("/", http.HandlerFunc(initRoot))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/github.com/Tok3n/OpenTok3n/webResources/static/"))))
}

func initRoot(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w,"Hello Open Tok3n, Just for the lols at port: %s",configData.Port)
}