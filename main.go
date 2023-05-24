package main

import (
	"github.com/vinitparekh17/go-mongo/config"
	"github.com/vinitparekh17/go-mongo/server"
)

func main() {
	config.Init()
	server.Init()
}
