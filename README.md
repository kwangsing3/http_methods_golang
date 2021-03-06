# http_method_golang
&emsp; **http_methods_golang** is a tiny wrapper for implementing HTTP request methods with Golang, and designed for reuse in every projects which don't have to implement those **METHODS**, when a simple communication is in need.

# Methods
&emsp;HTTP Protocol defined serveal request methods for gaining data and contents using Network.
| METHOD  |  STATUS |   
| --------|---------|
| GET     |√|
| POST    |√|
| HEAD    |--|
| PUT     |√|
| DELETE  |√|
| CONNECT |--|
| OPTIONS |--|
| TRACE   |--|
| PATCH   |--| 

| OPERATE |  STATUS |   
| --------|---------|
| READ  FILE   |√|
| WRITE FILE   |√|

# Install
```sh
go get github.com/kwangsing3/http_methods_golang
```



# Usage

* ## GET
1. Would return HTML file if request target is website. <br/>
2. Would return bytes if request a data.
``` go
//GET: To get data from GET method, need to wait for respone
//<url>: request address.
func GET(url string) ([]byte, error)
```
* ## POST
1. Mostly use for encrypt request or requested using ```<form>```.
2. May using for change server status or special request (Depends on the design).
``` go
//POST: To get data from POST method, need to wait for respone
//<url>: request address.
//<query>: The query you want to do, mostly are json, but still depends on the server you request.
func POST(url string, query []byte) ([]byte, error) 
```

* ## DELETE
``` go
//DELETE: http Delete method request
func DELETE(url string) ([]byte, error) 
```
* ## PUT
``` go
//PUT: http Delete method request
func PUT(url string, content string) ([]byte, error) 
```
# Example
```go
package main

import (
	"encoding/json"
	"fmt"
	HMG "github.com/kwangsing3/http_methods_golang"
)

func main() {

	/***GET request***/
	dataGET, errG := HMG.GET("https://example.com")
	if errG != nil {
		fmt.Println(errG.Error())
	} 

	fmt.Println(string(dataGET))
	/*** POST request***/
	query := struct {
		Msg string
	}{
		Msg: `New Message`, //query Struct depends on the server you request.
	}
	bytequery, _ := json.Marshal(query)
	dataPOST, errP := HMG.POST("https://example.com", bytequery)
	if errP != nil {
		fmt.Println(errP.Error())
	} 

	fmt.Println(string(dataPOST))
    return
}

```
