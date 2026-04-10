package main

import "fmt"

func main() {
	//FIRST TASK
	/*
		var amount int64

		fmt.Print("Введите сумму оплаты: ")
		fmt.Scan(&amount)

		cashback := int64(22) // 2.2% = 22 / 1000
		nasiya := int64(1)

		cashBack := amount * cashback / 1000
		cashbackNasiya := amount * nasiya / 1000

		fmt.Println("Кэшбек за сервисы (2.2%):", cashBack)
		fmt.Println("Кэшбек за рассрочку (0.1%):", cashbackNasiya)

	*/

	// SECOND TASK

	//var amount int64
	//isAzo := true
	//
	//fmt.Print("Введите сумму перевода: ")
	//fmt.Scan(&amount)
	//
	//freeLimit := int64(1_000_000)
	//var paidZone int64
	//
	//if isAzo {
	//	if amount > freeLimit {
	//		paidZone = amount - freeLimit
	//	} else {
	//		paidZone = 0
	//	}
	//} else {
	//	paidZone = amount
	//}
	//
	//commission := paidZone * 45 / 10000
	//total := amount + commission
	//
	//fmt.Println("Комиссия: ", commission)
	//fmt.Println("Итого спишется: ", total)

	sender := "Sherali"
	receiver := "Alisher"
	amount := int64(500_000)
	isAzo := true

	fmt.Printf("Отправитель: %s\n", sender)
	fmt.Printf("Получатель: %s\n", receiver)
	fmt.Printf("Сумма: %d\n", amount)
	fmt.Printf("Статус Аъзо: %t\n", isAzo)
}
