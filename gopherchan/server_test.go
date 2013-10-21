package gopherchan_test

import (
  "reflect"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/ohrite/gopherchan/gopherchan"
)

var _ = Describe("Server", func() {
  var (
    server *Server
  )

  BeforeEach(func(){
    server = NewServer("localham", nil)
  })

  Describe("NewServer()", func() {
    Context("when port provided is empty", func() {
      BeforeEach(func(){
        server = NewServer("", nil)
      })

      It("sets the port to 70", func() {
        Expect(server.Port).To(Equal("70"))
      })
    })

    Context("the port provided is not empty", func() {
      BeforeEach(func(){
        server = NewServer("delicious tacos", nil)
      })

      It("returns the content of PORT", func() {
        Expect(server.Port).To(Equal("delicious tacos"))
      })
    })
  })

  Describe("Address()", func() {
    It("returns the proposed address of the gopher server", func() {
      Expect(server.Address()).To(Equal("localhost:localham"))
    })
  })

  Describe("GopherServer()", func() {
    It("returns a gopher server instance", func() {
      Expect(reflect.TypeOf(server.GopherServer()).String()).To(Equal("*gopher.Server"))
    })

    It("has an address set to the gopherchan address", func() {
      Expect(server.GopherServer().Address).To(Equal("localhost:localham"))
    })
  })
})
