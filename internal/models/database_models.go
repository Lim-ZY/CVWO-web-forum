package models

import "time"

type Topic struct {
  ID              int     `json:"id"`
  Name            string  `json:"name"`
  CreationTime    time.Time  `json:"creation_time"`
  CreatedBy       string  `json:"created_by"`
  Description     string  `json:"description"`
}

type Post struct {
  ID              int     `json:"id"`
  Name            string  `json:"name"`
  CreationTime    time.Time  `json:"creation_time"`
  CreatedBy       string  `json:"created_by"`
  RelatedTopicID  int     `json:"related_topic_id"`
  Content         string  `json:"content"`
  Votes           int     `json:"votes"`
}

type Comment struct {
  ID              int     `json:"id"`
  CreationTime    time.Time  `json:"creation_time"`
  CreatedBy       string  `json:"created_by"`
  RelatedPostID   int     `json:"related_post_id"`
  Content         string  `json:"content"`
  Votes           int     `json:"votes"`
}
