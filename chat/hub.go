package chat

import (
	"fmt"
)

//Hub is a struct that holds all the clients and the messages that are sent to them
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool
	//Unregistered clients.
	unregister chan *Client
	// Register requests from the clients.
	register chan *Client
	// Inbound messages from the clients.
	broadcast chan Message
}

//Message struct to hold message data
type Message struct {
	Type      string  `json:"type"`
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Content   string  `json:"content"`
	ID        string  `json:"id"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		unregister: make(chan *Client),
		register:   make(chan *Client),
		broadcast:  make(chan Message),
	}
}

//Core function to run the hub
func (h *Hub) Run() {
	for {
		select {
		// Register a client.
		case client := <-h.register:
			h.RegisterNewClient(client)
			// Unregister a client.
		case client := <-h.unregister:
			h.RemoveClient(client)
			// Broadcast a message to all clients.
		case message := <-h.broadcast:

			//Check if the message is a type of "message"
			h.HandleMessage(message)

		}
	}
}

//function check if room exists and if not create it and add client to it
func (h *Hub) RegisterNewClient(client *Client) {
	connections := h.clients[client.ID]
	if connections == nil {
		connections = make(map[*Client]bool)
		h.clients[client.ID] = connections
	}
	h.clients[client.ID][client] = true

	fmt.Println("Size of clients: ", len(h.clients[client.ID]))
}

//function to remvoe client from room
func (h *Hub) RemoveClient(client *Client) {
	if _, ok := h.clients[client.ID]; ok {
		delete(h.clients[client.ID], client)
		close(client.send)
		fmt.Println("Removed client")
	}
}

//function to handle message based on type of message
func (h *Hub) HandleMessage(message Message) {

	//Check if the message is a type of "message"
	if message.Type == "message" {
		clients := h.clients[message.ID]
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients[message.ID], client)
			}
		}
	}

	//Check if the message is a type of "notification"
	if message.Type == "notification" {
		fmt.Println("Notification: ", message.Content)
		clients := h.clients[message.Recipient]
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients[message.Recipient], client)
			}
		}
	}

}
