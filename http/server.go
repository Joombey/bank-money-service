package http

import (
	cts "farukh.go/money/constants"
	hr "farukh.go/money/http/handlers"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET(cts.TransferMoneyRoute, hr.TransferMoney)
	router.GET(cts.CreateCardRoute, hr.CreateNewCardHandler)
	router.POST(cts.GetValueRoute, hr.GetValueHandler)
	router.Run("localhost:8081")
}
