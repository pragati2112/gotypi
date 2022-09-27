const WS_URL = 'ws://127.0.0.1:8080/ws/'

let socket = undefined

const setupWebSocket = (roomId = undefined, editor, lastChanges) => {
    if(!roomId) {
        roomId = document.querySelector("meta[name='roomid']").content
    }
    socket = new WebSocket(WS_URL + roomId)

    socket.onmessage = (event) => {
        let resp = JSON.parse(event.data)

        console.log('*****',resp, lastChanges)

        if (lastChanges !== resp){
            if(resp.action==='insert'){
                editor.session.insert(resp.start, resp.lines[0]);
            }

        }

    }

    socket.onerror = (error) => {
        console.error(error)
        socket.close()
    }

    socket.onopen = () => {
        clearInterval(this.timerId)
        console.log('connected to: ' + WS_URL)

        socket.onclose = ()  => {
            this.timerId = setInterval(() => {
                setupWebSocket()
            }, 1000)
        }
    }
}


function sendToWS(val) {
    socket.send(val)
}