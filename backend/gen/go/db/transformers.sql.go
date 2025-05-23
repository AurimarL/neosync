// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transformers.sql

package db_queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pg_models "github.com/nucleuscloud/neosync/backend/sql/postgresql/models"
)

const createUserDefinedTransformer = `-- name: CreateUserDefinedTransformer :one
INSERT INTO neosync_api.transformers (
  name, description, source, account_id, transformer_config, created_by_id, updated_by_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, created_at, updated_at, name, description, account_id, transformer_config, created_by_id, updated_by_id, source
`

type CreateUserDefinedTransformerParams struct {
	Name              string
	Description       string
	Source            int32
	AccountID         pgtype.UUID
	TransformerConfig *pg_models.TransformerConfig
	CreatedByID       pgtype.UUID
	UpdatedByID       pgtype.UUID
}

func (q *Queries) CreateUserDefinedTransformer(ctx context.Context, db DBTX, arg CreateUserDefinedTransformerParams) (NeosyncApiTransformer, error) {
	row := db.QueryRow(ctx, createUserDefinedTransformer,
		arg.Name,
		arg.Description,
		arg.Source,
		arg.AccountID,
		arg.TransformerConfig,
		arg.CreatedByID,
		arg.UpdatedByID,
	)
	var i NeosyncApiTransformer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.AccountID,
		&i.TransformerConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
		&i.Source,
	)
	return i, err
}

const deleteUserDefinedTransformerById = `-- name: DeleteUserDefinedTransformerById :exec
DELETE FROM neosync_api.transformers WHERE id = $1
`

func (q *Queries) DeleteUserDefinedTransformerById(ctx context.Context, db DBTX, id pgtype.UUID) error {
	_, err := db.Exec(ctx, deleteUserDefinedTransformerById, id)
	return err
}

const getUserDefinedTransformerById = `-- name: GetUserDefinedTransformerById :one
SELECT id, created_at, updated_at, name, description, account_id, transformer_config, created_by_id, updated_by_id, source from neosync_api.transformers WHERE id = $1
`

func (q *Queries) GetUserDefinedTransformerById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiTransformer, error) {
	row := db.QueryRow(ctx, getUserDefinedTransformerById, id)
	var i NeosyncApiTransformer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.AccountID,
		&i.TransformerConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
		&i.Source,
	)
	return i, err
}

const getUserDefinedTransformersByAccount = `-- name: GetUserDefinedTransformersByAccount :many
SELECT t.id, t.created_at, t.updated_at, t.name, t.description, t.account_id, t.transformer_config, t.created_by_id, t.updated_by_id, t.source from neosync_api.transformers t
INNER JOIN neosync_api.accounts a ON a.id = t.account_id
WHERE a.id = $1
ORDER BY t.name ASC
`

func (q *Queries) GetUserDefinedTransformersByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiTransformer, error) {
	rows, err := db.Query(ctx, getUserDefinedTransformersByAccount, accountid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiTransformer
	for rows.Next() {
		var i NeosyncApiTransformer
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.AccountID,
			&i.TransformerConfig,
			&i.CreatedByID,
			&i.UpdatedByID,
			&i.Source,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isTransformerNameAvailable = `-- name: IsTransformerNameAvailable :one
SELECT count(t.id) from neosync_api.transformers t
INNER JOIN neosync_api.accounts a ON a.id = t.account_id
WHERE a.id = $1 and t.name = $2
`

type IsTransformerNameAvailableParams struct {
	AccountId       pgtype.UUID
	TransformerName string
}

func (q *Queries) IsTransformerNameAvailable(ctx context.Context, db DBTX, arg IsTransformerNameAvailableParams) (int64, error) {
	row := db.QueryRow(ctx, isTransformerNameAvailable, arg.AccountId, arg.TransformerName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateUserDefinedTransformer = `-- name: UpdateUserDefinedTransformer :one
UPDATE neosync_api.transformers
SET
  name = $1,
  description = $2,
  transformer_config = $3,
  updated_by_id = $4
WHERE id = $5
RETURNING id, created_at, updated_at, name, description, account_id, transformer_config, created_by_id, updated_by_id, source
`

type UpdateUserDefinedTransformerParams struct {
	Name              string
	Description       string
	TransformerConfig *pg_models.TransformerConfig
	UpdatedByID       pgtype.UUID
	ID                pgtype.UUID
}

func (q *Queries) UpdateUserDefinedTransformer(ctx context.Context, db DBTX, arg UpdateUserDefinedTransformerParams) (NeosyncApiTransformer, error) {
	row := db.QueryRow(ctx, updateUserDefinedTransformer,
		arg.Name,
		arg.Description,
		arg.TransformerConfig,
		arg.UpdatedByID,
		arg.ID,
	)
	var i NeosyncApiTransformer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.AccountID,
		&i.TransformerConfig,
		&i.CreatedByID,
		&i.UpdatedByID,
		&i.Source,
	)
	return i, err
}
