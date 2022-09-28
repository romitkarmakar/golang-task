
let backendUrl = process.env.REACT_APP_API_URL
let socket = new WebSocket("ws://" + backendUrl + '/ws')

function connect(){
    console.log("Creating Connection on domain: "+backendUrl)

    socket.onopen = () => {
        console.log("Connection Successful")
    }

    socket.onmessage = (msg:MessageEvent<any>) => {
        console.log(msg)
    }

    socket.onclose = (close:CloseEvent) => {
        console.log("Connection Closed : ",close)
    }

    socket.onerror = (error:Event) => {
        console.log("Connection Web Socket Error : ",error)
    }

}

function send(msg:string){
    console.log("Message Sending : ",msg)
    socket.send(msg)
}

export {connect, send}
