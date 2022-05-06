package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ara-Chobanyan/school-manager/school-parent/app"
	"github.com/jackc/pgx/v4"
)

// Create a struct of the pgx connection
type Conn struct {
	Conn *pgx.Conn
}

//Open up the database with a link later on need to add env
func openConn() (*pgx.Conn, error) {
	dsn := "postgres://user:password@localhost/gokit?sslmode=disable"

	// have context to make sure we are not stuck in a connection port
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// call the cancel to release resources
	defer cancel()

	// Open up a pgx connection
	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
		return nil, err
	}

	// return a pointer to the pgx connection
	return conn, err
}

//create our repo and return a interface to inject into later
func NewPostgresRepository() (app.ParentRepository, error) {
	// initlized our struct that takes in a pgx connecting
	repo := &Conn{}

	// open up a connection
	client, err := openConn()

	// check error if we could not open up the connection for any reason
	if err != nil {
		log.Fatalf("Could not get the clien %v", err)
		return nil, err
	}

	// set the client to our struct
	repo.Conn = client

	// return our interface
	return repo, nil
}

// Adds a new student to the database, creates a table if there isn't any
func (r *Conn) CreateStudent(student app.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	createSQL := `
  create table if not exists students(
  id SERIAL PRIMARY KEY,
  first_name text,
  last_name text,
  grade float,
  behavior float,
  satisfaction float,
  critical bool,
  exceptional bool,
  tutelary bool,
  staff bool
  );
  `

	_, err := r.Conn.Exec(ctx, createSQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	tx, err := r.Conn.Begin(ctx)

	if err != nil {
		log.Fatalf("r.Conn.Being failed %v", err)
	}

	_, err = tx.Exec(ctx, `insert into students(
  first_name,
  last_name,
  grade,
  behavior,
  satisfaction,
  critical,
  exceptional,
  tutelary,
  staff) values ($1,$2,$3,$4,$5,$6,$7,$8,$9);
    `, student.FirstName, student.LastName, student.Grade, student.Behavior, student.Satisfaction, student.Critical, student.Exceptional, student.Tutelary, student.Staff)

	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("tx.Commit failed: %v", err)
	}

	return nil
}

// takes in a id and deletes a student with that id
func (r *Conn) DeleteStudent(id int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.Conn.Query(ctx, "DELETE from students WHERE id= $1;", id)

	if err != nil {
		log.Fatalf("r.Conn.QueryRow failed: %v", err)
		return err
	}

	defer rows.Close()

	return nil
}

// takes in a student payload and makes edit with that payload to the related id
func (r *Conn) EditStudent(student app.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.Conn.Query(ctx, `UPDATE students SET first_name = $1, last_name = $2, 
    grade = $3, behavior = $4, satisfaction = $5, critical = $6, exceptional = $7, tutelary = $8 WHERE id = $9;`, student.FirstName, student.LastName, student.Grade, student.Behavior, student.Satisfaction, student.Critical, student.Exceptional, student.Tutelary, student.Id)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
