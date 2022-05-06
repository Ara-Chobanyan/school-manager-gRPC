package api

import (
	"context"
	"log"
	"net"

	"github.com/Ara-Chobanyan/school-manager/school-data/api/proto"
	"github.com/Ara-Chobanyan/school-manager/school-data/app"
	"google.golang.org/grpc"
)

type APIPort interface {
	GetStudentById(ctx context.Context, student *pb.StudentId) (*pb.Student, error)
	GetStudentByLastName(ctx context.Context, student *pb.StudentLastName) (*pb.StudentList, error)
	GetAllStudents(ctx context.Context, student *pb.StudentParams) (*pb.StudentList, error)
	GetPopulation(ctx context.Context, in *pb.StudentParams) (*pb.Population, error)
	Run() error
}

// Implement gRPC port interface
type Adapter struct {
	api app.DataService
	pb.UnimplementedDataManagmentServer
}

func NewAdapter(api app.DataService) APIPort {
	return &Adapter{api: api}
}

func (r *Adapter) Run() error {
	lis, err := net.Listen("tcp", ":5001")

	log.Println("Listening on :5001")

	if err != nil {
		log.Fatalf("failed to listen on port :5001: %v", err)
		return err
	}
	s := grpc.NewServer()

	pb.RegisterDataManagmentServer(s, r)

	return s.Serve(lis)
}
