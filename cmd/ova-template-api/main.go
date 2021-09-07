package main

import (
	my_api "ova-animal-api/internal/api"
	"ova-animal-api/internal/config"
)

const (
	grpcAddr = ":8080"
	httpAddr = ":8081"
	network  = "tcp"
)

func main() {
	server := my_api.NewServer(
		config.Settings{
			GrpcAddr: grpcAddr,
			HttpAddr: httpAddr,
			Network:  network,
			Db: config.Db{
				Login:    "ova-animal-api",
				Password: "ova-animal-api",
				Name:     "ova-animal-api",
				Host:     "localhost",
				Port:     55432,
			},
		},
	)
	server.Run()
}
