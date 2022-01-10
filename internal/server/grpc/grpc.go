package grpc

import(
  "context"
  pbgameengine "github.com/denialtorres/m-apis/m-game-engine/v1"
  "github.com/denialtorres/m-game-engine/server/logic"
  "google.golang.org/grpc"
  "context"
  "github.com/rs/zerolog/log"
  "net"
  "github.com/pkg/errors"
)


type Grpc struct {
  address string
  srv *grpc.Server
}

// NewServer creates a new instance of a gRPC server
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// GetHighScore returns the highscore from the HighScore variable
func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in ms-game-engine called")
	sizeLogic := logic.GetSize()
	return &pbgameengine.GetSizeResponse{
		Size: sizeLogic, // For now this is a test size to see if connection happens correctly when the frontend calls it
	}, nil

}

// SetScore saves a score in ms-game-engine
func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("GetSize in ms-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil

}

// ListenAndServe starts the gRPC server on the given address
func (g *Grpc) ListenAndServe() error {
	// open tcp port to listen for incoming connections on
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open socket")
	}

	serverOpts := []grpc.ServerOption{}

	// create the server with the specified options
	g.srv = grpc.NewServer(serverOpts...)

	pbgameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Str("addr", g.address).Msg("starting gRPC server for m-game-engine microservice")

	// start listening on the given address
	if err := g.srv.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-game-engine microservice")
	}

	return nil
}
