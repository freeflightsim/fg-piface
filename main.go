

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
	client.AddNodes( conf.GetOutputNodes() )

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
			//fmt.Println("nodes", node_vals)


		// Buttons Pressed
		case inp_ev := <- board.InputChan:
			//fmt.Println(" INNN = ", inp_ev)

			if true  {
				// find the value from config
				in_def := conf.GetInputFromPin(inp_ev.Pin)

				curr_val := node_vals[in_def.Node]


				send_val := ""
				if curr_val == in_def.On {
					send_val = in_def.Off

				} else if curr_val == in_def.Off {
					send_val = in_def.On

				} else {
					fmt.Println("ARGS=", curr_val)
				}
				fmt.Println(in_def.Id, "curr=", curr_val, " on=", in_def.On, "off = " , in_def.Off, "send = ",  send_val)

				client.SendValue(in_def.Node, send_val)
				/*
				send_val := ip.Off
				if inp_ev.State == true {
					send_val = ip.On
				}
				*/
				//if send_val == "" {

				//}
				//fmt.Println(" n /v = ", ip, send_val)
				//client.SendValue(ip.Node, send_val)
			}


		}
	}


}


