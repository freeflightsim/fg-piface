
package ardio

import (
	"fmt"
	"bytes"
	"io"
	serial "github.com/tarm/goserial"

)

const (
	NL = 10 // \n character
)

type ArduinoBoard struct {
	Serial *io.ReadWriteCloser
	ReadChan chan string
}

func NewArduinoBoard() *ArduinoBoard {
	b := new(ArduinoBoard)
	b.ReadChan = make(chan string)
	return b
}

func (me *ArduinoBoard) Start() {

	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}

	var err error
	me.Serial, err = serial.OpenPort(c)
	if err != nil {
		//log.Fatal(err)
	}

	//n, err := s.Write([]byte("test"))
	//if err != nil {
		//log.Fatal(err)
	//}

	var lbuff bytes.Buffer
	buf := make([]byte, 128)
	for {
		n, err = s.Read(buf)
		if err != nil {
			//log.Fatal(err)
		} else {
			char := buf[:n]
			if char[0] == NL {
				s := lbuff.String()
				fmt.Println( "======", s  )
				lbuff.Reset()
			} else {
				lbuff.Write(char)
			}


		}


	}
}
