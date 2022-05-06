package api

import (
	"context"
	"log"

	"github.com/Ara-Chobanyan/school-manager/school-data/api/proto"
)

func (r *Adapter) GetStudentById(ctx context.Context, in *pb.StudentId) (*pb.Student, error) {
	id := in.Id

	// We would do something with the student data, from another interface
	student, err := r.api.GetStudentById(id)

	if err != nil {
		log.Fatalf("Could not retrived student from id: %v", err)
		return nil, err
	}

	res := pb.Student{
		Id:           int32(student.Id),
		FirstName:    student.FirstName,
		LastName:     student.LastName,
		Grade:        student.Grade,
		Behavior:     student.Behavior,
		Satisfaction: student.Satisfaction,
		Critical:     student.Critical,
		Exceptional:  student.Exceptional,
		Tutelary:     student.Tutelary,
		Staff:        student.Staff,
	}

	return &res, nil
}

func (r *Adapter) GetStudentByLastName(ctx context.Context, in *pb.StudentLastName) (*pb.StudentList, error) {
	var res *pb.StudentList = &pb.StudentList{}

	lastName := in.LastName

	student, err := r.api.GetStudentByLastName(lastName)

	if err != nil {
		log.Fatalf("Could not retrive a list of students by last name: %v", err)
		return nil, err
	}

	for _, x := range student {
		test := &pb.Student{
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

		res.Students = append(res.Students, test)
	}

	return res, err
}

func (r *Adapter) GetAllStudents(ctx context.Context, in *pb.StudentParams) (*pb.StudentList, error) {
	var res *pb.StudentList = &pb.StudentList{}

	students, err := r.api.GetAllStudents()

	if err != nil {
		log.Fatalf("Could not get a list of students: %v", err)
		return nil, err
	}

	for _, x := range students {
		student := &pb.Student{
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

		res.Students = append(res.Students, student)
	}

	return res, nil
}

func (r *Adapter) GetPopulation(ctx context.Context, in *pb.StudentParams) (*pb.Population, error) {

	teachers, students, err := r.api.GetPopulation()

	if err != nil {
		log.Fatalf("Could not get population: %v", err)
		return nil, err
	}

	res := pb.Population{
		Teacher: int32(teachers),
		Student: int32(students),
	}
	return &res, nil
}
