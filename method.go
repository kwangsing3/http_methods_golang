package http_methods_golang

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//http method:  GET、POST、PUT、DELETE, change here if needed

//GET: To get data from GET method, need to wait for respone
//<url>: request address.
func GET(url string, header map[string]string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("GET empty URL")
	} else if url == "empty" { //return null
		return nil, nil
	}

	/*
		var expression string = `https?:\/\/?[-a-zA-Z0-9@:%._\+~#=]{1,256}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
		matched, err := regexp.Match(expression, []byte(url))
		if !matched {
			return nil, errors.New("Input an Illegal URL:" + url)
		}
	*/ //Check if is an vaild URL

	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	for key, value := range header {
		req.Header.Add(key, value)
	}

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
//<query>: The query you want to do, mostly are json, depends on the server you request.
func POST(url string, header map[string]string, query []byte) ([]byte, error) {

	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	for key, value := range header {
		req.Header.Add(key, value)
	}

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

//DELETE: http Delete method request
func DELETE(url string, header map[string]string, query []byte) ([]byte, error) {
	if url == "" {
		return nil, errors.New("GET empty URL")
	} else if url == "empty" { //return null
		return nil, nil
	}

	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(query))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	for key, value := range header {
		req.Header.Add(key, value)
	}
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

//PUT: http Delete method request
func PUT(url string, header map[string]string, content string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("GET empty URL")
	} else if url == "empty" { //return null
		return nil, nil
	}
	context := strings.NewReader(content)

	client := &http.Client{}
	defer client.CloseIdleConnections()
	req, err := http.NewRequest("PUT", url, context)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	for key, value := range header {
		req.Header.Add(key, value)
	}
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

//ReadFile: Standard Read file operate. <return> []byte.
func ReadFile(inputPath string) ([]byte, error) {

	if path.IsAbs(inputPath) {
		return nil, errors.New("read file path: emtpy")
	}
	file, err := os.Open(path.Join(inputPath))
	if err != nil {
		return nil, errors.New(`read file path: ` + err.Error())
	}
	defer file.Close()
	byteData, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return nil, errors.New(`read file path: ` + readErr.Error())
	}
	return byteData, nil
}

//WriteFile:
func WriteFile(inputPath string, data []byte) error {

	file, err := os.Create(path.Join(inputPath))
	if err != nil {
		return errors.New(`write file path: ` + err.Error())
	}
	defer file.Close()
	file.Write(data)
	return nil
}
