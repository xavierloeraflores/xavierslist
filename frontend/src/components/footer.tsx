import Link from "next/link";

const Footer: React.FC = () => {
  return (
    <footer className="flex flex-col items-center justify-center ">
      <Link href="/">
        <span className="font-bold">Xavier&apos;s List</span>
      </Link>
      <span className="text-center">The Internet&apos;s Classifieds</span>
      <span className="text-center">Created by Xavier Loera Flores</span>
      <span className="text-center">
        Built with Go, Fiber, Next.js, TRPC, and Tailwind{" "}
      </span>
    </footer>
  );
};

export default Footer;
