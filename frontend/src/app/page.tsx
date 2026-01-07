import Image from "next/image";

export default function Home() {
  return (
    <div className="flex min-h-screen items-center justify-center bg-desk dark:bg-dark-desk mt-4">
      <main className="flex min-h-screen w-full max-w-3xl flex-col items-center justify-between py-32 px-16 bg-white dark:bg-dark-paper sm:items-start">
        <Image
          className="dark:invert"
          src="/next.svg"
          alt="Next.js logo"
          width={100}
          height={20}
          priority
        />
        <div className="flex flex-col items-center gap-6 text-center sm:items-start sm:text-left">
          <h1 className="max-w-xs text-3xl font-semibold leading-10 tracking-tight text-ink dark:text-dark-ink">
            Login Page
          </h1>
          <p className="max-w-md text-lg leading-8 text-ink dark:text-dark-ink">
            Login Stuff
          </p>
        </div>
        <div className="flex flex-col gap-4 text-base font-medium sm:flex-row">
          <a className="buttonOutline" href="/">
            Login
          </a>
        </div>
      </main>
    </div>
  );
}
