
package fgio


import (
	"fmt"
	"encoding/json"

	"golang.org/x/net/websocket"
)

type Client struct {

	Host string
	Port string

	Nodes map[string]bool

	Ws *websocket.Conn
	WsChan chan MessageFrame


}

// Creates a new FlightGear client
func NewClient(host string, port string) *Client{

	c := new(Client)
	c.Host = host
	c.Port = port
	c.Nodes = make(map[string]bool)
	c.WsChan = make(chan MessageFrame)

	return c
}

// Update nodes to listen on
func (me *Client) AddNodes(nodes []string){

	// first we make all current nodes, if any false
	//for k, _ := range me.Nodes {
	//	me.Nodes[k] = false
	//}

	// next we add the nodes
	for _, n := range nodes {
		me.Nodes[n] = true
	}

	// TODO, do the remove listener

}


func (me *Client) WsListen(){

	var bits = make([]byte, 512)
	var n int
	var err error
	//var m map[string]interface{}
	var fra MessageFrame
	for {
		n, err = me.Ws.Read(bits)
		if err != nil {
			fmt.Println("Read err", n, err)
		} else {
			//#fmt.Printf("Received: %s.\n", msg[:n])

			//fmt.Println("rcv", string(bits[:n]))
			err := json.Unmarshal(bits[:n], &fra)
			if err != nil {
				fmt.Println("decode error", err)
			} else {
				me.WsChan <- fra
			}
			//fmt.Println(m)
		}
	}
}

func (me *Client) Start() error {

	err := me.Connect()
	if err != nil {
		fmt.Println("Fatal, cannot start", err)
	}

	go me.WsListen()

	for node, _ := range me.Nodes {
		me.AddListener(node)
	}

	return nil
}

func (me *Client) AddListener(node string){

	fmt.Println(" + AddListener", node)
}

func (me *Client) Connect() error {

	// keeping adhoc creation of hosts etc in case
	// we can change ip et all on the fly
	// TODO make a reconnect on drop etc..
	origin := "http://" + me.Host + ":" + me.Port
	url := "ws://" + me.Host + ":" +  me.Port + "/PropertyListener"

	var err error
	me.Ws, err = websocket.Dial(url, "", origin)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("fatal", err)
		return err
	}

	return nil
}

func (me *Client) SendValue(node string, value string) {

	me.SendCommand( SetCmd(node, value) )
}


func (me *Client) SendCommand(comm interface{}) error {
	fmt.Println("SendCommand", comm)
	bits, err := json.Marshal(comm)
	if err != nil {
		fmt.Println("jsonerror", err)
		return err
	}
	//fmt.Println("bits", string(bits))
	if _, err := me.Ws.Write(bits); err != nil {
		//log.Fatal(err)
		fmt.Println("written", err)
	}
	return nil
}

