package main

import (
	"github.com/ymsht/mosaic-backend/router"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":1080"))
}
