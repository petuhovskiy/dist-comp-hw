package main

import (
	"log"
	"sms/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := config.EnvConfig()
	if err != nil {
		log.Panic("failed to read config, err=", err)
	}

}
