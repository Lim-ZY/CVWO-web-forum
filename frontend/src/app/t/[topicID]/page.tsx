import React from "react";
import PostCard from "@/components/PostCard/PostCard";
import TopicCard from "@/components/TopicCard/TopicCard";

interface ApiTopicResponse {
  payload: {
    data: Topic;
  };
  messages: string[];
  errorCode: number;
}

interface ApiResponse {
  payload: {
    data: Post[];
  };
  messages: string[];
  errorCode: number;
}

export default async function Posts({params}: {params: Promise<{ topicID: string }>}) {
  const { topicID } = await params;
  const topicResponse = await fetch(`http://localhost:8000/t${topicID}`);
  const topicResult : ApiTopicResponse = await topicResponse.json();
  const topic : Topic = topicResult.payload.data;
  const postsResponse = await fetch(`http://localhost:8000/t/${topicID}`);
  const postsResult : ApiResponse = await postsResponse.json();
  const posts : Post[] = postsResult.payload.data;

  return (
    <>
      <div className="flex py-6 px-0 justify-center">
        <div className="flex w-11/12">
          {/* Left Content */}
          <div className="flex flex-col min-h-screen w-9/12 items-center text-ink dark:text-dark-ink mr-6">
            <div className="w-full">
              <strong><h1 className="mb-4 text-xl"> Letters in {posts[0].topic_name} </h1></strong>
              {posts.map((p: Post) => p.id >= 0 ? <PostCard post={p} key={p.id}/> : null)}
              <a className="buttonSolid mt-4" href={`/t/${topicID}/create`}>
                Write a Letter
              </a>
            </div>
          </div>
          {/* Right Content */}
          <div className="flex flex-col flex-grow">
            <TopicCard topic={topic}/>
          </div>
        </div>
      </div>
    </>
  );
}
