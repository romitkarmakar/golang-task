package sockets

type MemberError struct {
	Member *Member
	Error string
}

type Pool struct {
	Register chan *Member
	Unregister chan *Member
	Broadcast chan Message
	Error chan MemberError
	AssignName chan *Member
	Members []*Member
}

func CreatePool() *Pool {
	return &Pool{
		Register: make(chan *Member),
		Unregister: make(chan *Member),
		Broadcast: make(chan Message),
		Members: make([]*Member, 0),
		Error: make(chan MemberError),
		AssignName: make(chan *Member),
	}
}
