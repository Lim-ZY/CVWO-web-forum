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
    <>
      <div className="flex min-h-screen justify-center py-6 px-0">
        <div className="flex flex-col w-11/12 items-center">
          <div className="mb-4 text-center text-ink dark:text-dark-ink">
            <strong><h1 className="text-4xl"> Topics </h1></strong>
            <h2 className="italic text-md"> Pick a mailbox to read letters! </h2>
          </div>

          <div className="flex grid lg:grid-cols-3 sm:grid-cols-2 gap-6">
            {topics.map((t: Topic) => (<TopicCard key={t.id} topic={t} />))}
          </div>
          <a className="buttonSolid mt-4" href="/t/create">
            Create Topic
          </a>
        </div>
      </div>
    </>
  );
}
