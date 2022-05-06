package api

import (
	"context"
	"log"
	"net"

	"github.com/Ara-Chobanyan/school-manager/school-parent/api/proto"

	"github.com/Ara-Chobanyan/school-manager/school-parent/app"
	"google.golang.org/grpc"
)

// create our api interface
type APIPort interface {
	CreateStudent(ctx context.Context, student *proto.NewStudent) (*proto.Student, error)
	DeleteStudent(ctx context.Context, student *proto.StudentId) (*proto.StudentId, error)
	EditStudent(ctx context.Context, student *proto.Student) (*proto.Student, error)
	Run() error
}

//Implements gRPC port interface
type Adapter struct {
	// changed from repo service to parent service
	api app.ParentService
	proto.UnimplementedParentManagmentServer
}

// Implement our interface by taking in injection of the main services
func NewAdapter(api app.ParentService) APIPort {
	return &Adapter{api: api}
}

// start the server
func (r Adapter) Run() error {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen on port :5000: %v", err)
		return err
	}

	s := grpc.NewServer()

	proto.RegisterParentManagmentServer(s, r)

	log.Printf("Server listening at %v", listen.Addr())

	return s.Serve(listen)
}
