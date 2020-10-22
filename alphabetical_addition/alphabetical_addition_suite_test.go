package alphabetical_addition

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAlphabeticalAddition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AlphabeticalAddition Suite")

	AddLetters([]rune{'r', 'n', 'r', 'y', 'j', 'l', 'g', 'w', 'c'})
}
