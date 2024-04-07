package database

import (
	"context"
	"fmt"
	"log"
	"strings"
)

type SearchParams struct {
	Name      string
	Subject   string
	Class     string
	FeeStatus string
}

func (q *Queries) Search(ctx context.Context, arg SearchParams) ([]Student, error) {

	queryParams := make([]interface{}, 0)
	queryBuilder := strings.Builder{}

	// Start with the base query
	queryBuilder.WriteString("Select id, created_at, updated_at, name, subject, class, fees, fee_status FROM students WHERE 1=1")

	// Dynamically add conditions
	if arg.Name != "" {
		queryParams = append(queryParams, arg.Name)
		queryBuilder.WriteString(fmt.Sprintf(" AND name = $%d", len(queryParams)))
	}
	if arg.Subject != "" {
		queryParams = append(queryParams, arg.Subject)
		queryBuilder.WriteString(fmt.Sprintf(" AND subject = $%d", len(queryParams)))
	}
	if arg.Class != "" {
		queryParams = append(queryParams, arg.Class)
		queryBuilder.WriteString(fmt.Sprintf(" AND class = $%d", len(queryParams)))
	}
	if arg.FeeStatus != "" {
		queryParams = append(queryParams, arg.FeeStatus)
		queryBuilder.WriteString(fmt.Sprintf(" AND fee_status = $%d", len(queryParams)))
	}

	log.Print(queryBuilder.String())
	rows, err := q.db.QueryContext(ctx, queryBuilder.String(), queryParams...)

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
