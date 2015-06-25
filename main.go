

package main

import (
	//"os"
	"fmt"
	"time"
	//"reflect"

	"github.com/freeflightsim/fg-piface/config"
	"github.com/freeflightsim/fg-piface/fgio"
	"github.com/freeflightsim/fg-piface/piio"
	"github.com/freeflightsim/fg-piface/vstate"
)



func main() {

	//TODO, flags for host, port, config.yaml

	conf, err := config.Load("protocol/787.yaml")
	if err != nil {
		fmt.Println(" oops= ", err)
		return
	}
	fmt.Println(" conf= ", conf)

	//= Initialise some local sotre and state
	state := vstate.NewVState()
	state.AddNodes(  conf.GetOutputNodes() )

	eng_node := "/controls/engines/engine[1]/throttle"
	state.AddNode(eng_node)

	// initialise Piface
	board := piio.NewPifaceBoard()
	board.Init()
	if board.Enabled == false {
		// On a pc with no piface, we fake inputs with timers
		board.PretendInputs( conf.InputDefs )
	}

	// initialise the websocket client
	client := fgio.NewFgClient("192.168.50.153", "7777")
	client.AddNodes( state.GetNodes() )
	go client.Start()

	timer := time.NewTicker(time.Second)

	// Loop around the messages from channels
	for {
		select {

		case t := <- timer.C:
			//fmt.Println("t=", t.Second() )
			sec := float64((t.Second() % 10)) * 0.1
			//fmt.Println( sec )
			v := fmt.Sprintf( "%0.1f", sec )
			v2 := fmt.Sprintf( "%0.1f", 1.0 - sec )
			client.WsSet(eng_node, v)
			client.WsSet("/controls/engines/engine[0]/throttle", v2)


		// Messages from Flightgear
		case msg := <- client.WsChan:
			//fmt.Printf("#%s#\n", msg.Node)

			if msg.Node == eng_node {
				fmt.Println("eng", msg.RawValue )
			}

			state.Update( msg.Node, msg.StrValue() )

			for _, out_p := range conf.OutputDefs {

				if out_p.Node == msg.Node {
					//fmt.Println("        YES = ", led)
					on := out_p.IsOn( msg.StrValue() )
					board.SetOutput(out_p.Pin, on)

				}
			}
			//fmt.Println("nodes", node_vals)


		// Buttons Pressed
		case inp_ev := <- board.InputChan:
			//fmt.Println(" INNN = ", inp_ev)

			if true  {
				// find the value from config
				in_def := conf.GetInputFromPin(inp_ev.Pin)

				curr_val := state.GetNodeVal(in_def.Node)


				send_val := ""
				if curr_val == in_def.On {
					send_val = in_def.Off

				} else if curr_val == in_def.Off {
					send_val = in_def.On

				} else {
					fmt.Println("ARGS=", curr_val)
				}
				fmt.Println(in_def.Id, "curr=", curr_val, " on=", in_def.On, "off = " , in_def.Off, "send = ",  send_val)

				client.WsSet(in_def.Node, send_val)
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


