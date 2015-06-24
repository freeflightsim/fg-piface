


package piio

import (
	//"os"
	"fmt"

	"time"


	"github.com/luismesas/goPi/piface"
	"github.com/luismesas/goPi/spi"

	"github.com/freeflightsim/fg-piface/config"
)

type Board struct {
	Enabled bool
	Pfd      *piface.PiFaceDigital
	InputChan chan InputState
	States   map[int]bool
}

func NewPifaceBoard() *Board {
	b := new(Board)
	b.Enabled = false
	b.Pfd = piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)
	b.InputChan = make(chan InputState)
	b.States = make(map[int]bool)
	for i := 0; i < 8; i++ {
		b.States[i] = false
	}
	return b
}

func (me *Board) Init() error {
	//fmt.Println("Board init()")
	err := me.Pfd.InitBoard()
	if err != nil {
		fmt.Println("error initialising board", err)
		return err
	}
	me.Enabled = true
	go me.ScanButtons()
	return nil
}

// Desperate hack in a go routine that scans buttons for changes
// no interrupts yet from luis
func (me *Board) ScanButtons() {
	time.Sleep(2 * time.Second) // let things catch up
	fmt.Println("Board ScanButtons Enabled()")
	t := time.Tick(100 * time.Millisecond)
	for _ = range t {
		//fmt.Println(now)
		for i := 0; i < 8; i++ {
			v := me.Pfd.InputPins[i].Value() == 1

			// button pressed, but not previously
			if v == true && me.States[i] == false {
				me.States[i] = true
				inp := InputState{i, true}
				me.InputChan <- inp

			// button previouslt pressed has been released
			} else if v == false && me.States[i] == true {
				me.States[i] = false
				//me.ButtChan <- Button{i, false} // we only send presses...

			}
			//fmt.Println(i, v, v)
		}

	}

}

func (me *Board) SetOutput(no int, state bool) {
	if me.Enabled == false {
		//fmt.Println("      board -> SetOut", no, state)
		return
	}
	if state {
		me.Pfd.Leds[no].SetValue(1)
	} else {
		me.Pfd.Leds[no].SetValue(0)
	}

}


// fakes the input pins by sending messages randomly
func (me *Board) PretendInputs(  inputs []config.InputPin) {

	for _, inp := range inputs {

		fmt.Println("inp=", inp)
		if inp.Disabled == false {
			tick := time.NewTicker(time.Duration(inp.Pin + 2) * time.Second)
			go func(in_pin config.InputPin) {
				var state bool
				for {
					select {
					case <-tick.C:
						//fmt.Println("GO", tick, in_pin)
						state = !state
					me.InputChan <- InputState{Pin: in_pin.Pin, State: state}
					}
				}
			}(inp)
		}

	}

}
