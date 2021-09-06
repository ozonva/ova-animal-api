package api

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"ova-animal-api/pkg/ova-animal-api/github.com/ozonva/ova-animal-api/api"
	"sync"
	"time"
)

func NewServer(settings Settings) *Server {
	server := &Server{
		settings: settings,
		wg:       &(sync.WaitGroup{}),
	}
	return server
}

type Server struct {
	settings   Settings
	wg         *sync.WaitGroup
	grpcServer *grpc.Server
	httpServer *http.Server
}

func (this *Server) runHttpRestJson() {
	this.wg.Add(1)
	defer this.wg.Done()

	log.Info().Msgf("Starting HTTP Rest server on address %s", this.settings.HttpAddr)
	log.Info().Msgf("Swagger url is http://localhost%s/", this.settings.HttpAddr)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := api.RegisterAnimalApiHandlerFromEndpoint(ctx, mux, this.settings.GrpcAddr, opts); err != nil {
		log.Panic().Err(err)
	}

	this.httpServer = &http.Server{
		ReadTimeout: 5 * time.Second,
		Addr:        this.settings.HttpAddr,
		Handler:     mux,
	}

	if err := this.httpServer.ListenAndServe(); err != nil {
		log.Panic().Err(err)
	}
	log.Info().Msgf("HTTP server stopped")
}

func (this *Server) runGrpc() {
	this.wg.Add(1)
	defer this.wg.Done()

	log.Info().Msgf("Starting gRPC server on %s address %s", this.settings.Network, this.settings.GrpcAddr)

	listen, err := net.Listen(this.settings.Network, this.settings.GrpcAddr)

	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	this.grpcServer = grpc.NewServer()
	api.RegisterAnimalApiServer(this.grpcServer, AnimalApiServerImpl{})

	if err := this.grpcServer.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
	log.Info().Msgf("gRPC server stopped")
}

func (this *Server) runAsyncWithinWaitGroup(payload func()) {
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()
		payload()
	}()
}

func (this *Server) Run() {
	this.runAsyncWithinWaitGroup(this.runGrpc)
	this.runAsyncWithinWaitGroup(this.runHttpRestJson)
	this.wg.Wait()
}

func (this *Server) Shutdown() {
	log.Info().Msgf("Shutting down gRPC server...")
	this.grpcServer.Stop()

	log.Info().Msgf("Shutting down HTTP server...")
	if err := this.httpServer.Shutdown(context.Background()); err != nil {
		log.Panic().Err(err)
	}

	this.wg.Wait()
}
