package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {

	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1) //return non-null exit code
	}
	fmt.Printf("ntp: %v, currect time: %v", ntpTime, time.Now())
}
