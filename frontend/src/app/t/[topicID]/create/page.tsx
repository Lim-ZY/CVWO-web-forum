import React from "react";
import Form from "next/form";
import { redirect } from "next/navigation";
import PostCard from "@/components/PostCard/PostCard";

interface ApiResponse {
  payload: {
    data: Post[];
  };
  messages: string[];
  errorCode: number;
}

export default function Post({params}: {params: Promise<{ topicID: string }>}) {
  async function CreatePost(formData: FormData) {
    'use server'
    const { topicID } = await params;
    const t: Topic = {
      name: formData.get('title'),
      created_by: "mike",
      description: formData.get('content'),
    }
    const response = await fetch(`http://localhost:8000/t/${topicID}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(t),
    })
    if (response.ok) {
      const result: ApiResponse = await response.json();
      const postID: number = result.payload.data.id;
      redirect(`/t/${topicID}/${postID}`)
    }
  }

  return (
    <>
    <div className="flex flex-col flex-grow min-h-screen items-center">
      <strong><h1 className="mb-4 text-xl text-ink"> Create Post </h1></strong>
      <Form action={CreatePost} className="text-ink">
        <h2>New Post:</h2>
        <p><input type="text" name="title" defaultValue="Post Name"/></p>
        <h2>Content:</h2>
        <p><textarea rows="5" cols="20" name="content" defaultValue="Content of your post"></textarea></p>
        <p><input type="submit" value="Create" className="buttonSolid"/></p>
      </Form>
    </div>
    </>
  );
}
