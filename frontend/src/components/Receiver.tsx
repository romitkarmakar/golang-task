import { useEffect, useState } from "react";
import { connect } from "../api";
import { Chats } from "./chats";

let backendUrl = process.env.REACT_APP_API_URL;
let socket = new WebSocket("ws://" + backendUrl + "/ws/receiver");

function Receiver() {
  const [chatHistory, setChatHistory] = useState<Array<string>>([]);

  function updateHistory(msg: string) {
    const arr: Array<string> = [...chatHistory, msg];
    setChatHistory(arr);
  }

  useEffect(() => {
    connect(updateHistory,socket);
  }, [chatHistory]); // dependency should be chatHistory as the updateHistory function changes with chatHistory
  return (
    <div className="Receiver">
      <h1>Broadcasts</h1>
      <Chats chatHistory={chatHistory}/>
    </div>
  );
}

export {Receiver};
