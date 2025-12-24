import React from "react";
import Link from "next/link";

interface Topic {
  id: number;
  name: string;
  creation_time: string;
  created_by: string;
  description: string;
}

export default function TopicCard({topic}: { topic: Topic }) {
  return (
    <div className="card">
      <strong><Link href={`/t/${topic.id}`}> {topic.name} </Link></strong>
      <p> Description: {topic.description} </p>
      <p> Created on: {topic.creation_time} </p>
      <p> Created by: {topic.created_by} </p>
    </div>
  );
};
