package sockets

type MemberPrivate struct {
	Member *Member
	Desc string
}

type Pool struct {
	Register chan *Member
	Unregister chan *Member
	Broadcast chan MemberMessage
	Private chan MemberPrivate
	AssignName chan *Member
	Members []*Member
}

func CreatePool() *Pool {
	return &Pool{
		Register: make(chan *Member),
		Unregister: make(chan *Member),
		Broadcast: make(chan MemberMessage),
		Members: make([]*Member, 0),
		Private: make(chan MemberPrivate),
		AssignName: make(chan *Member),
	}
}
