package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
)

func main() {
	spec := "0/5, *, *, *, ?, *"
	c := cron.New()
	err := c.AddFunc(spec, callYourFunc)
	if err != nil {
		fmt.Println("some wrong")
		return
	}
	c.Start()
	select {}
}

func callYourFunc() {
	a := 0
	fmt.Println("callYourFunc come here.")
	fmt.Println("heheheh", 1/a)
}
