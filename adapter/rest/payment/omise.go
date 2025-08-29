package payment

type OmiseAPI interface{}

type omiseAPI struct {
	SecretKey string
	PublicKey string
}

func NewOmiseAPI(secretKey string, publicKey string) OmiseAPI {
	return &omiseAPI{
		SecretKey: secretKey,
		PublicKey: publicKey,
	}
}
