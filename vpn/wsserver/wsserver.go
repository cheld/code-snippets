package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
)


func main(){
	fmt.Println("hello")
	http.Handle("/",websocket.Handler(Echo))

	if err := http.ListenAndServe("127.0.0.1:1234", nil); err != nil {
		fmt.Println(err)
	}
}

func Echo(ws *websocket.Conn){
	var err error

	for {
		var reply string

		if err = websocket.Message.Send(ws, "hello...."); err != nil {
			fmt.Println("Can't send")
			break
		}

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from clientin: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to clientin: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
