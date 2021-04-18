package main

import "github.com/google/wire"

// wire.go
//+build wireinject

func InitializeServer(gophers map[string]gopher.Gopher) server.Server {
	wire.Build(server.New, inmem.NewRepository, fetching.NewService, adding.NewService, modifying.NewService, removing.NewService)
	return server.NewWire()
}
