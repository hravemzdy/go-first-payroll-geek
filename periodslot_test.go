package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PeriodSlot", func() {
	const (
		testPeriodCode17 int16 = 2017
		testPeriodCode19 int16 = 2019
	)

	Describe("Comparing PeriodSlot", func() {
		It("different periods slots are equal", func() {
			testPeriodOne := NewPeriodSlot(testPeriodCode17)
			testPeriodTwo := NewPeriodSlot(testPeriodCode17)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeTrue())
		})

		It("different period 2019 are equal", func() {
			testPeriodOne := NewPeriodSlot(testPeriodCode19)
			testPeriodTwo := NewPeriodSlot(testPeriodCode19)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeTrue())
		})

		It("different period 2019 are greater then 2017", func() {
			testPeriodOne := NewPeriodSlot(testPeriodCode19)
			testPeriodTwo := NewPeriodSlot(testPeriodCode17)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpGt(testPeriodTwo)).To(BeTrue())
		})

		It("different period slot 2017 are less then 2019", func() {
			testPeriodOne := NewPeriodSlot(testPeriodCode17)
			testPeriodTwo := NewPeriodSlot(testPeriodCode19)

			Expect(testPeriodOne.Equals(testPeriodTwo)).To(BeFalse())
			Expect(testPeriodOne.OpLt(testPeriodTwo)).To(BeTrue())
		})
	})

	Describe("PeriodSlot Year", func() {
		It("return period slot year 2017", func() {
			testPeriod := NewPeriodSlot(testPeriodCode17)

			Expect(testPeriod.Year()).To(Equal(int16(2017)))
		})

		It("return period slot year 2019", func() {
			testPeriod := NewPeriodSlot(testPeriodCode19)

			Expect(testPeriod.Year()).To(Equal(int16(2019)))
		})
	})
})
