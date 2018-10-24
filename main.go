package main

import "fmt"

func main() {
	var periodFrom = NewPeriod(2018, 1)
	var periodStop = NewPeriod(2019, 12)

	fmt.Printf("Interval periods from: %s to %s\n", periodFrom.Format(), periodStop.Format())
	fmt.Printf("Interval periods from: %s to %s\n", PeriodBeginDateFormat(periodFrom), PeriodEndDateFormat(periodStop))
}
