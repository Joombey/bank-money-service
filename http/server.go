package http

import (
	cts "farukh.go/money/constants"
	hr "farukh.go/money/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/Depado/ginprom"
)

func Run() {
	router := gin.Default()

	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(p.Instrument())

	router.POST(cts.TransferMoneyRoute, hr.TransferMoney)
	router.POST(cts.LoadMoneyRoute, hr.LoadMoneyHandler)
	router.GET(cts.GetValueRoute, hr.GetValueHandler)
	router.GET(cts.CreateCardRoute, hr.CreateNewCardHandler)
	router.GET(cts.DeleteCardRoute, hr.DeleteCardHandler)
	router.Run()
}
