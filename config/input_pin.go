

package config


import (


)

type InputPin struct {
	Pin int 		` yaml:"pin" `
	Id string 		` yaml:"id" `
	Node string 	` yaml:"node" `
	On string 		` yaml:"on" `
	Off string 		` yaml:"off" `
}


// Returns a list of unique input Nodes
func (me *Config) GetInputNodes() []string {

	nodes := make(map[string]bool)
	for _, p := range me.Inputs {
		_, found := nodes[p.Node]
		if found == false {
			nodes[p.Node] = true
		}
	}
	lst := make([]string, 0)
	for n, _ := range nodes {
		lst = append(lst, n)
	}
	return lst

}


// Returns a list of unique input Nodes
func (me *Config) GetInputFromPin(pin_no int) InputPin {

	var ip InputPin
	for _, p := range me.Inputs {
		if p.Pin == pin_no {
			return p
		}

	}
	return ip
}
