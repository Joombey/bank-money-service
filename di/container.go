package di

import (
	"farukh.go/money/db"
	"farukh.go/money/repos"
)

type BaseContainer struct {
	MoneyRepo repos.MoneyRepository
}

func (c *BaseContainer) new() {
	c.MoneyRepo = db.MoneyRepositoryImpl{}.New()
}

func GetContainer() *BaseContainer {
	if (container == BaseContainer{}) {
		(&container).new()
	}
	return &container
}

var container = BaseContainer{}
