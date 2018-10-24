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

	Describe("Comparing Period", func() {
		It("different periods 2017 are equal", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan17)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeJan17)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeTrue())
		})

		It("different period 2019 are equal", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeFeb19)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb19)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeTrue())
		})

		It("different period same year Feb are greater then Jan", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeFeb17)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeJan17)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpGt(testPeriodTwo)).To(BeTrue())
		})

		It("different period same year Jan are less then Feb", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan19)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb19)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpLt(testPeriodTwo)).To(BeTrue())
		})

		It("different period same month 2019 are greater then 2017", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan19)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeJan17)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpGt(testPeriodTwo)).To(BeTrue())
		})

		It("different period same month 2017 are less then 2019", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeFeb17)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb19)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpLt(testPeriodTwo)).To(BeTrue())
		})

		It("different period Jan 2019 are greater then Feb 2017", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeJan19)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeJan17)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpGt(testPeriodTwo)).To(BeTrue())
		})

		It("different period Feb 2017 are less then Jan 2019", func() {
			testPeriodOne := NewPeriodWithCode(testPeriodCodeFeb17)
			testPeriodTwo := NewPeriodWithCode(testPeriodCodeFeb19)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpLt(testPeriodTwo)).To(BeTrue())
		})
	})

	Describe("Period Year and Month", func() {
		It("return period year 2017 and month 2", func() {
			testPeriod := NewPeriodWithCode(testPeriodCodeFeb17)

			Expect(testPeriod.Year()).To(Equal(int16(2017)))
			Expect(testPeriod.Month()).To(Equal(int16(2)))
		})

		It("return period year 2019 and month 1", func() {
			testPeriod := NewPeriodWithCode(testPeriodCodeJan19)

			Expect(testPeriod.Year()).To(Equal(int16(2019)))
			Expect(testPeriod.Month()).To(Equal(int16(1)))
		})
	})
})
