package gopherchan

import (
  "net"
  "net/url"

  "github.com/ohrite/gopher"
)

type Server struct {
  Port string
  Dispatcher func(net.Conn, *gopher.Request)

  gopherServer *gopher.Server
}

func NewServer(port string, dispatcher func(net.Conn, *gopher.Request)) *Server {
  if port == "" {
    port = "70"
  }

  return &Server{
    Port: port,
    Dispatcher: dispatcher,
  }
}

func (server *Server) URL() *url.URL {
  return server.GopherServer().URL()
}

func (server *Server) Address() string {
  return ":" + server.Port
}

func (server *Server) GopherServer() *gopher.Server {
  if server.gopherServer == nil {
    server.gopherServer = gopher.NewServer(server.Address())
  }
  return server.gopherServer;
}

func (server *Server) Serve() (err error) {
  if err == nil {
    err = server.GopherServer().ListenAndServe(server.Dispatcher)
  }

  return err
}
