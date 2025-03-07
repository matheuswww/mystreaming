package upload_controller_util

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/matheuswww/mystream/src/logger"
)

func SendWsRes(msg any, conn *websocket.Conn) {
	err := conn.WriteJSON(msg)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying SendWsRes: %v", err))
		return
	}
	return
}