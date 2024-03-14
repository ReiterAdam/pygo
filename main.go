package main

import (
	"log"

	"github.com/ReiterAdam/pygo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
