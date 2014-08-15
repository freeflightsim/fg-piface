

package main

import (
	//"os"
	"fmt"
	"net"
	//"time"
)

func main() {

      

        port := "127.0.0.1:4567"

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

        defer conn.Close()

		//var buf []byte
		buf := make([]byte, 1024)

		c := 0
        for {

               // time.Sleep(100 * time.Millisecond)

                n, address, err := conn.ReadFromUDP(buf)

                if err != nil {
                        fmt.Println("error reading data from connection")
                        fmt.Println(err)
                        return
                }

                if address != nil {

                        fmt.Println(">", c, address, " with n = ", n, string(buf[0:n]))

                        if n > 0 {
                                fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
                        }
                }
                c += 1
        }
 

}
