package fix_string_case

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFixStringCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FixStringCase Suite")
}
