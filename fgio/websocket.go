
package fgio


import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Client struct {
	Host string
	Port string
	Ws *websocket.Conn
	Nodes []string
}

func (this Client) AddListener(node string){

	this.Nodes = append(this.Nodes, node)
	fmt.Println("AddListener", node)
}

func (this Client) Start(){

	fmt.Println("SStart")


}

//origin := "http://192.168.50.153:7777/"
//url := "ws://192.168.50.153:7777/PropertyListener"
/*
func Connect(host string, port string) error {

	origin := "http://" + host + ":" + port
	url := "ws://" + host + ":" + port + "/PropertyListener"

	var err error
	Ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("fatal", err)
		return err
	}
	return nil
}
*/

func NewClient(host string, port string) *Client{

	c := new(Client)
	c.Host = host
	c.Port = port
	c.Nodes = make([]string, 0)
	return c
}
