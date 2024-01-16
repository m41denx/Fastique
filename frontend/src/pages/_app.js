import '@/styles/globals.css'
import {useEffect, useState} from "react";
import {ConfigProvider} from "antd";

export default function App({ Component, pageProps }) {
  const [token, setToken] = useState("")
  useEffect(() => {
    const ltoken = localStorage.getItem("token") || "."
    ltoken&&setToken(ltoken)
  }, [])
  return <ConfigProvider theme={{

  }}>
    <Component {...pageProps} token={token} setToken={setToken} />
  </ConfigProvider>
}
