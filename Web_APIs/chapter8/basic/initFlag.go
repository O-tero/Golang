package main

import (
	"flag"
	"log"
)

var name string
var age int

func init() {
	flag.StringVar(&name, "name", "stranger", "your_wonderful_name")
	flag.IntVar(&age, "age", 0, "your_age")
}

func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years) Welcome to the command line world", *name, *age)
}
