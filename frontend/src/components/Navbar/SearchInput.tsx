import React from "react";
import Form from "next/form"

export default function SearchInput() {
  async function Search(formData: FormData) { 
    'use server'
    const query: string = formData.get('query');
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
    <Form action="/search" className="flex flex-grow mx-10 shadow-lg rounded-full bg-paper border-black border">
      <input name="query" placeholder="Search Letters" className="input"/>
    </Form>
  )
}
