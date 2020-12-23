export default function Button(props) {

  return (
    <div>
      <button onClick={props.onClick}>{props.children}</button>
      <style jsx> {
        `
          button {
            background-color: black;
            border: none;
            radius: 5px;
            font-size: 4em;
            color: #FFF970;
            padding: 0.25em 0.4em;
            font-weight: bold;
            text-align: center;
            border-radius: 0.5em;
            -webkit-box-shadow: 5px 6px 13px -1px rgba(0,0,0,0.53); 
            box-shadow: 5px 6px 13px -1px rgba(0,0,0,0.53);   
            cursor: pointer;
          }

          button:hover {
            -webkit-box-shadow: 5px 6px 13px -1px rgba(0,0,0,0.8); 
            box-shadow: 5px 6px 13px -1px rgba(0,0,0,0.8);  
          }
      `
      }
      </style>
    </div>
  )

}