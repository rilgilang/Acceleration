// websockets.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Connections struct {
	Token string
	Pools []*websocket.Conn
}

func broadcast(pools []*websocket.Conn, msgType int, msg []byte) {
	for _, pool := range pools {
		pool.WriteMessage(msgType, msg)
	}
}

func main() {

	//make new pool
	AllConn := []Connections{}

	groupOne := Connections{
		Token: "asu",
	}

	groupTwo := Connections{
		Token: "koe",
	}

	AllConn = append(AllConn, groupOne)
	AllConn = append(AllConn, groupTwo)

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for _, group := range AllConn {
			if token == group.Token {

			}
		}

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			//Print the message to the console
			//fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			//

			//Write message back to browser
			//if err = conn.WriteMessage(msgType, msg); err != nil {
			//	return
			//}

			reply := fmt.Sprintf(`dari server %s`, msg)

			broadcast(pools, msgType, []byte(reply))
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8081", nil)
}
