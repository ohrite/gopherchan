package controllers

import (
  "net"
  "github.com/ohrite/gopher"
  . "github.com/ohrite/gopherchan/models"
)

type PostCreateController struct {
  Host string
  Port string
}

func NewPostCreateController(host string, port string) (*PostCreateController) {
  return &PostCreateController{
    Host: host,
    Port: port,
  }
}

func (controller *PostCreateController) Handle(conn net.Conn, request *gopher.Request, params map[string]string) {
  post := Post{Body:request.Body}
  post.Save()

  response := BuildResponse(
    gopher.NewPromptResponseLine("New Post", "/new", controller.Host, controller.Port),
    gopher.NewCommentResponseLine(""),
  )
  AddPostResponseLines(response, controller.Host, controller.Port)
  response.WriteResponse(conn)
}
