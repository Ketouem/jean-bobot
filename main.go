package main

import (
  "fmt"
  "time"
  gdax "github.com/preichenberger/go-gdax"
  ws "github.com/gorilla/websocket"
)

func main() {
  var wsDialer ws.Dialer
  wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)
  if err != nil {
    println(err.Error())
  }

  subscribe := gdax.Message{
    Type:      "subscribe",
    Channels: []gdax.MessageChannel{
      gdax.MessageChannel{
        Name: "ticker",
        ProductIds: []string{
          "BTC-EUR",
          "LTC-EUR",
          "ETH-EUR",
        },
      },
    },
  }
  if err := wsConn.WriteJSON(subscribe); err != nil {
    println(err.Error())
  }

  message:= gdax.Message{}
  for true {
    if err := wsConn.ReadJSON(&message); err != nil {
      println(err.Error())
      break
    }

    if message.Type == "match" {
      println("Got a match")
    } else if message.Type == "ticker" {
      price := message.Price
      productId := message.ProductId
      time := message.Time.Time().Format(time.RFC3339)
      fmt.Printf("%v [%v] : %v\n", productId, time, price)
    }
  }
}
