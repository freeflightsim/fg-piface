
package ardio

import (
	"fmt"
	"bytes"
	"io"
	"strconv"
	"strings"
	serial "github.com/tarm/goserial"
	"time"

	//"github.com/freeflightsim/fg-piface/config"
)

const (
	NL = 10 // \n character
)

type AnalogVal struct {
	Pin int
	Val int64
}
type EncoderVal struct {
	Id int
	Val int64
}
type DOutPin struct {
	Node string
	Pin int
	Val bool
}


type ArduinoBoard struct {
	BoardId string
	Serial io.ReadWriteCloser
	AnalogChan chan AnalogVal
	EncoderChan chan EncoderVal
	Enabled bool
	Nodes map[string]bool
}

func NewArduinoBoard(board_id string) *ArduinoBoard {
	b := new(ArduinoBoard)
	b.BoardId = board_id
	b.AnalogChan = make(chan AnalogVal)
	b.EncoderChan = make(chan EncoderVal)
	b.Enabled = false
	b.Nodes = make(map[string]bool)
	return b
}

/*
func (me *ArduinoBoard) LoadConfig(conf *config.Config) {




}
*/
func (me *ArduinoBoard) Run() {

	// let things catch up
	time.Sleep(time.Second * 2)

	conf := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}

	var err error
	me.Serial, err = serial.OpenPort(conf)
	if err != nil {
		//log.Fatal(err)
		fmt.Println( "======", err  )
		return
	}
	me.Enabled = true

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
				//fmt.Println("ard Serial=", s, s[0:2] )

				if( len(s) > 4 && s[0:3] == "enc"){
					p := strings.Split(s, "=")
					v, verr := strconv.ParseInt(p[1], 10, 64)
					//fmt.Println(" -------l=", p, v, verr)
					if verr != nil {

					} else {
						//fmt.Println( "======", s  )
						me.EncoderChan <- EncoderVal{Id: 0, Val: v}
					}

				}
				if(false){
					i, oops := strconv.ParseInt(s, 10, 64)
					if oops != nil {

					} else {
						//fmt.Println( "======", s  )
						me.AnalogChan <- AnalogVal{Pin: 0, Val: i}
					}
				}
				lbuff.Reset()
			} else {
				lbuff.Write(char)
			}
		}
	}
}

func (me *ArduinoBoard) SendSerial(val string) {
	fmt.Println("SendSerial.n=", val)
	val = val + "\n"
	if me.Enabled == false {
		fmt.Println("SendSerial.backout <<")
		return
	}
	n, err := me.Serial.Write([]byte(val))
	if err != nil {
		fmt.Println("SendSerial.err=", err)
	}
	fmt.Println("SendSerial.senf=", n)
}
