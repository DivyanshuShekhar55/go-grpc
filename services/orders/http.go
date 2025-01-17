package main

// this is the place where we can handle the http services / routes ... just fopr example purposes
// we are using just the net/http lib but you might want to use chi (or smthing else) too

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	srv := &httpServer{addr: addr}
	return srv
}
