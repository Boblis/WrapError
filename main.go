package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"log"
	"os"
	"warpError/dao"
)

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"123456",
		"127.0.0.1",
		"test",
	)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("open database error: %v", err)
	}
	userId := "001101"
	amount, err := GetUserAmount(userId, db)
	if err != nil {
		log.Printf("error type: %T, %v\n", errors.Cause(err), errors.Cause(err))
		log.Printf("trace stack: \n%+v\n", err)
		os.Exit(1)
	}
	log.Printf("got the amount: %v of user: %v\n", amount, userId)
}

func GetUserAmount(userId string, db *sql.DB) (int, error) {
	if db == nil {
		return 0, errors.New("invalid db")
	}
	amount, err := dao.CalUserAmountFromDb(userId, db)
	if err != nil {
		return amount, errors.Wrapf(err, "get user amount error, user_id: %s", userId)
	}
	return amount, err
}
