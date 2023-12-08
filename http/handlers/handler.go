package handlers

import (
	"net/http"
	"strconv"

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
	var response = [2]model.ValueResponse {
		model.ValueResponse { 
			CardNumber: request.From,
			Value: fromValue,
		},
		model.ValueResponse { 
			CardNumber: request.To,
			Value: toValue,
		},
	}
	c.IndentedJSON(http.StatusOK, &response)
}

func CreateNewCardHandler(c *gin.Context) {

}
