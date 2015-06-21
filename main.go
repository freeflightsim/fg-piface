

package main

import (
	//"os"
	"fmt"
	//"net"
	//"time"
	//"encoding/json"
	//"log"
	//"golang.org/x/net/websocket"

	"github.com/freeflightsim/fg-piface/fgio"

)


//{"command":"get","node":"/instrumentation/comm/station-name"}
//{"command":"get","node":"/instrumentation/comm[1]/frequencies/selected-mhz"}
//{"command":"addListener","node":"/instrumentation/comm/station-name"}


func main() {



	bot := fgio.NewClient("192.168.50.153", "7777")

	//bot.AddListener("/autopilot/settings/target-altitude-ft")
	bot.AddListener("/autopilot/locks/altitude")
	bot.AddListener("/autopilot/locks/heading")

	go bot.Start()


	for {
		select {
		case msg := <-bot.MessChan:
			fmt.Println(" GOT = ", msg)
		}
	}

	/*

	m := Message{Cmd: "addListener", Node: "/autopilot/settings/target-altitude-ft"}
	bits, err := json.Marshal(m)
	fmt.Println("bits", string(bits))
	if _, err := ws.Write(bits); err != nil {
		//log.Fatal(err)
		fmt.Println("written", err)
	}
	var msg = make([]byte, 512)
	var n int
	for {
		n, err = ws.Read(msg)
		if err != nil {
			fmt.Println("Read err", n, err)
		} else {
			//#fmt.Printf("Received: %s.\n", msg[:n])
			fmt.Println("rcv", string(msg[:n]))
		}
	}
	*/
}


