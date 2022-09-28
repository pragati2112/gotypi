package wsThings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"strconv"
	"time"
)

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}


func WsConnectionHandler(c *gin.Context) {
	roomId := c.Param("roomId")
	fmt.Println(roomId)
	fmt.Println("***********************")

	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		fmt.Println(err)
	}

	// create wrapped connection and subscription instances
	connection := &WSConnection{conn, make(chan []byte, 1024), strconv.FormatInt(timestamp(), 10)}
	subscription := Subscription{connection, roomId}

	// add this subscription into EditingRoomManager
	EditingRoomManagerInstance.Register <- subscription

	// start reader and writer go routines on the subscription
	go subscription.WriteToSubscription()
	go subscription.ReadFromSubscription()
}
