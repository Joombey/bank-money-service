package model

type TransferRequest struct {
	From  int     `json:"from"`
	To    int     `json:"to"`
	Value float32 `json:"value"`
}

type InsertRequest struct {
	CardNumber int     `json:"card_number"`
	Value      float32 `json:"value"`
}
