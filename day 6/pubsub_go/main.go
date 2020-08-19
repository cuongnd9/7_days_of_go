package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/icrowley/fake"
)

type Pubsub struct {
	mux         sync.RWMutex
	subscribers map[string][]chan string
	closed      bool
}

func CreatePubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subscribers = make(map[string][]chan string)
	return ps
}

func (ps *Pubsub) Subscribe(topic string) chan string {
	ps.mux.Lock()
	defer ps.mux.Unlock()

	ch := make(chan string, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

// TODO: Unsubscribe method for Pubsub.

func (ps *Pubsub) Publish(topic string, message string) {
	ps.mux.Lock()
	defer ps.mux.Unlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subscribers[topic] {
		go func(ch chan string) {
			ch <- message
		}(ch)
	}
}

func (ps *Pubsub) Close() {
	ps.mux.Lock()
	defer ps.mux.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subscriber := range ps.subscribers {
			for _, ch := range subscriber {
				// built-in function closes a channel.
				close(ch)
			}
		}
	}
}

func main() {
	ps := CreatePubsub()
	topics := []string{"topic 1", "topic 2", "topic 3", "topic 4", "topic 5", "topic 6", "topic 7", "topic 8", "topic 9", "topic 10", "topic 11", "topic 12"}

	rand.Seed(time.Now().UnixNano())

	var channels [5]chan string
	for i := range channels {
		topic := topics[rand.Intn(len(topics))]
		channels[i] = ps.Subscribe(topic)
	}

	for _, topic := range topics {
		ps.Publish(topic, fake.FullName())
	}

	for _, channel := range channels {
		fmt.Printf("💅 %v\n", <-channel)
	}
}
