package models

import "time"

type User struct {
	ID              int     `json:"id"`
	Username        string  `json:"username"`
  CreationTime    time.Time  `json:"creation_time"`
  LastActive      time.Time  `json:"last_active"`
}

type Topic struct {
  ID              int     `json:"id"`
  Name            string  `json:"name"`
  CreationTime    time.Time  `json:"creation_time"`
  CreatedBy       string  `json:"created_by"`
  Description     string  `json:"description"`
  PostCount       int     `json:"post_count"`
}

type Post struct {
  ID              int     `json:"id"`
  Name            string  `json:"name"`
  CreationTime    time.Time  `json:"creation_time"`
  CreatedBy       string  `json:"created_by"`
  RelatedTopicID  int     `json:"related_topic_id"`
  Content         string  `json:"content"`
  Votes           int     `json:"votes"`
  TopicName       string  `json:"topic_name"`
}

type Comment struct {
  ID              int     `json:"id"`
  CreationTime    time.Time  `json:"creation_time"`
  CreatedBy       string  `json:"created_by"`
  RelatedPostID   int     `json:"related_post_id"`
  Content         string  `json:"content"`
  Votes           int     `json:"votes"`
}
