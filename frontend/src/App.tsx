import React, {useEffect} from 'react';
import {connect, send} from "./api"

function sendMsg(){
  send("Hello")
}

function App() {
  useEffect(connect,[])
  return (
    <div className="App">
      <button onClick={sendMsg}>Hit</button>
    </div>
  );
}

export default App;
