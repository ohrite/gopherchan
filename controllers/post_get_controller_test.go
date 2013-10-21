package controllers_test

import (
  "strconv"
  "net/url"
  "github.com/ohrite/gopher"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/ohrite/gopherchan/models"
  . "github.com/ohrite/gopherchan/controllers"
)

var _ = Describe("PostGetController", func(){
  var (
    controller *PostGetController
  )

  BeforeEach(func() {
    ClearPosts()
    controller = NewPostGetController("host", "port")
  })

  Describe("NewPostGetController", func() {
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
      post *Post
    )

    BeforeEach(func() {
      fakeConn = new(testConnection)
      urlObj = &url.URL{}
      request = gopher.NewRequest(urlObj, "")
      post = &Post{Body: "tasty cardboard"}
      post.Save()
      params = map[string]string{"id":strconv.Itoa(post.Id)}
    })

    It("prints a link to the new post path", func() {
      controller.Handle(fakeConn, request, params)
      Expect(fakeConn.WriteBuf.String()).To(Equal("tasty cardboard"))
    })
  })
})
