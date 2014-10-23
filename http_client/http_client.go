package main

import (
	"bytes"
	//	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	//	"net/url"
	//	"strings"
	//	"sync"
	"time"
)

func myDial(network, address string) (c net.Conn, err error) {
	return net.DialTimeout(network, address, 1 * time.Millisecond)
}

func main() {

	tr := &http.Transport{
		DisableKeepAlives: false,
		Dial:              myDial,
	}

	client := &http.Client{
		Transport: tr,
//		Timeout:   1 * time.Millisecond,
	}

	url := "http://ngs.ru/"

	var body []byte

	request, err := http.NewRequest("GET", url, bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Can't create request. Error: '%s'\n", err.Error())
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("HTTP request failed. Error: '%s'\n", err.Error())
		return

	}
	if response != nil {
		var response_body []byte
		defer response.Body.Close()
		if response.StatusCode > 304 {
			fmt.Printf("HTTP response error status: '%s'\n", response.Status)
			return
		}
		if response_body, err = ioutil.ReadAll(response.Body); err != nil {
			fmt.Printf("can't read HTTP response body: '%s'\n", err.Error())
			return
		}
		fmt.Printf("Body has size %d bytes\n", len(response_body))
	}
}
