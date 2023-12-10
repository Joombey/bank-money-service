package tests

import (
	"database/sql"
	"testing"
	"time"

	_ "farukh.go/money/db"
	"farukh.go/money/di"
	in "farukh.go/money/internal"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var repo = di.GetContainer().MoneyRepo
var cfg = in.ObtainConfig()
var data, err = sql.Open(cfg.DbConfig.Driver, cfg.DbConfig.Path)

func TestCardInsertion(t *testing.T) {
	testCard := 1
	cardId := repo.InsertCard(testCard)
	var value float32
	var resultCard int
	err = data.QueryRow("SELECT card_number, value FROM moneys WHERE id = ($1)", cardId).Scan(&resultCard, &value)
	printAll()
	if err != nil || value != 0 || resultCard != testCard {
		t.Errorf("expected %f %d got %f %d or err = %s", 0.0, testCard, value, resultCard, err.Error())
	}
}

func TestMoneyInsertion(t *testing.T) {
	var insertedValue float32
	testCard := 2

	repo.InsertCard(testCard)
	newValue := repo.InsertMoney(testCard, 1000)

	printAll()

	err = data.QueryRow("SELECT value FROM moneys WHERE card_number = ($1)", testCard).Scan(&insertedValue)

	if err != nil || insertedValue != newValue {
		t.Errorf("expected %f got %f or err = %t", newValue, insertedValue, err == nil)
	}
}

func TestGetValueByCard(t *testing.T) {
	testCard := 3
	var value float32 = 0.0
	repo.InsertCard(testCard)
	repo.InsertMoney(testCard, 1000)
	time.Sleep(time.Second * 2)
	repo.GetValueByCard(testCard)
	printAll()
	err = data.QueryRow("SELECT value FROM moneys WHERE card_number = ($1)", testCard).Scan(&value)

	if err != nil || value != 1000 {
		t.Errorf("expected value = %d got %f", 1000, value)
	}
}

func TestTransfer(t *testing.T) {
	testCard1 := 4
	testCard2 := 5

	repo.InsertCard(testCard1)
	repo.InsertCard(testCard2)

	repo.InsertMoney(testCard1, 1000)

	time.Sleep(time.Second * 2)

	tv1, tv2 := repo.TransferMoney(testCard1, testCard2, 500)

	time.Sleep(time.Second * 2)

	var (
		v1 float32
		v2 float32
	)

	data.QueryRow("SELECT value from moneys where card_number = ($1)", testCard1).Scan(&v1)
	data.QueryRow("SELECT value from moneys where card_number = ($1)", testCard2).Scan(&v2)

	if tv1 != 500 || tv2 != 500 || tv1 != v1 || tv2 != v2 || v1 != 500 || v2 != 500 {
		t.Errorf("expected 500 500 got %f %f", v1, v2)
	}
}

func TestGetLastCard(t *testing.T) {
	testCard := 6
	
	repo.InsertCard(testCard)

	lastCard, err1 := repo.GetLatestCardNumber()

	var dbCard int
	data.QueryRow("SELECT card_number FROM moneys GROUP BY card_number ORDER BY card_number DESC LIMIT 1").Scan(&dbCard)

	if err1 != nil || lastCard != testCard || dbCard != lastCard {
		t.Errorf("expected %d got %d or err = %t", testCard, dbCard, err1 == nil)
	}
}

func printAll() {
	rows, _ := data.Query("SELECT * FROM moneys")
	// println(errQ.Error())
	for rows.Next() {
		var (
			id    int
			card  int
			value float32
		)
		rows.Scan(&id, &card, &value)

		println(id, card, value)
	}
}
