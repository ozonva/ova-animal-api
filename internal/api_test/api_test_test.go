//+build integration

package api_test_test

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/ptypes/empty"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"ova-animal-api/internal/api"
	grpc_api "ova-animal-api/pkg/ova-animal-api/github.com/ozonva/ova-animal-api/api"
	"time"
)

const grpcAddr = ":8888"
const httpAddr = ":8889"

var _ = Describe("ApiTest", func() {
	var (
		server *api.Server
	)

	BeforeEach(func() {
		server = api.NewServer(settings)

		go server.Run()

		daoTester.ClearTable("animal")
		// Используется godbt - клон dbunit на go
		daoTester.ApplyDbunitData("data/animal.xml")

		// Эту проблему не удалось решить. Надо как-то дождаться запуска сервера, но как, пока не понятно
		time.Sleep(300 * time.Millisecond)
	})

	AfterEach(func() {
		server.Shutdown()
		log.Info().Msgf("Test completed")
	})

	When("REST Json http", func() {
		var (
			httpClient *resty.Client
			urlBase    = "http://localhost" + httpAddr + "/v1"
		)
		BeforeEach(func() {
			httpClient = resty.New()
		})

		When("list", func() {
			It("empty", func() {
				_, err := httpClient.R().Get(urlBase)
				Expect(err).To(BeNil())
			})
		})

		When("get", func() {
			It("empty", func() {
				_, err := httpClient.R().Get(urlBase + "/100")
				Expect(err).To(BeNil())
				//
				//respObj := grpc_api.Animal{}
				//err = json.Unmarshal(resp.Body(), &respObj)
				//Expect(err).To(BeNil())
				//
				//Expect(respObj.Id).To(Equal(0))
				//Expect(respObj.UserId).To(Equal(0))
				//Expect(respObj.Name).To(Equal(""))
				//Expect(respObj.Type).To(Equal(grpc_api.Animal_UNKNOWN))
			})
		})

		When("delete", func() {
			It("empty", func() {
				_, err := httpClient.R().Delete(urlBase + "/1")
				Expect(err).To(BeNil())
			})
		})

		When("add", func() {
			It("empty", func() {
				_, err := httpClient.R().
					SetBody(grpc_api.Animal{}).
					Post(urlBase)
				Expect(err).To(BeNil())
			})
		})
	})

	When("gRPC", func() {
		var (
			grpcClient grpc_api.AnimalApiClient
			pool       *grpc.ClientConn
			ctx        context.Context
		)

		BeforeEach(func() {
			var err error

			ctx = context.Background()
			pool, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
			if err != nil {
				log.Panic().Err(err)
			}
			grpcClient = grpc_api.NewAnimalApiClient(pool)
		})

		AfterEach(func() {
			if err := pool.Close(); err != nil {
				log.Panic().Err(err)
			}
		})

		When("list", func() {
			It("empty", func() {
				entities, err := grpcClient.ListEntities(ctx, &empty.Empty{})
				Expect(err).To(BeNil())
				Expect(len(entities.Animal)).To(Equal(2))
				Expect(entities.Animal[0].Id).To(Equal(uint64(100)))
				Expect(entities.Animal[0].UserId).To(Equal(uint64(1)))
				Expect(entities.Animal[0].Name).To(Equal("Tom"))
				Expect(entities.Animal[0].Type).To(Equal(grpc_api.Animal_AnimalType_CAT))
				Expect(entities.Animal[1].Id).To(Equal(uint64(200)))
				Expect(entities.Animal[1].UserId).To(Equal(uint64(1)))
				Expect(entities.Animal[1].Name).To(Equal("Jerry"))
				Expect(entities.Animal[1].Type).To(Equal(grpc_api.Animal_AnimalType_MOUSE))
			})
		})
		When("get", func() {
			It("empty", func() {
				animal, err := grpcClient.DescribeEntity(ctx, &grpc_api.IdRequest{Id: 100})
				Expect(err).To(BeNil())
				Expect(animal.Id).To(Equal(uint64(100)))
				Expect(animal.UserId).To(Equal(uint64(1)))
				Expect(animal.Name).To(Equal("Tom"))
				Expect(animal.Type).To(Equal(grpc_api.Animal_AnimalType_CAT))
			})
		})
		When("delete", func() {
			It("empty", func() {
				_, err := grpcClient.RemoveEntity(ctx, &grpc_api.IdRequest{Id: 100})
				Expect(err).To(BeNil())
				daoTester.TestDatabase("animal", "data/animal_deleted.xml")
			})
		})
		When("create", func() {
			It("empty", func() {
				_, err := grpcClient.CreateEntity(ctx, &grpc_api.Animal{
					Id:     300,
					Name:   "Spike",
					UserId: 3,
					Type:   grpc_api.Animal_AnimalType_DOG,
				})
				Expect(err).To(BeNil())
				daoTester.TestDatabase("animal", "data/animal_added.xml")
			})
		})
	})
})
