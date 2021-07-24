package http_methods_golang

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//http method:  GET、POST、PUT、DELETE, change here if needed

//GET: To get data from GET method, need to wait for respone
//<url>: request address.
func GET(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("GET empty URL")
	} else if url == "empty" { //return null
		return nil, nil
	}
	//Check if is an URL start with http
	/*
		var expression string = `https?:\/\/?[-a-zA-Z0-9@:%._\+~#=]{1,256}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
		matched, err := regexp.Match(expression, []byte(url))
		if !matched {
			return nil, errors.New("Input an Illegal URL:" + url)
		}
	*/
	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	sitemap, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return sitemap, nil
}

//POST: To get data from POST method, need to wait for respone
//<url>: request address.
//<query>: The query you want to do, mostly are json, but still depends on the server you request.
func POST(url string, query []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatal(err)
	}
	sitemap, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return sitemap, nil
}
