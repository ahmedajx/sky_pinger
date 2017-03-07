package messaging

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/ahmedajx/sky_pinger/env"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ev = env.Load()

var (
	ACCOUNTSID = ev.AuthToken
	AUTHTOKEN  = ev.AccountSid
	FROM       = ev.From
	TO         = ev.To
)

type ErrorMessage struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func SendText(heading string) bool {
	client := &http.Client{}
	v := url.Values{}
	v.Set("From", FROM)
	v.Add("To", TO)
	v.Add("Body", heading)

	request, _ := http.NewRequest(
		"POST",
		"https://api.twilio.com/2010-04-01/Accounts/AC0dffb22f9c25820d4fb6a6460987bea7/Messages.json",
		bytes.NewBufferString(v.Encode()),
	)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic "+basicAuth(ACCOUNTSID, AUTHTOKEN))
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var errorMessage ErrorMessage
	json.Unmarshal(body, &errorMessage)
	if errorMessage.Message == "" {
		return true
	}
	return false
}

func basicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString(
		[]byte(username + ":" + password),
	)
}
