package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)
	if argsLen == 0 {
		fmtFatal("Please enter arguments")
	}
	n, err := randInt(argsLen)
	if err != nil {
		fmtFatal("An error has occurred: ", err)
	}
	fmt.Println(args[n])
}

func randInt(max int) (n int, err error) {
	if max <= 0 {
		err = errors.New("max must be greater than zero")
		return
	}
	var r *big.Int
	r, err = rand.Int(rand.Reader, big.NewInt(int64(max)))
	n = int(r.Int64())
	return
}

func fmtFatal(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(1)
}
