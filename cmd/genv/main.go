package main

import (
	"log"
	"os/exec"

	"github.com/brendanjcarlson/genv"
)

func main() {
	if err := genv.Load("./testdata/.env"); err != nil {
		log.Fatalf("failed to load env: %v\n", err)
	}

	out, err := exec.Command("printenv").Output()
	if err != nil {
		log.Fatalf("failed to exec `printenv`: %v\n", err)
	}
	log.Println(string(out))
}
