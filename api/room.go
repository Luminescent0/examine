package api

type Room struct {
	Id      string
	Clients []*Client
}
