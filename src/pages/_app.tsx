import { useEffect } from "react";
// import Head from "next/head";
import { type AppType } from "next/dist/shared/lib/utils";
import "~/styles/globals.css";

declare global {
  interface Window {
    Telegram: {
      WebApp: {
        ready: () => void;
        MainButton: {
          setText: (text: string) => void;
          show: () => void;
        };
        onEvent: (event: string, callback: () => void) => void;
        sendData: (data: string) => void;
      };
    };
  }
}

const MyApp: AppType = ({ Component, pageProps }) => {
  useEffect(() => {
    // Ensure this runs only on the client-side
    if (
      typeof window !== "undefined" &&
      window.Telegram &&
      window.Telegram.WebApp
    ) {
      const tg = window.Telegram.WebApp;

      // Ensure tg is defined before calling methods
      if (tg) {
        tg.ready();
        tg.MainButton.setText("Send to Bot");
        tg.MainButton.show();

        // Handle the button click event
        tg.onEvent("mainButtonClicked", () => {
          // Send data back to the bot
          tg.sendData(JSON.stringify({ message: "Hello from Web App!" }));
        });
      }
    } else {
      console.warn("Telegram WebApp is not available.");
    }
  }, []);

  return (
    <>
      {/* <Head>
        <title>React Hackaton </title>
        <meta
          name="description"
          content="Hackaton web app written with React"
        />
        <link rel="icon" href="/favicon.ico" />
        <meta name="theme-color" content="#0A0" />
        <link rel="manifest" href="/app.webmanifest" />
      </Head> */}
      <Component {...pageProps} />
    </>
  );
};

export default MyApp;
