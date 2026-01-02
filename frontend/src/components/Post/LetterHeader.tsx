import React from "react";
import PostCard from "@/components/PostCard/PostCard";

export default function Header({ post }: { post: Post }) {
  const date: Date = new Date(post.creation_time);

  return (
    <div className="flex flex-col bg-paper dark:bg-dark-paper min-w-full h-50 mb-8">
      <div className="flex flex-col mt-4 bg-white border-rounded rounded-2xl dark:bg-white text-lg w-120 h-30 p-2">
        <strong><h1>{post.name}</h1></strong>
        <h1 className="italic">From: {post.created_by}</h1>
        <h1>{date.toLocaleDateString('en-US', {
            day: "numeric",
            month: "long",
            year: "numeric",
            weekday: "long"
          })
        }</h1>
        <h1>{date.toLocaleTimeString('en-US')}</h1>
      </div>
    </div>
  );
};
