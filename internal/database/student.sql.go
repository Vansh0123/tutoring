// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: student.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getAllStudents = `-- name: GetAllStudents :many
Select id, created_at, updated_at, name, subject, class, fees, fee_status FROM students
`

func (q *Queries) GetAllStudents(ctx context.Context) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getAllStudents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Subject,
			&i.Class,
			&i.Fees,
			&i.FeeStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudentByName = `-- name: GetStudentByName :one
Select id, created_at, updated_at, name, subject, class, fees, fee_status FROM students WHERE name = $1
`

func (q *Queries) GetStudentByName(ctx context.Context, name string) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByName, name)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Subject,
		&i.Class,
		&i.Fees,
		&i.FeeStatus,
	)
	return i, err
}

const registerStudent = `-- name: RegisterStudent :one
INSERT INTO students(id, created_at, updated_at, name,subject,class,fees,fee_status)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING id, created_at, updated_at, name, subject, class, fees, fee_status
`

type RegisterStudentParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Subject   string
	Class     string
	Fees      int32
	FeeStatus string
}

func (q *Queries) RegisterStudent(ctx context.Context, arg RegisterStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, registerStudent,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Subject,
		arg.Class,
		arg.Fees,
		arg.FeeStatus,
	)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Subject,
		&i.Class,
		&i.Fees,
		&i.FeeStatus,
	)
	return i, err
}
