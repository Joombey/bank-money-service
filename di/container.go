package di

import (
	"farukh.go/money/repos"
	"farukh.go/money/db"
)

func Init() {
	(&container).new()
}

type BaseContainer struct {
	MoneyRepo repos.MoneyRepository
}

func (c *BaseContainer) new() {
	var moneyRepo = db.MoneyRepositoryImpl {}
	c.MoneyRepo = moneyRepo.New()
}

func GetContainer() BaseContainer {
	return container
}

var container = BaseContainer{}
