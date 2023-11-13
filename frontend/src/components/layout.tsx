import Header from "./header";
import Footer from "./footer";

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="Page">
      <Header />
      {children}
      <Footer />
    </div>
  );
}
