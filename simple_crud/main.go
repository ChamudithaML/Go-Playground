package main

import (
	"fmt"
	"main/dbconfig"
)

func main() {
	fmt.Println("hi")

	dbconfig.SetupDb()
}
