package main

import (
	"fmt"
	"time"
)

func main() {

	cTime := time.Now()
	eTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	fmt.Println(cTime)
	fmt.Println(eTime)

	// Place your code here
}
