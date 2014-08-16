

package main

import (
	//"os"
	"fmt"
	"net"
	//"time"
	"encoding/json"
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

type AP_Packet struct {
	Ap int `json:"ap"`
	At int `json:"at"`
	LNav int `json:"lnav"`
	VNav int `json:"vnav"`
}

var Board *piface.PiFaceDigital

func main() {

      	Board = piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	err := Board.InitBoard()
	if err != nil {
		fmt.Println("error initialising board")
		return

	}
        port := "0.0.0.0:45670"

        udpAddress, err := net.ResolveUDPAddr("udp", port)

        if err != nil {
                fmt.Println("error resolving UDP address on ", port)
                fmt.Println(err)
                return
        }

        conn ,err := net.ListenUDP("udp",udpAddress)

        if err != nil {
                fmt.Println("error listening on UDP port ", port)
                fmt.Println(err)
                return
        }
	fmt.Println("linstening")
        defer conn.Close()

		//var buf []byte
		buf := make([]byte, 1024)

		c := 0
	var packet AP_Packet
        for {

               // time.Sleep(100 * time.Millisecond)

                n, address, err := conn.ReadFromUDP(buf)

                if err != nil {
                        fmt.Println("error reading data from connection", err)
                        fmt.Println(err)
                        return
                }

                if address != nil {

                        //fmt.Println(">", c, address, " with n = ", n, string(buf[0:n]))

                        if n > 0 {
				err_decode := json.Unmarshal(buf[0:n], &packet)
				if err_decode != nil {
					fmt.Println("decode err:", err_decode)
				} else {
	                                fmt.Println("from address", address, "got message:", string(buf[0:n]), n, packet)
				
					SetLed(L_AP, packet.Ap == 1)
					SetLed(L_AT, packet.At == 1)

					SetLed(L_HDG_HOLD, packet.LNav == 1)
					
					SetLed(L_VS, packet.VNav == 2)
					SetLed(L_ALT_HOLD, packet.VNav == 1)
				}
                        }
                }
                c += 1
        }
 

}

func SetLed(no int, state bool){
	if state {	
		Board.Leds[no].SetValue(1)
	} else {
		Board.Leds[no].SetValue(0)
	}

} 
