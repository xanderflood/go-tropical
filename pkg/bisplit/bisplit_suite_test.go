package bisplit_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBisplit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bisplit Suite")
}
