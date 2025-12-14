package topics

import (
  "fmt"
  "net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Create topic")
}

func List(w http.ResponseWriter, r *http.Request) {
  fmt.Println("List topics")
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Update topic name")
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Delete topic")
}
