package main

import "mod-apis/configs"

func main() {
	//...
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
