package postview

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

type ViewHandler struct {
  DB *pg.Database
}

func (h *ViewHandler) Get(w http.ResponseWriter, r *http.Request) {
  postIDParam := chi.URLParam(r, "postID")
  postID, _ := strconv.Atoi(postIDParam)

  post, err := h.DB.GetPostByID(postID)
  if err != nil {
    fmt.Println("Failed to get posts: %w", err)
    return
  }
  comments, err := h.DB.GetComments(postID)
  if err != nil {
    fmt.Println("Failed to get comments: %w", err)
    return
  }

  view := &model.View{
    PostID: post.ID,
    PostName: post.Name,
    PostCreationTime: post.CreationTime,
    PostCreatedBy: post.CreatedBy,
    PostRelatedTopicID: post.RelatedTopicID,
    PostContent: post.Content,
    PostVotes: post.Votes,
    Comments: comments,
  }
  jsonData, err := json.Marshal(view)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: jsonData,
    },
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *ViewHandler) AddComment(w http.ResponseWriter, r *http.Request) {
  postIDParam := chi.URLParam(r, "postID")
  postID, _ := strconv.Atoi(postIDParam)

  var body struct {
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

  comment := model.Comment{
    CreationTime:   creationTime,
    CreatedBy:      body.CreatedBy,
    RelatedPostID:  postID,
    Content:        body.Content,
  }

  err = h.DB.InsertComment(comment)
  if err != nil {
    fmt.Println("Failed to insert comment: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(comment)
  if err != nil {
    fmt.Println("Failed to marshal: %w\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
  }
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *ViewHandler) UpdateCommentByID(w http.ResponseWriter, r *http.Request) {
  commentIDParam := chi.URLParam(r, "commentID")
  commentID, _ := strconv.Atoi(commentIDParam)

  var body struct {
    Content     string  `json:"content"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  comment := model.Comment{
    ID:           commentID,
    Content:      body.Content,
  }

  updatedComment, err := h.DB.UpdateComment(comment)
  if err != nil {
    fmt.Println("Failed to update: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(updatedComment)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *ViewHandler) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
  commentIDParam := chi.URLParam(r, "commentID")
  commentID, _ := strconv.Atoi(commentIDParam)
  err := h.DB.DeleteCommentByID(commentID)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
}
