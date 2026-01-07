import React from "react";
import Link from "next/link";
import { Topic } from "@/types/models";

export default function TopicCard({ topic }: { topic: Topic }) {
  return (
    <Link href={`/t/${topic.id}`} className="relative rounded-lg shadow-xl border-4 border-stone-600 overflow-hidden">
      {/* Top Cover */}
      <div className="bg-gradient-to-r from-stone-400 to-stone-500 px-4 py-3">
        {/* Mail slot */}
        <div className="w-full h-1 bg-black rounded-full mb-3 shadow-inner"></div>
        
        {/* Title */}
        <h2 className="text-2xl text-white font-bold text-center leading-tight truncate">
          {topic.name}
        </h2>
        
        {/* Screws */}
        <div className="flex justify-between mt-2">
          <div className="w-2 h-2 bg-gray-800 rounded-full border border-gray-600"></div>
          <div className="w-2 h-2 bg-gray-800 rounded-full border border-gray-600"></div>
        </div>
      </div>

      {/* Bottom Plate */}
      <div className="bg-gradient-to-r from-stone-600 to-stone-700 px-4 py-3">
        {/* Information Plate */}
        <div className="bg-paper border-2 border-accent rounded px-2 py-2 shadow-inner min-h-[60px]">
          <p className="text-amber-900 font-serif text-md line-clamp-[2] leading-tight h-10 text-center mb-1">
            {topic.description}
          </p>
          <div className="text-center pt-1 border-t border-amber-800/30 mt-1">
            <div className="text-amber-800 text-sm font-mono">
              <p> Created on: {new Date(topic.creation_time).toLocaleString('default')} </p>
              <p> Created by: {topic.created_by} </p>
              <p> Letters: {topic.post_count} </p>
            </div>
          </div>
        </div>
      </div>
    </Link>
  );
};
