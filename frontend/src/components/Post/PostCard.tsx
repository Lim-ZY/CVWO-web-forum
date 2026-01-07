import React from "react";
import Link from "next/link";
import { Post } from "@/types/models";

export default function PostCard({post}: { post: Post }) {
  const date: Date = new Date(post.creation_time);

  return (
    <Link href={`/t/${post.related_topic_id}/${post.id}`}>
      {/* Envelope */}
      <div className="relative bg-paper p-6 shadow-lg mb-4 w-full h-70">
        {/* Top Border */}
        <div className="flex absolute top-0 left-2 right-2 h-4">
          {[...Array(28)].map((_, i) => (
            <div key={`top-${i}`} 
             className={`flex-1 ${i % 4 === 0 ? 'bg-red-600' : i % 4 === 1 ? 'bg-paper' : i % 4 === 2 ? 'bg-blue-700' : 'bg-paper'} -skew-x-45`}>
            </div>
          ))}
        </div>
        {/* Bottom Border */}
        <div className="flex absolute bottom-0 left-2 right-2 h-4">
          {[...Array(28)].map((_, i) => (
            <div key={`bottom-${i}`} 
             className={`flex-1 ${i % 4 === 0 ? 'bg-red-600' : i % 4 === 1 ? 'bg-paper' : i % 4 === 2 ? 'bg-blue-700' : 'bg-paper'} -skew-x-45`}>
            </div>
          ))}
        </div>
        {/* Left Border */}
        <div className="flex flex-col absolute top-2 left-0 bottom-2 w-4">
          {[...Array(10)].map((_, i) => (
            <div key={`left-${i}`} 
             className={`flex-1 ${i % 4 === 0 ? 'bg-red-600' : i % 4 === 1 ? 'bg-paper' : i % 4 === 2 ? 'bg-blue-700' : 'bg-paper'} -skew-y-45`}>
            </div>
          ))}
        </div>
        {/* Right Border */}
        <div className="flex flex-col absolute top-2 right-0 bottom-2 w-4">
          {[...Array(10)].map((_, i) => (
            <div key={`right-${i}`} 
             className={`flex-1 ${i % 4 === 0 ? 'bg-red-600' : i % 4 === 1 ? 'bg-paper' : i % 4 === 2 ? 'bg-blue-700' : 'bg-paper'} -skew-y-45`}>
            </div>
          ))}
        </div>

        {/* Date */}
        <div className="flex flex-col absolute top-6 right-6 w-20 h-20 border-2 border-accent rounded-full items-center justify-center text-lg">
          <div>{date.getDate()} {date.toLocaleString('default', {month: 'short' })}</div>
          <div>{date.getFullYear()}</div>
        </div>
        
        {/* Content */}
        <div className="py-2 pl-2 pr-20 overflow-hidden text-xl">
          <strong className="truncate"> {post.name} </strong>
          <h1 className="italic"> From: {post.created_by} </h1>
          <p> Votes: {post.votes} </p>
        </div>
      </div>
    </Link>
  );
};
