'use server'
import { redirect } from "next/navigation";
import { TopicRequest, Topic } from "@/types/models";
import { ApiResponse } from "@/types/api";

export async function CreateTopic(formData: FormData) {
  const topicName = formData.get('topicName')?.toString() || "";
  const topicDesc = formData.get('topicDescription')?.toString() || "";
  const t: TopicRequest = {
    name: topicName.trim(),
    created_by: "mike",
    description: topicDesc.trim(),
  }
  const response = await fetch("http://localhost:8000/t", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(t),
  })
  if (response.ok) {
    const result: ApiResponse<Topic> = await response.json();
    const id: number = result.payload.data.id;
    redirect(`/t/${id}`)
  }
}
