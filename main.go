package main

import (
	_ "time/tzdata"

	"github.com/Thapanut/struct-test/configs"
	"github.com/Thapanut/struct-test/routes"
)

func init() {
	//set up config
	configs.Setup()
}

func main() {
	routes.StartServer()
}
