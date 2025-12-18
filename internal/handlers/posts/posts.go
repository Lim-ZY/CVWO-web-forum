package posts

import (
  "encoding/json"
  "fmt"
  "net/http"
  "time"
  "strconv"

  "github.com/go-chi/chi/v5"
  "github.com/lim-zy/CVWO-web-forum/internal/api"
  pg "github.com/lim-zy/CVWO-web-forum/internal/database"
  model "github.com/lim-zy/CVWO-web-forum/internal/models"
)

type PostHandler struct {
  DB *pg.Database
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
  topicIDParam := chi.URLParam(r, "id")
  topicID, _ := strconv.Atoi(topicIDParam)

  var body struct {
    Name            string  `json:"name"`
    CreatedBy       string  `json:"created_by"`
    Content         string  `json:"content"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  sg, err := time.LoadLocation("Asia/Singapore")
  if err != nil {
    fmt.Printf("Failed to load timezone: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  creationTime := time.Now().In(sg)

  post := model.Post{
    Name:           body.Name,
    CreationTime:   creationTime,
    CreatedBy:      body.CreatedBy,
    RelatedTopicID: topicID,
    Content:        body.Content,
  }

  err = h.DB.InsertPost(post)
  if err != nil {
    fmt.Println("Failed to insert post: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(post)
  if err != nil {
    fmt.Println("Failed to marshal: %w\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
    Messages: []string{fmt.Sprintf("Create post %s successful", post.Name)},
  }
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
  topicIDParam := chi.URLParam(r, "id")
  topicID, _ := strconv.Atoi(topicIDParam)

  posts, err := h.DB.GetPosts(topicID)
  if err != nil {
    fmt.Println("Failed to get posts: %w", err)
    return
  }

  jsonData, err := json.Marshal(posts)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: jsonData,
    },
    Messages: []string{"List posts successful"},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
  postIDParam := chi.URLParam(r, "postID")
  postID, _ := strconv.Atoi(postIDParam)

  var body struct {
    Content     string  `json:"content"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  post := model.Post{
    ID:           postID,
    Content:      body.Content,
  }

  updatedPost, err := h.DB.UpdatePost(post)
  if err != nil {
    fmt.Println("Failed to update: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(updatedPost)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
    Messages: []string{fmt.Sprintf("Update post %s successful", updatedPost.Name)},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
  postIDParam := chi.URLParam(r, "postID")
  postID, _ := strconv.Atoi(postIDParam)
  err := h.DB.DeletePostByID(postID)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  response := &api.Response{
    Messages: []string{fmt.Sprintf("Delete post %d successful", postID)},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}
