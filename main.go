

package main

import (
	//"os"
	"fmt"

	//"reflect"

	"github.com/freeflightsim/fg-piface/config"
	"github.com/freeflightsim/fg-piface/fgio"
	"github.com/freeflightsim/fg-piface/piio"

)




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


	// initialise the websocket client
	client := fgio.NewClient("192.168.50.153", "7777")

	for _, p := range conf.Outputs {
		client.AddListener(p.Node)
	}

	go client.Start()

	// Loop around the messages from channels
	state := false
	for {
		select {

		// Messages from the client
		case msg := <- client.MessChan:
			//if msg.Node ==  "/instrumentation/flightdirector/autopilot-on" {
			//	fmt.Println(" GOT = ", msg.RawValue, reflect.TypeOf(msg.RawValue), msg.Type, msg.String())
			//}
			fmt.Printf("#%s#\n", msg.Node)
			for _, op := range conf.Outputs {


				if op.Node == msg.Node {
					//fmt.Println("        YES = ", led)
					on := op.IsOn(msg.String())
					//if msg.Node ==  "/instrumentation/flightdirector/autopilot-on" {
					//	fmt.Println("        COMP = ", on, led.On, msg.String(), reflect.TypeOf(led.On), reflect.TypeOf(msg.String()))
						//fmt.Println(" YES = ", on)
						//fmt.Printf("#%s#\n", led.On)
						//fmt.Printf("#%s#\n", msg.String())
					//}
					if on {
						fmt.Printf("  #%s# ON\n", op.Node)
					} else{
						fmt.Printf("  #%s# --\n", op.Node)
					}
					board.SetOutput(op.Pin, on)

				}
			}

			state = !state

		// Buttons Pressed
		case butt := <- board.ButtChan:
			fmt.Println(" BUtt = ", butt)
			client.SendValue("/autopilot/locks/heading", "fg-heading-hold")
		}
	}


}


