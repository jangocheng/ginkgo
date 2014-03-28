package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Failing Specs", func() {
	var pathToTest string

	BeforeEach(func() {
		pathToTest = tmpPath("failing")
		copyIn("fail_fixture", pathToTest)
	})

	It("should fail in all the possible ways", func() {
		output, err := runGinkgo(pathToTest, "--noColor")

		Ω(err).Should(HaveOccurred())
		Ω(output).ShouldNot(ContainSubstring("NEVER SEE THIS"))
		Ω(output).Should(ContainSubstring("a sync failure"))
		Ω(output).Should(MatchRegexp(`Test Panicked\n\s+a sync panic`))
		Ω(output).Should(ContainSubstring("a sync FAIL failure"))
		Ω(output).Should(ContainSubstring("async timeout [It]"))
		Ω(output).Should(ContainSubstring("Timed out"))
		Ω(output).Should(ContainSubstring("an async failure"))
		Ω(output).Should(MatchRegexp(`Test Panicked\n\s+an async panic`))
		Ω(output).Should(ContainSubstring("an async FAIL failure"))
		Ω(output).Should(ContainSubstring("a goroutine FAIL failure"))
		Ω(output).Should(ContainSubstring("a goroutine failure"))
		Ω(output).Should(MatchRegexp(`Test Panicked\n\s+a goroutine panic`))
		Ω(output).Should(ContainSubstring("a measure failure"))
		Ω(output).Should(ContainSubstring("a measure FAIL failure"))
		Ω(output).Should(MatchRegexp(`Test Panicked\n\s+a measure panic`))

		Ω(output).Should(ContainSubstring("0 Passed | 13 Failed"))
	})
})
