package groute_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGroute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Groute Suite")
}
