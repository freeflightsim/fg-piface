


package piio

import (
	//"os"
	"fmt"

	"time"

	"github.com/luismesas/goPi/piface"
	"github.com/luismesas/goPi/spi"
)

type Board struct {
	Pfd      *piface.PiFaceDigital
	ButtChan chan Button
	States   map[int]bool
}

func NewPifaceBoard() *Board {
	b := new(Board)
	b.Pfd = piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)
	b.ButtChan = make(chan Button)
	b.States = make(map[int]bool)
	for i := 0; i < 8; i++ {
		b.States[i] = false
	}
	return b
}

func (me *Board) Init() error {
	fmt.Println("Board init()")
	err := me.Pfd.InitBoard()
	if err != nil {
		fmt.Println("error initialising board", err)
		return err
	}
	go me.ScanButtons()
	return nil
}

func (me *Board) ScanButtons() {
	fmt.Println("Board ScanButtons()")
	t := time.Tick(100 * time.Millisecond)
	for _ = range t {
		//fmt.Println(now)
		for i := 0; i < 8; i++ {
			v := me.Pfd.InputPins[i].Value() == 1
			if v == true && me.States[i] == false {
				// button pressed, but not previous
				me.States[i] = true
				b := Button{i, true}
				me.ButtChan <- b
			} else if v == false && me.States[i] == true {
				me.States[i] = false
				//me.ButtChan <- Button{i, false}

			}
			//fmt.Println(i, v, v)
		}

	}

}

func (me *Board) SetOutput(no int, state bool) {
	if state {
		me.Pfd.Leds[no].SetValue(1)
	} else {
		me.Pfd.Leds[no].SetValue(0)
	}

}
