package main

import (
	"fmt"

	"github.com/heshanthenura/SentriGo/internal/common"
)

func main() {

	var index int

	devices := common.GetInterfaces()
	fmt.Print("Select Interface By Index :- ")
	fmt.Scanln(&index)
	common.SniffInterface(devices[index])
}
