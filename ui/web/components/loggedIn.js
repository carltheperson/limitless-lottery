

let loggedIn = undefined;

function setLoggedIn(value) {
  loggedIn = value
}

function getLoggedIn() {
  return loggedIn
}

export default function LoggedIn() {

  return [getLoggedIn, setLoggedIn]
}