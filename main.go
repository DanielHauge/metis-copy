package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)


func main() {
	runServer()
}

var Copy string

func runServer(){
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "new-value", func(s socketio.Conn, msg string) {
		fmt.Println("update:", msg)
		Copy = msg
		s.Emit("update", "have "+msg)
	})


	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./files")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}