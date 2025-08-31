package payment

import (
	"time"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type CreateTokenRequest struct {
	Name     string     `json:"name"`
	CCNumber string     `json:"cc_number"`
	ExpMonth time.Month `json:"exp_month"`
	ExpYear  int        `json:"exp_year"`
	CVV      string     `json:"cvv"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}

func (o *omiseAPI) CreateToken(req CreateTokenRequest) (CreateTokenResponse, error) {
	client, _ := omise.NewClient(
		o.PublicKey,
		o.SecretKey,
	)

	result := &omise.Card{}

	err := client.Do(result, &operations.CreateToken{
		Name:            req.Name,
		Number:          req.CCNumber,
		ExpirationMonth: req.ExpMonth,
		ExpirationYear:  req.ExpYear,
		SecurityCode:    req.CVV,
	})

	if err != nil {
		return CreateTokenResponse{}, err
	}

	return CreateTokenResponse{
		Token: result.ID,
	}, nil
}
