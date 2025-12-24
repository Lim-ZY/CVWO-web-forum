import React from "react";
import Link from "next/link";
import Image from "next/image";

export default function Footer() {
  return (
    <div className="flex justify-between p-4 bg-accent mt-8 dark:bg-dark-accent">
      <p className="text-xs">&copy; Letters. By Ken. All rights reserved.</p>
    </div>
  );
};
