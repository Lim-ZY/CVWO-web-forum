'use server'
import { redirect } from "next/navigation";
import { PostRequest, Post } from "@/types/models";
import { ApiResponse } from "@/types/api";

export async function CreatePost(topicID: string, formData: FormData) {
  const postSubject = formData.get('postSubject')?.toString() || "";
  const postContent = formData.get('postContent')?.toString() || "";
  const p: PostRequest = {
    name: postSubject.trim(),
    created_by: "mike",
    content: postContent.trim(),
  }
  const response = await fetch(`http://localhost:8000/t/${topicID}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(p),
  })
  if (response.ok) {
    const result: ApiResponse<Post> = await response.json();
    const id: number = result.payload.data.id;
    redirect(`/t/${topicID}/${id}`)
  }
}
