package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

// Определение структуры карты
type Card struct {
	value int    // значение карты
	suit  string // масть карты
}

// Главная функция
func main() {
	run(os.Stdin, os.Stdout) // Запуск основной функции с вводом и выводом
}

// Основная функция
func run(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r) // Создание нового буферизованного ридера
	writer := bufio.NewWriter(w) // Создание нового буферизованного райтера
	defer writer.Flush()         // Очистка буфера после завершения функции

	var t int
	fmt.Fscan(reader, &t) // Считывание количества наборов входных данных

	// Цикл по каждому набору входных данных
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n) // Считывание количества игроков

		// Создание двумерного массива для хранения карт каждого игрока
		hands := make([][]Card, n)
		for j := 0; j < n; j++ {
			hands[j] = make([]Card, 2) // Каждый игрок получает две карты
			for k := 0; k < 2; k++ {
				var card string
				fmt.Fscan(reader, &card)      // Считывание карты
				hands[j][k] = parseCard(card) // Преобразование строки в структуру карты
			}
			// Сортировка карт каждого игрока по убыванию значения
			sort.Slice(hands[j], func(a, b int) bool {
				return hands[j][a].value > hands[j][b].value
			})
		}

		// Поиск выигрышных карт
		winCards := findWinningCards(hands)
		fmt.Fprintln(writer, len(winCards)) // Вывод количества выигрышных карт
		for _, card := range winCards {
			// Вывод каждой выигрышной карты
			fmt.Fprintln(writer, valueToString(card.value)+card.suit)
		}
	}
}

// Функция для преобразования строки в структуру карты
func parseCard(card string) Card {
	// Словарь для преобразования символа в значение карты
	valueMap := map[string]int{
		"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
		"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
	}
	// Возвращение структуры карты
	return Card{value: valueMap[string(card[0])], suit: string(card[1])}
}

// Функция для преобразования значения карты в строку
func valueToString(value int) string {
	// Словарь для преобразования значения карты в символ
	stringMap := map[int]string{
		14: "A", 13: "K", 12: "Q", 11: "J", 10: "T",
		9: "9", 8: "8", 7: "7", 6: "6", 5: "5", 4: "4", 3: "3", 2: "2",
	}
	// Возвращение символа карты
	return stringMap[value]
}

// Функция для поиска выигрышных карт
func findWinningCards(hands [][]Card) []Card {
	winCards := []Card{}                  // Массив для хранения выигрышных карт
	suits := []string{"S", "C", "D", "H"} // Массив мастей
	// Цикл по каждому значению карты
	for i := 2; i <= 14; i++ {
		// Цикл по каждой масти
		for _, suit := range suits {
			card := Card{value: i, suit: suit} // Создание карты
			// Если карта является выигрышной, добавить ее в массив
			if isWinningCard(card, hands) {
				winCards = append(winCards, card)
			}
		}
	}
	// Возвращение массива выигрышных карт
	return winCards
}

// Функция для проверки, является ли карта выигрышной
func isWinningCard(card Card, hands [][]Card) bool {
	// Если карта уже в руках игроков, вернуть false
	if isCardInHands(card, hands) {
		return false
	}
	// Создание руки первого игрока с новой картой
	playerHand1 := []Card{card, hands[0][0], hands[0][1]}
	// Сортировка карт по убыванию значения
	sort.Slice(playerHand1, func(a, b int) bool {
		return playerHand1[a].value > playerHand1[b].value
	})
	// Вычисление очков первого игрока
	playerScore := scoreHand(playerHand1)
	// Цикл по рукам остальных игроков
	for i := 1; i < len(hands); i++ {
		// Создание руки i-го игрока с новой картой
		otherHand := []Card{card, hands[i][0], hands[i][1]}
		// Сортировка карт по убыванию значения
		sort.Slice(otherHand, func(a, b int) bool {
			return otherHand[a].value > otherHand[b].value
		})
		// Вычисление очков i-го игрока
		otherScore := scoreHand(otherHand)
		// Если очки i-го игрока больше очков первого игрока, вернуть false
		if otherScore > playerScore {
			return false
		}
	}
	// Если карта является выигрышной для первого игрока, вернуть true
	return true
}

// Функция для проверки, есть ли карта в руках игроков
func isCardInHands(card Card, hands [][]Card) bool {
	// Цикл по каждому игроку
	for _, hand := range hands {
		// Цикл по каждой карте в руках игрока
		for _, c := range hand {
			// Если карта найдена, вернуть true
			if c.value == card.value && c.suit == card.suit {
				return true
			}
		}
	}
	// Если карта не найдена, вернуть false
	return false
}

// Функция для вычисления очков руки
func scoreHand(hand []Card) int {
	// Если все карты имеют одно и то же значение, вернуть 300 + значение карты (сет)
	if hand[0].value == hand[1].value && hand[1].value == hand[2].value {
		return 300 + hand[0].value // Сет
		// Если две карты имеют одно и то же значение, вернуть 200 + максимальное значение карты (пара)
	} else if hand[0].value == hand[1].value || hand[1].value == hand[2].value {
		return 200 + hand[1].value // Пара
	} else if hand[0].value == hand[2].value {
		return 200 + hand[0].value // Пара
		// Иначе, вернуть 100 + максимальное значение карты (старшая карта)
	} else {
		return 100 + max(hand[0].value, max(hand[1].value, hand[2].value)) // Старшая карта
	}
}

// Функция для вычисления максимального из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
