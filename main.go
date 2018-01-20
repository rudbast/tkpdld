package main

import (
	"log"

	"github.com/rudbast/tkpdld/config"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[DEBUG] Loaded config: %+v\n", config.Get())
}
