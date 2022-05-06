package api

import (
	"context"
	"log"

	"github.com/Ara-Chobanyan/school-manager/school-parent/api/proto"
	"github.com/Ara-Chobanyan/school-manager/school-parent/app"
)

// Creates a new student
func (r Adapter) CreateStudent(ctx context.Context, in *proto.NewStudent) (*proto.Student, error) {
	// init a student payload
	var payload app.Student

	// get the request data so we can return it later
	res := proto.Student{
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Grade:        in.Grade,
		Behavior:     in.Behavior,
		Satisfaction: in.Satisfaction,
		Exceptional:  in.Exceptional,
		Tutelary:     in.Tutelary,
		Staff:        in.Staff,
	}

	// set the payload to the request data
	payload.FirstName = in.FirstName
	payload.LastName = in.LastName
	payload.Grade = in.Grade
	payload.Behavior = in.Behavior
	payload.Satisfaction = in.Satisfaction
	payload.Critical = in.Critical
	payload.Exceptional = in.Exceptional
	payload.Tutelary = in.Tutelary
	payload.Staff = in.Staff

	// send our payload to the CreateStudent repo function
	err := r.api.CreateStudent(payload)
	if err != nil {
		log.Fatalf("Something went wrong in api.CreateStudent: %v", err)
		return nil, err
	}

	// return the data that was send as the response to show the data that was made in the database
	return &res, nil
}

// Deletes a student
func (r Adapter) DeleteStudent(ctx context.Context, in *proto.StudentId) (*proto.StudentId, error) {
	// Get the request id
	a := in.GetId()
	res := proto.StudentId{
		Id: a,
	}

	// send the request id to the DeleteStudent repo function
	err := r.api.DeleteStudent(a)

	if err != nil {
		log.Fatalf("Something went wrong in api.DeleteStudent: %v", err)
		return nil, err
	}

	// return the id as a response to show what was deleted
	return &res, nil
}

// Make edits to a student in the database
func (r Adapter) EditStudent(ctx context.Context, in *proto.Student) (*proto.Student, error) {
	// init the payload
	var payload app.Student

	// create a response of the request data
	res := proto.Student{
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Grade:        in.Grade,
		Behavior:     in.Behavior,
		Satisfaction: in.Satisfaction,
		Exceptional:  in.Exceptional,
		Tutelary:     in.Tutelary,
		Staff:        in.Staff,
	}

	// set the request data to the payload
	payload.Id = int(in.Id)
	payload.FirstName = in.FirstName
	payload.LastName = in.LastName
	payload.Grade = in.Grade
	payload.Behavior = in.Behavior
	payload.Satisfaction = in.Satisfaction
	payload.Critical = in.Critical
	payload.Exceptional = in.Exceptional
	payload.Tutelary = in.Tutelary
	payload.Staff = in.Staff

	// send the payload to the EditStudent function
	err := r.api.EditStudent(payload)
	if err != nil {
		log.Fatalf("api.EditStudent failed: %v", err)
		return nil, err
	}

	// send the response as the data that was give as the new edit
	return &res, nil
}
