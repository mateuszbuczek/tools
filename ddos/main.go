package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/mateuszbuczek/tools/ddos/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	command := ReadFile()

	client := GetRestClient()

	for n := 0; n <= int(command.NumberOfCalls); n++ {
		body := GetAndTransformBody(command)
		request := GetRequest(command, body)
		call, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		fmt.Println(call)
	}
}

func GetRequest(command *Request, body interface{}) *http.Request {
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
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
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

func GetAndTransformBody(command *Request) interface{} {
	jsonBytes, err := json.Marshal(command.Body)
	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile("randomString\\((\\d+),(\\d+)\\)")
	replacedJsonString := reg.ReplaceAllStringFunc(string(jsonBytes), transform.ReplaceWithRandomString)

	var body interface{}
	if err = json.Unmarshal([]byte(replacedJsonString), &body); err != nil {
		panic(err)
	}
	return body
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
