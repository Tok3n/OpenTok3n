package main
import (
	"net/http"
	"io/ioutil"
	//"log"
	"net/url"
	"fmt"

	"code.google.com/p/rsc/qr"
)

func registerApiHandlers(){
	http.Handle("/api/v1/getNewUserQR", http.HandlerFunc(getNewUserSession_API))
	http.Handle("/api/v1/addingUserResponse", http.HandlerFunc(addingUserResponse_API))
}

var TOK3N_DOMAIN = "my.tok3n.com"

func getNewUserSession_API(w http.ResponseWriter, req *http.Request){
	resp, err := http.PostForm(fmt.Sprintf("http://%s/api/v1/openTok3n/getNewUserSession",TOK3N_DOMAIN) , url.Values{"kind": {"OTaccess"}, "secretKey": {configData.Tok3nAPISecret}})
	if err != nil{
		cod,err := qr.Encode(fmt.Sprintf("http://%s/api/v1/openTok3n/getIntegrationError",TOK3N_DOMAIN),qr.M)
		if err != nil{
			fmt.Fprintf(w,"Something went very wrong, verify OpenTok3n integration instalation")
			return
		}
		img := cod.PNG()
		w.Write(img)
		return 
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	cod,err := qr.Encode(fmt.Sprintf("http://%s/api/v1/openTok3n/validateNewUserSession?session=%s&key=%s",TOK3N_DOMAIN,body,configData.Tok3nAPIKey),qr.H)
	if err != nil{
		fmt.Fprintf(w,"Something whent very wrong, verify OpenTok3n integration instalation")
		return
	}
	img := cod.PNG()
	w.Header().Add("Content-Type","image/png")
	w.Write(img)
}

func addingUserResponse_API(w http.ResponseWriter, req *http.Request){
	req.FormValue("userkey")
	req.FormValue("session")
}