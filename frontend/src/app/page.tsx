"use client"
import Image from "next/image";
import Form from "next/form";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { UserRequest } from "@/types/models";

export default function Home() {
  const [loginFail, setLoginFail] = useState(false);
  const router = useRouter();

  const Login = async (formData: FormData) => {
    const username: string = formData.get('username')?.toString().trim() || "";
    const req: UserRequest = {
      username: username,
    };
    const response = await fetch("http://localhost:8000/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: 'include',
      body: JSON.stringify(req),
    });
    if (response.ok) {
      router.push(`/t`);
    } else {
      setLoginFail(true);
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-desk dark:bg-dark-desk mt-4">
      <main className="flex flex-col min-h-screen w-full max-w-3xl items-center justify-between py-32 px-16 bg-paper dark:bg-dark-paper sm:items-start">
        <h1 className="max-w-xs text-3xl font-semibold leading-10 tracking-tight text-ink dark:text-dark-ink">
          Login
        </h1>
        <Form action={Login} className="flex flex-col items-center gap-30 text-center sm:items-start sm:text-left">
          <div className="max-w-md">
            <label className="text-lg leading-8">Username: </label>
            <input 
              type="text" 
              name="username" 
              className="border border-ink shadow-sm rounded-lg mb-4 px-2"
              required>
            </input>
            <div className="text-sm mb-4">
              <p>Username is required to access Letters.</p>
              <p>New to Letters? Create a username.</p>
              <p>Returning? Use the same username.</p>
            </div>
            <button type="submit" className="buttonOutline text-md">Login</button>
          </div>
        </Form>
        {
          loginFail ? <p className="text-red-500">Login failed. Please try again</p> : <p></p>
        }
      </main>
    </div>
  );
}
