package handle

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{ //用於http連接升級為websocket,設定一些限制
	CheckOrigin: func(r *http.Request) bool {
		return true //允許所有跨域連接
	},
}

func WsPing(c *gin.Context) {

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil) //將http升級為websocket協議
	if err != nil {
		log.Printf("create websocket connection")
		return
	}
	defer ws.Close()
	for {
		if err != nil {
			break
		}

		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		switch mt { //判斷client來的訊息種類
		case websocket.TextMessage:
			log.Println("TextMessage type")
		case websocket.BinaryMessage:
			log.Println("BinaryMessage type")
		case websocket.PingMessage:
			log.Println("PingMessage type")
		case websocket.PongMessage:
			log.Println("PongMessage type")
		case websocket.CloseMessage:
			log.Println("CloseMessage type")
		default:
			log.Println("unknow type")

		}

		if err := ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("got it : time: %s", string(message)))); err != nil {
			log.Println("WriteMessage error")
			break
		}
	}
}
