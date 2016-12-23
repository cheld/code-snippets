package main

import (
	"fmt"
	"net"
	"log"
	"github.com/songgao/packets/ethernet"
	//"github.com/songgao/water"
	"time"
	"golang.org/x/net/websocket"
	"github.com/songgao/water"
)

func main() {
	ifceOut, err := water.NewTUN("out")
	if err != nil {
		log.Fatal(err)
	}

	origin := "http://10.0.2.15/"
	url := "ws://172.31.0.123:1234/"
	con, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connected")

	go conToTap(con, ifceOut)
	go tunToCon(ifceOut, con)

	time.Sleep(1000000000000)
}

func conToTap(con net.Conn, tun *water.Interface) {
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := con.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
		log.Println("con to tap")
		tun.Write(frame)
	}
}

func tunToCon(tun *water.Interface, con net.Conn){
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := tun.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		//log.Printf("Dst: %s\n", frame.Destination())
		//log.Printf("Src: %s\n", frame.Source())
		//log.Printf("Ethertype: % x\n", frame.Ethertype())
		//log.Printf("Payload: % x\n", frame.Payload())
		log.Println("tun to con")
		con.Write(frame)
	}
}