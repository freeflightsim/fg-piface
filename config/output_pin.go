

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
