package foo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestFoo(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "Foo Suite")
}
