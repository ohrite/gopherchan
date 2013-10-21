package controllers_test

import (
  "github.com/ohrite/gopher"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/ohrite/gopherchan/models"
  . "github.com/ohrite/gopherchan/controllers"
)

var _ = Describe("Controller Helpers", func(){
  var (
    line *gopher.ResponseLine
  )

  BeforeEach(func() {
    line = gopher.NewCommentResponseLine("crunchy")
  })

  Describe("BuildResponse", func() {
    var (
      fakeConn *testConnection
    )

    BeforeEach(func() {
      fakeConn = new(testConnection)
    })

    It("writes the string to the response", func() {
      BuildResponse(line).WriteResponse(fakeConn)
      Expect(fakeConn.WriteBuf.String()).To(ContainSubstring("crunchy"))
    })
  })

  Describe("PostPath", func() {
    It("prints a path to the post", func() {
      post := Post{Id:6}
      Expect(PostPath(post)).To(Equal("/6"))
    })
  })
})
