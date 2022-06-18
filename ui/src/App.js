import { useState, useEffect } from "react";

function getStatus() {
  return fetch("/api/healthz")
    .then((response) => response.json())
    .then((data) => data.status);
}

function App() {
  const [status, setStatus] = useState("Loading...");
  useEffect(() => {
    getStatus()
      .then((status) => setStatus(status))
      .catch((error) => setStatus("Error Loading Status"));
  })
  return (
    <div className="flex justify-center pt-16 text-center">
      <header>
        <h1>Hello World!</h1>
        <p>{status}</p>
      </header>
    </div>
  );
}

export default App;
