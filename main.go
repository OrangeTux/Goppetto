package main

import (
	"github.com/OrangeTux/Goppetto/app"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	app.Run()
}
