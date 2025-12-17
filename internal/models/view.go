package models

import "time"

type View struct {
  PostID              int     `json:"post_id"`
  PostName            string  `json:"post_name"`
  PostCreationTime    time.Time  `json:"post_creation_time"`
  PostCreatedBy       string  `json:"post_created_by"`
  PostRelatedTopicID  int     `json:"post_related_topic_id"`
  PostContent         string  `json:"post_content"`
  PostVotes           int     `json:"post_votes"`
  Comments            []Comment  `json:"comments"`
}
