package alphabetical_addition

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Tests", func() {
	It("Testing for []rune{'a', 'b', 'c'}", func() { Expect(AddLetters([]rune{'a', 'b', 'c'})).To(Equal('f')) })
	It("Testing for []rune{'z'}", func() { Expect(AddLetters([]rune{'z'})).To(Equal('z')) })
	It("Testing for []rune{'a', 'b'}", func() { Expect(AddLetters([]rune{'a', 'b'})).To(Equal('c')) })
	It("Testing for []rune{'c'}", func() { Expect(AddLetters([]rune{'c'})).To(Equal('c')) })
	It("Testing for []rune{'z', 'a'}", func() { Expect(AddLetters([]rune{'z', 'a'})).To(Equal('a')) })
	It("Testing for []rune{'y', 'c', 'b'}", func() { Expect(AddLetters([]rune{'y', 'c', 'b'})).To(Equal('d')) })
	It("Testing for []rune{}", func() { Expect(AddLetters([]rune{})).To(Equal('z')) })
})
