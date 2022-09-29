import React, { useEffect, useState } from "react";
import { connect, send } from "./api";
import { Chats } from "./components/chats";

function sendMsg() {
  send("Hello");
}

function App() {
  const [chatHistory, setChatHistory] = useState<Array<string>>([]);

  function updateHistory(msg: string) {
    const arr: Array<string> = [...chatHistory, msg];
    setChatHistory(arr);
  }

  useEffect(() => {
    connect(updateHistory);
  }, [chatHistory]); // dependency should be chatHistory as the updateHistory function changes with chatHistory
  return (
    <div className="App">
      <Chats chatHistory={chatHistory}></Chats>
      <button onClick={sendMsg}>Hit</button>
    </div>
  );
}

export default App;
