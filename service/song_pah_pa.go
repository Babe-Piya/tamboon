package service

import "Babe-Piya/tamboo/adapter/rest/payment"

type SongPahPaService interface {
	Donate(records [][]string)
}
type songPahPaService struct {
	PaymentAPI payment.OmiseAPI
}

func NewSongPahPaService(paymentAPI payment.OmiseAPI) SongPahPaService {
	return &songPahPaService{
		PaymentAPI: paymentAPI,
	}
}
