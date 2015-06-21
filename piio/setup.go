

package piio

import (
	//"os"
	"fmt"

	//"time"

	"github.com/luismesas/goPi/piface"
	"github.com/luismesas/goPi/spi"

)

const (
	L_AP = 0
	L_AT = 1
	L_LNAV = 2
	L_VNAV = 3
	L_HDG_HOLD = 4
	L_VS = 5
	L_ALT_HOLD = 6
	L_APP = 7

)

const (
	LOW = 0
	HIGHT = 1
)

type Board struct {
	Pfd *piface.PiFaceDigital
}

func NewPifaceBoard() *Board {
	b := new(Board)
	b.Pfd = piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)
	return b
}

func (me *Board) Init() error {

	err := me.Pfd.InitBoard()
	if err != nil {
		fmt.Println("error initialising board", err)
		return err
	}
	return nil
}

func (me *Board) SetOutput(no int, state bool) {
	if state {
		me.Pfd.Leds[no].SetValue(1)
	} else {
		me.Pfd.Leds[no].SetValue(0)
	}

}
