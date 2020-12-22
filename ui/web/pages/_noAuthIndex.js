import Image from 'next/image'
import Button from "../components/button"


export default function NoAuthIndex() {

  return (
    <div className="container">
      <div className="text">
        <p className="thin">
          A simulation of being able to buy as many lottery tickets as you want.
      </p>
        <p>
          Will you make more than you spend?
        </p>
        <div className="btn">
          <Button onClick={() => { }}>Sign Up</Button>
        </div>


      </div>

      <div className="img-container">
        <Image src="/images/tickets.png" alt="tickets" width="340px" height="630px" />
      </div>

      <style jsx> {
        `
        p {
          text-align: center;
          font-size: 4em;
        }

        .thin {
          font-weight: 400;
        }

        .container {
          display: flex;
          justify-content: space-between;
          width: 100%;
        }

        .text {
          width: 100%;
          padding: 10% 8% 0 8%;
        }

        img {
          display: block !important;
        }

        .img-container {
          width: 40%;
          text-align: right;
          display: block !important;
          padding-top: 5%;
        }

        .btn {
          text-align: center;
          padding-top: 3em;
        }
      `
      }
      </style>
    </div>
  )
}