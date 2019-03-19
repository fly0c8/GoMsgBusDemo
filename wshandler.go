package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader    = websocket.Upgrader{}
	wsWriteChan = make(chan string)
	wsReadChan  = make(chan string)
	wsmap       = make(map[*websocket.Conn]bool) //map[socket]validated

	register   = make(chan *websocket.Conn)
	unregister = make(chan *websocket.Conn)
	Broadcast  = make(chan string)
)

func StartWebsocketHub() {
	go func() {
		for {
			select {
			case wsconn := <-register:
				wsmap[wsconn] = true
			case wsconn := <-unregister:
				if _, ok := wsmap[wsconn]; ok {
					delete(wsmap, wsconn)
					wsconn.Close()
				}
			case msg := <-Broadcast:
				for wsbroadcastconn := range wsmap {
					wsbroadcastconn.WriteMessage(websocket.TextMessage, []byte(msg))
				}
			}
		}
	}()
}

func Wshandler(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		// allow all connections
		return true
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer func() { unregister <- ws }()

	register <- ws

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("Received:", string(msg))
		MsgbusChan <- string(msg)

	}

	return nil
}
