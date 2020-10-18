package main

import (
	"fmt"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"net/http"
)


func main() {
	runServer()
}

var Copy string
var Connections int

func runServer(){
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		c.Emit("new-copy", Copy)
		c.Join("copy")
		c.BroadcastTo("copy","new-count", 	server.Amount("copy"))
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		c.BroadcastTo("copy", "new-count", server.Amount("copy"))
	})

	server.On("update", func(c *gosocketio.Channel, msg string) string {
		Copy = msg
		c.BroadcastTo("copy", "new-copy", Copy)
		return "OK"
	})


	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./files")))
	http.HandleFunc("/cp", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":	fmt.Fprintf(w, "%s", Copy)
		case "POST":
			Copy =  r.RequestURI[4:]
		}
	})

	log.Println("Serving at localhost:80...")
	log.Fatal(http.ListenAndServe(":80", nil))

}