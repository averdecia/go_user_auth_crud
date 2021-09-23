package ws

import (
	"encoding/json"
	"fmt"
)

// MessageType is used to select the type of messages
type MessageType int

// Test types of messages
const (
	Test MessageType = iota
	Open
)

// Message struct to map each message from and to server
type Message struct {
	Type MessageType
	Data interface{}
}

// ToString func to convert from Message to string
func (m *Message) ToString() ([]byte, error) {
	return json.Marshal(m)
}

// GetMessageFromString convert strign to Message
func GetMessageFromString(data []byte) (Message, error) {
	var message Message

	return message, nil

}

// ExecMessage function to implement each kind of message types
func ExecMessage(m *Message, c *Client) {
	fmt.Printf("Execute message type: %d on client: %s", m.Type, c.UserID)
	switch m.Type {
	case Test:
		fmt.Println("Entro a test")
		c.Send <- Message{
			Type: Open,
			Data: "regresoooo!!",
		}
	case Open:
		fmt.Println("Entro a open")
	default:
		fmt.Println("Invalid operation type")
	}
}
