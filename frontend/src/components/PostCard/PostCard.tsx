import React from "react";
import Link from "next/link";

interface Post {
  id: number;
  name: string;
  creation_time: string;
  created_by: string;
  related_topic_id: number;
  content: string;
  votes: number;
}

export default function PostCard({post}: { post: Post }) {
  return (
    <div className="card">
      <strong><Link href={`/t/${post.related_topic_id}/${post.id}`}> {post.name} </Link></strong>
      <p> Content: {post.content} </p>
      <p> Created on: {post.creation_time} </p>
      <p> Created by: {post.created_by} </p>
      <p> Votes: {post.votes} </p>
    </div>
  );
};
