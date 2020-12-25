import AuthHeader from "../components/authHeader"
import { useState, useEffect } from "react"
import axios from "axios"
import Button from "../components/button"

export default function AuthIndex(props) {
  const [balance, setBalance] = useState()
  const [odds, setOdds] = useState([])
  const [picked, setPicked] = useState([true, false, false])

  useEffect(() => {
    axios.get(process.env.NEXT_PUBLIC_API_URL + "/api/balance", { withCredentials: true }).then((res) => {
      setBalance(res.data)
    })

    axios.get(process.env.NEXT_PUBLIC_API_URL + "/api/ticket-odds").then((res) => {
      setOdds(res.data.Odds.map((odd => odd.split('\n').map(str => <p>{str}</p>))))
    })

  }, [])

  const pickedText = (i) => {
    if (picked[i] == true) {
      return "Picked"
    }
    return "Pick"
  }

  const getButtonClass = (i) => {
    if (picked[i] == true) {
      return "btn"
    } else {
      return "btn not-in-focus"
    }
  }

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
            <div className="odds">
              <div className={getButtonClass(0)}>
                <Button onClick={() => setPicked([true, false, false])}>{pickedText(0)}</Button>
              </div>
              {odds[0]}
            </div>
          </div>
          <div>
            <div className="image" style={{ backgroundImage: 'url("/images/gol.png")', width: "321px", height: "401px" }} ></div>
            <div className="odds">
              <div className={getButtonClass(1)}>
                <Button onClick={() => setPicked([false, true, false])} > {pickedText(1)}</Button>
              </div>
              {odds[1]}
            </div>
          </div>
          <div>
            <div className="image" style={{ backgroundImage: 'url("/images/ins.png")', width: "321px", height: "401px" }} ></div>
            <div className="odds">
              <div className={getButtonClass(2)}>
                <Button onClick={() => setPicked([false, false, true])} > {pickedText(2)}</Button>
              </div>
              {odds[2]}
            </div>
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

        .odds {
          text-align: center;
          font-size: 2em;
        }

        .btn {
          margin-top: 15px;
          font-size: 8px !important;
        }
        .not-in-focus {
          opacity: 0.75;
        }
          `
      }
      </style>
    </div>
  )

}