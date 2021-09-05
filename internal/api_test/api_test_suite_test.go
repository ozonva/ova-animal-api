//+build integration

package api_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAnimal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api integration test suite")
}
