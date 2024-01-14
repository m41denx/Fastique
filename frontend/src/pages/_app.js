import '@/styles/globals.css'
import {useEffect, useState} from "react";

export default function App({ Component, pageProps }) {
  const [token, setToken] = useState("")
  useEffect(() => {
    const ltoken = localStorage.getItem("token") || "."
    ltoken&&setToken(ltoken)
  }, [])
  return <Component {...pageProps} token={token} setToken={setToken} />
}
