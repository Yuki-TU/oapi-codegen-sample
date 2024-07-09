// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package sqlc

import (
	"context"
)

const getByUserID = `-- name: GetByUserID :one
SELECT
  ` + "`" + `id` + "`" + `,
  ` + "`" + `email` + "`" + `
FROM
  ` + "`" + `users` + "`" + `
WHERE ` + "`" + `id` + "`" + ` = ?
`

type GetByUserIDRow struct {
	ID    int64
	Email string
}

func (q *Queries) GetByUserID(ctx context.Context, id int64) (GetByUserIDRow, error) {
	row := q.db.QueryRowContext(ctx, getByUserID, id)
	var i GetByUserIDRow
	err := row.Scan(&i.ID, &i.Email)
	return i, err
}
