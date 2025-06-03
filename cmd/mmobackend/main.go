package main

import "isonetric-mmo-backend/configs"

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	println(config.Server.Port)
}
