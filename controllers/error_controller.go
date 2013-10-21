package controllers

import (
  "net"
  "github.com/ohrite/gopher"
)

type ErrorController struct {
  Host string
  Port string
}

func NewErrorController(host string, port string) (*ErrorController) {
  return &ErrorController{
    Host: host,
    Port: port,
  }
}

func (controller *ErrorController) Handle(conn net.Conn, request *gopher.Request, params map[string]string) {
  BuildResponse(
    gopher.NewCommentResponseLine("Could not find: " + params["path"]),
    gopher.NewCommentResponseLine(""),
    gopher.NewDirectoryResponseLine("Index", "/", controller.Host, controller.Port),
  ).WriteResponse(conn)
}
