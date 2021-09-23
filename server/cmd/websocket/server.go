package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	u := &websocket.Upgrader{}
	c, err := u.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("cannot upgrade:%v \n", err)
		return
	}

	i := 0
	for {
		i++
		err := c.WriteJSON(map[string]string{
			"hello":  "websocket",
			"msg_id": strconv.Itoa(i),
		})
		if err != nil {
			fmt.Printf("cannot write json:%v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}

}
