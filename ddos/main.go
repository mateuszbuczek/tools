package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	command := ReadFile()

	fmt.Println(command)

	client := GetRestClient()
	request := GetRequest(command)

	for n := 0; n <= int(command.NumberOfCalls); n++ {
		call, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		fmt.Println(call)
	}
}

func GetRequest(command *Request) *http.Request {
	var request *http.Request

	switch command.Method {
	case "GET":
		req, err := http.NewRequest(http.MethodGet, command.Url, nil)
		if err != nil {
			panic(err)
		}
		request = req
		break
	case "POST":
		jsonBytes, err := json.Marshal(command.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonBytes))
		req, err := http.NewRequest(http.MethodPost, command.Url, bytes.NewReader(jsonBytes))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		request = req
		break
	default:
		panic(fmt.Sprint("Could not recognize method '", command.Method, "'. Available methods: 'GET', 'POST'."))
	}

	return request
}

func GetRestClient() http.Client {
	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := http.Client{Transport: transportConfig}
	return client
}

func ReadFile() *Request {
	file, err := ioutil.ReadFile("request.json")
	if err != nil {
		panic("Unable to find request.json file")
	}

	data := Request{}
	if err = json.Unmarshal(file, &data); err != nil {
		panic("Unable to parse json")
	}

	return &data
}

type Request struct {
	NumberOfCalls int64       `json:"number_of_calls"`
	Method        string      `json:"method"`
	Url           string      `json:"url"`
	Body          interface{} `json:"body"`
}
