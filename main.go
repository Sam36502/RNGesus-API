package main

import (
	"fmt"

	"github.com/Sam36502/RNGesus-API/api"
)

const (
	API_PORT = 777
)

func main() {
	portstring := fmt.Sprintf(":%d", API_PORT)

	e := api.InitRouter()
	err := e.Start(portstring)
	fmt.Printf("Server Crashed:\n  %s\n", err.Error())
}
