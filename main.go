package main

import (
	"fmt"
	"time"
	"flag"
)

func calc_remaining(left int, work int, velocity int) int {
	rest := 7 - work
	fmt.Println(left, work, rest)
	count := 0
	for left > 0 {
		workday := 0
		for workday < work {
			if left <= 0 {
				return count
			}
			left -= velocity
			count++
			workday++
			fmt.Println("work", count, workday, left)
		}
		restday := 0
		for restday < rest {
			if left <= 0 {
				return count
			}
			count++
			restday++
			fmt.Println("rest", count, restday, left)
		}
	}
	return count
}

func main() {
	left := flag.Int("left", 0, "left")
	work := flag.Int("work", 5, "work")
	velocity := flag.Int("velocity", 1, "velocity")
	flag.Parse()

	var remaining = calc_remaining(*left, *work, *velocity)
	var today = time.Now()
	var eta = today.AddDate(0,0,remaining)

	fmt.Println(
		today.Format("2006-01-02"),
		"-->",
		remaining,
		"-->",
		eta.Format("2006-01-02"),
		"(ETA)",
	)
}