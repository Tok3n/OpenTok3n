package main

import (
	//"fmt"
	"flag"
	"log"
	//"net/http"
	)


var portFlag = flag.String("port", "", "server port")
var addrFlag = flag.String("addr", "", "http service address") 
var dbAddrFlag = flag.String("dbaddr", "", "database server address")
var dbPortFlag = flag.String("dbport", "", "database server port")
var dbUserFlag = flag.String("dbuser", "", "database server user")
var dbPassFlag = flag.String("dbpass", "", "database server pass")
var dbDBNameFlag = flag.String("dbname", "", "database name")
var dbTablePrefixFlag = flag.String("dbprefix", "", "database table prefix")



func main() {
	flag.Parse()

	err := configure();
	if err != nil{
		log.Fatal(err)
	}

	/*var localUrl = fmt.Sprintf("%s:%s",*addr,*port)
	log.Print(localUrl)

	registerHandlers()
	
	err := http.ListenAndServe(localUrl, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }*/
}
/*
func registerHandlers(){
	http.Handle("/", http.HandlerFunc(initRoot))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/github.com/Tok3n/OpenTok3n/webResources/static/"))))
}

func initRoot(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w,"Hello Open Tok3n, Just for the lols at port: %s",*port)
}*/