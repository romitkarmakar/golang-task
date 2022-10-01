import { Sender } from "./components/Sender";
import { Receiver } from "./components/Receiver"
import { Link, Route, Routes } from "react-router-dom";

function Home() {
  return <div className="Home">
    <Link to={"sender"}> Sender </Link>
    <br/>
    <Link to={"receiver"}> Receiver </Link>
  </div>
}

function App() {
  return (
    <Routes>
      <Route path="/sender" element={<Sender />}></Route>
      <Route path="/receiver" element={<Receiver />}></Route>
      <Route path="/" element={<Home/>}></Route>
    </Routes>
  );
}

export default App;
