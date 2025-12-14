package posts

import (
  "fmt"
  "net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Create post in topic")
}

func List(w http.ResponseWriter, r *http.Request) {
  fmt.Println("List posts in topic")
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Update posts")
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Delete posts")
}
