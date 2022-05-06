package api

import (
	"context"
	"log"
	"net"

	"github.com/Ara-Chobanyan/school-manager/school-manager-service/api/proto"
	"github.com/Ara-Chobanyan/school-manager/school-manager-service/app"
	"google.golang.org/grpc"
)

type APIPort interface {
	Behavior(ctx context.Context, helper *pb.Helper) (*pb.StudentList, error)
	Exceptional(ctx context.Context, helper *pb.Helper) (*pb.StudentList, error)
	Run() error
}

type Adapter struct {
	api app.ServiceManager
	pb.UnimplementedSchoolManagerServer
}

func NewAdapter(api app.ServiceManager) APIPort {
	return &Adapter{api: api}
}

func (r *Adapter) Run() error {
	lis, err := net.Listen("tcp", ":5002")

	log.Println("Listening on :5002")

	if err != nil {
		log.Fatalf("failed to listen on port :5001: %v", err)
		return err
	}
	s := grpc.NewServer()

	pb.RegisterSchoolManagerServer(s, r)

	return s.Serve(lis)
}
