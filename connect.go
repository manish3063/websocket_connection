package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	getlink()

	r := gin.Default()
	setupRoutes(r)
	r.Run()

}

func setupRoutes(r *gin.Engine) {

	r.GET("/conn", registerClient)

}

type linkDetails struct {
	Q string `json:"q"`
	A string `json:"a"`
	H string `json:"h"`
}

var wsupgraders = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func registerClient(c *gin.Context) {

	wsupgraders.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsupgraders.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		msg := fmt.Sprintf("Failed to set websocket upgrade: %+v", err)
		fmt.Println(msg)
		return
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 5)
		res := getlink()
		// mType, mByte, err := conn.ReadMessage()
		// fmt.Println("mByte: ", string(mByte))
		// fmt.Println("mType: ", mType)
		// fmt.Println("err: ", err)

		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s", res)))
	}
	conn.Close()
}
