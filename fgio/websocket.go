
package fgio


import (
	"fmt"
	"encoding/json"

	"golang.org/x/net/websocket"
)

type Client struct {

	Host string

	Port string

	Ws *websocket.Conn

	Nodes []string

	MessChan chan Frame //chan map[string]interface{}


}

// Creates a new FlightGear client
func NewClient(host string, port string) *Client{

	c := new(Client)
	c.Host = host
	c.Port = port
	c.Nodes = make([]string, 0)
	c.MessChan = make(chan Frame) // make(chan map[string]interface{})

	return c
}


func (me *Client) AddListener(node string){


	me.Nodes = append(me.Nodes, node)
	fmt.Println(" + AddListener", node)
}


func (me *Client) Listen(){

	var bits = make([]byte, 512)
	var n int
	var err error
	//var m map[string]interface{}
	var fra Frame
	for {
		n, err = me.Ws.Read(bits)
		if err != nil {
			fmt.Println("Read err", n, err)
		} else {
			//#fmt.Printf("Received: %s.\n", msg[:n])

			fmt.Println("rcv", string(bits[:n]))
			err := json.Unmarshal(bits[:n], &fra)
			if err != nil {
				fmt.Println("decode error", err)
			} else {
				me.MessChan <- fra
			}
			//fmt.Println(m)
		}
	}
}

func (me *Client) Start() error {

	me.Connect()

	return nil
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
	//fmt.Println("Connected")

	// Start the websocket reader
	go me.Listen()

	//fmt.Println("ssssssss", me.Nodes)
	for _, n := range me.Nodes {
		fmt.Println("addNode", n)
		comm := Command{"addListener", n, ""}
		me.SendCommand(comm)
	}

	return nil
}

func (me *Client) SendValue(node string, value string) {
	comm := Command{"set", node, value}
	me.SendCommand(comm)
}


func (me *Client) SendCommand(comm Command) error {
	bits, err := json.Marshal(comm)
	if err != nil {
		fmt.Println("jsonerror", err)
		return err
	}
	fmt.Println("bits", string(bits))
	if _, err := me.Ws.Write(bits); err != nil {
		//log.Fatal(err)
		fmt.Println("written", err)
	}
	return nil
}

