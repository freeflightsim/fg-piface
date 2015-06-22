

package config


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Led struct {
	Index int 		` yaml:"index" `
	Ignore bool 		` yaml:"ignore" `
	Node string 	` yaml:"node" `
	On string 		` yaml:"on" `
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
	return conf, nil
}

