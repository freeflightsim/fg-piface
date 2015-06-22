

package main

import (
	//"os"
	"fmt"
	//"net"
	//"time"
	//"encoding/json"
	//"log"
	//"golang.org/x/net/websocket"
	"github.com/freeflightsim/fg-piface/config"
	"github.com/freeflightsim/fg-piface/fgio"
	"github.com/freeflightsim/fg-piface/piio"

)


//{"command":"get","node":"/instrumentation/comm/station-name"}
//{"command":"get","node":"/instrumentation/comm[1]/frequencies/selected-mhz"}
//{"command":"addListener","node":"/instrumentation/comm/station-name"}




func main() {

	conf, err := config.Load("protocol/787.yaml")
	if err != nil {
		fmt.Println(" oops= ", err)
		return
	}
	fmt.Println(" conf= ", conf)

	// initialise Piface
	board := piio.NewPifaceBoard()
	board.Init()


	// initialise the websocket clients
	client := fgio.NewClient("192.168.50.153", "7777")

	for _, led := range conf.Leds {
		client.AddListener(led.Node)
	}

	//bot.AddListener("/autopilot/settings/target-altitude-ft")
	//bot.AddListener("/autopilot/locks/altitude")
	//bot.AddListener("/autopilot/locks/heading")


	go client.Start()

	state := false
	for {
		select {
		case msg := <- client.MessChan:
			fmt.Println(" MSG = ", msg.Value)

			for _, led := range conf.Leds {
				if led.Node == msg.Node {
					fmt.Println(" YES = ", led)
					on := led.On == msg.Value
					board.SetOutput(led.Index, on)
					fmt.Println(" YES = ", on)
				}
			}

			state = !state
			/*
			board.SetOutput(0, state)
			board.SetOutput(2, state)
			board.SetOutput(5, state)
			board.SetOutput(7, state)
			board.SetOutput(8, state)
			*/
		case butt := <- board.ButtChan:
			fmt.Println(" BUtt = ", butt)
			client.SendValue("/autopilot/locks/heading", "fg-heading-hold")
		}
	}


}


