

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

	//for _, p := range conf.Outputs {
	//	client.AddListener(p.Node)
	//}
	client.UpdateNodes( conf.GetOutputNodes() )

	go client.Start()

	if board.Enabled == false {
		board.PretendInputs( conf.Inputs )
	}

	// Loop around the messages from channels
	//state := false
	for {
		select {

		// Messages from the client
		case msg := <- client.WsChan:
			//fmt.Printf("#%s#\n", msg.Node)
			for _, op := range conf.Outputs {

				if op.Node == msg.Node {
					//fmt.Println("        YES = ", led)
					on := op.IsOn( msg.StrValue() )
					board.SetOutput(op.Pin, on)

				}
			}



		// Buttons Pressed
		case input := <- board.InputChan:
			fmt.Println(" INNN = ", input)

			//if input.Pin == 0 {
				// find the value from config
				ip := conf.GetInputFromPin(input.Pin)
				fmt.Println(" n /v = ", ip)
				send_val := ip.Off
				if input.State == true {
					send_val = ip.On
				}
				client.SendValue(ip.Node, send_val)
			//}


		}
	}


}


