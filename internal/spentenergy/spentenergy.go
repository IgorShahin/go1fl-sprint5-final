package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be greater than zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than zero")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("number of steps must be greater than zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than zero")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength

	return distance / mInKm
}
