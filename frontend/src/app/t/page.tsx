import React from "react";
import TopicCard from "@/components/TopicCard/TopicCard";

interface ApiResponse {
  payload: {
    data: Topic[];
  };
  messages: string[];
  errorCode: number;
}

export default async function Topics() {
  const response = await fetch("http://localhost:8000/t");
  const result : ApiResponse = await response.json();
  const topics = result.payload.data;

  return (
    <div className="flex min-h-screen flex-col flex-grow items-center text-ink dark:text-dark-ink">
      <strong><h1 className="mb-4 text-xl"> Topics </h1></strong>
      <ul>
        {topics.map((t: Topic) => (
          <li key={t.id}>
            <TopicCard topic={t} />
          </li>
        ))}
      </ul>
    </div>
  );
}
