package postgres

import (
	"context"
	"log"
	"time"

	"github.com/Ara-Chobanyan/school-manager/global-grade-service/app"
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
		log.Fatalf("Could not connect: %v", err)
		return nil, err
	}

	return conn, err
}

func NewPostgresRepository() (app.GlobalRepo, error) {
	repo := &Conn{}

	client, err := openConn()

	if err != nil {
		log.Fatalf("Could not open connection: %v", err)
		return nil, err
	}

	repo.Conn = client

	return repo, nil
}

func (r *Conn) GlobalAverage() (app.GlobalStatus, error) {
	var grade app.GlobalStatus

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `SELECT AVG(grade),AVG(behavior) FROM students`

	rows := r.Conn.QueryRow(ctx, query)

	err := rows.Scan(&grade.Grade, &grade.Behaviour)

	if err != nil {
		log.Fatalf("Could not get GlobalStatus: %v", err)
	}

	return grade, nil
}
