package Content

import (
	"fmt"
	"time"
)

func Demo1() {
	fmt.Println("------------------------ Time Now --------------------------")
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%T\n", now)
}
func Demo2() {
	sometime := time.Date(2023, 07, 21, 11, 0, 0, 0, time.UTC)
	fmt.Println(sometime)
	fmt.Printf("%T\n", sometime)
}
