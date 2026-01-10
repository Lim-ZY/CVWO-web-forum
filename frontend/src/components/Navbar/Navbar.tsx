// import Link from "next/link";
// import SearchInput from "./SearchInput";

export default function Navbar() {
  return (
    <nav className="flex justify-between py-4 px-12 bg-accent dark:bg-dark-accent text-lg">
      <h1 className="font-semibold tracking-wide"> Letters </h1>
      {/* <SearchInput /> */}
      <div className="flex gap-4">
        {/* <Link href={"/t"}> Topics </Link> */}
        {/* <Link href={"/"}> Login </Link> */}
      </div>
    </nav>
  );
};
