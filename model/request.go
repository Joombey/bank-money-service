package model

type TransferRequest struct {
	From  int     `json:"from"`
	To    int     `json:"to"`
	Value float32 `json:"value"`
}
