package controllers_test

import (
  "strings"
  "net/url"
  "github.com/ohrite/gopher"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/ohrite/gopherchan/models"
  . "github.com/ohrite/gopherchan/controllers"
)

var _ = Describe("PostCreateController", func(){
  var (
    controller *PostCreateController
  )

  BeforeEach(func() {
    controller = NewPostCreateController("host", "port")
  })

  Describe("NewPostCreateController", func() {
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
      params = map[string]string{"body": "omg shoes"}
    })

    It("creates a new post with the given body", func() {
      controller.Handle(fakeConn, request, params)
      Expect(FindPost(PostId).Body).To(Equal("omg shoes"))
    })

    It("prints a link to the new post path", func() {
      controller.Handle(fakeConn, request, params)
      Expect(fakeConn.WriteBuf.String()).To(ContainSubstring("New Post"))
      Expect(fakeConn.WriteBuf.String()).To(ContainSubstring("/new"))
    })

    It("prints a file link to a post", func() {
      controller.Handle(fakeConn, request, params)
      Expect(fakeConn.WriteBuf.String()).To(ContainSubstring("omg shoes"))
      Expect(fakeConn.WriteBuf.String()).To(ContainSubstring("/1"))
    })

    Context("and another post exists", func() {
      BeforeEach(func() {
        post = &Post{Body: "Crunchy face fungus"}
        post.Save()
      })

      It("sorts them in reverse post order", func() {
        controller.Handle(fakeConn, request, params)
        output := fakeConn.WriteBuf.String()
        output = strings.Replace(output, "\r\n", " ", -1)
        Expect(output).To(MatchRegexp("/2.*/1"))
      })
    })
  })
})
