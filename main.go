package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net/http"
	"time"
)

func handler(c *gin.Context) {
	log.Println("upgrading")
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		log.Println(err)
	}

	log.Println("ready")

	go func() {
		tick := time.NewTicker(time.Second * 5)
		defer tick.Stop()
		defer conn.Close()

		for {
			select {
			case <-tick.C:
				log.Println("tick")
				data, _ := json.Marshal(map[string]interface{}{
					"msg": "hello",
				})

				err = wsutil.WriteServerMessage(conn, ws.OpText, data)
				if err != nil {
					log.Println(err)
				}

			}
		}
	}()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Static("/static", "./static")

	router.GET("/ws", handler)

	router.LoadHTMLGlob("./templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Posts",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
