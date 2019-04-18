// chapter 3
package main

import (
	"fmt"
	"time"
)

func main() {

	// Pull in time.
	currentTime := time.Now()

	fmt.Println("The year is", currentTime.Year(), "should we leap?")

	var year = currentTime.Year()
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("It's a leap year.")
	} else {
		fmt.Println("No, it's not a leap year.")
	}

}
