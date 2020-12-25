import AuthHeader from "../components/authHeader"
import { useState, useEffect } from "react"
import axios from "axios"
import Button from "../components/button"

export default function AuthIndex(props) {
  const [balance, setBalance] = useState()
  const [odds, setOdds] = useState([])
  const [picked, setPicked] = useState([true, false, false])
  const [checkAmount, setCheckAmount] = useState(true)
  const [amount, setAmount] = useState(1)
  const [results, setResults] = useState()

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

  const getPickedTicketId = () => {
    if (picked[0]) {
      return "scr"
    } else if (picked[1]) {
      return "gol"
    } else if (picked[2]) {
      return "ins"
    }
  }

  const fetchResults = () => {
    if (checkAmount) {
      axios.put(process.env.NEXT_PUBLIC_API_URL + "/api/checkticketamount", { ticketid: getPickedTicketId(), amount: parseInt(amount) }, { withCredentials: true }).then((res) => {
        setBalance(res.data.Balance)
        renderAmountResults(res.data)
      })
    } else {
      axios.put(process.env.NEXT_PUBLIC_API_URL + "/api/checkticketuntilwin", { ticketid: getPickedTicketId() }, { withCredentials: true }).then((res) => {
        setBalance(res.data.Balance)
        renderUntilWinResults(res.data)
      })
    }
  }

  const tableStyles = <style jsx> {
    `
    .table-container {
      background-color: white;
      margin: 0 5% 5% 5%;
      border-radius: 20px;
      padding: 2%;
      margin: auto;
      margin-bottom: 50px;
      width: 70%;
      padding-top: 5px;
    }

    table {
      margin: auto;
      border-collapse: collapse;
    }

    table, th, td {
      border: 2px solid black;
    }

    th {
      padding: 10px 35px;
      font-size: 26px;
    }

    .table-title {
      font-size: 30px;
      text-align: center;
      margin: 10px 0;
      margin-top: 30px;
    }

    .main-tr {
      background-color: lightgrey;
    }
    `}</style>

  const renderAmountResults = (data) => {
    let odds = data.Ct.Wins
    odds = odds.map((item) => {
      return <tr key={item.OutOfOdds}>
        <th>1 / {item.OutOfOdds}</th>
        <th>{item.Prize}</th>
        <th>{item.AmountThatWon}</th>
        <th>{item.TotalWinning}$</th>
      </tr>
    })

    setResults(
      <div className="table-container">
        <p className="table-title">Wins</p>
        <table>
          <tr className="main-tr">
            <th>Odds</th>
            <th>Prize</th>
            <th>Amount that won</th>
            <th>Win for odd</th>
          </tr>
          {odds}
        </table>
        <p className="table-title">Total</p>
        <table>
          <tr className="main-tr">
            <th>Amount won</th>
            <th>Amount deducted</th>
            <th>Profit</th>
          </tr>
          <tr>
            <th>{data.Ct.AmountWonTotal}$</th>
            <th>{data.Ct.AmountDeducted}$</th>
            <th>{data.Ct.AmountWonTotal - data.Ct.AmountDeducted}$</th>
          </tr>
        </table>

        {tableStyles}
      </div>
    )
  }

  const renderUntilWinResults = (data) => {
    setResults(
      <div className="table-container">
        <p className="table-title">Win</p>
        <table>
          <tr className="main-tr">
            <th>Amount bought</th>
            <th>Winning Odds</th>
            <th>Prize</th>
            <th>Profit</th>
          </tr>
          <tr>
            <th>{data.Ct.AmountBought}</th>
            <th>1 / {data.Ct.OutOfOdds}</th>
            <th>{data.Ct.Prize}$</th>
            <th>{data.Ct.Profit}$</th>
          </tr>
        </table>
        {tableStyles}
      </div>
    )
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

        <p className="sub-title">Select mode</p>

        <div className="modes">
          <div>
            <p className="sub-title small">Check amount</p>
            <div className={checkAmount ? "btn" : "btn not-in-focus"}>
              <Button onClick={() => setCheckAmount(true)}>{checkAmount ? "Picked" : "Pick"}</Button>
            </div>
            {checkAmount &&
              <div>
                <input type="number" onChange={(e) => setAmount(e.target.value)} value={amount} min={1} />
              </div>
            }

          </div>

          <div>
            <p className="sub-title small">Check until a win</p>
            <div className={checkAmount ? "btn not-in-focus" : "btn"}>
              <Button onClick={() => setCheckAmount(false)}>{checkAmount ? "Pick" : "Picked"}</Button>
            </div>
          </div>
        </div>

        <div className="buy-btn">
          <Button onClick={fetchResults}>Buy</Button>
        </div>

        {results ? <p className="sub-title">Results</p> : ""}
        {results}

      </div>
      <style jsx> {
        `
        .title {
          font-size: 7em;
          text-align: center;
          margin-bottom: 20px;
          margin-top: 200px;
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
          margin-top: 20px;
          font-size: 8px !important;
        }
        .not-in-focus {
          opacity: 0.75;
          z-index: -5;
        }

        .small {
          font-size: 3em;
          margin-bottom: 10px;
        }

        .modes {
          display: flex;
          justify-content: space-between;
          margin: auto;
          padding: 0 25%;
          text-align: center;
        }

        input {
          margin-top: 10px;
          border: black 1px solid;
          width: 200px;
          font-size: 2em;
          padding: 3px;
          border-radius: 5px;
        }

        .buy-btn {
          text-align: center;
          margin-top: 100px;
          font-size: 1.5em;
        }
          `
      }
      </style>
    </div >
  )

}