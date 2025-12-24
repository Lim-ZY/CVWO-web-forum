import React from "react";
import Link from "next/link";

export default function Navbar() {
  return (
    <div className="flex justify-between p-4 bg-accent mb-8 dark:bg-dark-accent">
      <Link href={"/"} className="font-semibold tracking-wide"> Letters </Link>
      <nav className="flex gap-4">
        <Link href={"/t"}> Topics </Link>
        <Link href={"/"}> Login </Link>
      </nav>
    </div>
  );
};
