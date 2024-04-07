-- +goose Up

CREATE TABLE students (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    subject TEXT NOT NULL,
    class TEXT NOT NULL,
    fees INTEGER NOT NULL,
    fee_status TEXT NOT NULL
);

-- +goose Down

DROP TABLE students;