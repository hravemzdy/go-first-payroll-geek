package main

import "time"

func NewPeriodBeginDate(period Period) time.Time {
	return time.Date(int(period.Year()), time.Month(period.Month()), 1, 0, 0, 0, 0, time.UTC)
}

func NewPeriodEndDate(period Period) time.Time {
	var begin = NewPeriodBeginDate(period)
	return begin.AddDate(0, 1, -1)
}

func NewPeriodZeroDate() time.Time {
	return time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func PeriodBeginDateFormat(period Period) string {
	var periodDate = NewPeriodBeginDate(period)
	return periodDate.Format("January 2006")
}

func PeriodEndDateFormat(period Period) string {
	var periodDate = NewPeriodEndDate(period)
	return periodDate.Format("January 2006")
}
