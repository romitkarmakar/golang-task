let backendUrl = process.env.REACT_APP_API_URL;
let socket = new WebSocket("ws://" + backendUrl + "/ws");

function connect(cb: (msg: string) => any) {
  console.log("Connection on domain: " + backendUrl);

  socket.onopen = () => {
    console.log("Connection Successful");
  };

  socket.onmessage = (msg: MessageEvent<any>) => {
    console.log(msg);
    cb(msg.data);
  };

  socket.onclose = (close: CloseEvent) => {
    console.log("Connection Closed : ", close);
  };

  socket.onerror = (error: Event) => {
    console.log("Connection Web Socket Error : ", error);
  };
}

function send(msg: string) {
  console.log("Message Sending : ", msg);
  socket.send(msg);
}

export { connect, send };
