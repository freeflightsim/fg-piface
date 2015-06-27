

package config


import (

	"strings"
)

type OutputPin struct {
	Pin int 		` yaml:"pin" `
	Id string 		` yaml:"id" `
	Node string 	` yaml:"node" `
	On string 		` yaml:"on" `
}


// Returns a list of unique output Nodes
func (me *Config) GetOutputNodes() []string {

	nodes := make(map[string]bool)
	for _, p := range me.DOutPins {
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




func (me OutputPin) IsOn(val string) bool {
	// TODO add >3 and comparison
	if strings.Contains(me.On, ","){
		for _, s := range strings.Split(me.On, ",") {
			if s == val {
				return true
			}
		}
	} else if me.On == val {
		return true
	}
	return false
}
