package duplicate_encoder

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDuplicateEncoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DuplicateEncoder Suite")
}
