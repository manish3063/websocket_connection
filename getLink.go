package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getlink() string {

	url := "https://zenquotes.io/api/random"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}
	//fmt.Println(string(body))
	x := []linkDetails{}
	json.Unmarshal(body, &x)

	return x[0].Q
}
