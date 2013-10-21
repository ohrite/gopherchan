package gopherchan_test

import (
  "net"
  "net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  . "github.com/ohrite/gopherchan/gopherchan"

  "github.com/ohrite/gopher"
  "github.com/ant0ine/go-urlrouter"
)

type simpleMockHandler struct {
  Handled bool
}

func (handler *simpleMockHandler) Handle(c net.Conn, r *gopher.Request, p map[string]string) {
  handler.Handled = true
}

type captureMockHandler struct {
  Conn net.Conn
  Request *gopher.Request
  Params map[string]string
}

func (handler *captureMockHandler) Handle(c net.Conn, r *gopher.Request, p map[string]string) {
  handler.Conn = c
  handler.Request = r
  handler.Params = p
}

var _ = Describe("Router", func() {
  var (
    router *Router
  )

  BeforeEach(func() {
    router = NewRouter()
  })

  Describe("NewRouter", func() {
    It("creates a new router", func() {
      Expect(router).NotTo(BeNil())
    })
  })

  Describe("AddRoute", func() {
    var (
      routes []urlrouter.Route
      handler simpleMockHandler
    )

    BeforeEach(func(){
      routes = router.AddRoute("/tacos/:kind", &handler)
    })

    It("adds a route", func() {
      Expect(routes).To(HaveLen(1))
    })

    It("sets the path on the route", func(){
      Expect(routes[0].PathExp).To(Equal("/tacos/:kind"))
    })

    It("sets the handler on the route", func(){
      routes[0].Dest.(GopherRouter).Handle(nil, nil, nil)
      Expect(handler.Handled).To(BeTrue())
    })
  })

  Describe("HandleRequest", func() {
    var (
      handler captureMockHandler
      request *gopher.Request
      urlObj *url.URL
      fakeConn net.Conn
    )

    BeforeEach(func() {
      fakeConn = new(testConnection)
      router.AddRoute("/tacos/:kind", &handler)
    })

    Context("when the url fails to match", func() {
      BeforeEach(func() {
        urlObj = &url.URL{Path: ""}
        request = gopher.NewRequest(urlObj, "")
      })

      It("does not set the connection object", func(){
        router.HandleRequest(fakeConn, request)
        Expect(handler.Conn).To(BeNil())
      })
    })

    Context("when the url matches", func() {
      BeforeEach(func() {
        urlObj = &url.URL{Path: "/tacos/delicious"}
        request = gopher.NewRequest(urlObj, "")
      })

      It("sets the connection object", func(){
        router.HandleRequest(fakeConn, request)
        Expect(handler.Conn).To(Equal(fakeConn))
      })

      It("sets the request object", func(){
        router.HandleRequest(fakeConn, request)
        Expect(handler.Request).To(Equal(request))
      })

      It("sets the parameters", func(){
        fakeParams := map[string]string{"kind":"delicious"}
        router.HandleRequest(fakeConn, request)
        Expect(handler.Params).To(Equal(fakeParams))
      })
    })
  })
})
