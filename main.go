

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

	//TODO, flags for host, port, config.yaml

	conf, err := config.Load("protocol/787.yaml")
	if err != nil {
		fmt.Println(" oops= ", err)
		return
	}
	fmt.Println(" conf= ", conf)

	node_vals := make( map[string]string )

	// initialise Piface
	board := piio.NewPifaceBoard()
	board.Init()


	// initialise the websocket client
	client := fgio.NewClient("192.168.50.153", "7777")
	client.UpdateListeners( conf.GetOutputNodes() )

	go client.Start()

	if board.Enabled == false {
		board.PretendInputs( conf.InputDefs )
	}

	// Loop around the messages from channels
	//state := false
	for {
		select {

		// Messages from the client
		case msg := <- client.WsChan:
			//fmt.Printf("#%s#\n", msg.Node)
			//v, found := node_vals[msg.Node]
			//if found == false {
			node_vals[msg.Node] = msg.StrValue()
			//} else {

			//}

			for _, op := range conf.OutputDefs {

				if op.Node == msg.Node {
					//fmt.Println("        YES = ", led)
					on := op.IsOn( msg.StrValue() )
					board.SetOutput(op.Pin, on)

				}
			}
			fmt.Println("nodes", node_vals)


		// Buttons Pressed
		case input := <- board.InputChan:
			//fmt.Println(" INNN = ", input)

			//if input.Pin == 0 {
				// find the value from config
				ip := conf.GetInputFromPin(input.Pin)

				send_val := ip.Off
				if input.State == true {
					send_val = ip.On
				}
			if send_val == "" {

			}
				//fmt.Println(" n /v = ", ip, send_val)
				//client.SendValue(ip.Node, send_val)
			//}


		}
	}


}


