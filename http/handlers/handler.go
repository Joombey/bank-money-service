package handlers

import (
	"net/http"
	"strconv"

	cts "farukh.go/money/constants"
	"farukh.go/money/di"
	"farukh.go/money/model"
	"github.com/gin-gonic/gin"
)

var repo = di.GetContainer().MoneyRepo

func GetValueHandler(c *gin.Context) {
	cardNumber, _ := strconv.Atoi(c.Param("num"))
	value := repo.GetValueByCard(cardNumber)
	var response = model.ValueResponse{
		CardNumber: cardNumber,
		Value:      value,
	}
	c.IndentedJSON(http.StatusOK, &response)
}

func TransferMoney(c *gin.Context) {
	var request model.TransferRequest
	c.BindJSON(&request)
	fromValue, toValue := repo.TransferMoney(request.From, request.To, request.Value)
	var response = [2]model.ValueResponse{
		{
			CardNumber: request.From,
			Value:      fromValue,
		},
		{
			CardNumber: request.To,
			Value:      toValue,
		},
	}
	c.IndentedJSON(http.StatusOK, &response)
}

func CreateNewCardHandler(c *gin.Context) {
	var newCard int
	latestCard, err := repo.GetLatestCardNumber()
	if err != nil {
		newCard = cts.CardNumberStart
	} else {
		newCard = latestCard + 1
	}
	go repo.InsertCard(newCard)
	c.IndentedJSON(http.StatusOK, model.ValueResponse{CardNumber: newCard, Value: 0})
}

func LoadMoneyHandler(c *gin.Context) {
	var request model.InsertRequest
	c.BindJSON(&request)
	newValue := repo.InsertMoney(request.CardNumber, request.Value)
	response := model.ValueResponse{CardNumber: request.CardNumber, Value: newValue}
	c.IndentedJSON(http.StatusOK, response)
}
