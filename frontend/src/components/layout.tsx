type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="Page">
      <div>Header</div>
      {children}
      <div>Footer</div>
    </div>
  );
}
