package main

import (
	"fmt"
	"net"
	"github.com/songgao/water"
	"log"
	"github.com/songgao/packets/ethernet"
	"time"
	"golang.org/x/net/websocket"
)

func main() {
	ifceIn, err := water.NewTUN("in")
	if err != nil {
		log.Fatal(err)
	}
	origin := "http://localhost/"
	url := "ws://192.168.1.109:1234/"
	con, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connected")

	go tunToCon(ifceIn,con)
	go conToTap(con, ifceIn)


	time.Sleep(1000000000000)
}




func tunToCon(tun *water.Interface, con net.Conn){
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		fmt.Println("waiting for tun")
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

func conToTap(con net.Conn, tun *water.Interface) {
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := con.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		//log.Printf("Dst: %s\n", frame.Destination())
		//log.Printf("Src: %s\n", frame.Source())
		//log.Printf("Ethertype: % x\n", frame.Ethertype())
		//log.Printf("Payload: % x\n", frame.Payload())
		log.Println("con to tap")
		tun.Write(frame)
	}
}