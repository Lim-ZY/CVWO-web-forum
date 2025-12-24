import React from "react";
import PostCard from "@/components/PostCard/PostCard";

interface ApiResponse {
  payload: {
    data: Post[];
  };
  messages: string[];
  errorCode: number;
}

export default async function Posts({params}: {params: Promise<{ topicID: string }>}) {
  const { topicID } = await params;
  const response = await fetch(`http://localhost:8000/t/${topicID}`);
  const result : ApiResponse = await response.json();
  const posts = result.payload.data;

  return (
    <div className="flex min-h-screen flex-col flex-grow items-center text-ink dark:text-dark-ink">
      <strong><h1 className="mb-4 text-xl"> Topic {topicID} Posts </h1></strong>
      <ul>
        {posts.map((p: Post) => (
          <li key={p.id}>
            <PostCard post={p} />
          </li>
        ))}
      </ul>
    </div>
  );
}
