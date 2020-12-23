import NoAuthHeader from "../components/noAuthHeader"

export default function Home() {

  return (
    <div className="container">
      <NoAuthHeader />
      <p className="title">Privacy</p>
      <p className="text">
        There are no trackers on this website. Nothing tracks your behavior, meaning no analytics.
        When you sign up for this website, only your username and password are kept in a database (your password is hashed). Game data like user balance is of course also kept. When you sign in, a session cookie is kept in your browser.
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