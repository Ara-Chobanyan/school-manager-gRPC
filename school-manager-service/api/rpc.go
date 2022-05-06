package api

import (
	"context"
	"log"

	pb "github.com/Ara-Chobanyan/school-manager/school-manager-service/api/proto"
)

const behavior = "behavior"
const exceptional = "exceptional"

func (r *Adapter) Behavior(ctx context.Context, helper *pb.Helper) (*pb.StudentList, error) {
	var res *pb.StudentList = &pb.StudentList{}

	student, err := r.api.Behavior()

	if err != nil {
		log.Fatalf("Could not retrieve behavior list: %v", err)
		return nil, err
	}

	for _, x := range student {
		students := &pb.Student{
			Id:           int32(x.Id),
			FirstName:    x.FirstName,
			LastName:     x.LastName,
			Grade:        x.Grade,
			Behavior:     x.Behavior,
			Satisfaction: x.Satisfaction,
			Critical:     x.Critical,
			Exceptional:  x.Exceptional,
			Tutelary:     x.Tutelary,
			Staff:        x.Staff,
		}

		res.Students = append(res.Students, students)
	}

	// log.Println(helper.Who)
	err = r.api.FlagHelper(helper.Status, behavior)
	if err != nil {
		log.Fatalf("Could not change status: %v", err)
	}

	return res, nil
}

func (r *Adapter) Exceptional(ctx context.Context, helper *pb.Helper) (*pb.StudentList, error) {
	var res *pb.StudentList = &pb.StudentList{}

	student, err := r.api.Exceptional()

	if err != nil {
		log.Fatalf("Could not retrieve exceptional list: %v", err)
		return nil, err
	}

	for _, x := range student {
		students := &pb.Student{
			Id:           int32(x.Id),
			FirstName:    x.FirstName,
			LastName:     x.LastName,
			Grade:        x.Grade,
			Behavior:     x.Behavior,
			Satisfaction: x.Satisfaction,
			Critical:     x.Critical,
			Exceptional:  x.Exceptional,
			Tutelary:     x.Tutelary,
			Staff:        x.Staff,
		}

		res.Students = append(res.Students, students)
	}
	err = r.api.FlagHelper(helper.Status, exceptional)
	if err != nil {
		log.Printf("Could not change status")
	}

	return res, nil
}
