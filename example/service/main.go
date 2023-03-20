package main

import (
	todopb "example/gen"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {

	// initialize structured logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// listener for service
	lis, err := net.Listen("tcp", "localhost:5080")
	if err != nil {
		log.Fatal().Err(err).Msg("listen")
	}

	// instantiate grpcServer
	grpcServer := grpc.NewServer()

	// register todo service with grpcServer
	todopb.RegisterTodoServiceServer(grpcServer, newTodoService())

	// wrap grpcServer with grpcWeb
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	// enable cors
	handler := cors.AllowAll().Handler(wrappedGrpc)

	log.Info().Msg("listening on localhost:5080")

	if err = http.Serve(lis, handler); err != nil {
		log.Fatal().Err(err).Msg("serve")
	}
}
