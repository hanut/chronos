package main

import (
	"fmt"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/json-iterator/go"
	"log"
	"net/http"
	"time"
)

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	//handle connected
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		//join them to room
		c.Join("chat")
	})
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("Connected")

		c.Emit("/message", Message{10, "main", "using emit"})

		c.Join("test")
		c.BroadcastTo("test", "/message", Message{10, "main", "using broadcast"})
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})

	server.On("/join", func(c *gosocketio.Channel, channel Channel) string {
		time.Sleep(2 * time.Second)
		log.Println("Client joined to ", channel.Channel)
		return "joined to " + channel.Channel
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(":3811", serveMux))
}

func jsonParse(json []byte) {
	fmt.Printf("Json Parse Test\n")
	user := Utd{}
	jsonData := []byte(`{"phone":"9717899743","balance":400.00,"cpm":4,"startedAt":"2017-08-29T23:16:15.917Z"}`)
	var tmp = jsoniter.Unmarshal(jsonData, &user)
	log.Println("Json Data-\n" + string(jsonData))
	log.Println("Parsed Data-\n" + user.String())
	log.Println("Balance : ", user.Balance)
	log.Println(tmp)
}
