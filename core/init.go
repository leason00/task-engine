package core

import "flag"

func init() {
	confPath := flag.String("config", "config.yaml", "config file")
	InitLog()
	if err := InitConfig(*confPath); err != nil {
		panic(err)
	}
	if err := InitRedis(); err != nil {
		panic(err)
	}

	Setup()

}
