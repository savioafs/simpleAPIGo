package main

import (
	"fmt"

	"github.com/savioafs/simpleAPIGo/configs"
)

func main() {
	config, _ := configs.LoadConfig("./")
	fmt.Println(config.DBDriver)
}
