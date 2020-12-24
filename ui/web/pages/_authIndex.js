import AuthHeader from "../components/authHeader"
import { useState, useEffect } from "react"
import axios from "axios"

export default function AuthIndex(props) {
  const [balance, setBalance] = useState()

  useEffect(() => {
    axios.get(process.env.NEXT_PUBLIC_API_URL + "/api/balance", { withCredentials: true }).then((res) => {
      setBalance(res.data)
    })

  }, [])

  return (
    <div>
      <AuthHeader balance={balance} />
      <div>
        <p className="title">Play mode</p>

        <p className="playing-as">Playing as: {props.username}</p>

        <p className="sub-title">Pick ticket</p>


        <div className="tickets">
          <div>
            <div className="image" style={{ backgroundImage: 'url("/images/scr.png")', width: "321px", height: "401px" }} ></div>
          </div>
          <div>
            <div className="image" style={{ backgroundImage: 'url("/images/gol.png")', width: "321px", height: "401px" }} ></div>
          </div>
          <div>
            <div className="image" style={{ backgroundImage: 'url("/images/ins.png")', width: "321px", height: "401px" }} ></div>
          </div>
        </div>

      </div>
      <style jsx> {
        `
        .title {
          font-size: 7em;
          text-align: center;
          margin-bottom: 20px;
        }

        .playing-as {
          font-size: 3em;
          text-align: center;
          font-weight: 500;
          text-decoration: underline;
          margin-bottom: 50px;
        }

        .sub-title {
          font-size: 5em;
          text-align: center;
        }
        
        .tickets {
          display: flex;
          justify-content: space-between;
          margin: auto;
          padding: 0 15%;
        }
        
        .image {
          -webkit-box-shadow: 0px 0px 15px 5px rgba(0,0,0,0.45); 
          box-shadow: 0px 0px 15px 5px rgba(0,0,0,0.45);
        }

          `
      }
      </style>
    </div>
  )

}