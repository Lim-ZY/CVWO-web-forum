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

func (h *TopicHandler) Create(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
  var body struct {
    Name            string  `json:"name"`
    CreatedBy       string  `json:"created_by"`
    Description     string  `json:"description"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return nil, err
  }

  sg, err := time.LoadLocation("Asia/Singapore")
  if err != nil {
    fmt.Printf("Failed to load timezone: %v\n", err)
    return nil, err
  }
  creationTime := time.Now().In(sg)

  topic := model.Topic{
    Name:         body.Name,
    CreationTime: creationTime,
    CreatedBy:    body.CreatedBy,
    Description:  body.Description,
  }

  err = h.DB.InsertTopic(topic)
  if err != nil {
    fmt.Printf("Failed to insert: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return nil, err
  }

  res, err := json.Marshal(topic)
  if err != nil {
    fmt.Println("Failed to marshal: %w\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return nil, err
  }

  w.WriteHeader(http.StatusCreated)
  return &api.Response{
    Payload: api.Payload{
      Data: res,
    },
  }, nil
}

func (h *TopicHandler) List(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
  topics, err := h.DB.GetTopics()
  if err != nil {
    return nil, fmt.Errorf("Failed to get topics: %w", err)
  }

  jsonData, err := json.Marshal(topics)
  if err != nil {
    return nil, fmt.Errorf("Failed to marshal: %w", err)
  }

  return &api.Response{
    Payload: api.Payload{
      Data: jsonData,
    },
  }, nil
}

func (h *TopicHandler) UpdateByID(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
  idParam := chi.URLParam(r, "id")
  id, _ := strconv.Atoi(idParam)

  var body struct {
    Description     string  `json:"description"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return nil, err
  }

  topic := model.Topic{
    ID:           id,
    Description:  body.Description,
  }

  updatedTopic, err := h.DB.UpdateTopic(topic)
  if err != nil {
    fmt.Printf("Failed to insert: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return nil, err
  }

  res, err := json.Marshal(updatedTopic)
  if err != nil {
    fmt.Println("Failed to marshal: %w\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return nil, err
  }

  return &api.Response{
    Payload: api.Payload{
      Data: res,
    },
  }, nil
}

func (h *TopicHandler) DeleteByID(w http.ResponseWriter, r *http.Request) error {
  idParam := chi.URLParam(r, "id")
  id, _ := strconv.Atoi(idParam)
  err := h.DB.DeleteTopicByID(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return err
  }
  return nil
}
