package database

import (
  "context"
  "os"
  "fmt"
  "github.com/jackc/pgx/v5/pgxpool"
  model "github.com/lim-zy/CVWO-web-forum/internal/models"
)

type Database struct {
  Ctx context.Context
  Pool *pgxpool.Pool
}

func GetDB(ctx context.Context) (*Database, error) {
  connStr := os.Getenv("DATABASE_URL")
  if connStr == "" {
    return nil, fmt.Errorf("DATABASE_URL env var not set")
  }
  dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
    return nil, fmt.Errorf("Unable to create connection pool: %w\n", err)
	}
  err = dbpool.Ping(ctx)
  if err != nil {
    dbpool.Close()
    return nil, fmt.Errorf("Unable to ping database: %w", err)
  }

  return &Database{Pool: dbpool, Ctx: ctx}, nil
}

func (db *Database) GetTopics() ([]model.Topic, error) {
  q := `SELECT 
          t.id,
          t.name,
          t.creation_time,
          u.username,
          t.description,
          COUNT(p.id) as post_count
        FROM topics t
        INNER JOIN users u ON t.created_by = u.id
        LEFT JOIN posts p ON t.id = p.related_topic_id
        GROUP BY t.id, t.name, t.creation_time, u.username, t.description
        ORDER BY t.creation_time DESC`
  rows, err := db.Pool.Query(db.Ctx, q)
  if err != nil {
    return nil, fmt.Errorf("Error querying topics: %w", err)
  }
  defer rows.Close()
  
  topics := []model.Topic{}
  for rows.Next() {
    var t model.Topic
    err := rows.Scan(&t.ID, &t.Name, &t.CreationTime, &t.CreatedBy, &t.Description, &t.PostCount)
    if err != nil {
      return nil, fmt.Errorf("Error scanning rows: %w", err)
    }
    topics = append(topics, t)
  }

  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("Error iterating rows: %w", err)
  }
  return topics, nil
}

func (db *Database) InsertTopic(topic *model.Topic) error {
  q := `INSERT INTO topics (name, creation_time, created_by, description) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`
  err := db.Pool.QueryRow(db.Ctx, q, topic.Name, topic.CreationTime, topic.CreatedBy, topic.Description).Scan(&topic.ID)
  if err != nil {
    return fmt.Errorf("Error inserting topic: %w", err)
  }
  return nil
}

func (db *Database) UpdateTopic(topic model.Topic) (*model.Topic, error) {
  t, err := db.FindTopicByID(topic.ID)
  if err != nil {
    return nil, fmt.Errorf("No topic found: %w", err)
  }

  q := `UPDATE topics 
        SET 
            description = $1
        WHERE
            id = $2`
  _, err = db.Pool.Exec(db.Ctx, q, topic.Description, topic.ID)
  if err != nil {
    return nil, fmt.Errorf("Error updating topic: %w", err)
  }
  t.Description = topic.Description
  return t, nil
}

func (db *Database) DeleteTopicByID(id int) error {
  q := `DELETE FROM topics WHERE id = $1`
  _, err := db.Pool.Exec(db.Ctx, q, id)
  if err != nil {
    return fmt.Errorf("Error deleting topic: %w", err)
  }
  return nil
}

func (db *Database) FindTopicByID(id int) (*model.Topic, error) {
  q := `SELECT 
          t.id,
          t.name,
          t.creation_time,
          u.username,
          t.description,
          COUNT(p.id) as post_count
        FROM topics t
        INNER JOIN users u ON t.created_by = u.id
        LEFT JOIN posts p ON t.id = p.related_topic_id
        WHERE t.id = $1
        GROUP BY t.id, t.name, t.creation_time, u.username, t.description`
  var t model.Topic
  err := db.Pool.QueryRow(db.Ctx, q, id).Scan(&t.ID, &t.Name, &t.CreationTime, &t.CreatedBy, &t.Description, &t.PostCount)
  if err != nil {
    return nil, fmt.Errorf("No topic found: %w", err)
  }
  return &t, nil
}

func (db *Database) GetPosts(topicID int) ([]model.Post, error) {
  q := `SELECT 
          p.id,
          p.name,
          p.creation_time,
          u.username,
          p.related_topic_id,
          p.content,
          p.votes,
          t.name as topic_name
        FROM posts p
        INNER JOIN users u ON p.created_by = u.id
        LEFT JOIN topics t ON p.related_topic_id = t.id
        WHERE related_topic_id = $1
        ORDER BY p.creation_time DESC`
  rows, err := db.Pool.Query(db.Ctx, q, topicID)
  if err != nil {
    return nil, fmt.Errorf("Error querying posts: %w", err)
  }
  defer rows.Close()
  
  posts := []model.Post{}
  for rows.Next() {
    var p model.Post
    err := rows.Scan(&p.ID, &p.Name, &p.CreationTime, &p.CreatedBy, &p.RelatedTopicID, &p.Content, &p.Votes, &p.TopicName)
    if err != nil {
      return nil, fmt.Errorf("Error scanning rows: %w", err)
    }
    posts = append(posts, p)
  }

  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("Error iterating rows: %w", err)
  }

  if len(posts) == 0 {
    t, _ := db.FindTopicByID(topicID)
    p := model.Post{
      ID: -1,
      TopicName: t.Name,
    }
    posts = append(posts, p)
  }

  return posts, nil
}

func (db *Database) GetPostByID(postID int) (*model.Post, error) {
  q := `SELECT 
          p.id,
          p.name,
          p.creation_time,
          u.username,
          p.related_topic_id,
          p.content,
          p.votes
        FROM posts p
        INNER JOIN users u ON p.created_by = u.id
        WHERE p.id = $1`
  var p model.Post
  err := db.Pool.QueryRow(db.Ctx, q, postID).Scan(&p.ID, &p.Name, &p.CreationTime, &p.CreatedBy, &p.RelatedTopicID, &p.Content, &p.Votes)
  if err != nil {
    return nil, fmt.Errorf("Error querying post: %w", err)
  }
  return &p, nil
}

func (db *Database) InsertPost(post *model.Post) error {
  q := `INSERT INTO posts (name, creation_time, created_by, related_topic_id, content) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id`
  err := db.Pool.QueryRow(db.Ctx, q, post.Name, post.CreationTime, post.CreatedBy, post.RelatedTopicID, post.Content).Scan(&post.ID)
  if err != nil {
    return fmt.Errorf("Error inserting post: %w", err)
  }
  return nil
}

func (db *Database) UpdatePost(post model.Post) (*model.Post, error) {
  p, err := db.FindPostByID(post.ID)
  if err != nil {
    return nil, fmt.Errorf("No post found: %w", err)
  }

  q := `UPDATE posts 
        SET 
            content = $1
        WHERE
            id = $2`
  _, err = db.Pool.Exec(db.Ctx, q, post.Content, post.ID)
  if err != nil {
    return nil, fmt.Errorf("Error updating post: %w", err)
  }
  p.Content = post.Content
  return p, nil
}

func (db *Database) DeletePostByID(postID int) error {
  q := `DELETE FROM posts WHERE id = $1`
  _, err := db.Pool.Exec(db.Ctx, q, postID)
  if err != nil {
    return fmt.Errorf("Error deleting post: %w", err)
  }
  return nil
}

func (db *Database) FindPostByID(postID int) (*model.Post, error) {
  q := `SELECT * FROM posts WHERE id = $1`
  var p model.Post
  err := db.Pool.QueryRow(db.Ctx, q, postID).Scan(&p.ID, &p.Name, &p.CreationTime, &p.CreatedBy, &p.RelatedTopicID, &p.Content, &p.Votes)
  if err != nil {
    return nil, fmt.Errorf("No post found: %w", err)
  }
  return &p, nil
}

func (db *Database) GetComments(postID int) ([]model.Comment, error) {
  q := `SELECT 
          c.id,
          c.creation_time,
          u.username,
          c.related_post_id,
          c.content,
          c.votes
        FROM comments c
        INNER JOIN users u ON c.created_by = u.id
        LEFT JOIN posts p ON c.related_post_id = p.id
        WHERE related_post_id = $1
        ORDER BY p.creation_time DESC`
  rows, err := db.Pool.Query(db.Ctx, q, postID)
  if err != nil {
    return nil, fmt.Errorf("Error querying comments: %w", err)
  }
  defer rows.Close()
  
  comments := []model.Comment{}
  for rows.Next() {
    var c model.Comment
    err := rows.Scan(&c.ID, &c.CreationTime, &c.CreatedBy, &c.RelatedPostID, &c.Content, &c.Votes)
    if err != nil {
      return nil, fmt.Errorf("Error scanning rows: %w", err)
    }
    comments = append(comments, c)
  }

  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("Error iterating rows: %w", err)
  }
  return comments, nil
}

func (db *Database) InsertComment(comment model.Comment) error {
  q := `INSERT INTO comments (creation_time, created_by, related_post_id, content) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`
  err := db.Pool.QueryRow(db.Ctx, q, comment.CreationTime, comment.CreatedBy, comment.RelatedPostID, comment.Content).Scan(&comment.ID)
  if err != nil {
    return fmt.Errorf("Error inserting comment: %w", err)
  }
  return nil
}

func (db *Database) UpdateComment(comment model.Comment) (*model.Comment, error) {
  c, err := db.FindCommentByID(comment.ID)
  if err != nil {
    return nil, fmt.Errorf("No comment found: %w", err)
  }

  q := `UPDATE comments 
        SET 
            content = $1
        WHERE
            id = $2`
  _, err = db.Pool.Exec(db.Ctx, q, comment.Content, comment.ID)
  if err != nil {
    return nil, fmt.Errorf("Error updating comment: %w", err)
  }
  c.Content = comment.Content
  return c, nil
}

func (db *Database) DeleteCommentByID(commentID int) error {
  q := `DELETE FROM comments WHERE id = $1`
  _, err := db.Pool.Exec(db.Ctx, q, commentID)
  if err != nil {
    return fmt.Errorf("Error deleting comment: %w", err)
  }
  return nil
}

func (db *Database) FindCommentByID(commentID int) (*model.Comment, error) {
  q := `SELECT * FROM comments WHERE id = $1`
  var c model.Comment
  err := db.Pool.QueryRow(db.Ctx, q, commentID).Scan(&c.ID, &c.CreationTime, &c.CreatedBy, &c.RelatedPostID, &c.Content, &c.Votes)
  if err != nil {
    return nil, fmt.Errorf("No comment found: %w", err)
  }
  return &c, nil
}
