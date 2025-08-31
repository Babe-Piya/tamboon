package payment

type OmiseAPI interface {
	CreateToken(req CreateTokenRequest) (CreateTokenResponse, error)
	Charge(req ChargeRequest) (ChargeResponse, error)
}

type omiseAPI struct {
	PublicKey string
	SecretKey string
}

func NewOmiseAPI(publicKey string, secretKey string) OmiseAPI {
	return &omiseAPI{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}
}
