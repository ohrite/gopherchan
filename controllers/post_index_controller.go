package controllers

import (
  "net"
  "github.com/ohrite/gopher"
)

type PostIndexController struct {
  Host string
  Port string
}

func NewPostIndexController(host string, port string) (*PostIndexController) {
  return &PostIndexController{
    Host: host,
    Port: port,
  }
}

func (controller *PostIndexController) Handle(conn net.Conn, request *gopher.Request, params map[string]string) {
  response := BuildResponse(
    gopher.NewPromptResponseLine("New Post", "/new", controller.Host, controller.Port),
    gopher.NewCommentResponseLine(""),
  )
  AddPostResponseLines(response, controller.Host, controller.Port)
  response.WriteResponse(conn)
}
