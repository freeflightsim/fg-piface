

package main

import (
	"os"
	"os/signal"
	"fmt"

	"github.com/freeflightsim/fg-piface/config"
	"github.com/freeflightsim/fg-piface/ardio"
	"github.com/freeflightsim/fg-piface/fgio"
	"github.com/freeflightsim/fg-piface/piio"
	"github.com/freeflightsim/fg-piface/vstate"
)



func main() {

	//TODO, flags for host, port, not sure where its gonna go yes, even config
	//host = flags.Foo("host")

	//= Handle ctrl+c to kill on terminal (required as were `multithread`)
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, os.Interrupt)

	//= Load Config (this also loads def in protocol/
	conf, err := config.Load("protocol/787.yaml")
	if err != nil {
		fmt.Println(" Config error= ", err)
		return
	}
	fmt.Println(" config = ", conf)


	//= Initialise the local State store thingi
	state := vstate.NewVState()
	state.AddNodes(  conf.GetOutputNodes() )

	// Some custom nodes in dev
	ias_node := "/instrumentation/airspeed-indicator/indicated-speed-kt"
	//state.AddNode(ias_node)

	hdg_bug := "/autopilot/settings/heading-bug-deg"
	state.AddNode(hdg_bug)

	eng_node := "/controls/engines/engine[1]/throttle"
	state.AddNode(eng_node)

	//= Initialise the flightgear client(s) later udp
	fg_client := fgio.NewFgClient("192.168.50.153", "7777")
	fg_client.AddNodes( state.GetNodes() )


	//=  Piface Digital IO board initialisaction (rpi only)
	pdf_board := piio.NewPifaceBoard()
	//pdf_board.Init()
	if pdf_board.Enabled == false {
		// On a pc with no piface, we fake inputs with timers
		//board.PretendInputs( conf.DInPins )
	}

	//= Arduino Board (current dev is Duemilanove .. olde)
	arduino_1 := ardio.NewArduinoBoard("ard_1")
	arduino_2 := ardio.NewArduinoBoard("ard_mega_1")
	//ard_board.LoadConfig(conf.FgNodes)


	go fg_client.Start()
	//go arduino_1.Run()
	//go arduino_2.Run()


	var last_v int64

	// Route all messages
	for {
		select {

		//= ctrl+c to kill
		case  <- killChan:
			// TODO gracefully shutdown things
			fmt.Println( "killed" )
			os.Exit(0)

		//= Analog pinInput from arduino ?
		case apin := <- arduino_1.AnalogChan:
		//case apin := <- arduino_2.AnalogChan:

			if last_v != apin.Val {
				vs := float64(apin.Val) / 100.0
				vsf := fmt.Sprintf("%0.2f", vs)
				fmt.Println("read", apin, vsf)
				fg_client.WsSet(eng_node, vsf)
				last_v = apin.Val
			}

		// Encoder pin from arduino
		case epin := <- arduino_1.EncoderChan:
			fmt.Println("Encoder Chan < ",   epin)
			chdg_val, enc_val_err := state.GetInt(hdg_bug)

			if enc_val_err == nil {
				rem := chdg_val % 5
				hdg_val := chdg_val - rem
				fmt.Println("   ",   chdg_val,  rem, hdg_val)
				if(epin.Val == -1){

					hdg_val = hdg_val + 5
				} else {
					hdg_val = hdg_val - 5
				}
				if(hdg_val > 360){
					hdg_val = hdg_val - 360
				}
				if(hdg_val < 0){
					hdg_val = hdg_val + 360
				}

				enc_val := fmt.Sprintf("%v", hdg_val)
				fmt.Println("encsend=",  hdg_bug, enc_val)
				fg_client.WsSet(hdg_bug, enc_val)


			}

		// Messages from Flightgear
		case msg := <- fg_client.WsChan:
			fmt.Println("", msg.Node, msg.StrValue())
			state.Update( msg.Node, msg.StrValue() )

			if msg.Node == eng_node {
				fmt.Println("eng", msg.StrValue() )
			}

			if msg.Node == ias_node {
				fmt.Println("iass", msg.StrValue() )
				arduino_1.SendSerial( msg.StrValue() )
				//num = ParseInt(msg.WayValue)
			}
			if msg.Node == ias_node {
				fmt.Println("iass", msg.StrValue() )
				arduino_2.SendSerial( msg.StrValue() )
				//num = ParseInt(msg.WayValue)
			}


			for _, out_p := range conf.DOutPins {

				if out_p.Node == msg.Node {
					fmt.Println("        YES = ", msg.Node)
					on := out_p.IsOn( msg.StrValue() )
					pdf_board.SetOutput(out_p.Pin, on)

				}
			}
			//fmt.Println("nodes", node_vals)


		// Buttons Pressed
		case inp_ev := <- pdf_board.InputChan:
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

				fg_client.WsSet(in_def.Node, send_val)
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


