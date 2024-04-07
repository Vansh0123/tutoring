-- name: RegisterStudent :one
INSERT INTO students(id, created_at, updated_at, name,subject,class,fees,fee_status)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetStudentByName :one
Select * FROM students WHERE name = $1;

-- name: GetAllStudents :many
Select * FROM students;