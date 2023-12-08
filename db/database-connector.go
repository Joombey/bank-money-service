package db

import (
	"database/sql"
	"fmt"
	"log"

	cts "farukh.go/money/constants"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", cts.MySQLConfig.FormatDSN())
	stmt, err := db.Prepare(cts.DatabaseSchema)
	if err != nil {
		log.Panicf("error creation tables %s", err.Error())
	}
	stmt.Exec()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

type MoneyRepositoryImpl struct {
	db *sql.DB
}

func (r MoneyRepositoryImpl) New() *MoneyRepositoryImpl {
	return &MoneyRepositoryImpl{db: db}
}

func (r *MoneyRepositoryImpl) InsertCard(cardNumber int) (value int) {
	r.db.Exec("INSERT INTO moneys (card_number) values(?)", cardNumber)
	r.db.QueryRow("SELECT * FROM moneys WHERE card_number = ?", cardNumber).Scan(&value)
	return value
}

func (r *MoneyRepositoryImpl) TransferMoney(from int, to int, value float32) (fromValue float32, toValue float32) {
	fromValue = r.GetValueByCard(from) - value
	toValue = r.GetValueByCard(to) + value

	go r.updateCardNumber(from, fromValue)
	go r.updateCardNumber(to, toValue)

	return
}

func (r *MoneyRepositoryImpl) GetValueByCard(cardNumber int) (value float32) {
	r.db.QueryRow("SELECT value FROM moneys WHERE card_number = ?", cardNumber).Scan(&value)
	return value
}

func (r *MoneyRepositoryImpl) updateCardNumber(cardNumber int, value float32) {
	r.db.Exec(`UPDATE moneys SET value = ? WHERE card_number = ?`, value, cardNumber)
}

