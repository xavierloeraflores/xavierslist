import { type AppType } from "next/app";
import Layout from "~/components/layout";

import { api } from "~/utils/api";

import "~/styles/globals.css";

const MyApp: AppType = ({ Component, pageProps }) => {
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  );
};

export default api.withTRPC(MyApp);
