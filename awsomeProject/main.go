package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type JsonRPC2 struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result, omitempty"`
	Id      *int        `json: "id, ommitempty"`
}

type SubscribeParams struct {
	Channel string `json:"channel"`
}

func main() {
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	log.Printf("connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalln("dial", err)
	}
	defer c.Close()

	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}}); err != nil {
		log.Fatalln("dial:", err)
		return
	}

	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Fatalln("read:", err)
		}

		if message.Method == "channelMessage" {
			log.Println(message.Params)
		}
	}
}
