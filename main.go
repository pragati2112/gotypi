package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
)

func handler(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				fmt.Println(err)
			}
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				fmt.Println(err)
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
