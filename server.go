package main

import (
  "os"
  "log"
  "net"

  "github.com/ohrite/gopher"
  "github.com/ohrite/gopherchan/gopherchan"
  "github.com/ohrite/gopherchan/controllers"
)

func main() {
  port := os.Getenv("PORT")
  router := gopherchan.NewRouter()
  server := gopherchan.NewServer(port, func(connection net.Conn, request *gopher.Request) {
    log.Printf("Handling request: %v", request.URL.String())
    router.HandleRequest(connection, request)
  })
  router.AddRoute("/new", &controllers.PostCreateController{Host: "localhost", Port: server.Port})
  router.AddRoute("/:id", &controllers.PostGetController{Host: "localhost", Port: server.Port})
  router.AddRoute("/", &controllers.PostIndexController{Host: "localhost", Port: server.Port})
  router.AddRoute("", &controllers.PostIndexController{Host: "localhost", Port: server.Port})
  router.AddRoute("*", &controllers.ErrorController{Host: "localhost", Port: server.Port})
  log.Printf("Serving at %v", server.URL())
  log.Fatal(server.Serve())
}
