package gopherchan

import (
  "net"

  "github.com/ohrite/gopher"
  "github.com/ant0ine/go-urlrouter"
)

type Router struct {
  urlRouter *urlrouter.Router
}

type GopherRouter interface {
  Handle(net.Conn, *gopher.Request, map[string]string)
}

func NewRouter() *Router {
  routes := []urlrouter.Route{}
  urlRouter := &urlrouter.Router{Routes: routes}

  return &Router{
    urlRouter: urlRouter,
  }
}

func (router *Router) AddRoute(path string, handler GopherRouter) ([]urlrouter.Route) {
  router.appendRoute(createRoute(path, handler))
  router.urlRouter.Start()
  return router.urlRouter.Routes
}

func createRoute(path string, handler GopherRouter) (*urlrouter.Route) {
  return &urlrouter.Route{
    PathExp: path,
    Dest: handler,
  }
}

func (router *Router) appendRoute(route *urlrouter.Route) {
  router.urlRouter.Routes = append(router.urlRouter.Routes, *route)
}

func (router *Router) HandleRequest(connection net.Conn, request *gopher.Request) {
  route, params := router.urlRouter.FindRouteFromURL(request.URL)
  if route != nil {
    route.Dest.(GopherRouter).Handle(connection, request, params)
  }
}
