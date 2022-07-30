const WS_URL = 'ws://127.0.0.1:8080/ws'

let socket = new WebSocket('ws://127.0.0.1:8080/ws');

function sendToWS(val) {
    socket.send(val);
}