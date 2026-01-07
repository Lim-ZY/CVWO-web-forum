import React from "react";
import { Comment } from "@/types/models";

export default function CommentCard({comment}: { comment: Comment }) {
  const date: Date = new Date(comment.creation_time);

  return (
    <div className="flex flex-col bg-paper dark:bg-dark-paper min-w-full max-h-60 mb-2 p-4 border-2 border-amber-900/20 border-dashed shadow-2xl text-lg">
      {/* Header */}
      <div className="flex italic mb-4 border-b-2 border-desk justify-between">
        <h2 className=""> From: {comment.created_by} </h2>
        <h2>{date.toLocaleDateString('en-US', {
            day: "numeric",
            month: "long",
            year: "numeric",
            weekday: "long"
          })
        }</h2>
      </div>
      {/* Body */}
      <div className="flex flex-col text-lg">
        <p> {comment.content} </p>
      </div>
      <p> Votes: {comment.votes} </p>
    </div>
  );
};
