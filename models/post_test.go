package models_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "github.com/ohrite/gopherchan/models"
)

var _ = Describe("Post", func() {
  var (
    post *Post
  )

  BeforeEach(func() {
    post = &Post{Body:"toenail torture"}
  })

  Describe("Save()", func() {
    It("saves a post", func() {
      Expect(Posts).To(HaveLen(0))
      post.Save()
      Expect(Posts).To(HaveLen(1))
    })

    Context("when there are more posts than the maximum", func() {
      var (
        lastPost *Post
        previousMaxPosts int
      )

      BeforeEach(func() {
        previousMaxPosts = MaxPosts
        MaxPosts = 1
        lastPost = &Post{Body:"i've never seen the ocean before"}
        lastPost.Save()
      })

      AfterEach(func() {
        MaxPosts = previousMaxPosts
      })

      It("does not change the number of posts", func() {
        Expect(Posts).To(HaveLen(1))
        post.Save()
        Expect(Posts).To(HaveLen(1))
      })

      It("removes the last post", func() {
        post.Save()
        Expect(Posts[0]).To(Equal(*post))
      })
    })
  })

  Describe("FindPost()", func() {
    Context("when the post does not exist", func() {
      It("returns nil", func() {
        Expect(FindPost(-1)).To(BeNil())
      })
    })

    Context("when the post exists", func() {
      BeforeEach(func() {
        post.Save()
      })

      It("returns the post", func() {
        Expect(FindPost(post.Id)).To(Equal(post))
      })
    })
  })
})
