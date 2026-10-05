package main

import (
	"fmt"
	"math"
)

func main() {
	var bannerWidth, bannerHeigth = 12, 8
	fmt.Println("Задание 1")
	bannerArea := (bannerWidth * bannerHeigth)
	fmt.Println(bannerArea)
	halfbannerArea := (bannerArea / 2)
	fmt.Println(halfbannerArea)
	bannerBorderLength := (bannerWidth + bannerHeigth) * 2
	fmt.Println(bannerBorderLength)
	fmt.Println() // новая строка, чтобы ответы задач не кучковались

	var BoxCount = 29
	fmt.Println("Задание 2")
	fmt.Println("Количество коробок на складе:", BoxCount)
	leftoverBoxes := BoxCount % 5
	fmt.Println("Остаток от деление:", leftoverBoxes)
	fmt.Println()

	var tempMorning, tempAfternoon, tempEvening = 15, 20, 25
	fmt.Println("Задание 3")
	totalTemp := tempMorning + tempAfternoon + tempEvening
	fmt.Println("После сложение всех температур получаем:", totalTemp)
	averageTemp := totalTemp / 3
	fmt.Println("Средняя температура за день:", averageTemp)
	fmt.Println()

	var knownWords, wordsGoal = 47, 120
	fmt.Println("Задание 4")
	progressPercent := int(float64(knownWords) / float64(wordsGoal) * 100)
	fmt.Println("Процент:", progressPercent, "%")
	fmt.Println()

	var coins = 0
	fmt.Println("Задание 5")
	fmt.Println("У меня на балансе:", coins)
	coins += 500
	fmt.Println("Получил за задание:", coins)
	coins += 1200
	fmt.Println("Получил бонус, итого:", coins)
	coins /= 2
	fmt.Println("Потратил половину, итого:", coins)
	coins *= 2
	fmt.Println("Увеличил в 2 раза, итого:", coins)
	coins -= 300
	fmt.Println("Потратил 300, итого:", coins)
	fmt.Println()

	var participants, groupCount = 42, 8
	fmt.Println("Задание 6")
	fmt.Println("Участников:", participants)
	fmt.Println("Групп:", groupCount)
	participantsPerGroup := (participants / groupCount)
	fmt.Println("Итоги разделения:", participantsPerGroup)
	fmt.Println()

	fmt.Println("Задание 7")
	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)
	// в первом ответе 8, ибо сначала идёт умножение
	// во втором сначала идёт то , что в скобках, а потом уже умножение
	fmt.Println()

	var squareValue = 81
	fmt.Println("Задание 8")
	result := math.Sqrt(float64(squareValue))
	fmt.Println(result)
	multiplier := 5
	exponent := 2
	result2 := math.Pow(float64(multiplier), float64(exponent))
	fmt.Println(result2)
}
