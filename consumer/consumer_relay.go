package consumer

import (
	store "github.com/AdityaMayukhSom/ruskin/store"

	"github.com/gorilla/websocket"
)

type ConsumerRelay interface {
	Relay() error
}

type WSConsumerRelay struct {
	topicStore *store.Store

	// Registered clients.
	clients map[*websocket.Conn]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Consumer

	// Unregister requests from clients.
	unregister chan *Consumer
}

func (wscr *WSConsumerRelay) Relay() error {
	// this should relay messages to all the connected consumers
	// before that consume message from

	msg, err := (*wscr.topicStore).Extract(10)
	if err != nil {
		return err
	}

	// sends the message to the corresponding consumer
	for consumer, isValid := range wscr.clients {
		if isValid {
			consumer.WriteMessage(websocket.BinaryMessage, msg)
		}
	}
	return nil
}