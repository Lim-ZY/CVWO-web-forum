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
  rows, err := db.Pool.Query(db.Ctx, "SELECT * FROM topics")
  if err != nil {
    return nil, fmt.Errorf("Error querying topics: %w", err)
  }
  defer rows.Close()
  
  topics := []model.Topic{}
  for rows.Next() {
    var t model.Topic
    err := rows.Scan(&t.ID, &t.Name, &t.CreationTime, &t.CreatedBy, &t.Description)
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

func (db *Database) InsertTopic(topic model.Topic) error {
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
  q := `SELECT * FROM topics WHERE id = $1`
  var t model.Topic
  err := db.Pool.QueryRow(db.Ctx, q, id).Scan(&t.ID, &t.Name, &t.CreationTime, &t.CreatedBy, &t.Description)
  if err != nil {
    return nil, fmt.Errorf("No topic found: %w", err)
  }
  return &t, nil
}

func (db *Database) GetPosts(topicID int) ([]model.Post, error) {
  q := `SELECT * FROM posts WHERE related_topic_id = $1`
  rows, err := db.Pool.Query(db.Ctx, q, topicID)
  if err != nil {
    return nil, fmt.Errorf("Error querying topics: %w", err)
  }
  defer rows.Close()
  
  posts := []model.Post{}
  for rows.Next() {
    var p model.Post
    err := rows.Scan(&p.ID, &p.Name, &p.CreationTime, &p.CreatedBy, &p.RelatedTopicID, &p.Content, &p.Votes)
    if err != nil {
      return nil, fmt.Errorf("Error scanning rows: %w", err)
    }
    posts = append(posts, p)
  }

  if err = rows.Err(); err != nil {
    return nil, fmt.Errorf("Error iterating rows: %w", err)
  }
  return posts, nil
}

func (db *Database) InsertPost(post model.Post) error {
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
