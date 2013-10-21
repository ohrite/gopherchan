package main

import (
  "os"
  "log"

  "github.com/ohrite/gopherchan/gopherchan"
  "github.com/ohrite/gopherchan/controllers"
)

func main() {
  port := os.Getenv("PORT")
  router := gopherchan.NewRouter()
  server := gopherchan.NewServer(port, router.HandleRequest)
  log.Printf("Serving at %v", server.URL())
  router.AddRoute("/new", &controllers.PostCreateController{Host: "localhost", Port: server.Port})
  router.AddRoute("/:id", &controllers.PostGetController{Host: "localhost", Port: server.Port})
  router.AddRoute("/", &controllers.PostIndexController{Host: "localhost", Port: server.Port})
  router.AddRoute("", &controllers.PostIndexController{Host: "localhost", Port: server.Port})
  router.AddRoute("*", &controllers.ErrorController{Host: "localhost", Port: server.Port})
  log.Fatal(server.Serve())
}
