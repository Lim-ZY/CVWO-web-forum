import React from "react";

interface Comment {
  id: number;
  creation_time: string;
  created_by: string;
  related_post_id: number;
  content: string;
  votes: number;
}

export default function CommentCard({comment}: { comment: Comment }) {
  return (
    <div className="card">
      <p> Content: {comment.content} </p>
      <p> Created on: {comment.creation_time} </p>
      <p> Created by: {comment.created_by} </p>
      <p> Votes: {comment.votes} </p>
    </div>
  );
};
