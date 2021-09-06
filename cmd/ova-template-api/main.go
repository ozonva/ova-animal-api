package main

import (
	my_api "ova-animal-api/internal/api"
)

const (
	grpcAddr = ":8080"
	httpAddr = ":8081"
	network  = "tcp"
)

func main() {
	server := my_api.NewServer(
		my_api.Settings{grpcAddr, httpAddr, network},
	)
	server.Run()
}
