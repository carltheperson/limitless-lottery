import Link from 'next/link'

export default function NoAuthHeader() {

  return (
    <div className="container">
      <p className="logo-text">Limitless Lottery</p>

      <div className="links">
        <Link href="">
          <a>About</a>
        </Link>
        <Link href="">
          <a>Privacy</a>
        </Link>
        <Link href="">
          <a>Login</a>
        </Link>
        <Link href="">
          <a>Sign up</a>
        </Link>
      </div>

      <style jsx>{`
        p {
          margin-top: 0;
        }

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

        .links {
          display: flex;
          justify-content: space-between;
          width: 50%;
          align-items: center;
          
        }

        a {
          font-size: 3em;
        }

      `
      }

      </style>
    </div>
  )
}