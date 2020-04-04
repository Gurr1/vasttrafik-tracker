package main

import (
	"fmt"
	"github.com/gurr1/vasttrafik-tracker/service"
)

func main() {
	fmt.Println("starting server")
	service.Connect()
}
