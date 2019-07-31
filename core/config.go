package core

import (
	"flag"
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

func init() {
	confPath := flag.String("config", "config.yaml", "config file")
	InitLog()
	if err := InitConfig(*confPath); err != nil {
		panic(err)
	}
	if err := InitRedis(); err != nil {
		panic(err)
	}

	if err := InitMysql(); err != nil {
		panic(err)
	}

}
