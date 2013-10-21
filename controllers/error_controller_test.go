package controllers_test

import (
  "net/url"
  "github.com/ohrite/gopher"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/ohrite/gopherchan/controllers"
)

var _ = Describe("ErrorController", func(){
  var (
    controller *ErrorController
  )

  BeforeEach(func() {
    controller = NewErrorController("host", "port")
  })

  Describe("NewErrorController", func() {
    It("creates a new error controller", func() {
      Expect(controller).ToNot(BeNil())
    })

    It("assigns the host", func() {
      Expect(controller.Host).To(Equal("host"))
    })

    It("assigns the port", func() {
      Expect(controller.Port).To(Equal("port"))
    })
  })

  Describe("Handle", func() {
    var (
      request *gopher.Request
      urlObj *url.URL
      fakeConn *testConnection
      params map[string]string
    )

    BeforeEach(func() {
      fakeConn = new(testConnection)
      urlObj = &url.URL{Path: "/ham"}
      request = gopher.NewRequest(urlObj, "")
      params = map[string]string{"path": urlObj.Path}
    })

    It("reports a path error", func() {
      controller.Handle(fakeConn, request, params)
      Expect(fakeConn.WriteBuf.String()).To(ContainSubstring("Could not find: /ham"))
    })
  })
})
