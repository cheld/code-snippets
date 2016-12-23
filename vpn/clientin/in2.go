package main

import (
	"fmt"
	"net"
	"github.com/songgao/water"
	"log"
	"github.com/songgao/packets/ethernet"
	"time"
)

func main2() {
	ifceIn, err := water.NewTUN("in")
	if err != nil {
		log.Fatal(err)
	}
	targetAddr,err := net.ResolveTCPAddr("tcp4","172.31.0.123:7777")
	if err != nil {
		fmt.Println(err)
	}
	con,err := net.DialTCP("tcp4",nil,targetAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connected")

	go tunToCon(ifceIn,con)
	go conToTap(con, ifceIn)

	time.Sleep(1000000000000)
}




func tunToCon2(tun *water.Interface, con net.Conn){
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

func conToTap2(con net.Conn, tun *water.Interface) {
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