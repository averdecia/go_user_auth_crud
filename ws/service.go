package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Clients map to save all client connections
var Clients map[string]*Client = make(map[string]*Client)

// Reader channel to receive
var Reader chan Message

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Serve function to convert http to websocket
func Serve(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	userID := "wdfdg"

	client := &Client{
		Clients:       &Clients,
		UserID:        userID,
		Conn:          conn,
		Send:          make(chan Message, 256),
		MessageReader: Reader,
	}

	Clients[userID] = client

	go client.readMessages()
	go client.sendMessage()

}
