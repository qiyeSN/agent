package main

import (
	"errors"
	"github.com/agent/deal"
	"log"
)

func main() {
	/*err := getErr()
	if err != nil {
		fmt.Println("some wrong")
	}*/
	deal.CheckFile()
	//file()
	err := getErr()
	log.Println(err)
}

func getErr() error {
	//a := 1
	b := 0
	if b == 0 {
		return errors.New("some wrong")
	}
	return nil
}
