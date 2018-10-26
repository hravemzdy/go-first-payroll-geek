package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	)

var _ = Describe("bundle version repository", func() {
	Describe("add version to bundle repository", func() {
		It("should show up new version in bundle", func() {
			version := "first version"
			bundle := newMapVersionBundle()
			err := bundle.addVersion(version)

			Expect(err).To(BeNil(), "Got an error adding a version to bundle, should not have.")

			versions, err := bundle.getAllVersions()

			Expect(err).To(BeNil(), "Unexpected error in getAllVersions()", err)

			Expect(len(versions)).To(Equal(1), "Expected to have 1 payroll run in the repository")

			Expect(versions[0]).To(Equal("first version"), "version should have been <first version>")
		})
	})
	Describe("get version from bundle repository", func() {
		It("should retrieve proper version from bundle", func() {
			version := "first version"
			bundle := newMapVersionBundle()
			err := bundle.addVersion(version)

			Expect(err).To(BeNil(), "Got an error adding a version to bundle, should not have.")

			target, err := bundle.getVersion(0)
			Expect(err).To(BeNil(), "Got an error when retrieving version from bundle instead of success.")

			Expect(target).To(Equal("first version"), "Got the wrong version. Expected <first version>")

		})
	})
	Describe("empty bundle repository", func() {
		It("should get no version from bundle", func() {
			bundle := newMapVersionBundle()

			versions, err := bundle.getAllVersions()
			Expect(err).To(BeNil(), "Unexpected error in getAllVersions()")

			Expect(len(versions)).To(Equal(0), "Expected to have 0 versions in the bundle")
		})
	})
})

