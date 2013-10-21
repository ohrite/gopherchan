package controllers

import (
  "fmt"
  "github.com/ohrite/gopher"
  "github.com/ohrite/gopherchan/models"
)

func BuildResponse(lines... *gopher.ResponseLine) (response *gopher.Response) {
  response = new(gopher.Response)
  for _, line := range lines {
    response.AddResponseLine(line)
  }
  return response
}

func PostPath(post models.Post) string {
  return fmt.Sprintf("/%d", post.Id)
}

func PresentPost(post models.Post, host string, port string) (*gopher.ResponseLine) {
  return gopher.NewFileResponseLine(
    post.Body,
    PostPath(post),
    host,
    port,
  )
}

func AddPostResponseLines(response *gopher.Response, host string, port string) {
  for p := len(models.Posts) - 1; p >= 0; p-- {
    response.AddResponseLine(PresentPost(models.Posts[p], host, port))
  }
}
