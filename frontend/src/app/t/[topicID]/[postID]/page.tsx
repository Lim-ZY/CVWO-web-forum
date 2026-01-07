import React from "react";
import TopicCard from "@/components/Topic/TopicCard";
import CommentCard from "@/components/PostView/CommentCard";
import Letter from "@/components/PostView/Letter";
import { Topic, Post } from "@/types/models";
import { ApiResponse } from "@/types/api";

interface ApiMainResponse {
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
  const topicResponse = await fetch(`http://localhost:8000/t${topicID}`);
  const topicResult: ApiResponse<Topic> = await topicResponse.json();
  const topic: Topic = topicResult.payload.data;
  const response = await fetch(`http://localhost:8000/t/${topicID}/${postID}`);
  const result: ApiMainResponse = await response.json();
  const data = result.payload.data;
  const comments = data.comments;
  const post: Post = {
    id: data.post_id,
    name: data.post_name,
    creation_time: data.post_creation_time,
    created_by: data.post_created_by,
    related_topic_id: data.post_related_topic_id,
    content: data.post_content,
    votes: data.post_votes,
  }

  return (
    <>
    <div className="flex px-6 py-6 justify-center">
      <div className="flex w-11/12">
        {/* Left Content */}
        <div className="flex flex-col min-h-screen w-9/12 mr-6">
          {/* Letter */}
          <Letter post={post} />
          {/* Replies */}
          <div className="flex flex-col flex-grow min-w-full">
            <div className="mb-4">
              <h1 className="text-xl"> Replies </h1>
            </div>
            <div className="ml-10">
              {comments.map((c: Comment) => ( <CommentCard comment={c} key={c.id} /> ))}
            </div>
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
