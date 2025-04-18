package main

import (
	"fmt"
	"log"

	"github.com/irony0egoist/maldev/src/redteam"
)

func main() {
	hashes, err := redteam.AutoDumpHashes() // Retrieve hashes as raw structs
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range hashes {
		fmt.Println(h.Format()) // Print formatted hash
	}
}
