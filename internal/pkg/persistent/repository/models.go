package repository

import "github.com/jackc/pgx/v5/pgtype"

type Product struct {
	ID pgtype.Int4 `db:"notification_status"`
}
