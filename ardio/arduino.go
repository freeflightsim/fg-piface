
package ardio

import (
	"fmt"
	"bytes"
	"io"
	"strconv"
	serial "github.com/tarm/goserial"
	"time"
)

const (
	NL = 10 // \n character
)

type AnalogPin struct {
	Pin int
	Val int64
}



type ArduinoBoard struct {
	Serial io.ReadWriteCloser
	AnalogChan chan AnalogPin
}

func NewArduinoBoard() *ArduinoBoard {
	b := new(ArduinoBoard)
	b.AnalogChan = make(chan AnalogPin)
	return b
}

func (me *ArduinoBoard) Run() {

	// let things catch up
	time.Sleep(time.Second * 2)

	conf := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}

	var err error
	me.Serial, err = serial.OpenPort(conf)
	if err != nil {
		//log.Fatal(err)
		fmt.Println( "======", err  )
	}

	//n, err := s.Write([]byte("test"))
	//if err != nil {
		//log.Fatal(err)
	//}
	var n int
	var lbuff bytes.Buffer
	buf := make([]byte, 128)
	for {
		n, err = me.Serial.Read(buf)
		if err != nil {
			//log.Fatal(err)
		} else {
			char := buf[:n]
			if char[0] == NL {
				s := lbuff.String()
				i, oops := strconv.ParseInt(s, 10, 64)
				if oops != nil {

				} else {
					//fmt.Println( "======", s  )
					me.AnalogChan <- AnalogPin{Pin: 0, Val: i}
				}
				lbuff.Reset()
			} else {
				lbuff.Write(char)
			}


		}


	}
}