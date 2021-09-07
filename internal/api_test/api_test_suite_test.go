//+build integration

package api_test_test

import (
	"ova-animal-api/internal/api_test"
	"ova-animal-api/internal/config"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAnimal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api integration test suite")
}

var settings = config.Settings{
	GrpcAddr: grpcAddr,
	HttpAddr: httpAddr,
	Network:  "tcp",
	Db: config.Db{
		Login:    "ova-animal-api",
		Password: "ova-animal-api",
		Name:     "ova-animal-api",
		Host:     "localhost",
		Port:     55433,
	},
}

var daoTester *api_test.DaoTester

var _ = BeforeSuite(func() {
	daoTester = &api_test.DaoTester{}
	daoTester.Init(&settings.Db)
})

var _ = AfterSuite(func() {
	daoTester.Close()
})
