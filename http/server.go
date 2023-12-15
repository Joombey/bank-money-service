package http

import (
	cts "farukh.go/money/constants"
	hr "farukh.go/money/http/handlers"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.POST(cts.TransferMoneyRoute, hr.TransferMoney)
	router.POST(cts.LoadMoneyRoute, hr.LoadMoneyHandler)
	router.GET(cts.GetValueRoute, hr.GetValueHandler)
	router.GET(cts.CreateCardRoute, hr.CreateNewCardHandler)
	router.GET(cts.DeleteCardRoute, hr.DeleteCardHandler)
	router.Run("0.0.0.0:8081")
}
