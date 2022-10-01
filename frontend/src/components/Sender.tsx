
import { useEffect, useState } from "react";
import { connect, send } from "../api";
import { Chats } from "./chats";
import { FieldButton } from "./fieldButton";

let backendUrl = process.env.REACT_APP_API_URL;
let socket = new WebSocket("ws://" + backendUrl + "/ws/sender");

function Sender() {
    const [chatHistory, setChatHistory] = useState<Array<string>>([]);

    function updateHistory(msg: string) {
      const arr: Array<string> = [...chatHistory, msg];
      setChatHistory(arr);
    }
    useEffect(() => {
        connect(updateHistory,socket);
      }, [chatHistory]); // dep
    return <div className="Sender">
        <Chats chatHistory={chatHistory}/>
        <FieldButton title="name" send={(str)=>{send(str,socket)}}/>
        <FieldButton title="message" send={(str)=>{send(str,socket)}}/>
    </div>
}

export {Sender}