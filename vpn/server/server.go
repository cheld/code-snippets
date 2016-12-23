package main

import (
	"fmt"
	"net"
	"github.com/songgao/packets/ethernet"
	"log"
	"net/http"

	"time"
	"golang.org/x/net/websocket"
)

var conIn *websocket.Conn
var conOut *websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func main(){


	http.HandleFunc("/",Pipe)
	fmt.Println("server starting")

	if err := http.ListenAndServe("0.0.0.0:1234", nil); err != nil {
		fmt.Println(err)
	}
}



func Pipe(w http.ResponseWriter, r *http.Request){
	fmt.Println("new handler")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	fmt.Println("upgraded")

	if conIn == nil {
		conIn = c
	} else if conOut == nil {
		conOut = c
		go inToOut()
		go outToIn()
	}
	time.Sleep(1000000000000)
}

func inToOut(){
	//var frame ethernet.Frame
	for {
		//frame.Resize(1500)
		//var n int
		var err error
		//var data []byte
		mt, message, err := conIn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(mt)
		log.Println(message)
		//frame = data
		//frame = frame[:n]
		conOut.WriteMessage(mt,message)
		//conOut.Write(frame)
		//fmt.Printf("Dst: a%s\n", frame.Destination())
		//fmt.Printf("Src: %s\n", frame.Source())
		//fmt.Printf("Ethertype: % x\n", frame.Ethertype())
		//fmt.Printf("Payload: % x\n", frame.Payload())
	}
}

func outToIn(){
	var frame ethernet.Frame
	for {
		frame.Resize(1500)
		var err error
		//var data []byte
		mt, message, err := conOut.ReadMessage();
		if err != nil {
			log.Fatal(err)
		}
		log.Println(mt)
		log.Println(message)
		//frame = data
		conIn.WriteMessage(mt,message)
		//conIn.Write(frame)
		//fmt.Printf("Dst: a%s\n", frame.Destination())
		//fmt.Printf("Src: %s\n", frame.Source())
		//fmt.Printf("Ethertype: % x\n", frame.Ethertype())
		//fmt.Printf("Payload: % x\n", frame.Payload())
	}
}


func main3() {
	fmt.Println("server starting")
	serverAddr,_ := net.ResolveTCPAddr("tcp","0.0.0.0:7777")
	fmt.Println(serverAddr)
	listener,err := net.ListenTCP("tcp",serverAddr)
	if err != nil {
		log.Fatal(err)
	}


	for {
		conIn,err := listener.Accept()
		if err != nil {
			fmt.Print(err)
		}
		log.Println("in connected")
		conOut,err := listener.Accept()
		if err != nil {
			fmt.Print(err)
		}
		log.Println("out connected")

		go inToOut3(conIn,conOut)
		go outToIn3(conOut,conIn)

	}

}


func inToOut3(conIn, conOut net.Conn){
	var frame ethernet.Frame
	for {
		frame.Resize(1500)
		n, err := conIn.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("inToOut")
		frame = frame[:n]
		fmt.Printf("Dst: a%s\n", frame.Destination())
		fmt.Printf("Src: %s\n", frame.Source())
		fmt.Printf("Ethertype: % x\n", frame.Ethertype())
		fmt.Printf("Payload: % x\n", frame.Payload())
		conOut.Write([]byte(frame))
	}
}

func outToIn3(conOut, conIn net.Conn){
	var frame ethernet.Frame
	for {
		frame.Resize(1500)
		n, err := conOut.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("outToIn")
		frame = frame[:n]
		fmt.Printf("Dst: a%s\n", frame.Destination())
		fmt.Printf("Src: %s\n", frame.Source())
		fmt.Printf("Ethertype: % x\n", frame.Ethertype())
		fmt.Printf("Payload: % x\n", frame.Payload())
		conIn.Write([]byte(frame))
	}
}