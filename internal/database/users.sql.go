// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createrUser = `-- name: CreaterUser :one
INSERT INTO users(id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING id, created_at, updated_at, name, api_key
`

type CreaterUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreaterUser(ctx context.Context, arg CreaterUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createrUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getUserByAPIKey = `-- name: GetUserByAPIKey :one
SELECT id, created_at, updated_at, name, api_key FROM users WHERE api_key = $1
`

func (q *Queries) GetUserByAPIKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPIKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
