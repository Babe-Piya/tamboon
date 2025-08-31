package service

import (
	"fmt"
	"strconv"
	"time"

	"Babe-Piya/tamboo/adapter/rest/payment"
)

type SongPahPa struct {
	Name           string
	AmountSubunits int
	CCNumber       string
	CVV            string
	ExpMonth       int
	ExpYear        int
}

func (s *songPahPaService) Donate(records [][]string) {
	fmt.Println("performing donations...")
	var sum, successDonate, faultyDonate float64
	var totalDonated, firstAmount, secondAmount, thirdAmount int
	var firstName, secondName, thirdName string
	now := time.Now()

	for _, record := range records {
		if len(record) != 6 {
			fmt.Println("Skip wrong record:", record)
			continue
		}

		amount, errAtoi := strconv.Atoi(record[1])
		if errAtoi != nil {
			fmt.Printf("Error converting amount for record %v: %v\n", record, errAtoi)
			continue
		}

		expMonth, errAtoi := strconv.Atoi(record[4])
		if errAtoi != nil {
			fmt.Printf("Error converting exp month for record %v: %v\n", record, errAtoi)
			continue
		}

		expYear, errAtoi := strconv.Atoi(record[5])
		if errAtoi != nil {
			fmt.Printf("Error converting exp year for record %v: %v\n", record, errAtoi)
			continue
		}

		sum += float64(amount)

		if now.Year() > expYear {
			continue
		} else if now.Year() == expYear && int(now.Month()) > expMonth {
			continue
		}

		card, err := s.PaymentAPI.CreateToken(payment.CreateTokenRequest{
			Name:     record[0],
			CCNumber: record[2],
			ExpMonth: time.Month(expMonth),
			ExpYear:  expYear,
			CVV:      record[3],
		})
		if err != nil {
			fmt.Printf("Error creating token for record %v: %v\n", record, err)
			continue
		}

		if _, err = s.PaymentAPI.Charge(payment.ChargeRequest{
			Token:    card.Token,
			Amount:   int64(amount),
			Currency: "thb",
		}); err != nil {
			fmt.Printf("Error charging record %v: %v\n", record, err)
			continue
		}

		if amount > firstAmount {
			thirdAmount = secondAmount
			thirdName = secondName

			secondAmount = firstAmount
			secondName = firstName

			firstAmount = amount
			firstName = record[0]
		} else if amount > secondAmount && amount != firstAmount {
			thirdAmount = secondAmount
			thirdName = secondName

			secondAmount = amount
			secondName = record[0]
		} else if amount > thirdAmount && amount != firstAmount && amount != secondAmount {
			thirdAmount = amount
			thirdName = record[0]
		}

		successDonate += float64(amount)
		totalDonated = totalDonated + 1
	}

	sum = sum / float64(100)
	successDonate = successDonate / float64(100)
	faultyDonate = sum - successDonate
	average := successDonate / float64(totalDonated)

	fmt.Println("done.")
	fmt.Printf("\ntotal received: THB %.2f\n", sum)
	fmt.Printf("successfully donated: THB %.2f\n", successDonate)
	fmt.Printf("faulty donation: THB %.2f\n", faultyDonate)
	fmt.Printf("\naverage per person: THB %.2f\n", average)
	fmt.Printf("top donors: %s\n\t%s\n\t%s", firstName, secondName, thirdName)

}
