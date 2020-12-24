import Router from 'next/router'
import axios from "axios"
import LoggedIn from "./loggedIn"

export default function AuthHeader(props) {
  const [getLoggedIn, setLoggedIn] = LoggedIn()


  const logOut = () => {
    axios.delete(process.env.NEXT_PUBLIC_API_URL + "/api/session", { withCredentials: true })
    setLoggedIn(false)
    Router.reload("/")
  }



  return (
    <div className="container">
      <p className="logo-text">Limitless Lottery</p>
      <p className="balance">Balance: {props.balance}$</p>
      <p onClick={logOut} className="logout">Log out</p>

      <style jsx>{`

      .container {
        background-color: #76C6FF;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 2em;
      }

      .logo-text {
        font-family: 'Montserrat', sans-serif;
        font-weight: bold;
        font-size: 5em;
        margin: 0;
      }

      .balance, .logout {
        font-size: 3em;
      }

      .logout {
        cursor: pointer;
      }

      `}
      </style>

    </div>
  )
}


function deleteCookie(sKey, sPath, sDomain) {
  document.cookie = encodeURIComponent(sKey) +
    "=; expires=Thu, 01 Jan 1970 00:00:00 GMT" +
    (sDomain ? "; domain=" + sDomain : "") +
    (sPath ? "; path=" + sPath : "");
}