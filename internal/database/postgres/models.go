// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Student struct {
	ID        int64
	Name      string
	Clas      pgtype.Text
	Scool     pgtype.Text
	OrderDay  pgtype.Int2
	OrderTime pgtype.Time
	OrderCost pgtype.Int2
}
