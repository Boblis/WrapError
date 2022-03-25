package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func CalUserAmountFromDb(userId string, db *sql.DB) (int, error) {
	queryCmd := `select amount from bill where user_id = ?`
	row := db.QueryRow(queryCmd, userId)

	var amount int
	if err := row.Scan(&amount); err != nil && err == sql.ErrNoRows {
		return amount, errors.Wrap(err, "table bill has no record to meets condition")
	}
	if row.Err() != nil {
		return amount, errors.Wrapf(row.Err(), "row error, %v", queryCmd)
	}
    return amount, nil
}
