import NoAuthHeader from "../components/noAuthHeader"
import NoAuthIndex from "./_noAuthIndex";

export default function Home() {

  return (
    <div className="container">
      <NoAuthHeader />

      <NoAuthIndex />

      <style jsx> {
        `

      `
      }
      </style>
    </div>
  )
}
