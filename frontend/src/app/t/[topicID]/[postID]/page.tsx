import React from "react";
import PostCard from "@/components/PostCard/PostCard";
import CommentCard from "@/components/CommentCard/CommentCard";
import Header from "@/components/Post/LetterHeader";

interface ApiResponse {
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
  const response = await fetch(`http://localhost:8000/t/${topicID}/${postID}`);
  const result: ApiResponse = await response.json();
  const data = result.payload.data;
  const comments = data.comments;
  const post: PostCard = {
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
        <div className="relative flex min-h-screen w-9/12 items-center mr-6 bg-paper">
          <div className="absolute top-2 flex flex-col items-center pl-8">
            <Header post={post} />
            <div className="flex flex-col flex-grow min-w-full items-start text-ink dark:text-dark-ink bg-paper">
              <section className="flex flex-col mb-4">
                <h1>{post.content}</h1>
                <h1 className="buttonSolid">Votes: {post.votes} </h1>
              </section>
              <section className="flex flex-col items-center mb-4">
              <h1> Comments </h1>
              <ul>
                {comments.map((c: Comment) => (
                  <li key={c.id}>
                    <CommentCard comment={c} />
                  </li>
                ))}
              </ul>
              </section>
            </div>
          </div>
        </div>
        {/* Right Content */}
        <div className="flex flex-col flex-grow bg-accent border-4 border-red">
        </div>
      </div>
    </div>
    </>
  );
}
