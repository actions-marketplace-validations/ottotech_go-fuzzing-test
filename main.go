package main

import (
	"flag"
	"fmt"
	"github.com/ottotech/go-fuzzing-test/types"
	"log"
	"os"
)

func main() {
	var nFlag = flag.Int64("n", 0, "n is just a number you can pass")
	flag.Parse()

	if *nFlag == 0 {
		fmt.Printf("%q flag not given\n", "n")
		fmt.Printf("Usage of %s:\n", "go-fuzzing-test")
		flag.PrintDefaults()
		os.Exit(1)
	}

	n, err := types.OnlyCertainNumbers(*nFlag)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Yay your number is:", n)
}
