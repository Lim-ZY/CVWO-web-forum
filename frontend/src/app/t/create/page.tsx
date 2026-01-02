import React from "react";
import Form from "next/form";
import { redirect } from "next/navigation";
import TopicCard from "@/components/TopicCard/TopicCard";

interface ApiResponse {
  payload: {
    data: Topic[];
  };
  messages: string[];
  errorCode: number;
}

export default function Topics() {
  async function CreateTopic(formData: FormData) {
    'use server'
    const t: Topic = {
      name: formData.get('topic'),
      created_by: "mike",
      description: formData.get('description'),
    }
    const response = await fetch("http://localhost:8000/t", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(t),
    })
    if (response.ok) {
      const result: ApiResponse = await response.json();
      const id: number = result.payload.data.id;
      redirect(`/t/${id}`)
    }
  }

  return (
    <>
    <div className="flex flex-col flex-grow min-h-screen items-center">
      <strong><h1 className="mb-4 text-xl text-ink"> Create Topic </h1></strong>
      <Form action={CreateTopic} className="text-ink">
        <h2>Topic:</h2>
        <p><input type="text" name="topic" defaultValue="Topic Name"/></p>
        <h2>Description:</h2>
        <p><textarea rows="5" cols="20" name="description" defaultValue="Description of your topic"></textarea></p>
        <p><input type="submit" value="Create" className="buttonSolid"/></p>
      </Form>
    </div>
    </>
  );
}
