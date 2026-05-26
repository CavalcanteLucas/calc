package main

import (
	"fmt"
	"time"
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
	var total = 299
	var answered = 37
	var left = total - answered

	var work = 5  // working days per week
	var velocity = 2  // questions per day

	var remaining = calc_remaining(left, work, velocity)
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