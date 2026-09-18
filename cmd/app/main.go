package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/kama1ovv/lab4-variant03/pkg/salarycalc"
)

func main() {
	// Использование внешнего пакета uuid для генерации ID сеанса
	sessionID := uuid.New()

	// Использование внешнего пакета color для цветного вывода
	color.Cyan("=== Система расчета зарплаты (Сессия: %s) ===", sessionID.String())

	reader := bufio.NewReader(os.Stdin)

	// Ввод имени сотрудника
	fmt.Print("Введите ФИО сотрудника: ")
	employee, _ := reader.ReadString('\n')
	employee = strings.TrimSpace(employee)

	// Ввод часов
	var hours float64
	fmt.Print("Введите количество отработанных часов (например, 160): ")
	_, err := fmt.Scan(&hours)
	if err != nil {
		log.Fatalf("Ошибка ввода часов: %v", err)
	}

	// Ввод часовой ставки
	var rate float64
	fmt.Print("Введите часовую ставку ($): ")
	_, err = fmt.Scan(&rate)
	if err != nil {
		log.Fatalf("Ошибка ввода ставки: %v", err)
	}

	// Ввод процента налога
	var tax float64
	fmt.Print("Введите процент налога (0-100): ")
	_, err = fmt.Scan(&tax)
	if err != nil {
		log.Fatalf("Ошибка ввода налога: %v", err)
	}

	// Ввод процента бонуса
	var bonus float64
	fmt.Print("Введите процент бонуса (>= 0): ")
	_, err = fmt.Scan(&bonus)
	if err != nil {
		log.Fatalf("Ошибка ввода бонуса: %v", err)
	}

	fmt.Println("\n----------------------------------------")

	// 1. Вычисление GrossSalary
	gross, err := salarycalc.GrossSalary(hours, rate)
	if err != nil {
		color.Red("Ошибка GrossSalary: %v", err)
		return
	}
	color.Green("✓ Валовая зарплата успешно рассчитана.")

	// 2. Вычисление NetSalary
	net, err := salarycalc.NetSalary(gross, tax)
	if err != nil {
		color.Red("Ошибка NetSalary: %v", err)
		return
	}
	color.Green("✓ Чистая зарплата после налогов успешно рассчитана.")

	// 3. Применение бонуса через указатель
	err = salarycalc.ApplyBonus(&net, bonus)
	if err != nil {
		color.Red("Ошибка ApplyBonus: %v", err)
		return
	}
	color.Green("✓ Бонус успешно применен по указателю.")

	// 4. Форматирование и вывод отчета
	report, err := salarycalc.FormatSalaryReport(employee, gross, net)
	if err != nil {
		color.Red("Ошибка FormatSalaryReport: %v", err)
		return
	}

	color.Yellow("\n%s", report)
}
