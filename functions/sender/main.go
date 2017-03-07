package main

import (
	"encoding/json"
	"github.com/ahmedajx/sky_pinger/messaging"
	"github.com/ahmedajx/sky_pinger/sky"
	"github.com/apex/go-apex"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		heading := sky.LatestChelseaNews()
		response, output := messaging.SendText(heading)
		if response {
			return nil, nil
		}
		message := Message{output}
		b, _ := json.Marshal(message)
		return string(b), nil
	})
}
