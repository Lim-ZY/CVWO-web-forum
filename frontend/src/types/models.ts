export interface User {
  id: number;
  username: string;
  creation_time: string;
  last_active: string;
}

export interface UserRequest {
  username: string;
}

export interface Topic {
  id: number;
  name: string;
  creation_time: string;
  created_by: string;
  description: string;
  post_count: number;
}

export interface TopicRequest {
  name: string;
  created_by: string;
  description: string;
}

export interface Post {
  id: number;
  name: string;
  creation_time: string;
  created_by: string;
  related_topic_id: number;
  content: string;
  votes: number;
  topic_name: string;
}

export interface PostRequest {
  name: string;
  created_by: string;
  content: string;
}

export interface Comment {
  id: number;
  creation_time: string;
  created_by: string;
  related_post_id: number;
  content: string;
  votes: number;
}
