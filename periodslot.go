package main

import "fmt"

type PeriodSlot struct {
	Code int16
}

//NewPeriod - creates a new instance of type Period with year and month parameters
func NewPeriodSlot(year int16) PeriodSlot {
	return PeriodSlot{year}
}

func (p PeriodSlot) Year() int16 {
	if p.Code == 0 {
		return 0
	}
	return int16(p.Code)
}

func (p PeriodSlot) CompareTo(other PeriodSlot) int {
	if p.Code == other.Code {
		return 0
	}
	if p.Code > other.Code {
		return 1
	}
	return -1
}

func (p PeriodSlot) Equals(other PeriodSlot) bool {
	return p.CompareTo(other) == 0
}

func (p PeriodSlot) OpGt(other PeriodSlot) bool {
	return p.CompareTo(other) > 0
}

func (p PeriodSlot) OpLt(other PeriodSlot) bool {
	return p.CompareTo(other) < 0
}

func (p PeriodSlot) Format() string {
	if p.Code == 0 {
		return ""
	}
	return fmt.Sprintf("%d", p.Code)
}
