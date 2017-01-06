package main

import (
	"fmt"
	"net"
	"github.com/songgao/packets/ethernet"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/websocket"
	"strings"
)

//var conIn *websocket.Conn
//var conOut *websocket.Conn

var con map[string]*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func main(){
	con = make(map[string]*websocket.Conn)


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

	remoteAddr := strings.Split(c.UnderlyingConn().RemoteAddr().String(),":")[0]
	fmt.Println(remoteAddr)

	if con["10.1.0.10"] == nil {
		fmt.Println("first connection recieves 10.1.0.10")
		con["10.1.0.10"] = c
	} else if con["10.1.0.11"] == nil {
		fmt.Println("second connection recieves 10.1.0.11")
		con["10.1.0.11"] = c
	}


	go inToOut(c)

	time.Sleep(1000000000000)
}

func inToOut(conIn *websocket.Conn){
	fmt.Println("thread started")
	var frame ethernet.Frame
	for {
		frame.Resize(1500)
		//var n int
		var err error
		//var data []byte
		mt, message, err := conIn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		frame = message

		//fmt.Printf("Payload: % x\n", frame.Payload())
		//frame = frame[:n]

		payload := frame.Payload()

		//t := fmt.Sprintf("test %v",int(payload[3]))
		dest := net.IPv4(payload[2],payload[3],payload[4],payload[5])

		fmt.Printf("sending data to %v",dest)
		conOut := con[dest.String()]


		conOut.WriteMessage(mt,message)
		//conOut.Write(frame)
		//fmt.Printf("Dst: a%s\n", frame.Destination())
		//fmt.Printf("Src: %s\n", frame.Source())
		//fmt.Printf("Ethertype: % x\n", frame.Ethertype())
		//fmt.Printf("Payload: % x\n", frame.Payload())
	}
}

//func outToIn(){
//	var frame ethernet.Frame
//	for {
//		frame.Resize(1500)
//		var err error
		//var data []byte
//		mt, message, err := conOut.ReadMessage();
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Println(mt)
//		log.Println(message)
		//frame = data
		//conIn.WriteMessage(mt,message)
		//conIn.Write(frame)
		//fmt.Printf("Dst: a%s\n", frame.Destination())
		//fmt.Printf("Src: %s\n", frame.Source())
		//fmt.Printf("Ethertype: % x\n", frame.Ethertype())
		//fmt.Printf("Payload: % x\n", frame.Payload())
//	}
//}


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