package main

import (
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
		message := Message{"Successfully sent."}
		b, _ := json.Marshal(message)
		if response := messaging.SendText(heading); response {
			return string(b), nil
		}
		return nil, nil
	})
}
