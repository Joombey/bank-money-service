package constants

import "github.com/go-sql-driver/mysql"

var MySQLConfig = mysql.Config{
	User:                 "root",
	Passwd:               "root",
	DBName:               "maria_db",
	Net:                  "tcp",
	Addr:                 "bank-db:3306",
	AllowNativePasswords: true,
	CheckConnLiveness:    true,
	MaxAllowedPacket:     64 << 20,
}

const (
	BaseBankApi        string = "http://" + RunAddress
	RunAddress         string = "localhost:8081"
	CreateCardRoute    string = "/new-card"      // GET
	TransferMoneyRoute string = "/transfer"      // POST
	GetValueRoute      string = "/get-card/:num" // GET with route argument :num
	LoadMoneyRoute     string = "/load-money"    // POST
	CardNumberStart    int    = 22_000_000
	LocalConfigPath    string = "I:/dev/go-projects/bank-money-service/configs/local.yaml"
)

const DatabaseSchema = `
CREATE TABLE IF NOT EXISTS moneys(
	id Serial primary key,
	card_number integer not null,
	value numeric default 0.0
);
`
