package lg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lg Suite")
}
