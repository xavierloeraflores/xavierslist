const Sandbox: React.FC = () => {
  const env = process.env.NODE_ENV;
  if (env === "production") {
    return <></>;
  }
  return (
    <div className="sandbox">
      <h1>React Sandbox</h1>
    </div>
  );
};

export default Sandbox;
