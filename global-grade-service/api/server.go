package api

import (
	"context"
	"log"
	"net"

	"github.com/Ara-Chobanyan/school-manager/global-grade-service/api/proto"
	"github.com/Ara-Chobanyan/school-manager/global-grade-service/app"
	"google.golang.org/grpc"
)

type APIPort interface {
	GetGlobalAverage(ctx context.Context, in *pb.GlobalParams) (*pb.GlobalAverage, error)
	Run() error
}

type Adapter struct {
	api app.GlobalService
	pb.UnimplementedGlobalServer
}

func NewAdapter(api app.GlobalService) APIPort {
	return &Adapter{api: api}
}

func (r *Adapter) Run() error {
	lis, err := net.Listen("tcp", ":5002")

	log.Println("Listening on port :5002")

	if err != nil {
		log.Fatalf("Failed to lisen on port :5002: %v", err)
		return err
	}

	s := grpc.NewServer()

	pb.RegisterGlobalServer(s, r)

	return s.Serve(lis)
}
