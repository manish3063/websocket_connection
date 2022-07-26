package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sendWp(quote string) {

	apiurl := "https://api.twilio.com/2010-04-01/Accounts/AC12e8e2d86e1eb3c88cf7bb9e678fd9e7/Messages.json"
	method := "POST"

	data := url.Values{}
	data.Set("To", "whatsapp:+919039875487")
	data.Set("From", "whatsapp:+14155238886")
	data.Set("MessagingServiceSid", "MG1bafff1c3e6f459379f22b27c086efb8")
	data.Set("Body", quote)
	client := &http.Client{}
	req, err := http.NewRequest(method, apiurl, strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic QUMxMmU4ZTJkODZlMWViM2M4OGNmN2JiOWU2NzhmZDllNzpiODYxZGMyNzg1NDE2YjcwZjllYTcwYjNkODNkNDk2MA==")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
