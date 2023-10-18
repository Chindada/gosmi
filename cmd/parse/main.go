package main

import (
	"log"
	"os"

	"github.com/Chindada/gosmi/parser"
	"github.com/alecthomas/repr"
)

func main() {
	module, err := parser.ParseFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	_ = module
	repr.Println(module)
}
