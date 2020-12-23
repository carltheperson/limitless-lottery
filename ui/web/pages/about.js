import NoAuthHeader from "../components/noAuthHeader"

export default function Home() {

  return (
    <div className="container">
      <NoAuthHeader />
      <p className="title">About</p>
      <p className="text">
        Limitless Lottery is a free to play and open source game. It allows you to buy lottery tickets with a fictitious currency. There is no real objective in the game, but a fun way to play is to try and win more than you spend. You have different options in how you play. You can  for example choose between multiple kinds of lottery tickets, each with their own odds and price. Have fun!
      </p>

      <style jsx> {
        `
        .container {
          text-align: center;
        }
        
        .title {
          font-size: 6em;
        }
        
        .text {
          width: 80%;
          margin: auto;
          font-size: 3.5em;
          font-weight: 400;
        }
      `
      }
      </style>
    </div>
  )
}