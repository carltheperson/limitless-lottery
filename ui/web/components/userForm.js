import Button from "./button"

export default function UserForm(props) {

  return (
    <div className="center-div">
      <div className="container">

        <p>Username</p>
        <br />
        <input />

        <br />

        <p>Password</p>
        <br />
        <input type="password" className="password" />

        <br />
        <div className="btn">
          <Button>{props.buttonText}</Button>
        </div>


      </div>

      <style jsx> {
        `
        .container {
          background-color: white;
          padding: 3em;
          border-radius: 6em;
          margin: 3%;
          display: inline-block;
          -webkit-box-shadow: 5px 6px 13px -1px rgba(0,0,0,0.35); 
          box-shadow: 5px 6px 13px -1px rgba(0,0,0,0.35);
          margin-top: 8%;
        }
        
        p {
          font-size: 4.5em;
          margin: 20px 0;
          display: inline-block;
        }

        input {
          border: none;
          border-radius: 0;
          background-color: #FFF970;
          font-size: 3em;
          font-weight: 600;
          padding: 5px;
          padding-left: 10px;
          margin-right: 75px;
        }

        .password {
          margin-bottom: 40px;
        }

        .btn {
          display: inline-block;
        }

        .center-div {
          display: flex;
          justify-content: center;
          align-items: center;
        }
      `
      }
      </style>
    </div>
  )

}