import NoAuthHeader from "../components/noAuthHeader"
import NoAuthIndex from "./_noAuthIndex";
import LoggedIn from "../components/loggedIn"
import { useState, useEffect } from "react"
import axios from "axios";

export default function Home() {
  const [username, setUsername] = useState()
  const [getLoggedIn, setLoggedIn] = LoggedIn()



  useEffect(() => {

      if (getLoggedIn() == undefined || getLoggedIn() == true) {
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
        (username != undefined) && <p>
          you are logged in  : 

          
          {username}
        </p>
      }

      {
        (username == undefined) &&

        <div>
          <NoAuthHeader />

          <NoAuthIndex />
        </ div>

      }

    </div>
  )
}
