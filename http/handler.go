package http

type Handler interface {
	OnRequest(*Context)
}
