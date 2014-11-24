package main

import (
	"github.com/OrangeTux/Goppetto/goppetto"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	goppetto.Run()
}
