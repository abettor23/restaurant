package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ресторан.")
	fmt.Println("Определим итоговую сумму за стол к оплате.")

	fmt.Println("Введите день недели (1-7):")
	var weekDay int
	fmt.Scan(&weekDay)

	fmt.Println("Введите число гостей:")
	var guestsAmount int
	fmt.Scan(&guestsAmount)

	fmt.Println("Введите сумму чека:")
	var checkSum int
	fmt.Scan(&checkSum)

	//надбавка за обслуживание в 10% за количество гостей больше 5 человек.
	serviceCharge := (checkSum * 10) / 100

	//скидка по понедельникам (1) в 10%
	mondayDiscount := (checkSum * 10) / 100

	//скидка по пятницам (5) в 5%, если сумма чека превышает 10000
	fridayDiscount := (checkSum * 5) / 100

	if weekDay > 7 {
		fmt.Println("Введён неверный день недели. Попробуйте заново.")
	} else if weekDay == 5 {
		if checkSum > 10000 {
			if guestsAmount > 5 {
				fmt.Println("Надбавка за обслуживание:", serviceCharge)
				fmt.Println("Скидка по пятницам:", fridayDiscount)
				fmt.Println("Итоговая сумма к оплате:", (checkSum+serviceCharge)-fridayDiscount)
			} else if guestsAmount > 0 {
				fmt.Println("Скидка по пятницам:", fridayDiscount)
				fmt.Println("Итоговая сумма к оплате:", checkSum-fridayDiscount)
			} else {
				fmt.Println("Введено неверное число гостей. Попробуйте заново.")
			}
		} else if checkSum > 0 {
			if guestsAmount > 5 {
				fmt.Println("Надбавка за обслуживание:", serviceCharge)
				fmt.Println("Итоговая сумма к оплате:", checkSum+serviceCharge)
			} else if guestsAmount > 0 {
				fmt.Println("Итоговая сумма к оплате:", checkSum)
			} else {
				fmt.Println("Введено неверное число гостей. Попробуйте заново.")
			}
		} else {
			fmt.Println("Введена неверная сумма чека. Попробуйте заново.")
		}
	} else if weekDay == 1 {
		if checkSum > 0 {
			if guestsAmount > 5 {
				fmt.Println("Надбавка за обслуживание:", serviceCharge)
				fmt.Println("Скидка по понедельникам:", mondayDiscount)
				fmt.Println("Итоговая сумма к оплате:", (checkSum+serviceCharge)-mondayDiscount)
			} else if guestsAmount > 0 {
				fmt.Println("Скидка по понедельникам:", mondayDiscount)
				fmt.Println("Итоговая сумма к оплате:", checkSum-mondayDiscount)
			} else {
				fmt.Println("Введено неверное число гостей. Попробуйте заново.")
			}
		} else {
			fmt.Println("Введена неверная сумма чека. Попробуйте заново.")
		}
	} else if weekDay > 0 {
		if checkSum > 0 {
			if guestsAmount > 5 {
				fmt.Println("Надбавка за обслуживание:", serviceCharge)
				fmt.Println("Итоговая сумма к оплате:", checkSum+serviceCharge)
			} else if guestsAmount > 0 {
				fmt.Println("Итоговая сумма к оплате:", checkSum)
			} else {
				fmt.Println("Введено неверное число гостей. Попробуйте заново.")
			}
		} else {
			fmt.Println("Введена неверная сумма чека. Попробуйте заново.")
		}
	} else {
		fmt.Println("Введён неверный день недели. Попробуйте заново.")
	}
}
