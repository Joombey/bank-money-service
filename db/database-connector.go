package db

import (
	"database/sql"
	"log"

	cts "farukh.go/money/constants"
	in "farukh.go/money/internal"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var localDb *sql.DB

func init() {
	cfg := in.ObtainConfig()
	db, err := sql.Open(cfg.DbConfig.Driver, cfg.DbConfig.Path)
	if cfg.Env == "test-db-repo" {
		db.Exec("DROP TABLE IF EXISTS moneys")
	}
	defer func() { localDb = db }()
	if err != nil || db == nil {
		log.Panicf("error opening db %s", err.Error())
	}
	stmt, err := db.Prepare(cts.DatabaseSchema)
	if err != nil {
		log.Panicf("error creation tables %s", err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	}
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
}

type MoneyRepositoryImpl struct {
	db *sql.DB
}

func (r MoneyRepositoryImpl) New() *MoneyRepositoryImpl {
	return &MoneyRepositoryImpl{db: localDb}
}

func (r *MoneyRepositoryImpl) InsertCard(cardNumber int) (value int) {
	_, err := r.db.Exec("INSERT INTO moneys (card_number) values ($1)", cardNumber)
	if err != nil {
		panic(err.Error())
	}
	err = r.db.QueryRow("SELECT id FROM moneys WHERE card_number = ($1)", cardNumber).Scan(&value)
	if err != nil {
		panic(err.Error())
	}
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
	err := r.db.QueryRow("SELECT value FROM moneys WHERE card_number = ($1)", cardNumber).Scan(&value)
	if err != nil {
		panic(err.Error())
	}
	return value
}

func (r *MoneyRepositoryImpl) GetLatestCardNumber() (cardNumber int, err error) {
	err = r.db.QueryRow(`SELECT card_number FROM moneys ORDER BY card_number DESC LIMIT 1`).Scan(&cardNumber)
	return
}

func (r *MoneyRepositoryImpl) InsertMoney(cardNumber int, money float32) float32 {
	value := r.GetValueByCard(cardNumber) + money
	go r.updateCardNumber(cardNumber, value)
	return value
}

func (r *MoneyRepositoryImpl) updateCardNumber(cardNumber int, value float32) {
	r.db.Exec(`UPDATE moneys SET value = ($1) WHERE card_number = ($2)`, value, cardNumber)
}
