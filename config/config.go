

package config


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"
	"strings"
)

type Led struct {
	Index int 		` yaml:"index" `
	Ignore bool 		` yaml:"ignore" `
	Node string 	` yaml:"node" `
	On string 		` yaml:"on" `
}

func (me Led) IsOn(val string) bool {

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


type Config struct {
	Model string 	` yaml:"model" `
	Leds []Led		` yaml:"leds" `
	//LedMap map[int]string
}

func Load(file_path string) (*Config, error) {



	contents, err_file := ioutil.ReadFile(file_path)
	if err_file != nil {
		return nil, err_file
	}

	conf := new(Config)
	err_yaml := yaml.Unmarshal(contents, &conf)
	if err_yaml != nil {
		return nil, err_yaml
	}
	/*
	conf.LedMap = make(map[string]int)
	for _, led := range conf.Leds {
		conf.LedMap[led.Node] = led.Index
	}
	*/
	err := conf.Validate()
	if err != nil {
		return conf, err
	}

	return conf, nil
}


func (me *Config) Validate() error {

	exists := make(map[int]bool)
	mess := ""

	for _, led := range me.Leds {
		if led.Index > 7 {
			mess +=  "Led " + led.Node + " has index > 7\n"
		}
		_, found := exists[led.Index]
		if found {
			mess +=  "Led " + led.Node + " has duplicate index\n"
		}
		exists[led.Index] = true
	}
	if mess == "" {
		return nil
	}
	return errors.New(mess)

}
