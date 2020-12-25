import UserForm from "../components/userForm"
import NoAuthHeader from "../components/noAuthHeader"
import axios from "axios"
import { useState } from "react"
import Router from 'next/router'
import LoggedIn from "../components/loggedIn"

export default function Home() {
  const [getLoggedIn, setLoggedIn] = LoggedIn()
  const [errorText, setErrorText] = useState()

  const signUp = (data) => {
    axios.post(process.env.NEXT_PUBLIC_API_URL + "/api/signup", data, { withCredentials: true }).then((res) => {
      document.cookie = res.data;
      setErrorText("")
      Router.push("/")
      setLoggedIn(true)
    }).catch((err) => {
      console.log(err)
      setErrorText(err.response.data.Errors[0].Message)
    })
  }

  return (
    <div className="container">
      <NoAuthHeader />

      <UserForm onSubmit={signUp} buttonText="Sign Up" />

      <div container="error-text">
        <p>
          {errorText}

        </p>
      </div>


      <style jsx> {
        `
        .error-text {
          background-color: white;
        }
        
        p {
          text-align: center;
          font-size: 4em;
          margin-top: -10px;
          color: darkred;

        }
      `
      }
      </style>
    </div>

  )
}