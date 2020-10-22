package fix_string_case

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(solve("a")).To(Equal("a"))
		Expect(solve("Z")).To(Equal("Z"))
		Expect(solve("coDe")).To(Equal("code"))
		Expect(solve("CODe")).To(Equal("CODE"))
		Expect(solve("coDE")).To(Equal("code"))
		Expect(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).To(Equal("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
	})
})