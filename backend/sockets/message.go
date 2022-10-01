package sockets

type MemberMessage struct {
	Member *Member `json:"member"`
	Type int `json:"type"`
	Body string `json:"body"`
}