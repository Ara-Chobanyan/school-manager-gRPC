package postgres

import (
	"context"
	"log"
	"time"

	"github.com/Ara-Chobanyan/school-manager/school-manager-service/app"
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
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	return conn, nil
}

func NewPostgresRepository() (app.ManagerRepository, error) {
	repo := &Conn{}

	client, err := openConn()

	if err != nil {
		log.Fatalf("Could not get client: %v", err)
		return nil, err
	}

	repo.Conn = client

	return repo, nil
}

func (r *Conn) Behavior() ([]*app.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT * from students where grade < 60 and behavior < 50`

	rows, err := r.Conn.Query(ctx, query)

	if err != nil {
		log.Fatalf("Could not retrive student list: %v", err)
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
			log.Fatalf("Could not scan student row: %v", err)
			return nil, err
		}

		students = append(students, &student)
	}

	return students, nil
}

func (r *Conn) Exceptional() ([]*app.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT * from students where grade > 80 and behavior > 70`

	rows, err := r.Conn.Query(ctx, query)

	if err != nil {
		log.Fatalf("Could not retrive student list: %v", err)
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
			log.Fatalf("Could not scan student row: %v", err)
			return nil, err
		}

		students = append(students, &student)
	}

	return students, nil
}

func (r *Conn) FlagHelper(status bool, who string) error {
	if !status {
		return nil
	}

	switch {
	case who == "behavior":

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		query := `UPDATE students SET critical = true WHERE grade < 60 AND behavior < 50`

		_, err := r.Conn.Exec(ctx, query)

		if err != nil {
			log.Fatalf("could not change status: %v", err)
			return err
		}

	case who == "exceptional":

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		query := `UPDATE students SET exceptional = true WHERE grade > 80 AND behavior > 70`

		_, err := r.Conn.Exec(ctx, query)

		if err != nil {
			log.Fatalf("could not change status: %v", err)
			return err
		}

	default:
		return nil
	}

	return nil
}
