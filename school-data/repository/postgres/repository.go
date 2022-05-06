package postgres

import (
	"context"
	"log"
	"time"

	"github.com/Ara-Chobanyan/school-manager/school-data/app"
	"github.com/jackc/pgx/v4"
)

type Conn struct {
	Conn *pgx.Conn
}

func openConn() (*pgx.Conn, error) {
	dsn := "postgres://user:password@localhost/gokit?sslmode=disable"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalf("Could not connect on postgres.openConn: %v", err)
		return nil, err
	}

	return conn, err
}

func NewPostgresRepository() (app.DataRepository, error) {
	repo := &Conn{}

	client, err := openConn()

	if err != nil {
		log.Fatalf("Could not open up client on postgres.NewPostgresRepository: %v", err)
		return nil, err
	}

	repo.Conn = client

	return repo, nil
}

func (r *Conn) GetStudentById(id int32) (*app.Student, error) {
	var student app.Student

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM students WHERE id = $1;"

	rows := r.Conn.QueryRow(ctx, query, id)

	err := rows.Scan(
		&student.Id,
		&student.FirstName,
		&student.LastName,
		&student.Grade,
		&student.Behavior,
		&student.Satisfaction,
		&student.Critical,
		&student.Exceptional,
		&student.Tutelary,
		&student.Staff,
	)

	if err != nil {
		log.Fatalf("Could not get data postgres.GetStudentById %v", err)
		return nil, err
	}

	return &student, nil
}

func (r *Conn) GetStudentByLastName(last_name string) ([]*app.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT * FROM students WHERE last_name = $1`

	rows, err := r.Conn.Query(ctx, query, last_name)

	if err != nil {
		log.Fatalf("Something went wrong with postgres.GetStudentByLastName: %v", err)
		return nil, err
	}

	var students []*app.Student

	for rows.Next() {
		var student app.Student

		err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.Grade,
			&student.Behavior,
			&student.Satisfaction,
			&student.Critical,
			&student.Exceptional,
			&student.Tutelary,
			&student.Staff,
		)

		if err != nil {
			log.Fatalf("Could not retrieve a [] students: %v", err)
			return nil, err
		}

		students = append(students, &student)
	}

	return students, nil
}

func (r *Conn) GetAllStudents() ([]*app.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT * FROM students`

	rows, err := r.Conn.Query(ctx, query)

	if err != nil {
		log.Fatalf("Something went wrong with postgres.GetAllStudents: %v", err)
		return nil, err
	}

	var students []*app.Student
	for rows.Next() {
		var student app.Student

		err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.Grade,
			&student.Behavior,
			&student.Satisfaction,
			&student.Critical,
			&student.Exceptional,
			&student.Tutelary,
			&student.Staff,
		)

		if err != nil {
			log.Fatalf("Could not retrieve a [] students: %v", err)
			return nil, err
		}

		students = append(students, &student)
	}

	return students, nil
}

func (r *Conn) GetPopulation() (int, int, error) {
	var students int
	var teachers int

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	queryStudent := `select count(*) from students where staff = false`

	rows := r.Conn.QueryRow(ctx, queryStudent)

	err := rows.Scan(&students)

	if err != nil {
		log.Fatalf("Could not query the amount of students: %v", err)
		return 0, 0, err
	}

	queryTeacher := `select count(*) from teachers where staff = true`

	rows = r.Conn.QueryRow(ctx, queryTeacher)

	err = rows.Scan(&teachers)

	if err != nil {
		log.Fatalf("Could not query the amount of teachers: %v", err)
		return 0, 0, err
	}

	return teachers, students, nil
}
