

package main

import (
	//"os"
	"fmt"
	//"net"
	//"time"
	"encoding/json"
	//"log"
	"golang.org/x/net/websocket"

)

type Message struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
}

//{"command":"get","node":"/instrumentation/comm/station-name"}
//{"command":"get","node":"/instrumentation/comm[1]/frequencies/selected-mhz"}
//{"command":"addListener","node":"/instrumentation/comm/station-name"}


func main() {


	origin := "http://192.168.50.153:7777/"
	url := "ws://192.168.50.153:7777/PropertyListener"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("fatal", err)
		return
	}

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



}


