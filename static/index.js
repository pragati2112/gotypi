const WS_URL = 'ws://127.0.0.1:8080/ws/'

let socket = undefined

function sendToWS(val) {
    socket.send(val);
}

function connectToWebSocket(roomId = null) {
    if(!roomId) {
        roomId = document.querySelector("meta[name='roomid']").content
    }
    socket = new WebSocket(WS_URL + roomId);
}

function sendToRoom(){
// get a valid room id from server
// navigate user to this valid room id
    window.location.href = "/someRoomId"
}