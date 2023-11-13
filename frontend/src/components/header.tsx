import Link from "next/link";

const Header: React.FC = () => {
  return (
    <header className="flex flex-col items-center justify-center">
      <Link href="/">
        <h1 className="text-6xl font-bold">Xavier&apos;s List</h1>
      </Link>
      <h2 className="text-center text-2xl font-bold">
        The Internet&apos;s Classifieds
      </h2>
      <h3 className="text-center">Created by Xavier Loera Flores</h3>
    </header>
  );
};

export default Header;
