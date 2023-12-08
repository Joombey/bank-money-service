package repos

type MoneyRepository interface {
	InsertCard(cardNumber int) int
	GetValueByCard(cardNumber int) float32
	TransferMoney(from int, to int, value float32) (float32, float32)
}