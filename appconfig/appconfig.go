package appconfig

type AppConfig struct {
	Omise OmiseConfig
}

type OmiseConfig struct {
	PublicKey string
	SecretKey string
}
