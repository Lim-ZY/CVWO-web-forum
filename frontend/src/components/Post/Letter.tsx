import React from "react";
import PostCard from "@/components/PostCard/PostCard";

export default function Letter({ post }: { post: Post }) {
  const date: Date = new Date(post.creation_time);

  return (
    <>
      <div className="flex flex-col bg-paper dark:bg-dark-paper min-w-full max-h-full mb-2 p-4 border-2 border-amber-900/20 border-dashed shadow-2xl">
        {/* Header */}
        <div className="flex flex-col dark:bg-white text-lg w-full max-h-full p-2">
          <h1 className="italic mb-4">From: {post.created_by}</h1>
          <h1>{date.toLocaleDateString('en-US', {
              day: "numeric",
              month: "long",
              year: "numeric",
              weekday: "long"
            })
          }</h1>
          <h1 className="mb-4">{date.toLocaleTimeString('en-US')}</h1>
          <strong><h1 className="text-3xl pb-2 border-b-2 border-desk">{post.name}</h1></strong>
        </div>
        {/* Body */}
        <div className="flex flex-col p-2 text-lg">
          <h2 className="mb-2"> Dear Community, </h2>
          <h2 className="mb-10"> {post.content} </h2>
          <h2 className="mb-2 italic"> Yours Sincerely, </h2>
          <h2 className="mb-2 italic"> {post.created_by} </h2>
        </div>
      </div>
      <div className="mb-8">
        <h1 className="buttonOutline">Votes: {post.votes} </h1>
      </div>
    </>
  );
};
