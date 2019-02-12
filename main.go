package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func main() {
	args := os.Args[1:]
	argsLen := int64(len(args))
	r, err := rand.Int(rand.Reader, big.NewInt(argsLen))
	if err != nil {
		fmt.Println("An error has occurred: ", err)
		os.Exit(1)
	}
	fmt.Println(args[r.Int64()])
}
