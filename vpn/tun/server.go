package main

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

func main() {
	ifceIn, err := water.NewTUN("in")
	if err != nil {
		log.Fatal(err)
	}
	ifceOut, err := water.NewTUN("out")
	if err != nil {
		log.Fatal(err)
	}
	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := ifceIn.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
		ifceOut.Write(frame)
	}
}