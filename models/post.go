package models


type Post struct {
  Id int
  Body string
}

var PostId int = 0
var MaxPosts int = 10
var Posts []Post = []Post{}

func (post *Post) Save() {
  if len(Posts) > MaxPosts - 1 {
    Posts = Posts[1:MaxPosts]
  }
  PostId++
  post.Id = PostId
  Posts = append(Posts, *post)
}

func FindPost(id int) (post *Post) {
  for _, post := range Posts {
    if post.Id == id {
      return &post
    }
  }

  return nil
}
