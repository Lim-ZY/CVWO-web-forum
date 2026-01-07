import React from "react";
import PostCard from "@/components/Post/PostCard";
import TopicCard from "@/components/Topic/TopicCard";
import { Topic, Post } from "@/types/models";
import { ApiResponse } from "@/types/api";
import CreatePostButton from "@/components/Post/CreatePostButton";

export default async function Posts({params}: {params: Promise<{ topicID: string }>}) {
  const { topicID } = await params;
  const topicResponse = await fetch(`http://localhost:8000/t${topicID}`);
  const topicResult: ApiResponse<Topic> = await topicResponse.json();
  const topic: Topic = topicResult.payload.data;
  const postsResponse = await fetch(`http://localhost:8000/t/${topicID}`);
  const postsResult: ApiResponse<Post[]> = await postsResponse.json();
  const posts: Post[] = postsResult.payload.data;

  return (
    <>
      <div className="flex py-6 px-0 justify-center">
        <div className="flex w-11/12">
          {/* Left Content */}
          <div className="flex flex-col min-h-screen min-w-9/12 items-center text-ink dark:text-dark-ink mr-6">
            <div className="relative w-full mb-4">
              <div className="absolute top-0 right-0">
                <CreatePostButton />
              </div>
              <strong><h1 className="mb-4 text-xl"> Letters in {posts[0].topic_name} </h1></strong>
            </div>
            <div className="w-full">
              {posts.map((p: Post) => p.id >= 0 ? <PostCard post={p} key={p.id} /> : null)}
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
