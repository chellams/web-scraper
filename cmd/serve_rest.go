package main

type RestServer struct {
}

func NewRestServer() WebServer {
	return RestServer{}
}

func (r RestServer) Serve(address string) {
	panic("implement me")
}
