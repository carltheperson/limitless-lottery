import NoAuthIndex from "./_noAuthIndex";
import AuthIndex from "./_authIndex";
import LoggedIn from "../components/loggedIn"
import { useState, useEffect } from "react"
import axios from "axios";

export default function Home() {
  const [username, setUsername] = useState()
  const [showAuth, setShowAuth] = useState(true)
  const [getLoggedIn, setLoggedIn] = LoggedIn()



  useEffect(() => {

    if (getLoggedIn() == undefined || getLoggedIn() == true) {
      setShowAuth(false)
      axios.get(process.env.NEXT_PUBLIC_API_URL + "/api/session-username", { withCredentials: true }).then((res) => {
        setLoggedIn(true)
        setUsername(res.data)
      }).catch(() => {
        setLoggedIn(false)
      })
    }
  }, [])



  return (
    <div className="container">
      {
        (username != undefined) && <div>
          <AuthIndex username={username} />

        </div>
      }

      {
        (username == undefined) &&

        <div>
          <NoAuthIndex />
        </ div>

      }



    </div>
  )
}
