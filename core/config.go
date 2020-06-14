package core

import (
	"github.com/olebedev/config"
	"io/ioutil"
)

var Conf *config.Config

func InitConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	Conf, err = config.ParseYaml(string(file))
	if err != nil {
		return err
	}
	return nil
}
