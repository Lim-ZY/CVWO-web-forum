import React from "react";
import PostCard from "@/components/PostCard/PostCard";
import CommentCard from "@/components/CommentCard/CommentCard";

interface ApiResponse {
  payload: {
    data: {
      post_id: number;
      post_name: string;
      post_creation_time: string;
      post_created_by: string;
      post_related_topic_id: number;
      post_content: string;
      post_votes: number;
      comments: Comment[];
    }
  };
  messages: string[];
  errorCode: number;
}

export default async function PostView({params}: {params: Promise<{ topicID: string, postID: string }>}) {
  const { topicID, postID } = await params;
  const response = await fetch(`http://localhost:8000/t/${topicID}/${postID}`);
  const result: ApiResponse = await response.json();
  const data = result.payload.data;
  const comments = data.comments;
  const post: PostCard = {
    id: data.post_id,
    name: data.post_name,
    creation_time: data.post_creation_time,
    created_by: data.post_created_by,
    related_topic_id: data.post_related_topic_id,
    content: data.post_content,
    votes: data.post_votes,
  }

  return (
    <div className="flex min-h-screen flex-col flex-grow items-center text-ink dark:text-dark-ink">
      <strong><h1 className="mb-4 text-xl"> Topic {topicID} Post {postID} </h1></strong>
      <section className="flex flex-col items-center mb-4">
        <h1> Post </h1>
        <PostCard post={post} />
      </section>
      <section className="flex flex-col items-center mb-4">
      <h1> Comments </h1>
      <ul>
        {comments.map((c: Comment) => (
          <li key={c.id}>
            <CommentCard comment={c} />
          </li>
        ))}
      </ul>
      </section>
    </div>
  );
}
