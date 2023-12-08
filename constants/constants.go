package constants

import "github.com/go-sql-driver/mysql"

var MySQLConfig = mysql.Config{
	User:                 "root",
	Passwd:               "root",
	DBName:               "maria_db",
	Net:                  "tcp",
	Addr:                 "localhost:3306",
	AllowNativePasswords: true,
	CheckConnLiveness:    true,
	MaxAllowedPacket:     64 << 20,
}

const (
	BaseBankApi        string = "http://localhost:8081"
	CreateCardRoute    string = "/new-card"      // GET
	TransferMoneyRoute string = "/transfer"      // POST
	GetValueRoute      string = "/get-card/:num" // GET with route argument :num
	CardNumberStart    int    = 22
)

const DatabaseSchema = `
CREATE TABLE IF NOT EXISTS moneys(
	id int primary key auto_increment,
	card_number int not null,
	value float default 0.0
);
`
