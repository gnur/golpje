package command

import (
	"fmt"
	"net"
	"time"

	"github.com/gnur/golpje/events"
	pb "github.com/gnur/golpje/golpje"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":3222"
)

// ServerCommand basic setup
type ServerCommand struct {
	Test string
}

// server is a stub
type server struct{}

// Help returns the help for this command
func (c *ServerCommand) Help() string {
	return "Super awesome help for this ServerCommand"
}

// Run actually runs the command
func (c *ServerCommand) Run(args []string) int {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("IM NOTLISTENING!")
	}
	s := grpc.NewServer()
	pb.RegisterGolpjeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve")
	}

	return 0
}

// Synopsis returns a show description
func (c *ServerCommand) Synopsis() string {
	return "This does shit"
}

func (s *server) GetEvents(ctx context.Context, in *pb.EventRequest) (*pb.ProtoEvents, error) {
	var ev []events.Event
	var err error
	if in.All {
		ev, err = events.All()
	} else {
		ev, err = events.After(time.Unix(0, in.Since))
	}
	if err != nil {
		fmt.Println(err.Error())
		return &pb.ProtoEvents{}, err
	}

	var retEvents []*pb.ProtoEvent
	for _, event := range ev {
		apEvent := event.ToProto()
		retEvents = append(retEvents, &apEvent)
	}
	return &pb.ProtoEvents{
		Events: retEvents,
	}, nil
}

func (s *server) GetShows(ctx context.Context, in *pb.ShowRequest) (*pb.ProtoShows, error) {
	return &pb.ProtoShows{}, nil
}

func (s *server) AddShow(ctx context.Context, in *pb.ProtoShow) (*pb.AddShowResponse, error) {
	return &pb.AddShowResponse{}, nil
}

func (s *server) GetEpisodes(ctx context.Context, in *pb.EpisodeRequest) (*pb.ProtoEpisodes, error) {
	return &pb.ProtoEpisodes{}, nil
}

func (s *server) AddEpisode(ctx context.Context, in *pb.ProtoEpisode) (*pb.AddEpisodeResponse, error) {
	return &pb.AddEpisodeResponse{}, nil
}
