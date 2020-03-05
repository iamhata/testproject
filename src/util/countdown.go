package util

import (
	"fmt"
	"time"
)

func Countdown(goal chan string) {
	ints := [] int { 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for _, i:= range ints {
		fmt.Println(i)
	}

	time.Sleep(1* time.Second)
	goal <- "countdown "
}