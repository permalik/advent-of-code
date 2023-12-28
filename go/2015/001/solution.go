package main

import (
	"log"
	"os"
)

func main() {
    puzzleInput, err := os.ReadFile("./input.txt")
    if err != nil {
        log.Fatal("Panicked reading file", err)
    }

    os.Stdout.Write(puzzleInput)
}

