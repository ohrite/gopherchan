package controllers

import (
  "net"
  "strconv"
  "github.com/ohrite/gopher"
  . "github.com/ohrite/gopherchan/models"
)

type PostGetController struct {
  Host string
  Port string
}

func NewPostGetController(host string, port string) (*PostGetController) {
  return &PostGetController{
    Host: host,
    Port: port,
  }
}

func (controller *PostGetController) Handle(conn net.Conn, request *gopher.Request, params map[string]string) {
  postId, _ := strconv.ParseInt(params["id"], 10, 64)
  post := FindPost(int(postId))
  conn.Write([]byte(post.Body))
}
