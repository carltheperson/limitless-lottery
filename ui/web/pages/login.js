import UserForm from "../components/userForm"
import NoAuthHeader from "../components/noAuthHeader"

export default function Home() {

  return (
    <div className="container">
      <NoAuthHeader />

      <UserForm buttonText="Log In" />

    </div>

  )
}