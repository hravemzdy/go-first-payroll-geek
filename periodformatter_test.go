package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Periods", func() {
	const (
		testPeriodCodeJan17 int32 = 201701
		testPeriodCodeFeb17 int32 = 201702
		testPeriodCodeJan19 int32 = 201901
		testPeriodCodeFeb19 int32 = 201902
	)

	Describe("Period Formats With Date", func() {
		It("return January 2017 and February 2017", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan17)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb17)

			Expect(PeriodBeginDateFormat(testPeriodOne)).To(Equal("January 2017"))
			Expect(PeriodBeginDateFormat(testPeriodTwo)).To(Equal("February 2017"))
		})

		It("return January 2019 and February 2019", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan19)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb19)

			Expect(PeriodBeginDateFormat(testPeriodOne)).To(Equal("January 2019"))
			Expect(PeriodBeginDateFormat(testPeriodTwo)).To(Equal("February 2019"))
		})
	})
	Describe("Period Formats With Year and Month", func() {
		It("return 1/2017 and 2/2017", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan17)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb17)

			Expect(testPeriodOne.Format()).To(Equal("1/2017"))
			Expect(testPeriodTwo.Format()).To(Equal("2/2017"))
		})

		It("return 1/2019 and 2/2019", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan19)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb19)

			Expect(testPeriodOne.Format()).To(Equal("1/2019"))
			Expect(testPeriodTwo.Format()).To(Equal("2/2019"))
		})
	})
})
