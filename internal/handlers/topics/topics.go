package topics

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

type TopicHandler struct {
  DB *pg.Database
}

func (h *TopicHandler) Create(w http.ResponseWriter, r *http.Request) {
  var body struct {
    Name            string  `json:"name"`
    CreatedBy       string  `json:"created_by"`
    Description     string  `json:"description"`
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

  topic := model.Topic{
    Name:         body.Name,
    CreationTime: creationTime,
    CreatedBy:    body.CreatedBy,
    Description:  body.Description,
  }

  err = h.DB.InsertTopic(&topic)
  if err != nil {
    fmt.Printf("Failed to insert: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(topic)
  if err != nil {
    fmt.Println("Failed to marshal: %w\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
    Messages: []string{fmt.Sprintf("Create topic %s successful", topic.Name)},
  }
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *TopicHandler) List(w http.ResponseWriter, r *http.Request) {
  topics, err := h.DB.GetTopics()
  if err != nil {
    fmt.Println("Failed to get topics: %w", err)
    return
  }

  jsonData, err := json.Marshal(topics)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: jsonData,
    },
    Messages: []string{"List topics successful"},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *TopicHandler) GetByID(w http.ResponseWriter, r *http.Request) {
  topicIDParam := chi.URLParam(r, "id")
  topicID, _ := strconv.Atoi(topicIDParam)

  topic, err := h.DB.FindTopicByID(topicID)
  if err != nil {
    fmt.Println("Failed to get topics: %w", err)
    return
  }

  jsonData, err := json.Marshal(topic)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: jsonData,
    },
    Messages: []string{"Get topic successful"},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *TopicHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
  idParam := chi.URLParam(r, "id")
  id, _ := strconv.Atoi(idParam)

  var body struct {
    Description     string  `json:"description"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  topic := model.Topic{
    ID:           id,
    Description:  body.Description,
  }

  updatedTopic, err := h.DB.UpdateTopic(topic)
  if err != nil {
    fmt.Println("Failed to insert: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(updatedTopic)
  if err != nil {
    fmt.Println("Failed to marshal: %w", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
    Messages: []string{fmt.Sprintf("Update topic %s successful", updatedTopic.Name)},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *TopicHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
  idParam := chi.URLParam(r, "id")
  id, _ := strconv.Atoi(idParam)
  err := h.DB.DeleteTopicByID(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  response := &api.Response{
    Messages: []string{fmt.Sprintf("Delete topic %d successful", id)},
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}
