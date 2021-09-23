package ws

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Maximum message size allowed from peer.
	maxMessageSize = 512
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
)

// Client struct to map users connections
type Client struct {
	Clients       *map[string]*Client
	UserID        string
	Conn          *websocket.Conn
	Send          chan Message
	MessageReader chan Message
}

func (c *Client) readMessages() {
	defer func() {
		delete(*c.Clients, c.UserID)
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	for {
		_, message, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		parsedMessage, err := GetMessageFromString(message)
		if err != nil {
			log.Printf("Invalid message format: %v", err)
		}
		ExecMessage(&parsedMessage, c)
	}
}

func (c *Client) sendMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:

			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			messageStr, err := message.ToString()
			if err != nil {
				fmt.Println("Unable to parse message")
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Println("Writer not found")
				return
			}

			fmt.Printf("hello : %s", messageStr)
			w.Write(messageStr)
		}
	}
}
