package main

import "fmt"

type Period struct {
	Code int32
}

//NewPeriod - creates a new instance of type Period with year and month parameters
func NewPeriod(year int16, month int16) Period {
	var periodCode int32 = int32(year)*100 + int32(month)
	return Period{periodCode}
}

func NewPeriodWithCode(periodCode int32) Period {
	return Period{periodCode}
}

func (p Period) Year() int16 {
	if p.Code == 0 {
		return 0
	}
	return int16(p.Code / 100)
}

func (p Period) Month() int16 {
	if p.Code == 0 {
		return 0
	}
	return int16(p.Code % 100)
}

func (p Period) CompareTo(other Period) int {
	if p.Code == other.Code {
		return 0
	}
	if p.Code > other.Code {
		return 1
	}
	return -1
}

func (p Period) Equals(other Period) bool {
	return p.CompareTo(other) == 0
}

func (p Period) OpGt(other Period) bool {
	return p.CompareTo(other) > 0
}

func (p Period) OpLt(other Period) bool {
	return p.CompareTo(other) < 0
}

func (p Period) Format() string {
	if p.Code == 0 {
		return ""
	}
	var m int16 = p.Month()
	var y int16 = p.Year()
	return fmt.Sprintf("%d/%d", m, y)
}
