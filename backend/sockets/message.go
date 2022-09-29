package sockets

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}