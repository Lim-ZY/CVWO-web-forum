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
    <>
      <div className="flex border-4 border-red-700 py-6 px-0 justify-center">
        <div className="flex border-4 border-orange-600 w-11/12">
          {/* Left Content */}
          <div className="flex flex-col min-h-screen w-9/12 items-center text-ink dark:text-dark-ink mr-6 border-4 border-yellow-600">
            <div className="w-full">
              <strong><h1 className="mb-4 text-xl"> Letters in {posts[0].topic_name} </h1></strong>
              {posts.map((p: Post) => p.id >= 0 ? <PostCard post={p} key={p.id}/> : null)}
              <a className="buttonSolid mt-4" href={`/t/${topicID}/create`}>
                Create Post
              </a>
            </div>
          </div>
          {/* Right Content */}
          <div className="flex flex-col flex-grow border-4 border-green-600">
          </div>
        </div>
      </div>
    </>
  );
}
