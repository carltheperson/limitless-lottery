import Link from 'next/link'

export default function NoAuthHeader() {

  return (
    <div className="container">
      <Link href="/">
        <p className="logo-text">Limitless Lottery</p>
      </Link>

      <div className="links">
        <Link href="/about">
          <a>About</a>
        </Link>
        <Link href="/privacy">
          <a>Privacy</a>
        </Link>
        <Link href="/login">
          <a>Login</a>
        </Link>
        <Link href="/sign-up">
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
          cursor: pointer;
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