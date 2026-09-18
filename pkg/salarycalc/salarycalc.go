// Package salarycalc предоставляет инструменты для расчета заработной платы сотрудников,
// включая вычисление валовой и чистой зарплаты, применение бонусов через указатели
// и генерацию отчетов.
package salarycalc

import (
	"fmt"
)

// GrossSalary вычисляет валовую зарплату (Gross) на основе отработанных часов и часовой ставки.
// Возвращает ошибку, если часы или ставка имеют отрицательное значение.
func GrossSalary(hours, rate float64) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("ошибка: количество часов не может быть отрицательным (%.2f)", hours)
	}
	if rate < 0 {
		return 0, fmt.Errorf("ошибка: часовая ставка не может быть отрицательной (%.2f)", rate)
	}
	return hours * rate, nil
}

// NetSalary вычисляет чистую зарплату (Net) путем вычета налогов из валовой зарплаты.
// Возвращает ошибку, если валовая зарплата отрицательная или процент налога вне диапазона [0, 100].
func NetSalary(gross, taxPercent float64) (float64, error) {
	if gross < 0 {
		return 0, fmt.Errorf("ошибка: валовая зарплата не может быть отрицательной (%.2f)", gross)
	}
	if taxPercent < 0 || taxPercent > 100 {
		return 0, fmt.Errorf("ошибка: процент налога должен быть от 0 до 100 (получено: %.2f)", taxPercent)
	}
	taxAmount := gross * (taxPercent / 100)
	return gross - taxAmount, nil
}

// ApplyBonus изменяет значение чистой зарплаты по указателю, добавляя к ней процент бонуса.
// Возвращает ошибку, если передан nil-указатель или процент бонуса отрицательный.
func ApplyBonus(net *float64, bonusPercent float64) error {
	if net == nil {
		return fmt.Errorf("ошибка: указатель на чистую зарплату не может быть nil")
	}
	if bonusPercent < 0 {
		return fmt.Errorf("ошибка: процент бонуса не может быть отрицательным (%.2f)", bonusPercent)
	}
	bonusAmount := (*net) * (bonusPercent / 100)
	*net += bonusAmount
	return nil
}

// FormatSalaryReport формирует текстовый отчет по зарплате сотрудника с использованием fmt.Sprintf.
// Возвращает ошибку, если имя сотрудника пустое или суммы отрицательные.
func FormatSalaryReport(employee string, gross, net float64) (string, error) {
	if employee == "" {
		return "", fmt.Errorf("ошибка: имя сотрудника не может быть пустым")
	}
	if gross < 0 || net < 0 {
		return "", fmt.Errorf("ошибка: суммы зарплат не могут быть отрицательными")
	}

	return fmt.Sprintf("Отчет по сотруднику: %s\n  Валовая зарплата (Gross): $%.2f\n  Итоговая чистая зарплата (Net + бонус): $%.3f",
		employee, gross, net), nil
}
