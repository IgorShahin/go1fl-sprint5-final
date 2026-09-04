package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slicesData := strings.Split(datastring, ",")
	if len(slicesData) != 2 {
		return fmt.Errorf("ожидалось 2 элемента, получено: %d", len(slicesData))
	}

	steps, err := strconv.Atoi(slicesData[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("некорректное количество шагов: %d", steps)
	}

	timeData, err := time.ParseDuration(slicesData[1])
	if err != nil {
		return err
	}

	if timeData <= 0 {
		return fmt.Errorf("некорректная продолжительность: %s", timeData)
	}

	ds.Steps = steps
	ds.Duration = timeData

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
