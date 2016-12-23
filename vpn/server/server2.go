package main

import (
	"fmt"
	"net"
	"github.com/songgao/packets/ethernet"
	"log"
)

func main2() {
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

		go inToOut(conIn,conOut)
		go outToIn(conOut,conIn)

	}

}


func inToOut2(conIn, conOut net.Conn){
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

func outToIn2(conOut, conIn net.Conn){
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