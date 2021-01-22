package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"url-shortener/common"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	Rebrandly_API_Key = <YOUR-REBRANDLY-API-KEY> //create your rebrandly api key from https://www.rebrandly.com/
	Rebrandly_API     = "https://api.rebrandly.com/v1/links"
)

func ShortenUrlHandler(res http.ResponseWriter, req *http.Request) {

	//read original long url from Post Request JSON Body
	log.Info("reading request body...")
	decoder := json.NewDecoder(req.Body)

	var rb common.RequestBody
	err := decoder.Decode(&rb)
	if err != nil {
		log.Fatalf("cannot decode request body %+v", rb)
	}

	log.Infof("original long url: %s", rb.OriginalUrl)

	log.Info("creating request to send to rebrandly api...")
	//create a new http client
	client := http.Client{}

	//create the request body to call Rebrandly API
	domain := common.DomainName{
		FullName: "rebrand.ly",
	}

	brb := common.BrandlyRequestBody{
		Destination: rb.OriginalUrl,
		Domain:      domain,
	}

	//Marshal the request body to convert it to JSON
	brbjson, err := json.Marshal(brb)
	if err != nil {
		log.Fatalf("cannot marshal request body:%+v", err)
	}

	//create the request to call Rebrandly API
	breq, err := http.NewRequest("POST", Rebrandly_API, bytes.NewBuffer(brbjson))
	if err != nil {
		log.Fatalf("cannot create request %+v", err)
	}

	breq.Header.Set("Content-Type", "application/json")
	breq.Header.Set("apikey", Rebrandly_API_Key)

	log.Info("sending request to rebrandly api...")
	resp, err := client.Do(breq)
	if err != nil {
		log.Fatalf("cannot send request to Rebrandly API %+v", err)
	}

	//check the response Status
	log.Info("getting response from rebrandly api...")
	log.Infof("checking response status... %s ", resp.Status)
	if resp.StatusCode == 400 {
		log.Fatalf("Bad Request:%+v", err)
	}

	//read JSON body from response
	log.Info("reading response body...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("cannot read body %+v", err)
	}
	defer resp.Body.Close() //close the response body after reading

	//unmarshal the rebrandly response body
	var brandlyresponsebody common.BrandlyResponseBody
	err = json.Unmarshal(body, &brandlyresponsebody)
	if err != nil {
		log.Fatalf("cannot unmarshal response body: %+v", err)
	}

	//extract the short url from the rebrandly response
	log.Info("extracting shorten url...")
	short_url := brandlyresponsebody.ShortURL

	log.Infof("original url : %s", rb.OriginalUrl)
	log.Infof("shorten url : %s", short_url)

	res.Write([]byte(fmt.Sprintf("original url : %s\n", rb.OriginalUrl)))
	res.Write([]byte(fmt.Sprintf("shorten url : %s\n", short_url)))

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", ShortenUrlHandler).Methods("Post")
	log.Info("starting web server on 8080...")
	http.ListenAndServe(":8080", router)

}
