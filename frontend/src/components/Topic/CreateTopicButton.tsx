"use client"
import { useState, useEffect } from "react";
import Form from "next/form";
import { CreateTopic } from "@/actions/CreateTopic";

export default function CreateTopicButton() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [count, setCount] = useState(0);
  const maxChars = 500;

  const countChars = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const text = e.target.value;
    setCount(text.length);
  };

  return (
    <>
      <button onClick={() => setIsModalOpen(true)} className="buttonOutline">
        Create Mailbox
      </button>

      {useEffect(() => {
        if (isModalOpen) {
          document.body.style.overflow = 'hidden';
        } else {
          document.body.style.overflow = 'unset';
        }
        return () => { document.body.style.overflow = 'unset'; };
      }, [isModalOpen])}

      {isModalOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-8">
          <div className="absolute inset-0 bg-slate-900/40 backdrop-blur-sm"></div>
          <Form action={CreateTopic} className="relative z-10 flex flex-col w-full max-w-2xl max-h-120 bg-paper shadow-xl border-4 border-accent rounded-xl text-start p-4">
            <button onClick={() => setIsModalOpen(false)} className="absolute top-0 right-2 p-2 mt-2 flex items-center justify-center w-8 h-8 font-bold shadow-lg rounded-lg hover:bg-accent">
              X
            </button>
            {/* Header */}
            <div className="mb-4">
              <h1 className="text-xl font-bold">
                New Mailbox
              </h1>
            </div>
            {/* Body */}
            <div className="mb-4">
              <label> Topic Name: </label>
              <input 
                type="text" 
                name="topicName"
                placeholder="e.g., NUS, Butterflies, Weather Phenomenon..." 
                required
                className="border border-ink shadow-sm rounded-lg mb-4 px-2">
              </input>
              <h2> Topic Description: </h2>
              <textarea 
                rows={5}
                name="topicDescription"
                placeholder="Take your time... Slowly describe the purpose of this mailbox. What kind of letters would you want to receive?" 
                className="border border-ink shadow-sm rounded-lg w-full p-2" 
                onChange={countChars} 
                required
                maxLength={maxChars}>
              </textarea>
              <p className="text-sm justify-self-end">Characters Left: {count} / {maxChars}</p>
            </div>
            {/* End */}
            <div className="mt-8">
              <button className="buttonOutline justify-self-end"> Create! </button>
            </div>
          </Form>
        </div>
      )}
    </>
  );
}
