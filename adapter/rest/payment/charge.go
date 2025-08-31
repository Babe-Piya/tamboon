package payment

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type ChargeRequest struct {
	Token    string `json:"token"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type ChargeResponse struct {
	Status string `json:"status"`
	Amount int64  `json:"amount"`
}

func (o *omiseAPI) Charge(req ChargeRequest) (ChargeResponse, error) {
	client, _ := omise.NewClient(
		o.PublicKey,
		o.SecretKey,
	)

	result := &omise.Charge{}

	err := client.Do(result, &operations.CreateCharge{
		Card:     req.Token,
		Amount:   req.Amount,
		Currency: req.Currency,
	})

	if err != nil {
		return ChargeResponse{}, err
	}

	return ChargeResponse{
		Status: string(result.Status),
		Amount: result.Amount,
	}, nil
}
