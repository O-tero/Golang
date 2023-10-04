package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your_wonderful_name")

func main() {
	flag.Parse()
	log.Printf("Hello %s, Welcome to the command line world", *name)
}
